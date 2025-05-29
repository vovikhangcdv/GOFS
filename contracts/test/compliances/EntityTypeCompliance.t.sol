// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {Entity, EntityType} from "../../src/interfaces/ITypes.sol";
import {EntityLibrary} from "../../src/libraries/EntityLibrary.sol";
import {EntityTypeCompliance} from "../../src/compliances/EntityTypeCompliance.sol";
import {EntityRegistry} from "../../src/EntityRegistry.sol";
import {IEntityRegistry} from "../../src/interfaces/IEntityRegistry.sol";

contract EntityTypeComplianceTest is Test {
    using EntityLibrary for Entity;
    EntityTypeCompliance public compliance;
    EntityRegistry public entityRegistry;

    // Test roles and addresses
    address public admin = makeAddr("admin");
    address public alice = makeAddr("alice");
    address public bob = makeAddr("bob");
    address private _verifier;
    uint256 private constant VERIFIER_PRIVATE_KEY = 0x123;

    // Test events
    event TransferPolicySet(
        EntityType indexed fromType,
        EntityType indexed toType,
        bool allowed
    );

    function setUp() public {
        // Setup verifier from private key
        _verifier = vm.addr(VERIFIER_PRIVATE_KEY);

        // Deploy and set up EntityRegistry
        vm.startPrank(admin);

        entityRegistry = new EntityRegistry();

        // Deploy EntityTypeCompliance with EntityRegistry
        compliance = new EntityTypeCompliance(address(entityRegistry));

        // Setup roles
        compliance.grantRole(compliance.DEFAULT_ADMIN_ROLE(), admin);
        compliance.grantRole(compliance.COMPLIANCE_ADMIN_ROLE(), admin);
        vm.stopPrank();
    }

    // Helper function to register an entity
    function _registerEntity(
        address entityAddress,
        uint8 entityTypeId
    ) internal {
        // Add verifier if not already added
        EntityType[] memory allowedTypes = new EntityType[](5);
        for (uint8 i = 0; i < 5; i++) {
            allowedTypes[i] = EntityType.wrap(i + 1);
        }

        // Add verifier only if not already added
        vm.startPrank(admin);
        try entityRegistry.addVerifier(_verifier, allowedTypes) {} catch {}
        vm.stopPrank();

        // Create the entity
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: EntityType.wrap(entityTypeId),
            entityData: "",
            verifier: _verifier
        });

        // Get domain separator
        bytes32 domainSeparator = entityRegistry.domainSeparator();

        // Hash the entity struct
        bytes32 structHash = keccak256(
            abi.encode(
                entityRegistry.ENTITY_TYPE_HASH(),
                entity.entityAddress,
                entity.entityType,
                entity.entityData,
                entity.verifier
            )
        );

        // Create digest
        bytes32 digest = keccak256(
            abi.encodePacked("\x19\x01", domainSeparator, structHash)
        );

        // Have the verifier sign the digest
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(VERIFIER_PRIVATE_KEY, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        // Register the entity
        vm.startPrank(entityAddress);
        entityRegistry.register(entity, signature);
        vm.stopPrank();
    }

    function test_Initialization() public {
        // Test correct initialization
        assertEq(address(compliance.entityRegistry()), address(entityRegistry));
        assertTrue(
            compliance.hasRole(compliance.COMPLIANCE_ADMIN_ROLE(), admin)
        );
    }

    function test_SetSingleTransferPolicy() public {
        EntityType fromType = EntityType.wrap(1);
        EntityType toType = EntityType.wrap(2);

        _registerEntity(address(1), 1); // Type 1
        _registerEntity(address(2), 2); // Type 2
        vm.startPrank(admin);

        // Test setting policy to true
        vm.expectEmit(true, true, true, true);
        emit TransferPolicySet(fromType, toType, true);
        compliance.setTransferPolicy(fromType, toType, true);

        // Verify policy was set
        (bool success, ) = compliance.canTransferWithFailureReason(
            address(1),
            address(2),
            1000
        );
        assertTrue(success);

        // Test setting policy to false
        vm.expectEmit(true, true, true, true);
        emit TransferPolicySet(fromType, toType, false);
        compliance.setTransferPolicy(fromType, toType, false);

        // Verify policy was updated
        (success, ) = compliance.canTransferWithFailureReason(
            address(1),
            address(2),
            1000
        );
        assertFalse(success);

        vm.stopPrank();
    }

    function test_SetBatchTransferPolicy() public {
        // Create test addresses
        address sender = address(0x1111);
        address receiver1 = address(0x3333);
        address receiver2 = address(0x4444);

        // Register entities with their types
        _registerEntity(sender, 1); // Type 1 sender
        _registerEntity(receiver1, 3); // Type 3 receiver
        _registerEntity(receiver2, 4); // Type 4 receiver

        EntityType fromType = EntityType.wrap(1);
        EntityType[] memory toTypes = new EntityType[](2);
        toTypes[0] = EntityType.wrap(3);
        toTypes[1] = EntityType.wrap(4);

        bool[] memory allowed = new bool[](2);
        allowed[0] = true;
        allowed[1] = false;

        vm.startPrank(admin);
        // Test batch setting with expected events
        vm.expectEmit(true, true, true, true);
        emit TransferPolicySet(fromType, toTypes[0], allowed[0]);
        vm.expectEmit(true, true, true, true);
        emit TransferPolicySet(fromType, toTypes[1], allowed[1]);
        compliance.setTransferPolicy(fromType, toTypes, allowed);
        vm.stopPrank();

        // Verify policies were set correctly
        (bool success, ) = compliance.canTransferWithFailureReason(
            sender,
            receiver1,
            1000
        );
        assertTrue(success);

        (success, ) = compliance.canTransferWithFailureReason(
            sender,
            receiver2,
            1000
        );
        assertFalse(success);

        vm.stopPrank();
    }

    function test_OnlyAdminCanSetPolicy() public {
        EntityType fromType = EntityType.wrap(1);
        EntityType toType = EntityType.wrap(2);

        // Try setting policy from non-admin address
        vm.startPrank(alice);
        vm.expectRevert();
        compliance.setTransferPolicy(fromType, toType, true);
        vm.stopPrank();
    }

    function test_CanTransfer() public {
        EntityType fromType = EntityType.wrap(1);
        EntityType toType = EntityType.wrap(2);

        address sender = address(0x1111);
        address receiver = address(0x2222);

        // Register entities with their types
        _registerEntity(sender, 1); // Type 1
        _registerEntity(receiver, 2); // Type 2
        address otherReceiver = address(0x3333);
        _registerEntity(otherReceiver, 3); // Type 3

        vm.startPrank(admin);
        compliance.setTransferPolicy(fromType, toType, true);
        vm.stopPrank();

        // Test allowed transfer
        assertTrue(compliance.canTransfer(sender, receiver, 1000));

        // Test disallowed transfer (policy not set)
        assertFalse(compliance.canTransfer(sender, otherReceiver, 1000));
    }

    function test_CanTransferWithFailureReason() public {
        EntityType fromType = EntityType.wrap(1);
        EntityType toType = EntityType.wrap(2);

        address sender = address(0x1111);
        address receiver = address(0x2222);

        // Register entities with their types
        _registerEntity(sender, 1); // Type 1
        _registerEntity(receiver, 2); // Type 2

        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(sender, receiver, 1000);
        assertFalse(success);
        assertEq(reason, "Transfer between these entity types is not allowed");

        vm.prank(admin);
        compliance.setTransferPolicy(fromType, toType, true);

        (success, reason) = compliance.canTransferWithFailureReason(
            sender,
            receiver,
            1000
        );
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_MintingAndBurning() public {
        address receiver = address(0x1111);
        address sender = address(0x2222);

        // Register entities
        _registerEntity(receiver, 1);
        _registerEntity(sender, 2);

        // Test minting (from zero address)
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(address(0), receiver, 1000);
        assertTrue(success, "Minting should be allowed");
        assertEq(reason, "");

        // Test burning (to zero address)
        (success, reason) = compliance.canTransferWithFailureReason(
            sender,
            address(0),
            1000
        );
        assertTrue(success, "Burning should be allowed");
        assertEq(reason, "");
    }

    // Test minting (from zero address)
    function test_CanTransferFromZeroAddress() public {
        // Minting (from zero address) should be allowed regardless of recipient's type
        _registerEntity(alice, 1);
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(address(0), alice, 1);
        assertTrue(success);
        assertEq(reason, "");
    }

    // Test burning (to zero address)
    function test_CanTransferToZeroAddress() public {
        // Burning (to zero address) should be allowed regardless of sender's type
        _registerEntity(alice, 1);
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(alice, address(0), 1);
        assertTrue(success);
        assertEq(reason, "");
    }

    // Test invalid array lengths
    function test_SetTransferPolicyInvalidLengths() public {
        EntityType fromType = EntityType.wrap(1);
        EntityType[] memory toTypes = new EntityType[](2);
        toTypes[0] = EntityType.wrap(2);
        toTypes[1] = EntityType.wrap(3);
        bool[] memory alloweds = new bool[](1);
        alloweds[0] = true;

        vm.startPrank(admin);
        vm.expectRevert(EntityTypeCompliance.InvalidArrayLengths.selector);
        compliance.setTransferPolicy(fromType, toTypes, alloweds);
        vm.stopPrank();
    }

    // Test invalid entity registry address
    function test_ConstructorZeroAddress() public {
        vm.expectRevert(
            EntityTypeCompliance.InvalidEntityRegistryAddress.selector
        );
        new EntityTypeCompliance(address(0));
    }

    // Test disallowed transfer between types
    function test_CanTransferDisallowed() public {
        // Register entities of different types
        _registerEntity(alice, 1);
        _registerEntity(bob, 2);

        // No transfer policy set, should be disallowed by default
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(alice, bob, 1);
        assertFalse(success);
        assertEq(reason, "Transfer between these entity types is not allowed");
    }

    // Test entity registry immutability
    function test_EntityRegistryImmutability() public {
        address storedRegistry = address(compliance.entityRegistry());
        assertEq(storedRegistry, address(entityRegistry));
    }

    function test_IsTransferAllowed() public {
        EntityType fromType = EntityType.wrap(1);
        EntityType toType = EntityType.wrap(2);

        _registerEntity(address(1), 1); // Type 1
        _registerEntity(address(2), 2); // Type 2

        // Initially should not be allowed
        assertFalse(compliance.isTransferAllowed(fromType, toType));

        // Set policy to true and verify
        vm.startPrank(admin);
        compliance.setTransferPolicy(fromType, toType, true);
        vm.stopPrank();

        assertTrue(compliance.isTransferAllowed(fromType, toType));

        // Set policy back to false and verify
        vm.startPrank(admin);
        compliance.setTransferPolicy(fromType, toType, false);
        vm.stopPrank();

        assertFalse(compliance.isTransferAllowed(fromType, toType));
    }
}
