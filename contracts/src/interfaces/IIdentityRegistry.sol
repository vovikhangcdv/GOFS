// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IIdentityRegistry
 * @dev Interface for the Identity Registry contract which manages entities
 * and their types in the GOFS (Governable-Onchain-Finance-System).
 */
interface IIdentityRegistry {
    /**
     * @dev Custom type for entity type (uint8)
     */
    type EntityType is uint8;

    /**
     * @dev Entity structure representing an authorized participant
     * @param entityAddress Address of the entity (owner)
     * @param entityType Type of entity (individual, organization, etc.)
     * @param isSuspended Flag indicating if the entity is suspended
     * @param additionalData Reserved for future extensions
     */
    struct Entity {
        address entityAddress;
        EntityType entityType;
        bool isSuspended;
        bytes additionalData;
    }

    /**
     * @dev Events emitted by the contract
     */
    event EntityRegistered(
        address indexed entityAddress,
        EntityType entityType,
        uint256 tokenId
    );
    event EntitySuspensionChanged(
        address indexed entityAddress,
        bool isSuspended
    );

    /**
     * @dev Register a new entity in the system (only callable by governor)
     * @param entityAddress Address of the entity to register
     * @param entityType Type code for the entity
     * @param isSuspended Initial suspension state
     * @return tokenId The ID of the entity token minted
     */
    function registerEntity(
        address entityAddress,
        EntityType entityType,
        bool isSuspended
    ) external returns (uint256 tokenId);

    /**
     * @dev Get entity information by address
     * @param entityAddress The address to query
     * @return Entity struct with entity details
     */
    function getEntity(
        address entityAddress
    ) external view returns (Entity memory);

    /**
     * @dev Set suspension status for an entity
     * @param entityAddress Address of the entity
     * @param isSuspended New suspension state
     * @return success Boolean indicating if the operation was successful
     */
    function setEntitySuspension(
        address entityAddress,
        bool isSuspended
    ) external returns (bool success);

    /**
     * @dev Check if an address is registered as an entity
     * @param entityAddress Address to check
     * @return isRegistered True if the address is a registered entity
     */
    function isRegisteredEntity(
        address entityAddress
    ) external view returns (bool isRegistered);

    /**
     * @dev Check if an entity is active (registered and not suspended)
     * @param entityAddress Address of the entity
     * @return isActive True if the entity is active
     */
    function isActiveEntity(
        address entityAddress
    ) external view returns (bool isActive);

    /**
     * @dev Get the type of an entity
     * @param entityAddress Address of the entity
     * @return entityType The type code of the entity
     */
    function getEntityType(
        address entityAddress
    ) external view returns (EntityType entityType);

    /**
     * @dev Check if an entity is suspended
     * @param entityAddress Address of the entity
     * @return isSuspended True if the entity is suspended
     */
    function isEntitySuspended(
        address entityAddress
    ) external view returns (bool isSuspended);

    /**
     * @dev Get the number of registered entities
     * @return count Total number of entities registered
     */
    function getEntityCount() external view returns (uint256 count);

    /**
     * @dev Check if an entity is of a specific type
     * @param entityAddress Address of the entity
     * @param entityType Type to check against
     * @return isOfType True if the entity is of the specified type
     */
    function isEntityOfType(
        address entityAddress,
        EntityType entityType
    ) external view returns (bool isOfType);
}
