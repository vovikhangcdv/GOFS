// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {EntityInfo, EntityType} from "./ITypes.sol";

/// @notice Interface for entity registration and verification
/// @dev Handles the registration process and verification status of entities
interface IEntityRegistry {
    /// @notice Registers an entity with verification from an approved verifier
    /// @dev Performs multiple checks:
    /// - Only allow self-registration, to prevent unauthorized registrations
    /// - Validates verifier's authorization
    /// - Confirms entityType is valid for the verifier
    /// - Verifies EIP-712 compliant signature with entityInfo and verifier address
    /// - Validates signature against verifier's public key
    /// @param entityInfo Struct containing entity details
    /// @param verifierSignature EIP-712 signature from an approved verifier
    function register(
        EntityInfo calldata entityInfo,
        bytes memory verifierSignature
    ) external;

    /// @notice Checks if an entity is verified
    /// @param entityAddress Address of the entity to check
    /// @return bool True if the entity is verified by an approved verifier for its type, false otherwise
    function isVerifiedEntity(
        address entityAddress
    ) external view returns (bool);
}
