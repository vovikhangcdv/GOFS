// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IEntityRegistry} from "./interfaces/IEntityRegistry.sol";
import {IEntityVerifierRegistry} from "./interfaces/IEntityVerifierRegistry.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import {Entity, EntityType} from "./interfaces/ITypes.sol";
import {EntityLibrary} from "./libraries/EntityLibrary.sol";
import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract EntityRegistry is
    IEntityRegistry,
    IEntityVerifierRegistry,
    AccessControlUpgradeable,
    EIP712Upgradeable
{
    using EntityLibrary for Entity;

    error OnlySelfRegistrationAllowed();
    error EntityAlreadyRegistered();
    error UnauthorizedVerifier();
    error VerifierNotAllowedForEntityType();
    error InvalidVerifierSignature();
    error InvalidVerifierAddress();
    error EmptyEntityTypes();
    error VerifierDoesNotExist();
    error VerifierAlreadyExists();

    function domainSeparator() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    bytes32 public constant ENTITY_REGISTRY_ADMIN_ROLE =
        keccak256("ENTITY_REGISTRY_ADMIN_ROLE");
    bytes32 public ENTITY_TYPE_HASH; // Changed from constant to allow initialization
    bytes32 public constant PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE =
        keccak256("PROACTIVE_REGISTRY_FORWARDER_ADMIN_ROLE");

    mapping(address => Entity) private _entities;
    mapping(address => EntityType[]) private _verifierAllowedTypes;

    event VerifierAdded(address indexed verifier, EntityType[] entityTypes);
    event VerifierUpdated(address indexed verifier, EntityType[] entityTypes);
    event VerifierRemoved(address indexed verifier);
    event EntityRegistered(address indexed entityAddress, Entity entity);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __AccessControl_init();
        __EIP712_init("EntityRegistry", "1.0.0");

        ENTITY_TYPE_HASH = EntityLibrary.ENTITY_TYPEHASH;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ENTITY_REGISTRY_ADMIN_ROLE, msg.sender);
        _grantRole(PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE, msg.sender);
    }

    function getEntity(
        address entityAddress
    ) external view returns (Entity memory) {
        return _entities[entityAddress];
    }

    function getVerifierAllowedTypes(
        address verifier
    ) external view returns (EntityType[] memory) {
        return _verifierAllowedTypes[verifier];
    }

    function register(
        Entity calldata entity,
        bytes memory verifierSignature
    ) external {
        if (entity.entityAddress != msg.sender) {
            revert OnlySelfRegistrationAllowed();
        }
        _register(entity, verifierSignature);
    }

    function forwardRegister(
        Entity calldata entity,
        bytes memory verifierSignature
    ) external onlyRole(PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE) {
        _register(entity, verifierSignature);
    }

    function _register(
        Entity calldata entity,
        bytes memory verifierSignature
    ) internal {
        if (_entities[entity.entityAddress].entityAddress != address(0)) {
            revert EntityAlreadyRegistered();
        }

        if (_verifierAllowedTypes[entity.verifier].length == 0) {
            revert UnauthorizedVerifier();
        }

        if (!_isAllowedEntityType(entity.verifier, entity.entityType)) {
            revert VerifierNotAllowedForEntityType();
        }

        bytes32 hash = entity.hash(_domainSeparatorV4());
        address signer = ECDSA.recover(hash, verifierSignature);
        if (signer != entity.verifier) {
            revert InvalidVerifierSignature();
        }

        _entities[entity.entityAddress] = entity;

        emit EntityRegistered(entity.entityAddress, entity);
    }

    function isVerifiedEntity(
        address entityAddress
    ) external view returns (bool) {
        Entity memory entity = _entities[entityAddress];

        if (entity.entityAddress == address(0)) {
            return false;
        }

        return
            _verifierAllowedTypes[entity.verifier].length > 0 &&
            _isAllowedEntityType(entity.verifier, entity.entityType);
    }

    function addVerifier(
        address verifier,
        EntityType[] memory entityTypes
    ) external onlyRole(ENTITY_REGISTRY_ADMIN_ROLE) {
        if (verifier == address(0)) {
            revert InvalidVerifierAddress();
        }
        if (entityTypes.length == 0) {
            revert EmptyEntityTypes();
        }

        if (_verifierAllowedTypes[verifier].length > 0) {
            revert VerifierAlreadyExists();
        }

        _verifierAllowedTypes[verifier] = entityTypes;

        emit VerifierAdded(verifier, entityTypes);
    }

    function updateVerifier(
        address verifier,
        EntityType[] memory entityTypes
    ) external onlyRole(ENTITY_REGISTRY_ADMIN_ROLE) {
        if (_verifierAllowedTypes[verifier].length == 0) {
            revert VerifierDoesNotExist();
        }
        if (entityTypes.length == 0) {
            revert EmptyEntityTypes();
        }

        _verifierAllowedTypes[verifier] = entityTypes;

        emit VerifierUpdated(verifier, entityTypes);
    }

    function removeVerifier(
        address verifier
    ) external onlyRole(ENTITY_REGISTRY_ADMIN_ROLE) {
        if (_verifierAllowedTypes[verifier].length == 0) {
            revert VerifierDoesNotExist();
        }

        delete _verifierAllowedTypes[verifier];

        emit VerifierRemoved(verifier);
    }

    /// @notice Verifies if a given info (hashed) is part of the Entity's info root
    /// @param entity The Entity struct
    /// @param hashedInfo The hashed info to verify
    /// @param proof The Merkle proof to verify the info against
    /// @return True if the info is valid, false otherwise
    function verifyInfo(
        Entity calldata entity,
        bytes32 hashedInfo,
        bytes32[] calldata proof
    ) external view returns (bool) {
        return MerkleProof.verify(proof, entity.getInfoRoot(), hashedInfo);
    }

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

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
