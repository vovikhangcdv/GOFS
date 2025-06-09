// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICompliance} from "../interfaces/ICompliance.sol";
import {EntityRegistry} from "../EntityRegistry.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {TxType} from "../interfaces/ITypes.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

/**
 * @title VerificationCompliance
 * @dev Compliance module that checks if both sender and receiver are verified entities
 * in the EntityRegistry before allowing transfers.
 */
contract VerificationCompliance is ICompliance, AccessControlUpgradeable {
    // The reference to the Entity Registry (changed from immutable)
    EntityRegistry public entityRegistry;

    error InvalidEntityRegistryAddress();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev Initialize the contract with EntityRegistry address
     * @param _entityRegistry The address of the EntityRegistry contract
     */
    function initialize(address _entityRegistry) public initializer {
        if (_entityRegistry == address(0))
            revert InvalidEntityRegistryAddress();
        entityRegistry = EntityRegistry(_entityRegistry);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
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
        override(IERC165, AccessControlUpgradeable)
        returns (bool)
    {
        return interfaceId == type(ICompliance).interfaceId;
    }

    /**
     * @dev Check if both addresses are verified entities
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
     * @dev Check if both addresses are verified entities and provide failure reason
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
        // For minting, only check recipient
        if (_from == address(0)) {
            if (!entityRegistry.isVerifiedEntity(_to)) {
                return (
                    false,
                    "Recipient for minting is not a verified entity"
                );
            }
            return (true, "");
        }

        // For burning, only check sender
        if (_to == address(0)) {
            if (!entityRegistry.isVerifiedEntity(_from)) {
                return (false, "Sender for burning is not a verified entity");
            }
            return (true, "");
        }

        // Check sender for regular transfers
        if (!entityRegistry.isVerifiedEntity(_from)) {
            return (false, "Sender is not a verified entity");
        }

        // Check recipient for regular transfers
        if (!entityRegistry.isVerifiedEntity(_to)) {
            return (false, "Recipient is not a verified entity");
        }

        return (true, "");
    }

    /**
     * @dev Check if both addresses are verified entities for typed transfers
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
     * @dev Check if both addresses are verified entities for typed transfers and provide failure reason
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
        // For this compliance module, transaction type doesn't affect entity verification
        // So we just delegate to the standard compliance check
        return this.canTransferWithFailureReason(_from, _to, _amount);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
