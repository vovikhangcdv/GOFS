// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../../src/compliances/TransactionTypeCompliance.sol";
import "../../src/EntityRegistry.sol";
import {Entity, EntityType, TxType} from "../../src/interfaces/ITypes.sol";
import {EntityLibrary} from "../../src/libraries/EntityLibrary.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {ICompliance} from "../../src/interfaces/ICompliance.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
contract TransactionTypeComplianceTest is Test {
    using EntityLibrary for Entity;

    TransactionTypeCompliance public compliance;
    EntityRegistry public entityRegistry;

    address public admin;
    address public txTypeAdmin;
    address public user1;
    address public user2;
    address public user3;
    address private _verifier;
    uint256 private constant VERIFIER_PRIVATE_KEY = 0x123;

    EntityType public constant BANK = EntityType.wrap(1);
    EntityType public constant CORPORATE = EntityType.wrap(2);
    EntityType public constant INDIVIDUAL = EntityType.wrap(3);

    TxType public constant LOAN_TYPE = TxType.wrap(1);
    TxType public constant PAYMENT_TYPE = TxType.wrap(2);
    TxType public constant INVESTMENT_TYPE = TxType.wrap(3);

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

    event AllowedFromEntityTypeAdded(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event AllowedFromEntityTypeRemoved(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event AllowedToEntityTypeAdded(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event AllowedToEntityTypeRemoved(
        TxType indexed txType,
        EntityType indexed entityType
    );
    event TransactionTypeCleared(TxType indexed txType);

    function setUp() public {
        admin = makeAddr("admin");
        txTypeAdmin = makeAddr("txTypeAdmin");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");

        vm.startPrank(admin);

        // Deploy EntityRegistry
        entityRegistry = EntityRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(new EntityRegistry()),
                    address(new ProxyAdmin(admin)),
                    abi.encodeWithSelector(
                        EntityRegistry.initialize.selector,
                        admin
                    )
                )
            )
        );

        // Deploy compliance contract
        compliance = TransactionTypeCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(new TransactionTypeCompliance()),
                    address(new ProxyAdmin(admin)),
                    abi.encodeWithSelector(
                        TransactionTypeCompliance.initialize.selector,
                        address(entityRegistry)
                    )
                )
            )
        );

        // Grant TX_TYPE_ADMIN_ROLE to txTypeAdmin
        compliance.grantRole(compliance.TX_TYPE_ADMIN_ROLE(), txTypeAdmin);

        vm.stopPrank();

        // Setup verifier from private key
        _verifier = vm.addr(VERIFIER_PRIVATE_KEY);

        // Register entities in the registry
        _registerEntity(user1, EntityType.unwrap(BANK));
        _registerEntity(user2, EntityType.unwrap(CORPORATE));
        _registerEntity(user3, EntityType.unwrap(INDIVIDUAL));
    }

    function testSupportsInterface() public view {
        // Test ICompliance interface
        assertTrue(compliance.supportsInterface(type(ICompliance).interfaceId));

        // Test AccessControl interface
        assertTrue(
            compliance.supportsInterface(type(IAccessControl).interfaceId)
        );

        // Test IERC165 interface
        assertTrue(compliance.supportsInterface(type(IERC165).interfaceId));

        // Test unsupported interface
        assertFalse(compliance.supportsInterface(bytes4(0x12345678)));
    }

    function testAddAllowedFromEntityTypes() public {
        EntityType[] memory entityTypes = new EntityType[](2);
        entityTypes[0] = BANK;
        entityTypes[1] = CORPORATE;

        vm.startPrank(txTypeAdmin);

        vm.expectEmit(true, true, false, false);
        emit AllowedFromEntityTypeAdded(LOAN_TYPE, BANK);
        vm.expectEmit(true, true, false, false);
        emit AllowedFromEntityTypeAdded(LOAN_TYPE, CORPORATE);

        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);

        vm.stopPrank();

        // Verify entity types were added
        assertTrue(compliance.isAllowedFromEntityType(LOAN_TYPE, BANK));
        assertTrue(compliance.isAllowedFromEntityType(LOAN_TYPE, CORPORATE));
        assertFalse(compliance.isAllowedFromEntityType(LOAN_TYPE, INDIVIDUAL));

        // Check count
        assertEq(compliance.getAllowedFromEntityTypesCount(LOAN_TYPE), 2);
    }

    function testAddAllowedFromEntityTypesEmptyArray() public {
        EntityType[] memory entityTypes = new EntityType[](0);

        vm.startPrank(txTypeAdmin);
        vm.expectRevert(
            TransactionTypeCompliance.EmptyEntityTypesList.selector
        );
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);
        vm.stopPrank();
    }

    function testAddAllowedFromEntityTypesAlreadyExists() public {
        EntityType[] memory entityTypes = new EntityType[](1);
        entityTypes[0] = BANK;

        vm.startPrank(txTypeAdmin);
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);

        // Try to add the same entity type again
        vm.expectRevert(
            abi.encodeWithSelector(
                TransactionTypeCompliance.EntityTypeAlreadyAllowed.selector,
                BANK
            )
        );
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);
        vm.stopPrank();
    }

    function testAddAllowedFromEntityTypesUnauthorized() public {
        EntityType[] memory entityTypes = new EntityType[](1);
        entityTypes[0] = BANK;

        vm.startPrank(user1);
        vm.expectRevert();
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);
        vm.stopPrank();
    }

    function testAddAllowedToEntityTypes() public {
        EntityType[] memory entityTypes = new EntityType[](2);
        entityTypes[0] = CORPORATE;
        entityTypes[1] = INDIVIDUAL;

        vm.startPrank(txTypeAdmin);

        vm.expectEmit(true, true, false, false);
        emit AllowedToEntityTypeAdded(LOAN_TYPE, CORPORATE);
        vm.expectEmit(true, true, false, false);
        emit AllowedToEntityTypeAdded(LOAN_TYPE, INDIVIDUAL);

        compliance.addAllowedToEntityTypes(LOAN_TYPE, entityTypes);

        vm.stopPrank();

        // Verify entity types were added
        assertTrue(compliance.isAllowedToEntityType(LOAN_TYPE, CORPORATE));
        assertTrue(compliance.isAllowedToEntityType(LOAN_TYPE, INDIVIDUAL));
        assertFalse(compliance.isAllowedToEntityType(LOAN_TYPE, BANK));

        // Check count
        assertEq(compliance.getAllowedToEntityTypesCount(LOAN_TYPE), 2);
    }

    function testAddAllowedToEntityTypesEmptyArray() public {
        EntityType[] memory entityTypes = new EntityType[](0);

        vm.startPrank(txTypeAdmin);
        vm.expectRevert(
            TransactionTypeCompliance.EmptyEntityTypesList.selector
        );
        compliance.addAllowedToEntityTypes(LOAN_TYPE, entityTypes);
        vm.stopPrank();
    }

    function testAddAllowedToEntityTypesAlreadyExists() public {
        EntityType[] memory entityTypes = new EntityType[](1);
        entityTypes[0] = CORPORATE;

        vm.startPrank(txTypeAdmin);
        compliance.addAllowedToEntityTypes(LOAN_TYPE, entityTypes);

        // Try to add the same entity type again
        vm.expectRevert(
            abi.encodeWithSelector(
                TransactionTypeCompliance.EntityTypeAlreadyAllowed.selector,
                CORPORATE
            )
        );
        compliance.addAllowedToEntityTypes(LOAN_TYPE, entityTypes);
        vm.stopPrank();
    }

    function testRemoveAllowedFromEntityTypes() public {
        // First add some entity types
        EntityType[] memory entityTypes = new EntityType[](2);
        entityTypes[0] = BANK;
        entityTypes[1] = CORPORATE;

        vm.startPrank(txTypeAdmin);
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);

        // Now remove one
        EntityType[] memory toRemove = new EntityType[](1);
        toRemove[0] = BANK;

        vm.expectEmit(true, true, false, false);
        emit AllowedFromEntityTypeRemoved(LOAN_TYPE, BANK);

        compliance.removeAllowedFromEntityTypes(LOAN_TYPE, toRemove);

        vm.stopPrank();

        // Verify removal
        assertFalse(compliance.isAllowedFromEntityType(LOAN_TYPE, BANK));
        assertTrue(compliance.isAllowedFromEntityType(LOAN_TYPE, CORPORATE));
        assertEq(compliance.getAllowedFromEntityTypesCount(LOAN_TYPE), 1);
    }

    function testRemoveAllowedFromEntityTypesNotExists() public {
        EntityType[] memory entityTypes = new EntityType[](1);
        entityTypes[0] = BANK;

        vm.startPrank(txTypeAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                TransactionTypeCompliance.EntityTypeNotAllowed.selector,
                BANK
            )
        );
        compliance.removeAllowedFromEntityTypes(LOAN_TYPE, entityTypes);
        vm.stopPrank();
    }

    function testRemoveAllowedToEntityTypes() public {
        // First add some entity types
        EntityType[] memory entityTypes = new EntityType[](2);
        entityTypes[0] = CORPORATE;
        entityTypes[1] = INDIVIDUAL;

        vm.startPrank(txTypeAdmin);
        compliance.addAllowedToEntityTypes(LOAN_TYPE, entityTypes);

        // Now remove one
        EntityType[] memory toRemove = new EntityType[](1);
        toRemove[0] = CORPORATE;

        vm.expectEmit(true, true, false, false);
        emit AllowedToEntityTypeRemoved(LOAN_TYPE, CORPORATE);

        compliance.removeAllowedToEntityTypes(LOAN_TYPE, toRemove);

        vm.stopPrank();

        // Verify removal
        assertFalse(compliance.isAllowedToEntityType(LOAN_TYPE, CORPORATE));
        assertTrue(compliance.isAllowedToEntityType(LOAN_TYPE, INDIVIDUAL));
        assertEq(compliance.getAllowedToEntityTypesCount(LOAN_TYPE), 1);
    }

    function testSetTransactionTypePolicy() public {
        EntityType[] memory fromTypes = new EntityType[](2);
        fromTypes[0] = BANK;
        fromTypes[1] = CORPORATE;

        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = INDIVIDUAL;

        vm.startPrank(txTypeAdmin);
        compliance.setTransactionTypePolicy(PAYMENT_TYPE, fromTypes, toTypes);
        vm.stopPrank();

        // Verify the policy was set
        assertTrue(compliance.isAllowedFromEntityType(PAYMENT_TYPE, BANK));
        assertTrue(compliance.isAllowedFromEntityType(PAYMENT_TYPE, CORPORATE));
        assertTrue(compliance.isAllowedToEntityType(PAYMENT_TYPE, INDIVIDUAL));

        // Check counts
        assertEq(compliance.getAllowedFromEntityTypesCount(PAYMENT_TYPE), 2);
        assertEq(compliance.getAllowedToEntityTypesCount(PAYMENT_TYPE), 1);

        // Verify it's usable
        assertTrue(compliance.isTransactionTypeUsable(PAYMENT_TYPE));
    }

    function testSetTransactionTypePolicyOverwrite() public {
        // First set a policy
        EntityType[] memory fromTypes1 = new EntityType[](1);
        fromTypes1[0] = BANK;

        EntityType[] memory toTypes1 = new EntityType[](1);
        toTypes1[0] = CORPORATE;

        vm.startPrank(txTypeAdmin);
        compliance.setTransactionTypePolicy(PAYMENT_TYPE, fromTypes1, toTypes1);

        // Verify initial policy
        assertTrue(compliance.isAllowedFromEntityType(PAYMENT_TYPE, BANK));
        assertTrue(compliance.isAllowedToEntityType(PAYMENT_TYPE, CORPORATE));

        // Now overwrite with new policy
        EntityType[] memory fromTypes2 = new EntityType[](1);
        fromTypes2[0] = CORPORATE;

        EntityType[] memory toTypes2 = new EntityType[](1);
        toTypes2[0] = INDIVIDUAL;

        compliance.setTransactionTypePolicy(PAYMENT_TYPE, fromTypes2, toTypes2);
        vm.stopPrank();

        // Verify old policy was cleared and new policy is set
        assertFalse(compliance.isAllowedFromEntityType(PAYMENT_TYPE, BANK));
        assertTrue(compliance.isAllowedFromEntityType(PAYMENT_TYPE, CORPORATE));
        assertFalse(compliance.isAllowedToEntityType(PAYMENT_TYPE, CORPORATE));
        assertTrue(compliance.isAllowedToEntityType(PAYMENT_TYPE, INDIVIDUAL));
    }

    function testClearTransactionType() public {
        // Set up a policy first
        EntityType[] memory fromTypes = new EntityType[](2);
        fromTypes[0] = BANK;
        fromTypes[1] = CORPORATE;

        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = INDIVIDUAL;

        vm.startPrank(txTypeAdmin);
        compliance.setTransactionTypePolicy(
            INVESTMENT_TYPE,
            fromTypes,
            toTypes
        );

        // Verify policy is set
        assertTrue(compliance.isTransactionTypeUsable(INVESTMENT_TYPE));

        // Clear the transaction type
        vm.expectEmit(true, false, false, false);
        emit TransactionTypeCleared(INVESTMENT_TYPE);

        compliance.clearTransactionType(INVESTMENT_TYPE);
        vm.stopPrank();

        // Verify everything is cleared
        assertFalse(compliance.isAllowedFromEntityType(INVESTMENT_TYPE, BANK));
        assertFalse(
            compliance.isAllowedFromEntityType(INVESTMENT_TYPE, CORPORATE)
        );
        assertFalse(
            compliance.isAllowedToEntityType(INVESTMENT_TYPE, INDIVIDUAL)
        );
        assertEq(compliance.getAllowedFromEntityTypesCount(INVESTMENT_TYPE), 0);
        assertEq(compliance.getAllowedToEntityTypesCount(INVESTMENT_TYPE), 0);
        assertFalse(compliance.isTransactionTypeUsable(INVESTMENT_TYPE));
    }

    function testIsTransactionTypeUsable() public {
        // Initially not usable (empty)
        assertFalse(compliance.isTransactionTypeUsable(LOAN_TYPE));

        vm.startPrank(txTypeAdmin);

        // Add only from types - still not usable
        EntityType[] memory fromTypes = new EntityType[](1);
        fromTypes[0] = BANK;
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, fromTypes);
        assertFalse(compliance.isTransactionTypeUsable(LOAN_TYPE));

        // Add to types - now usable
        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = INDIVIDUAL;
        compliance.addAllowedToEntityTypes(LOAN_TYPE, toTypes);
        assertTrue(compliance.isTransactionTypeUsable(LOAN_TYPE));

        // Remove from types - not usable again
        compliance.removeAllowedFromEntityTypes(LOAN_TYPE, fromTypes);
        assertFalse(compliance.isTransactionTypeUsable(LOAN_TYPE));

        vm.stopPrank();
    }

    function testGetAllowedEntityTypes() public {
        EntityType[] memory fromTypes = new EntityType[](2);
        fromTypes[0] = BANK;
        fromTypes[1] = CORPORATE;

        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = INDIVIDUAL;

        vm.startPrank(txTypeAdmin);
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, fromTypes);
        compliance.addAllowedToEntityTypes(LOAN_TYPE, toTypes);
        vm.stopPrank();

        // Test getting all allowed from entity types
        EntityType[] memory retrievedFromTypes = compliance
            .getAllowedFromEntityTypes(LOAN_TYPE);
        assertEq(retrievedFromTypes.length, 2);
        // Note: Order might not be preserved, so we check for presence
        bool bankFound = false;
        bool corporateFound = false;
        for (uint i = 0; i < retrievedFromTypes.length; i++) {
            if (
                EntityType.unwrap(retrievedFromTypes[i]) ==
                EntityType.unwrap(BANK)
            ) {
                bankFound = true;
            }
            if (
                EntityType.unwrap(retrievedFromTypes[i]) ==
                EntityType.unwrap(CORPORATE)
            ) {
                corporateFound = true;
            }
        }
        assertTrue(bankFound);
        assertTrue(corporateFound);

        // Test getting all allowed to entity types
        EntityType[] memory retrievedToTypes = compliance
            .getAllowedToEntityTypes(LOAN_TYPE);
        assertEq(retrievedToTypes.length, 1);
        assertEq(
            EntityType.unwrap(retrievedToTypes[0]),
            EntityType.unwrap(INDIVIDUAL)
        );
    }

    function testCanTransfer() public view {
        // Standard transfers should always return true
        assertTrue(compliance.canTransfer(user1, user2, 100));
        assertTrue(compliance.canTransfer(address(0), user2, 100)); // minting
        assertTrue(compliance.canTransfer(user1, address(0), 100)); // burning
    }

    function testCanTransferWithFailureReason() public view {
        // Standard transfers should always return (true, "")
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(user1, user2, 100);
        assertTrue(success);
        assertEq(reason, "");
    }

    function testCanTransferWithTypeSuccessful() public {
        // Set up a valid transaction type policy
        EntityType[] memory fromTypes = new EntityType[](1);
        fromTypes[0] = BANK; // user1 is BANK

        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = CORPORATE; // user2 is CORPORATE

        vm.startPrank(txTypeAdmin);
        compliance.setTransactionTypePolicy(LOAN_TYPE, fromTypes, toTypes);
        vm.stopPrank();

        // Test successful transfer
        assertTrue(
            compliance.canTransferWithType(user1, user2, 100, LOAN_TYPE)
        );

        (bool success, string memory reason) = compliance
            .canTransferWithTypeAndFailureReason(user1, user2, 100, LOAN_TYPE);
        assertTrue(success);
        assertEq(reason, "");
    }

    function testCanTransferWithTypeMintingAndBurning() public {
        // Minting and burning should always be allowed regardless of transaction type policy
        assertTrue(
            compliance.canTransferWithType(address(0), user1, 100, LOAN_TYPE)
        );
        assertTrue(
            compliance.canTransferWithType(user1, address(0), 100, LOAN_TYPE)
        );

        (bool success, string memory reason) = compliance
            .canTransferWithTypeAndFailureReason(
                address(0),
                user1,
                100,
                LOAN_TYPE
            );
        assertTrue(success);
        assertEq(reason, "");

        (success, reason) = compliance.canTransferWithTypeAndFailureReason(
            user1,
            address(0),
            100,
            LOAN_TYPE
        );
        assertTrue(success);
        assertEq(reason, "");
    }

    function testCanTransferWithTypeNotUsable() public {
        // Transaction type is not usable (no policy set)
        assertFalse(
            compliance.canTransferWithType(user1, user2, 100, LOAN_TYPE)
        );

        (bool success, string memory reason) = compliance
            .canTransferWithTypeAndFailureReason(user1, user2, 100, LOAN_TYPE);
        assertFalse(success);
        assertEq(
            reason,
            "Transaction type is not usable (empty from or to entity type lists)"
        );
    }

    function testCanTransferWithTypeSenderNotAllowed() public {
        // Set up policy where sender's entity type is not allowed
        EntityType[] memory fromTypes = new EntityType[](1);
        fromTypes[0] = CORPORATE; // user1 is BANK, not CORPORATE

        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = INDIVIDUAL;

        vm.startPrank(txTypeAdmin);
        compliance.setTransactionTypePolicy(PAYMENT_TYPE, fromTypes, toTypes);
        vm.stopPrank();

        assertFalse(
            compliance.canTransferWithType(user1, user3, 100, PAYMENT_TYPE)
        );

        (bool success, string memory reason) = compliance
            .canTransferWithTypeAndFailureReason(
                user1,
                user3,
                100,
                PAYMENT_TYPE
            );
        assertFalse(success);
        assertEq(
            reason,
            "Sender entity type is not allowed for this transaction type"
        );
    }

    function testCanTransferWithTypeReceiverNotAllowed() public {
        // Set up policy where receiver's entity type is not allowed
        EntityType[] memory fromTypes = new EntityType[](1);
        fromTypes[0] = BANK; // user1 is BANK

        EntityType[] memory toTypes = new EntityType[](1);
        toTypes[0] = CORPORATE; // user3 is INDIVIDUAL, not CORPORATE

        vm.startPrank(txTypeAdmin);
        compliance.setTransactionTypePolicy(PAYMENT_TYPE, fromTypes, toTypes);
        vm.stopPrank();

        assertFalse(
            compliance.canTransferWithType(user1, user3, 100, PAYMENT_TYPE)
        );

        (bool success, string memory reason) = compliance
            .canTransferWithTypeAndFailureReason(
                user1,
                user3,
                100,
                PAYMENT_TYPE
            );
        assertFalse(success);
        assertEq(
            reason,
            "Receiver entity type is not allowed for this transaction type"
        );
    }

    function testPublicMappingsAccess() public {
        // Set up some data
        EntityType[] memory fromTypes = new EntityType[](1);
        fromTypes[0] = BANK;

        vm.startPrank(txTypeAdmin);
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, fromTypes);
        vm.stopPrank();

        // Test that we can access public mappings
        // Note: We can't directly test EnumerableSet through public mappings in this way
        // But we can verify the state through other functions
        assertTrue(compliance.isAllowedFromEntityType(LOAN_TYPE, BANK));
        assertEq(compliance.getAllowedFromEntityTypesCount(LOAN_TYPE), 1);
    }

    function testRoleAccess() public {
        // Test that only TX_TYPE_ADMIN_ROLE can call admin functions
        EntityType[] memory entityTypes = new EntityType[](1);
        entityTypes[0] = BANK;

        // user1 doesn't have TX_TYPE_ADMIN_ROLE
        vm.startPrank(user1);

        vm.expectRevert();
        compliance.addAllowedFromEntityTypes(LOAN_TYPE, entityTypes);

        vm.expectRevert();
        compliance.addAllowedToEntityTypes(LOAN_TYPE, entityTypes);

        vm.expectRevert();
        compliance.removeAllowedFromEntityTypes(LOAN_TYPE, entityTypes);

        vm.expectRevert();
        compliance.removeAllowedToEntityTypes(LOAN_TYPE, entityTypes);

        vm.expectRevert();
        compliance.setTransactionTypePolicy(
            LOAN_TYPE,
            entityTypes,
            entityTypes
        );

        vm.expectRevert();
        compliance.clearTransactionType(LOAN_TYPE);

        vm.stopPrank();
    }
}
