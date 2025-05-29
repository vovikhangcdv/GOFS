// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ICompliance} from "../interfaces/ICompliance.sol";
import {EntityRegistry} from "../EntityRegistry.sol";
import {Entity, EntityType} from "../interfaces/ITypes.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {BitMaps} from "@openzeppelin/contracts/utils/structs/BitMaps.sol";

/**
 * @title EntityTypeCompliance
 * @dev Compliance module that manages which entity types are allowed to transfer to which other entity types
 */
contract EntityTypeCompliance is ICompliance, AccessControl {
    // Role for managing compliance constraints
    bytes32 public constant COMPLIANCE_ADMIN_ROLE =
        keccak256("COMPLIANCE_ADMIN_ROLE");

    // The immutable reference to the Entity Registry
    EntityRegistry public immutable entityRegistry;

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

    /**
     * @dev Constructor to set the EntityRegistry address and grant admin roles
     * @param _entityRegistry The address of the EntityRegistry contract
     */
    constructor(address _entityRegistry) {
        if (_entityRegistry == address(0))
            revert InvalidEntityRegistryAddress();
        entityRegistry = EntityRegistry(_entityRegistry);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(COMPLIANCE_ADMIN_ROLE, msg.sender);
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
}
