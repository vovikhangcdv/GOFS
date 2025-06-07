// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {VerificationCompliance} from "../../src/compliances/VerificationCompliance.sol";
import {EntityRegistry} from "../../src/EntityRegistry.sol";
import {Entity, EntityType, TxType} from "../../src/interfaces/ITypes.sol";
import {ICompliance} from "../../src/interfaces/ICompliance.sol";

contract VerificationComplianceTest is Test {
    VerificationCompliance public complianceModule;
    EntityRegistry public entityRegistry;

    address public admin;
    address public verifier;
    address public verifiedEntity1;
    address public verifiedEntity2;
    address public unverifiedEntity;

    uint256 private verifierKey;
    uint256 private entity1Key;
    uint256 private entity2Key;

    function setUp() public {
        // Set up accounts
        admin = makeAddr("admin");
        verifierKey = 0x123;
        verifier = vm.addr(verifierKey);
        entity1Key = 0x456;
        entity2Key = 0x789;
        verifiedEntity1 = vm.addr(entity1Key);
        verifiedEntity2 = vm.addr(entity2Key);
        unverifiedEntity = makeAddr("unverified");

        // Deploy contracts
        vm.startPrank(admin);
        entityRegistry = new EntityRegistry();

        // Setup allowed entity types for verifier
        EntityType[] memory allowedTypes = new EntityType[](1);
        allowedTypes[0] = EntityType.wrap(0);
        entityRegistry.addVerifier(verifier, allowedTypes);

        // Deploy compliance module
        complianceModule = new VerificationCompliance(address(entityRegistry));
        vm.stopPrank();

        // Register two verified entities
        _registerEntity(verifiedEntity1);
        _registerEntity(verifiedEntity2);
    }

    function _registerEntity(address entityAddress) internal {
        Entity memory entity = Entity({
            entityAddress: entityAddress,
            entityType: EntityType.wrap(0),
            entityData: bytes("Test Entity Data"),
            verifier: verifier
        });

        // Pre-validate all required fields are set correctly
        require(
            entity.entityAddress != address(0),
            "Entity address cannot be zero"
        );
        require(
            entity.verifier != address(0),
            "Verifier address cannot be zero"
        );
        require(entity.entityData.length > 0, "Entity data cannot be empty");

        // Sign the entity registration using EIP712 standard
        bytes32 structHash = keccak256(
            abi.encode(
                entityRegistry.ENTITY_TYPE_HASH(),
                entity.entityAddress,
                entity.entityType,
                entity.entityData,
                entity.verifier
            )
        );

        bytes32 digest = keccak256(
            abi.encodePacked(
                "\x19\x01",
                entityRegistry.domainSeparator(),
                structHash
            )
        );

        (uint8 v, bytes32 r, bytes32 s) = vm.sign(verifierKey, digest);
        bytes memory signature = abi.encodePacked(r, s, v);

        vm.prank(entityAddress);
        entityRegistry.register(entity, signature);
    }

    function test_Constructor() public {
        // Test constructor with zero address
        vm.expectRevert(
            VerificationCompliance.InvalidEntityRegistryAddress.selector
        );
        new VerificationCompliance(address(0));

        // Test constructor success
        assertEq(
            address(complianceModule.entityRegistry()),
            address(entityRegistry)
        );
    }

    function test_CanTransfer_MintingAndBurning() public {
        // Test minting (from = address(0))
        assertTrue(
            complianceModule.canTransfer(address(0), verifiedEntity1, 100)
        );

        // Test burning (to = address(0))
        assertTrue(
            complianceModule.canTransfer(verifiedEntity1, address(0), 100)
        );
    }

    function test_CanTransfer_BothVerified() public {
        assertTrue(
            complianceModule.canTransfer(verifiedEntity1, verifiedEntity2, 100)
        );
    }

    function test_CanTransfer_UnverifiedSender() public {
        assertFalse(
            complianceModule.canTransfer(unverifiedEntity, verifiedEntity1, 100)
        );
    }

    function test_CanTransfer_UnverifiedRecipient() public {
        assertFalse(
            complianceModule.canTransfer(verifiedEntity1, unverifiedEntity, 100)
        );
    }

    function test_CanTransferWithFailureReason_MintingAndBurning() public {
        // Test minting
        (bool canMint, string memory mintReason) = complianceModule
            .canTransferWithFailureReason(address(0), verifiedEntity1, 100);
        assertTrue(canMint);
        assertEq(mintReason, "");

        // Test burning
        (bool canBurn, string memory burnReason) = complianceModule
            .canTransferWithFailureReason(verifiedEntity1, address(0), 100);
        assertTrue(canBurn);
        assertEq(burnReason, "");
    }

    function test_CanTransferWithFailureReason_BothVerified() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(
                verifiedEntity1,
                verifiedEntity2,
                100
            );
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithFailureReason_UnverifiedSender() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(
                unverifiedEntity,
                verifiedEntity1,
                100
            );
        assertFalse(success);
        assertEq(reason, "Sender is not a verified entity");
    }

    function test_CanTransferWithFailureReason_UnverifiedRecipient() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(
                verifiedEntity1,
                unverifiedEntity,
                100
            );
        assertFalse(success);
        assertEq(reason, "Recipient is not a verified entity");
    }

    function test_CanTransferWithFailureReason_UnverifiedMinting() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(address(0), unverifiedEntity, 100);
        assertFalse(success);
        assertEq(reason, "Recipient for minting is not a verified entity");
    }

    function test_CanTransferWithFailureReason_UnverifiedBurning() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(unverifiedEntity, address(0), 100);
        assertFalse(success);
        assertEq(reason, "Sender for burning is not a verified entity");
    }

    function test_SupportsInterface() public {
        // Test ICompliance interface
        assertTrue(
            complianceModule.supportsInterface(type(ICompliance).interfaceId)
        );

        // Test invalid interface
        assertFalse(complianceModule.supportsInterface(0x12345678));
    }

    function test_CanTransferWithType_BothVerified() public {
        assertTrue(
            complianceModule.canTransferWithType(
                verifiedEntity1,
                verifiedEntity2,
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithType_UnverifiedSender() public {
        assertFalse(
            complianceModule.canTransferWithType(
                unverifiedEntity,
                verifiedEntity1,
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithType_UnverifiedRecipient() public {
        assertFalse(
            complianceModule.canTransferWithType(
                verifiedEntity1,
                unverifiedEntity,
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithType_MintingAndBurning() public {
        // Test typed minting (from zero address)
        assertTrue(
            complianceModule.canTransferWithType(
                address(0),
                verifiedEntity1,
                100,
                TxType.wrap(1)
            )
        );

        // Test typed burning (to zero address)
        assertTrue(
            complianceModule.canTransferWithType(
                verifiedEntity1,
                address(0),
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithTypeAndFailureReason_BothVerified() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithTypeAndFailureReason(
                verifiedEntity1,
                verifiedEntity2,
                100,
                TxType.wrap(1)
            );
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithTypeAndFailureReason_UnverifiedSender()
        public
    {
        (bool success, string memory reason) = complianceModule
            .canTransferWithTypeAndFailureReason(
                unverifiedEntity,
                verifiedEntity1,
                100,
                TxType.wrap(1)
            );
        assertFalse(success);
        assertEq(reason, "Sender is not a verified entity");
    }

    function test_CanTransferWithTypeAndFailureReason_UnverifiedRecipient()
        public
    {
        (bool success, string memory reason) = complianceModule
            .canTransferWithTypeAndFailureReason(
                verifiedEntity1,
                unverifiedEntity,
                100,
                TxType.wrap(1)
            );
        assertFalse(success);
        assertEq(reason, "Recipient is not a verified entity");
    }

    function test_CanTransferWithTypeAndFailureReason_MintingAndBurning()
        public
    {
        // Test typed minting
        (bool canMint, string memory mintReason) = complianceModule
            .canTransferWithTypeAndFailureReason(
                address(0),
                verifiedEntity1,
                100,
                TxType.wrap(1)
            );
        assertTrue(canMint);
        assertEq(mintReason, "");

        // Test typed burning
        (bool canBurn, string memory burnReason) = complianceModule
            .canTransferWithTypeAndFailureReason(
                verifiedEntity1,
                address(0),
                100,
                TxType.wrap(1)
            );
        assertTrue(canBurn);
        assertEq(burnReason, "");
    }
}
