// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {Entity, EntityType} from "./ITypes.sol";

/// @notice Interface for entity registration and verification
/// @dev Handles the registration process and verification status of entities
interface IEntityRegistry {
    /// @notice Registers an entity with verification from an approved verifier
    /// @dev Performs multiple checks:
    /// - Only allow self-registration, to prevent unauthorized registrations
    /// - Validates verifier's authorization
    /// - Confirms entityType is valid for the verifier
    /// - Verifies EIP-712 compliant signature with entity and verifier address
    /// - Validates signature against verifier's public key (should use openzeppelin's ECDSA library)
    /// @param entity Struct containing entity details
    /// @param verifierSignature EIP-712 signature from an approved verifier
    function register(
        Entity calldata entity,
        bytes memory verifierSignature
    ) external;

    /// @notice Checks if an entity is verified and its verifier is still valid
    /// @dev Performs multiple checks:
    /// - Verifies the entity exists in the registry
    /// - Confirms the entity's verifier is still authorized
    /// - Validates that the verifier is still allowed to verify the entity's type
    /// @param entityAddress Address of the entity to check
    /// @return bool True if the entity exists and its verifier is still authorized for its type, false otherwise
    function isVerifiedEntity(
        address entityAddress
    ) external view returns (bool);
}
