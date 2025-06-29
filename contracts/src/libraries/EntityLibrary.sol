// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Entity, EntityType} from "../interfaces/ITypes.sol";
import {EntityTypeCompliance} from "../compliances/EntityTypeCompliance.sol";

/// @title EntityLibrary
/// @notice Library for Entity-related constants and hashing functions
/// @dev Provides reusable EIP712 constants and hashing functions for Entity operations
library EntityLibrary {
    // Entity Type constants based on the enum
    EntityType constant UNKNOWN = EntityType.wrap(0);
    EntityType constant INDIVIDUAL = EntityType.wrap(1);
    EntityType constant PRIVATE_ENTERPRISE = EntityType.wrap(2);
    EntityType constant LLC_ONE_MEMBER = EntityType.wrap(3);
    EntityType constant LLC_MULTI_MEMBER = EntityType.wrap(4);
    EntityType constant JOINT_STOCK_COMPANY = EntityType.wrap(5);
    EntityType constant PARTNERSHIP = EntityType.wrap(6);
    EntityType constant COOPERATIVE = EntityType.wrap(7);
    EntityType constant UNION_OF_COOPERATIVES = EntityType.wrap(8);
    EntityType constant COOPERATIVE_GROUP = EntityType.wrap(9);
    EntityType constant HOUSEHOLD_BUSINESS = EntityType.wrap(10);
    EntityType constant FOREIGN_INDIVIDUAL_INVESTOR = EntityType.wrap(11);
    EntityType constant FOREIGN_ORGANIZATION_INVESTOR = EntityType.wrap(12);
    EntityType constant FOREIGN_INVESTED_ECONOMIC_ORGANIZATION =
        EntityType.wrap(13);
    EntityType constant EXCHANGE_PORTAL = EntityType.wrap(14);

    bytes32 public constant ENTITY_TYPEHASH =
        keccak256(
            "Entity(address entityAddress,uint8 entityType,bytes entityData,address verifier)"
        );

    /// @notice Computes the EIP712 hash for an Entity struct
    /// @param entity The Entity struct to hash
    /// @param domainSeparator The EIP712 domain separator
    /// @return The EIP712-compatible hash of the Entity
    function hash(
        Entity memory entity,
        bytes32 domainSeparator
    ) internal pure returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                ENTITY_TYPEHASH,
                entity.entityAddress,
                entity.entityType,
                entity.entityData,
                entity.verifier
            )
        );

        return
            keccak256(
                abi.encodePacked("\x19\x01", domainSeparator, structHash)
            );
    }

    /// @notice Returns the info root of an Entity
    /// @param entity The Entity struct
    /// @return The info root of the Entity
    function getInfoRoot(Entity memory entity) internal pure returns (bytes32) {
        (, bytes32 infoRoot) = abi.decode(entity.entityData, (string, bytes32));
        return infoRoot;
    }
}
