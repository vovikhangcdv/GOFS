// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {EntityInfo, EntityType} from "./ITypes.sol";

/// @notice Interface for managing entity verifiers
/// @dev Handles the addition, removal, and update of verifiers who can validate entities
interface IEntityVerifierRegistry {
    /// @notice Adds a new verifier with their allowed entity types
    /// @param verifier Address of the verifier to add
    /// @param entityTypes Array of entity types the verifier is allowed to verify
    function addVerifier(
        address verifier,
        EntityType[] memory entityTypes
    ) external;

    /// @notice Updates the allowed entity types for an existing verifier
    /// @param verifier Address of the verifier to update
    /// @param entityTypes New array of entity types the verifier is allowed to verify
    function updateVerifier(
        address verifier,
        EntityType[] memory entityTypes
    ) external;

    /// @notice Removes a verifier from the registry
    /// @param verifier Address of the verifier to remove
    function removeVerifier(address verifier) external;
}
