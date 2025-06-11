// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICompliance} from "../interfaces/ICompliance.sol";
import {EntityRegistry} from "../EntityRegistry.sol";
import {Entity, EntityType, TxType} from "../interfaces/ITypes.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {BitMaps} from "@openzeppelin/contracts/utils/structs/BitMaps.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/**
 * @title EntityTypeCompliance
 * @dev Compliance module that manages which entity types are allowed to transfer to which other entity types
 */
contract EntityTypeCompliance is ICompliance, AccessControlUpgradeable {
    // Role for managing compliance constraints
    bytes32 public constant COMPLIANCE_ADMIN_ROLE =
        keccak256("COMPLIANCE_ADMIN_ROLE");

    // The reference to the Entity Registry (changed from immutable)
    EntityRegistry public entityRegistry;

    // Use BitMaps for storing transfer policies
    // Each EntityType has its own bitmap where each bit represents whether it can transfer to that EntityType
    using BitMaps for BitMaps.BitMap;
    mapping(uint8 => BitMaps.BitMap) private _transferPolicies;

    // Events
    event TransferPolicySet(
        EntityType indexed fromType,
        EntityType indexed toType,
        bool allowed
    );

    error InvalidEntityRegistryAddress();
    error InvalidArrayLengths();
    error UnauthorizedTransferBetweenTypes(
        EntityType fromType,
        EntityType toType
    );
    error ZeroAddress();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function name() external pure override returns (string memory) {
        return "EntityTypeCompliance";
    }

    /**
     * @dev Initialize the contract with EntityRegistry address and grant admin roles
     * @param _entityRegistry The address of the EntityRegistry contract
     */
    function initialize(address _entityRegistry) public initializer {
        if (_entityRegistry == address(0))
            revert InvalidEntityRegistryAddress();

        __AccessControl_init();

        entityRegistry = EntityRegistry(_entityRegistry);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(COMPLIANCE_ADMIN_ROLE, msg.sender);
    }

    function setEntityRegistry(
        address _entityRegistry
    ) public onlyRole(DEFAULT_ADMIN_ROLE) {
        entityRegistry = EntityRegistry(_entityRegistry);
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
     * @dev Set whether transfers are allowed from one entity type to another entity type
     * @param fromType The EntityType that can transfer
     * @param toType The EntityType that may receive
     * @param allowed Boolean indicating whether transfer is allowed
     */
    function setTransferPolicy(
        EntityType fromType,
        EntityType toType,
        bool allowed
    ) public onlyRole(COMPLIANCE_ADMIN_ROLE) {
        uint8 fromTypeId = EntityType.unwrap(fromType);
        uint8 toTypeId = EntityType.unwrap(toType);

        if (allowed) {
            _transferPolicies[fromTypeId].set(toTypeId);
        } else {
            _transferPolicies[fromTypeId].unset(toTypeId);
        }

        emit TransferPolicySet(fromType, toType, allowed);
    }

    /**
     * @dev Set whether transfers are allowed from one entity type to multiple entity types
     * @param fromType The EntityType that can transfer
     * @param toTypes Array of EntityTypes that may receive
     * @param alloweds Array of booleans indicating whether transfer to each type is allowed
     */
    function setTransferPolicy(
        EntityType fromType,
        EntityType[] calldata toTypes,
        bool[] calldata alloweds
    ) public onlyRole(COMPLIANCE_ADMIN_ROLE) {
        if (toTypes.length != alloweds.length) revert InvalidArrayLengths();

        for (uint256 i = 0; i < toTypes.length; i++) {
            setTransferPolicy(fromType, toTypes[i], alloweds[i]); // Set subsequent policies
        }
    }

    /**
     * @dev Check if transfers are allowed between two entity types
     * @param fromType The EntityType of the sender
     * @param toType The EntityType of the receiver
     * @return bool Whether transfers between these types are allowed
     */
    function isTransferAllowed(
        EntityType fromType,
        EntityType toType
    ) public view returns (bool) {
        return
            _transferPolicies[EntityType.unwrap(fromType)].get(
                EntityType.unwrap(toType)
            );
    }

    /**
     * @dev Check if a transfer is allowed based on both addresses' entity types
     * @param _from The sender's address
     * @param _to The recipient's address
     * @param _amount The amount to transfer (unused in this module)
     * @return bool Whether the transfer is allowed
     */
    function canTransfer(
        address _from,
        address _to,
        uint256 _amount
    ) public view override returns (bool) {
        (bool success, ) = canTransferWithFailureReason(_from, _to, _amount);
        return success;
    }

    /**
     * @dev Check if a transfer is allowed based on both addresses' entity types and provide failure reason
     * @param _from The sender's address
     * @param _to The recipient's address
     * @param _amount The amount to transfer (unused in this module)
     * @return (bool, string) Whether the transfer is allowed and failure reason if not
     */
    function canTransferWithFailureReason(
        address _from,
        address _to,
        uint256 _amount
    ) public view override returns (bool, string memory) {
        // For minting (_from is zero address), we only need to check recipient's entity type
        if (_from == address(0)) {
            return (true, ""); // Allow minting to any entity type
        }

        // For burning (_to is zero address), we only need to check sender's entity type
        if (_to == address(0)) {
            return (true, ""); // Allow burning from any entity type
        }

        // Get the entity types for both addresses
        Entity memory fromEntity = entityRegistry.getEntity(_from);
        Entity memory toEntity = entityRegistry.getEntity(_to);

        // Check if transfer is allowed between these entity types
        if (
            !_transferPolicies[EntityType.unwrap(fromEntity.entityType)].get(
                EntityType.unwrap(toEntity.entityType)
            )
        ) {
            return (
                false,
                "Transfer between these entity types is not allowed"
            );
        }

        return (true, "");
    }

    /**
     * @dev Check if a typed transfer is allowed based on both addresses' entity types
     * @param _from The sender's address
     * @param _to The recipient's address
     * @param _amount The amount to transfer (unused in this module)
     * @param _txType The type of transaction (unused in this module)
     * @return bool Whether the transfer is allowed
     */
    function canTransferWithType(
        address _from,
        address _to,
        uint256 _amount,
        TxType _txType
    ) external view override returns (bool) {
        (bool success, ) = this.canTransferWithTypeAndFailureReason(
            _from,
            _to,
            _amount,
            _txType
        );
        return success;
    }

    /**
     * @dev Check if a typed transfer is allowed based on both addresses' entity types and provide failure reason
     * @param _from The sender's address
     * @param _to The recipient's address
     * @param _amount The amount to transfer (unused in this module)
     * @param _txType The type of transaction (unused in this module)
     * @return (bool, string) Whether the transfer is allowed and failure reason if not
     */
    function canTransferWithTypeAndFailureReason(
        address _from,
        address _to,
        uint256 _amount,
        TxType _txType
    ) external view override returns (bool, string memory) {
        // For this compliance module, transaction type doesn't affect entity type checks
        // So we just delegate to the standard compliance check
        return canTransferWithFailureReason(_from, _to, _amount);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[47] private __gap;
}
