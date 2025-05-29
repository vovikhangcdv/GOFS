// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IEntityRegistry} from "./interfaces/IEntityRegistry.sol";
import {IEntityVerifierRegistry} from "./interfaces/IEntityVerifierRegistry.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {EIP712} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {Entity, EntityType} from "./interfaces/ITypes.sol";
import {EntityLibrary} from "./libraries/EntityLibrary.sol";

contract EntityRegistry is
    IEntityRegistry,
    IEntityVerifierRegistry,
    AccessControl,
    EIP712
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

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public ENTITY_TYPE_HASH = EntityLibrary.ENTITY_TYPEHASH;

    mapping(address => Entity) private _entities;
    mapping(address => EntityType[]) private _verifierAllowedTypes;

    event VerifierAdded(address indexed verifier, EntityType[] entityTypes);
    event VerifierUpdated(address indexed verifier, EntityType[] entityTypes);
    event VerifierRemoved(address indexed verifier);
    event EntityRegistered(address indexed entityAddress, Entity entity);

    constructor() EIP712("EntityRegistry", "1.0.0") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
    }

    function getEntity(
        address entityAddress
    ) external view returns (Entity memory) {
        return _entities[entityAddress];
    }

    function register(
        Entity calldata entity,
        bytes memory verifierSignature
    ) external {
        if (entity.entityAddress != msg.sender) {
            revert OnlySelfRegistrationAllowed();
        }

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
    ) external onlyRole(ADMIN_ROLE) {
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
    ) external onlyRole(ADMIN_ROLE) {
        if (_verifierAllowedTypes[verifier].length == 0) {
            revert VerifierDoesNotExist();
        }
        if (entityTypes.length == 0) {
            revert EmptyEntityTypes();
        }

        _verifierAllowedTypes[verifier] = entityTypes;

        emit VerifierUpdated(verifier, entityTypes);
    }

    function removeVerifier(address verifier) external onlyRole(ADMIN_ROLE) {
        if (_verifierAllowedTypes[verifier].length == 0) {
            revert VerifierDoesNotExist();
        }

        delete _verifierAllowedTypes[verifier];

        emit VerifierRemoved(verifier);
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
}
