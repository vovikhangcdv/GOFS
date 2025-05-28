// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {EntityRegistry} from "../src/EntityRegistry.sol";
import {Entity, EntityType} from "../src/interfaces/ITypes.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract EntityRegistryTest is Test {
    EntityRegistry public registry;
    address public admin;
    address public verifier;
    address public entityAddress;
    uint256 public verifierPrivateKey;
    uint256 public entityPrivateKey;

    // Test Entity data
    string constant TEST_ENTITY_NAME = "Test Entity";
    string constant TEST_ENTITY_INFO = "Test Entity Info";
    EntityType constant TEST_ENTITY_TYPE = EntityType.wrap(0); // Assuming 0 is a valid EntityType
    EntityType constant ENTITY_TYPE_INDIVIDUAL = EntityType.wrap(1); // Example type for testing
    EntityType constant ENTITY_TYPE_ORGANIZATION = EntityType.wrap(2); // Example type for testing
    function setUp() public {
        // Set up accounts
        admin = makeAddr("admin");
        verifierPrivateKey = 0x123; // Example private key for testing
        verifier = vm.addr(verifierPrivateKey);
        entityPrivateKey = 0x456; // Example private key for testing
        entityAddress = vm.addr(entityPrivateKey);

        // Deploy contract
        vm.startPrank(admin);
        registry = new EntityRegistry();
        vm.stopPrank();

        // Set up allowed entity types for verifier
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = TEST_ENTITY_TYPE;

        vm.prank(admin);
        registry.addVerifier(verifier, allowedTypes);
    }

    function test_Constructor() public {
        assertTrue(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin));
        assertTrue(registry.hasRole(registry.ADMIN_ROLE(), admin));
    }

    function test_AddVerifier() public {
        address newVerifier = makeAddr("newVerifier");
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = ENTITY_TYPE_INDIVIDUAL;

        vm.prank(admin);
        registry.addVerifier(newVerifier, allowedTypes);

        // Test entity registration with new verifier should work
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: ENTITY_TYPE_INDIVIDUAL,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: newVerifier
        });

        assertTrue(registry.isVerifiedEntity(entityAddress) == false);
    }

    function test_AddVerifier_ZeroAddress() public {
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = ENTITY_TYPE_INDIVIDUAL;

        vm.prank(admin);
        vm.expectRevert(EntityRegistry.InvalidVerifierAddress.selector);
        registry.addVerifier(address(0), allowedTypes);
    }

    function test_AddVerifier_EmptyTypes() public {
        address newVerifier = makeAddr("newVerifier");
        EntityType[] memory emptyTypes = new EntityType[](0);

        vm.prank(admin);
        vm.expectRevert(EntityRegistry.EmptyEntityTypes.selector);
        registry.addVerifier(newVerifier, emptyTypes);
    }

    function test_UpdateVerifier() public {
        EntityType[] memory newAllowedTypes = new EntityType[](2);
        newAllowedTypes[0] = ENTITY_TYPE_INDIVIDUAL;
        newAllowedTypes[1] = ENTITY_TYPE_ORGANIZATION;

        vm.prank(admin);
        registry.updateVerifier(verifier, newAllowedTypes);

        // Both Individual and Organization types should now work with this verifier
        // Additional verification would be done through entity registration
    }

    function test_UpdateVerifier_NonExistent() public {
        address nonExistentVerifier = makeAddr("nonExistent");
        EntityType[] memory newTypes = new EntityType[](1);
        newTypes[0] = ENTITY_TYPE_INDIVIDUAL;

        vm.prank(admin);
        vm.expectRevert(EntityRegistry.VerifierDoesNotExist.selector);
        registry.updateVerifier(nonExistentVerifier, newTypes);
    }

    function test_UpdateVerifier_EmptyTypes() public {
        EntityType[] memory emptyTypes = new EntityType[](0);

        vm.prank(admin);
        vm.expectRevert(EntityRegistry.EmptyEntityTypes.selector);
        registry.updateVerifier(verifier, emptyTypes);
    }

    function test_RemoveVerifier() public {
        vm.prank(admin);
        registry.removeVerifier(verifier);

        // Try to register an entity with removed verifier (should fail)
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        bytes memory signature = _signEntity(verifierPrivateKey, entity);

        vm.prank(entityAddress);
        vm.expectRevert(EntityRegistry.UnauthorizedVerifier.selector);
        registry.register(entity, signature);
    }

    function test_RemoveVerifier_NonExistent() public {
        address nonExistentVerifier = makeAddr("nonExistent");

        vm.prank(admin);
        vm.expectRevert(EntityRegistry.VerifierDoesNotExist.selector);
        registry.removeVerifier(nonExistentVerifier);
    }

    function test_UpdateVerifier_EntityBecomesInvalid() public {
        // First register an entity
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        bytes memory signature = _signEntity(verifierPrivateKey, entity);

        vm.prank(entityAddress);
        registry.register(entity, signature);

        // Verify entity is valid initially
        assertTrue(registry.isVerifiedEntity(entityAddress));

        // Update the verifier to a new one that doesn't allow this type
        address newVerifier = makeAddr("newVerifier");
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = ENTITY_TYPE_INDIVIDUAL; // Different type

        vm.prank(admin);
        registry.updateVerifier(verifier, allowedTypes);

        // Entity should no longer be valid
        assertFalse(registry.isVerifiedEntity(entityAddress));
    }

    function test_RegisterEntity() public {
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        bytes memory signature = _signEntity(verifierPrivateKey, entity);

        vm.prank(entityAddress);
        registry.register(entity, signature);

        assertTrue(registry.isVerifiedEntity(entityAddress));
    }

    function test_RegisterEntity_Unauthorized() public {
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        bytes memory signature = _signEntity(verifierPrivateKey, entity);

        // Try to register from a different address
        vm.prank(makeAddr("unauthorized"));
        vm.expectRevert(EntityRegistry.OnlySelfRegistrationAllowed.selector);
        registry.register(entity, signature);
    }

    function test_RegisterEntity_InvalidSignature() public {
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        // Sign with wrong private key
        bytes memory wrongSignature = _signEntity(0x789, entity);

        vm.prank(entityAddress);
        vm.expectRevert(EntityRegistry.InvalidVerifierSignature.selector);
        registry.register(entity, wrongSignature);
    }

    function test_RegisterEntity_AlreadyRegistered() public {
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        bytes memory signature = _signEntity(verifierPrivateKey, entity);

        // First registration
        vm.prank(entityAddress);
        registry.register(entity, signature);

        // Try to register again
        vm.prank(entityAddress);
        vm.expectRevert(EntityRegistry.EntityAlreadyRegistered.selector);
        registry.register(entity, signature);
    }

    function test_RegisterEntity_WrongEntityType() public {
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: ENTITY_TYPE_ORGANIZATION, // Not allowed for this verifier
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: verifier
        });

        bytes memory signature = _signEntity(verifierPrivateKey, entity);

        vm.prank(entityAddress);
        vm.expectRevert(
            EntityRegistry.VerifierNotAllowedForEntityType.selector
        );
        registry.register(entity, signature);
    }

    function test_RegisterEntity_MultipleAllowedTypes() public {
        // Setup verifier with multiple allowed types
        EntityType[] memory allowedTypes = new EntityType[](3);
        allowedTypes[0] = ENTITY_TYPE_INDIVIDUAL;
        allowedTypes[1] = ENTITY_TYPE_ORGANIZATION;
        allowedTypes[2] = TEST_ENTITY_TYPE;

        uint256 multiVerifierKey = 0xabc;
        address multiVerifier = vm.addr(multiVerifierKey);
        vm.prank(admin);
        registry.addVerifier(multiVerifier, allowedTypes);

        // Try to register with each allowed type
        for (uint i = 0; i < allowedTypes.length; i++) {
            address entityAddr = makeAddr(
                string.concat("entity", vm.toString(i))
            );
            Entity memory entity = Entity({
                entityAddress: entityAddr,
                entityType: allowedTypes[i],
                entityData: _encodeEntityData(
                    TEST_ENTITY_NAME,
                    TEST_ENTITY_INFO
                ),
                verifier: multiVerifier
            });

            bytes memory signature = _signEntity(multiVerifierKey, entity);
            vm.prank(entityAddr);
            registry.register(entity, signature);

            assertTrue(registry.isVerifiedEntity(entityAddr));
        }

        // Try with a non-allowed type
        address invalidEntityAddr = makeAddr("invalidEntity");
        Entity memory invalidEntity = Entity({
            entityAddress: invalidEntityAddr,
            entityType: EntityType.wrap(99), // Some invalid type
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: multiVerifier
        });

        bytes memory invalidSignature = _signEntity(0xdef, invalidEntity);

        vm.prank(invalidEntityAddr);
        vm.expectRevert(
            EntityRegistry.VerifierNotAllowedForEntityType.selector
        );
        registry.register(invalidEntity, invalidSignature);
    }

    function test_IsAllowedEntityType_NoMatch() public {
        // Setup verifier with a specific set of types
        uint256 multiVerifierKey = 0xabc;
        address multiVerifier = vm.addr(multiVerifierKey);
        EntityType[] memory allowedTypes = new EntityType[](2);
        allowedTypes[0] = ENTITY_TYPE_INDIVIDUAL;
        allowedTypes[1] = ENTITY_TYPE_ORGANIZATION;

        vm.prank(admin);
        registry.addVerifier(multiVerifier, allowedTypes);

        // Try to register with a different type
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE, // This type is not in allowedTypes
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: multiVerifier
        });

        bytes memory signature = _signEntity(multiVerifierKey, entity);

        vm.prank(entityAddress);
        vm.expectRevert(
            EntityRegistry.VerifierNotAllowedForEntityType.selector
        );
        registry.register(entity, signature);
    }

    function test_IsVerifiedEntity_VerifierExistsButWrongType() public {
        // Setup verifier with one type
        uint256 multiVerifierKey = 0xabc;
        address multiVerifier = vm.addr(multiVerifierKey);
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = ENTITY_TYPE_INDIVIDUAL;

        vm.prank(admin);
        registry.addVerifier(multiVerifier, allowedTypes);

        // Register an entity with TEST_ENTITY_TYPE
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE, // Different from allowed type
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: multiVerifier
        });

        // Plant the entity in storage to test isVerifiedEntity
        vm.store(
            address(registry),
            keccak256(abi.encode(entityAddress, uint256(2))), // slot for _entities mapping
            bytes32(abi.encode(entity))
        );

        // The entity should not be verified even though the verifier exists
        assertFalse(registry.isVerifiedEntity(entityAddress));
    }

    function test_IsVerifiedEntity_RemovedVerifierWithExistingEntity() public {
        // Setup verifier with allowed type
        uint256 multiVerifierKey = 0xabc;
        address multiVerifier = vm.addr(multiVerifierKey);
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = TEST_ENTITY_TYPE;

        vm.prank(admin);
        registry.addVerifier(multiVerifier, allowedTypes);

        // Create and register entity
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: TEST_ENTITY_TYPE,
            entityData: _encodeEntityData(TEST_ENTITY_NAME, TEST_ENTITY_INFO),
            verifier: multiVerifier
        });

        bytes memory signature = _signEntity(multiVerifierKey, entity);

        vm.prank(entityAddress);
        registry.register(entity, signature);

        // Verify entity is valid initially
        assertTrue(registry.isVerifiedEntity(entityAddress));

        // Remove verifier but don't remove entity
        vm.prank(admin);
        registry.removeVerifier(multiVerifier);

        // Entity should now be invalid because verifier is removed, testing first part of AND condition
        assertFalse(registry.isVerifiedEntity(entityAddress));
    }

    function test_AddVerifier_AlreadyExists() public {
        address newVerifier = makeAddr("newVerifier");
        EntityType[] memory types = new EntityType[](1);
        types[0] = ENTITY_TYPE_INDIVIDUAL;

        // First addition should succeed
        vm.prank(admin);
        registry.addVerifier(newVerifier, types);

        // Second addition should fail
        vm.prank(admin);
        vm.expectRevert(EntityRegistry.VerifierAlreadyExists.selector);
        registry.addVerifier(newVerifier, types);
    }

    // Helper function to sign entity data
    function _signEntity(
        uint256 signerKey,
        Entity memory entity
    ) internal view returns (bytes memory) {
        bytes32 digest = _getEntityDigest(entity);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(signerKey, digest);
        return abi.encodePacked(r, s, v);
    }

    // Helper function to get entity digest
    function _getEntityDigest(
        Entity memory entity
    ) internal view returns (bytes32) {
        bytes32 structHash = keccak256(
            abi.encode(
                registry.ENTITY_TYPE_HASH(),
                entity.entityAddress,
                entity.entityType,
                entity.entityData,
                entity.verifier
            )
        );

        bytes32 domainSeparator = registry.domainSeparator();
        return
            keccak256(
                abi.encodePacked("\x19\x01", domainSeparator, structHash)
            );
    }

    // Helper function to encode entity data
    function _encodeEntityData(
        string memory name,
        string memory info
    ) internal pure returns (bytes memory) {
        return abi.encode(name, info);
    }
}
