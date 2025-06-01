// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Entity} from "../interfaces/ITypes.sol";

/// @title EntityLibrary
/// @notice Library for Entity-related constants and hashing functions
/// @dev Provides reusable EIP712 constants and hashing functions for Entity operations
library EntityLibrary {
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
        (, bytes32 infoRoot) = abi.decode(
            entity.entityData,
            (string, bytes32)
        );
        return infoRoot;
    }
}
