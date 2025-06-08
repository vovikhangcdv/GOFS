// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICompliance} from "../interfaces/ICompliance.sol";
import {EntityRegistry} from "../EntityRegistry.sol";
import {Entity, EntityType, TxType} from "../interfaces/ITypes.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title TransactionTypeCompliance
 * @dev Compliance module that manages which entity types are allowed to send/receive for specific transaction types
 */
contract TransactionTypeCompliance is ICompliance, AccessControlUpgradeable {
    using EnumerableSet for EnumerableSet.UintSet;

    // Role for managing transaction type policies
    bytes32 public constant TX_TYPE_ADMIN_ROLE =
        keccak256("TX_TYPE_ADMIN_ROLE");

    // The reference to the Entity Registry (changed from immutable)
    EntityRegistry public entityRegistry;

    // Mappings to store allowed entity types for each transaction type
    mapping(uint8 => EnumerableSet.UintSet) private _allowedFromEntityTypes;
    mapping(uint8 => EnumerableSet.UintSet) private _allowedToEntityTypes;

    // Events
    event AllowedFromEntityTypeAdded(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event AllowedFromEntityTypeRemoved(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event AllowedToEntityTypeAdded(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event AllowedToEntityTypeRemoved(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event TransactionTypeCleared(TxType indexed txType);

    // Custom errors
    error InvalidEntityRegistryAddress();
    error EntityTypeAlreadyAllowed(EntityType entityType);
    error EntityTypeNotAllowed(EntityType entityType);
    error EmptyEntityTypesList();
    error InvalidArrayLengths();
    error TransactionTypeNotUsable(TxType txType);
    error ZeroAddress();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address _entityRegistry) public initializer {
        if (_entityRegistry == address(0))
            revert InvalidEntityRegistryAddress();

        __AccessControl_init();

        entityRegistry = EntityRegistry(_entityRegistry);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(TX_TYPE_ADMIN_ROLE, msg.sender);
    }

    /**
     * @dev Implements IERC165 interface detection
     */
    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        virtual
        override(AccessControlUpgradeable, IERC165)
        returns (bool)
    {
        return
            interfaceId == type(ICompliance).interfaceId ||
            interfaceId == type(AccessControlUpgradeable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @dev Add allowed from entity types for a transaction type
     * @param txType The transaction type
     * @param entityTypes The entity types to allow as senders
     */
    function addAllowedFromEntityTypes(
        TxType txType,
        EntityType[] memory entityTypes
    ) external onlyRole(TX_TYPE_ADMIN_ROLE) {
        if (entityTypes.length == 0) revert EmptyEntityTypesList();

        uint8 txTypeId = TxType.unwrap(txType);
        for (uint256 i = 0; i < entityTypes.length; i++) {
            uint8 entityTypeId = EntityType.unwrap(entityTypes[i]);

            if (!_allowedFromEntityTypes[txTypeId].add(entityTypeId)) {
                revert EntityTypeAlreadyAllowed(entityTypes[i]);
            }
            emit AllowedFromEntityTypeAdded(txType, entityTypes[i]);
        }
    }

    /**
     * @dev Add allowed to entity types for a transaction type
     * @param txType The transaction type
     * @param entityTypes The entity types to allow as receivers
     */
    function addAllowedToEntityTypes(
        TxType txType,
        EntityType[] memory entityTypes
    ) external onlyRole(TX_TYPE_ADMIN_ROLE) {
        if (entityTypes.length == 0) revert EmptyEntityTypesList();

        uint8 txTypeId = TxType.unwrap(txType);
        for (uint256 i = 0; i < entityTypes.length; i++) {
            uint8 entityTypeId = EntityType.unwrap(entityTypes[i]);

            if (!_allowedToEntityTypes[txTypeId].add(entityTypeId)) {
                revert EntityTypeAlreadyAllowed(entityTypes[i]);
            }
            emit AllowedToEntityTypeAdded(txType, entityTypes[i]);
        }
    }

    /**
     * @dev Remove allowed from entity types for a transaction type
     * @param txType The transaction type
     * @param entityTypes The entity types to remove from allowed senders
     */
    function removeAllowedFromEntityTypes(
        TxType txType,
        EntityType[] memory entityTypes
    ) external onlyRole(TX_TYPE_ADMIN_ROLE) {
        if (entityTypes.length == 0) revert EmptyEntityTypesList();

        uint8 txTypeId = TxType.unwrap(txType);
        for (uint256 i = 0; i < entityTypes.length; i++) {
            uint8 entityTypeId = EntityType.unwrap(entityTypes[i]);

            if (!_allowedFromEntityTypes[txTypeId].remove(entityTypeId)) {
                revert EntityTypeNotAllowed(entityTypes[i]);
            }
            emit AllowedFromEntityTypeRemoved(txType, entityTypes[i]);
        }
    }

    /**
     * @dev Remove allowed to entity types for a transaction type
     * @param txType The transaction type
     * @param entityTypes The entity types to remove from allowed receivers
     */
    function removeAllowedToEntityTypes(
        TxType txType,
        EntityType[] memory entityTypes
    ) external onlyRole(TX_TYPE_ADMIN_ROLE) {
        if (entityTypes.length == 0) revert EmptyEntityTypesList();

        uint8 txTypeId = TxType.unwrap(txType);
        for (uint256 i = 0; i < entityTypes.length; i++) {
            uint8 entityTypeId = EntityType.unwrap(entityTypes[i]);

            if (!_allowedToEntityTypes[txTypeId].remove(entityTypeId)) {
                revert EntityTypeNotAllowed(entityTypes[i]);
            }
            emit AllowedToEntityTypeRemoved(txType, entityTypes[i]);
        }
    }

    /**
     * @dev Set transfer policy for a transaction type (batch operation)
     * @param txType The transaction type
     * @param fromEntityTypes Array of entity types allowed to send
     * @param toEntityTypes Array of entity types allowed to receive
     */
    function setTransactionTypePolicy(
        TxType txType,
        EntityType[] memory fromEntityTypes,
        EntityType[] memory toEntityTypes
    ) external onlyRole(TX_TYPE_ADMIN_ROLE) {
        // Clear existing policies first
        clearTransactionType(txType);

        uint8 txTypeId = TxType.unwrap(txType);

        // Add new from entity types
        if (fromEntityTypes.length > 0) {
            for (uint256 i = 0; i < fromEntityTypes.length; i++) {
                uint8 entityTypeId = EntityType.unwrap(fromEntityTypes[i]);

                if (!_allowedFromEntityTypes[txTypeId].add(entityTypeId)) {
                    revert EntityTypeAlreadyAllowed(fromEntityTypes[i]);
                }
                emit AllowedFromEntityTypeAdded(txType, fromEntityTypes[i]);
            }
        }

        // Add new to entity types
        if (toEntityTypes.length > 0) {
            for (uint256 i = 0; i < toEntityTypes.length; i++) {
                uint8 entityTypeId = EntityType.unwrap(toEntityTypes[i]);

                if (!_allowedToEntityTypes[txTypeId].add(entityTypeId)) {
                    revert EntityTypeAlreadyAllowed(toEntityTypes[i]);
                }
                emit AllowedToEntityTypeAdded(txType, toEntityTypes[i]);
            }
        }
    }

    /**
     * @dev Clear all allowed entity types for a transaction type
     * @param txType The transaction type to clear
     */
    function clearTransactionType(
        TxType txType
    ) public onlyRole(TX_TYPE_ADMIN_ROLE) {
        uint8 txTypeId = TxType.unwrap(txType);

        // Clear all from entity types
        uint256[] memory fromEntityTypeIds = _allowedFromEntityTypes[txTypeId]
            .values();
        for (uint256 i = 0; i < fromEntityTypeIds.length; i++) {
            _allowedFromEntityTypes[txTypeId].remove(fromEntityTypeIds[i]);
            emit AllowedFromEntityTypeRemoved(
                txType,
                EntityType.wrap(uint8(fromEntityTypeIds[i]))
            );
        }

        // Clear all to entity types
        uint256[] memory toEntityTypeIds = _allowedToEntityTypes[txTypeId]
            .values();
        for (uint256 i = 0; i < toEntityTypeIds.length; i++) {
            _allowedToEntityTypes[txTypeId].remove(toEntityTypeIds[i]);
            emit AllowedToEntityTypeRemoved(
                txType,
                EntityType.wrap(uint8(toEntityTypeIds[i]))
            );
        }

        emit TransactionTypeCleared(txType);
    }

    /**
     * @dev Check if an entity type is allowed as sender for a transaction type
     * @param txType The transaction type
     * @param entityType The entity type to check
     * @return bool True if the entity type is allowed as sender
     */
    function isAllowedFromEntityType(
        TxType txType,
        EntityType entityType
    ) public view returns (bool) {
        return
            _allowedFromEntityTypes[TxType.unwrap(txType)].contains(
                EntityType.unwrap(entityType)
            );
    }

    /**
     * @dev Check if an entity type is allowed as receiver for a transaction type
     * @param txType The transaction type
     * @param entityType The entity type to check
     * @return bool True if the entity type is allowed as receiver
     */
    function isAllowedToEntityType(
        TxType txType,
        EntityType entityType
    ) public view returns (bool) {
        return
            _allowedToEntityTypes[TxType.unwrap(txType)].contains(
                EntityType.unwrap(entityType)
            );
    }

    /**
     * @dev Check if a transaction type is usable (has both from and to entity types defined)
     * @param txType The transaction type to check
     * @return bool True if the transaction type is usable
     */
    function isTransactionTypeUsable(TxType txType) public view returns (bool) {
        uint8 txTypeId = TxType.unwrap(txType);
        return
            _allowedFromEntityTypes[txTypeId].length() > 0 &&
            _allowedToEntityTypes[txTypeId].length() > 0;
    }

    /**
     * @dev Get all allowed from entity types for a transaction type
     * @param txType The transaction type
     * @return EntityType[] Array of allowed sender entity types
     */
    function getAllowedFromEntityTypes(
        TxType txType
    ) external view returns (EntityType[] memory) {
        uint256[] memory entityTypeIds = _allowedFromEntityTypes[
            TxType.unwrap(txType)
        ].values();
        EntityType[] memory result = new EntityType[](entityTypeIds.length);
        for (uint256 i = 0; i < entityTypeIds.length; i++) {
            result[i] = EntityType.wrap(uint8(entityTypeIds[i]));
        }
        return result;
    }

    /**
     * @dev Get all allowed to entity types for a transaction type
     * @param txType The transaction type
     * @return EntityType[] Array of allowed receiver entity types
     */
    function getAllowedToEntityTypes(
        TxType txType
    ) external view returns (EntityType[] memory) {
        uint256[] memory entityTypeIds = _allowedToEntityTypes[
            TxType.unwrap(txType)
        ].values();
        EntityType[] memory result = new EntityType[](entityTypeIds.length);
        for (uint256 i = 0; i < entityTypeIds.length; i++) {
            result[i] = EntityType.wrap(uint8(entityTypeIds[i]));
        }
        return result;
    }

    /**
     * @dev Get the count of allowed from entity types for a transaction type
     * @param txType The transaction type
     * @return uint256 Count of allowed sender entity types
     */
    function getAllowedFromEntityTypesCount(
        TxType txType
    ) external view returns (uint256) {
        return _allowedFromEntityTypes[TxType.unwrap(txType)].length();
    }

    /**
     * @dev Get the count of allowed to entity types for a transaction type
     * @param txType The transaction type
     * @return uint256 Count of allowed receiver entity types
     */
    function getAllowedToEntityTypesCount(
        TxType txType
    ) external view returns (uint256) {
        return _allowedToEntityTypes[TxType.unwrap(txType)].length();
    }

    /**
     * @dev Standard compliance check (does not enforce transaction type restrictions)
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @param amount Amount of tokens to transfer
     * @return bool Always returns true for standard transfers
     */
    function canTransfer(
        address from,
        address to,
        uint256 amount
    ) external pure override returns (bool) {
        // Standard transfers don't have transaction type constraints
        return true;
    }

    /**
     * @dev Standard compliance check with failure reason (does not enforce transaction type restrictions)
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @param amount Amount of tokens to transfer
     * @return (bool, string) Always returns (true, "") for standard transfers
     */
    function canTransferWithFailureReason(
        address from,
        address to,
        uint256 amount
    ) external pure override returns (bool, string memory) {
        // Standard transfers don't have transaction type constraints
        return (true, "");
    }

    /**
     * @dev Compliance check for typed transfers
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     * @return bool True if the transfer is allowed
     */
    function canTransferWithType(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external view override returns (bool) {
        (bool success, ) = this.canTransferWithTypeAndFailureReason(
            from,
            to,
            amount,
            txType
        );
        return success;
    }

    /**
     * @dev Compliance check for typed transfers with failure reason
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     * @return (bool, string) Tuple of (success, failure reason)
     */
    function canTransferWithTypeAndFailureReason(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external view override returns (bool, string memory) {
        // Skip checks for minting/burning
        if (from == address(0) || to == address(0)) {
            return (true, "");
        }

        // Check if transaction type is usable
        if (!isTransactionTypeUsable(txType)) {
            return (
                false,
                "Transaction type is not usable (empty from or to entity type lists)"
            );
        }

        // Get the entity types for both addresses
        Entity memory fromEntity = entityRegistry.getEntity(from);
        Entity memory toEntity = entityRegistry.getEntity(to);

        // Check if sender's entity type is allowed for this transaction type
        if (!isAllowedFromEntityType(txType, fromEntity.entityType)) {
            return (
                false,
                "Sender entity type is not allowed for this transaction type"
            );
        }

        // Check if receiver's entity type is allowed for this transaction type
        if (!isAllowedToEntityType(txType, toEntity.entityType)) {
            return (
                false,
                "Receiver entity type is not allowed for this transaction type"
            );
        }

        return (true, "");
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[47] private __gap;
}
