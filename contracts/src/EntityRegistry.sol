// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IEntityRegistry} from "./interfaces/IEntityRegistry.sol";
import {IEntityVerifierRegistry} from "./interfaces/IEntityVerifierRegistry.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {Entity, EntityType} from "./interfaces/ITypes.sol";
import {EntityLibrary} from "./libraries/EntityLibrary.sol";

/// @title EntityRegistry
/// @notice Manages entity registration and verifier authorization
/// @dev Implements IEntityRegistry and IEntityVerifierRegistry with EIP-712 signature verification
contract EntityRegistry is
    IEntityRegistry,
    IEntityVerifierRegistry,
    AccessControl,
    EIP712
{
    using EntityLibrary for Entity;
    // Custom Errors
    error OnlySelfRegistrationAllowed();
    error EntityAlreadyRegistered();
    error UnauthorizedVerifier();
    error VerifierNotAllowedForEntityType();
    error InvalidVerifierSignature();
    error InvalidVerifierAddress();
    error VerifierAlreadyExists();
    error EmptyEntityTypes();
    error VerifierDoesNotExist();

    // Public getters for EIP-712 domain values
    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public ENTITY_TYPE_HASH = EntityLibrary.ENTITY_TYPEHASH;

    // Storage for entity and verifier data
    mapping(address => Entity) private _entities;
    mapping(address => bool) private _verifiers;
    mapping(address => EntityType[]) private _verifierAllowedTypes;

    // Events
    event VerifierAdded(address indexed verifier, EntityType[] entityTypes);
    event VerifierUpdated(address indexed verifier, EntityType[] entityTypes);
    event VerifierRemoved(address indexed verifier);
    event EntityRegistered(address indexed entityAddress, Entity entity);

    constructor() EIP712("EntityRegistry", "1.0.0") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    /// @inheritdoc IEntityRegistry
    function register(
        Entity calldata entity,
        bytes memory verifierSignature
    ) external {
        // Ensure self-registration
        if (entity.entityAddress != msg.sender) {
            revert OnlySelfRegistrationAllowed();
        }

        // Verify that the entity hasn't been registered
        if (_entities[entity.entityAddress].entityAddress != address(0)) {
            revert EntityAlreadyRegistered();
        }

        // Check if verifier is authorized
        if (!_verifiers[entity.verifier]) {
            revert UnauthorizedVerifier();
        }

        // Check if verifier is allowed to verify this entity type
        if (!_isAllowedEntityType(entity.verifier, entity.entityType)) {
            revert VerifierNotAllowedForEntityType();
        }

        // Verify signature using EntityLibrary
        bytes32 hash = entity.hash(_domainSeparatorV4());
        address signer = ECDSA.recover(hash, verifierSignature);
        if (signer != entity.verifier) {
            revert InvalidVerifierSignature();
        }

        // Store entity info
        _entities[entity.entityAddress] = entity;

        emit EntityRegistered(entity.entityAddress, entity);
    }

    /// @inheritdoc IEntityRegistry
    function isVerifiedEntity(
        address entityAddress
    ) external view returns (bool) {
        Entity memory entity = _entities[entityAddress];

        // Check if entity exists
        if (entity.entityAddress == address(0)) {
            return false;
        }

        // Check if verifier is still authorized and can verify this entity type
        return
            _verifiers[entity.verifier] &&
            _isAllowedEntityType(entity.verifier, entity.entityType);
    }

    /// @inheritdoc IEntityVerifierRegistry
    function addVerifier(
        address verifier,
        EntityType[] memory entityTypes
    ) external onlyRole(ADMIN_ROLE) {
        if (verifier == address(0)) {
            revert InvalidVerifierAddress();
        }
        if (_verifiers[verifier]) {
            revert VerifierAlreadyExists();
        }
        if (entityTypes.length == 0) {
            revert EmptyEntityTypes();
        }

        _verifiers[verifier] = true;
        _verifierAllowedTypes[verifier] = entityTypes;

        emit VerifierAdded(verifier, entityTypes);
    }

    /// @inheritdoc IEntityVerifierRegistry
    function updateVerifier(
        address verifier,
        EntityType[] memory entityTypes
    ) external onlyRole(ADMIN_ROLE) {
        if (!_verifiers[verifier]) {
            revert VerifierDoesNotExist();
        }
        if (entityTypes.length == 0) {
            revert EmptyEntityTypes();
        }

        _verifierAllowedTypes[verifier] = entityTypes;

        emit VerifierUpdated(verifier, entityTypes);
    }

    /// @inheritdoc IEntityVerifierRegistry
    function removeVerifier(address verifier) external onlyRole(ADMIN_ROLE) {
        if (!_verifiers[verifier]) {
            revert VerifierDoesNotExist();
        }

        delete _verifiers[verifier];
        delete _verifierAllowedTypes[verifier];

        emit VerifierRemoved(verifier);
    }

    /// @notice Checks if a verifier is allowed to verify a specific entity type
    /// @param verifier Address of the verifier
    /// @param entityType Type of entity to check
    /// @return bool True if the verifier is allowed to verify this entity type
    function _isAllowedEntityType(
        address verifier,
        EntityType entityType
    ) private view returns (bool) {
        EntityType[] memory allowedTypes = _verifierAllowedTypes[verifier];
        for (uint256 i = 0; i < allowedTypes.length; i++) {
            if (allowedTypes[i] == entityType) {
                return true;
            }
        }
        return false;
    }
}
