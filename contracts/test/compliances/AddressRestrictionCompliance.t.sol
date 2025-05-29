// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {AddressRestrictionCompliance} from "../../src/compliances/AddressRestrictionCompliance.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";

contract AddressRestrictionComplianceTest is Test {
    bytes32 public constant BLACKLIST_ADMIN_ROLE =
        keccak256("BLACKLIST_ADMIN_ROLE");

    AddressRestrictionCompliance public complianceModule;

    address public admin;
    address public user1;
    address public user2;
    address public user3;
    address public nonAdmin;

    // Events
    event AddressBlacklistedFrom(address indexed account);
    event AddressUnblacklistedFrom(address indexed account);
    event AddressBlacklistedTo(address indexed account);
    event AddressUnblacklistedTo(address indexed account);
    event AddressBlacklisted(address indexed account);
    event AddressUnblacklisted(address indexed account);

    function setUp() public {
        admin = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");
        nonAdmin = makeAddr("nonAdmin");

        complianceModule = new AddressRestrictionCompliance();
    }

    function test_InitialState() public {
        assertTrue(
            complianceModule.hasRole(
                complianceModule.DEFAULT_ADMIN_ROLE(),
                admin
            )
        );
        assertTrue(complianceModule.hasRole(BLACKLIST_ADMIN_ROLE, admin));

        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));
    }

    function test_BlacklistFrom() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistFrom(users);
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedFrom(user2));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklistedTo(user2));
        assertFalse(complianceModule.isBlacklisted(user1));
        assertFalse(complianceModule.isBlacklisted(user2));
    }

    function test_BlacklistTo() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistTo(users);
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedFrom(user2));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklistedTo(user2));
        assertFalse(complianceModule.isBlacklisted(user1));
        assertFalse(complianceModule.isBlacklisted(user2));
    }

    function test_BlacklistBoth() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklist(users);
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedFrom(user2));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklistedTo(user2));
        assertTrue(complianceModule.isBlacklisted(user1));
        assertTrue(complianceModule.isBlacklisted(user2));
    }

    function test_BlacklistFromRevertEmptyList() public {
        address[] memory users = new address[](0);
        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklistFrom(users);
    }

    function test_BlacklistFromRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklistFrom(users);
    }

    function test_BlacklistToRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklistTo(users);
    }

    function test_BlacklistRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklist(users);
    }

    function test_BlacklistFromRevertAlreadyBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users);
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedFrom
                    .selector,
                user1
            )
        );
        complianceModule.blacklistFrom(users);
    }

    function test_BlacklistToRevertAlreadyBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistTo(users);
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedTo
                    .selector,
                user1
            )
        );
        complianceModule.blacklistTo(users);
    }

    function test_BlacklistRevertNonAdmin() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        vm.startPrank(nonAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                nonAdmin,
                BLACKLIST_ADMIN_ROLE
            )
        );
        complianceModule.blacklist(users);
        vm.stopPrank();
    }

    function test_UnblacklistFrom() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistFrom(users);
        complianceModule.unblacklistFrom(users);
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedFrom(user2));
    }

    function test_UnblacklistTo() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistTo(users);
        complianceModule.unblacklistTo(users);
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklistedTo(user2));
    }

    function test_UnblacklistBoth() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklist(users);
        complianceModule.unblacklist(users);
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedFrom(user2));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklistedTo(user2));
        assertFalse(complianceModule.isBlacklisted(user1));
        assertFalse(complianceModule.isBlacklisted(user2));
    }

    function test_UnblacklistFromRevertNotBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance.AddressNotBlacklistedFrom.selector,
                user1
            )
        );
        complianceModule.unblacklistFrom(users);
    }

    function test_UnblacklistToRevertNotBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance.AddressNotBlacklistedTo.selector,
                user1
            )
        );
        complianceModule.unblacklistTo(users);
    }

    function test_CanTransferAllowedAddresses() public {
        assertTrue(complianceModule.canTransfer(user1, user2, 100));
    }

    function test_CanTransferBlockedSender() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
    }

    function test_CanTransferBlockedRecipient() public {
        address[] memory users = new address[](1);
        users[0] = user2;

        complianceModule.blacklistTo(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
    }

    function test_CanTransferCompletelyBlocked() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklist(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
        assertFalse(complianceModule.canTransfer(user2, user1, 100));
    }

    function test_CanTransferWithFailureReasonAllowedAddresses() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithFailureReasonBlockedSender() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users);
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(reason, "Sender address is blacklisted from sending");
    }

    function test_CanTransferWithFailureReasonBlockedRecipient() public {
        address[] memory users = new address[](1);
        users[0] = user2;

        complianceModule.blacklistTo(users);
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(reason, "Recipient address is blacklisted from receiving");
    }

    function test_Events() public {
        address[] memory fromUsers = new address[](1);
        fromUsers[0] = user1;

        address[] memory toUsers = new address[](1);
        toUsers[0] = user2;

        address[] memory bothUsers = new address[](1);
        bothUsers[0] = user3;

        vm.expectEmit(true, false, false, false);
        emit AddressBlacklistedFrom(user1);
        complianceModule.blacklistFrom(fromUsers);

        vm.expectEmit(true, false, false, false);
        emit AddressUnblacklistedFrom(user1);
        complianceModule.unblacklistFrom(fromUsers);

        vm.expectEmit(true, false, false, false);
        emit AddressBlacklistedTo(user2);
        complianceModule.blacklistTo(toUsers);

        vm.expectEmit(true, false, false, false);
        emit AddressUnblacklistedTo(user2);
        complianceModule.unblacklistTo(toUsers);

        vm.expectEmit(true, false, false, false);
        emit AddressBlacklistedFrom(user3);
        vm.expectEmit(true, false, false, false);
        emit AddressBlacklistedTo(user3);
        vm.expectEmit(true, false, false, false);
        emit AddressBlacklisted(user3);
        complianceModule.blacklist(bothUsers);

        vm.expectEmit(true, false, false, false);
        emit AddressUnblacklistedFrom(user3);
        vm.expectEmit(true, false, false, false);
        emit AddressUnblacklistedTo(user3);
        vm.expectEmit(true, false, false, false);
        emit AddressUnblacklisted(user3);
        complianceModule.unblacklist(bothUsers);
    }

    // Test double blacklisting
    function test_DoubleBlacklistFromFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users);
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedFrom
                    .selector,
                user1
            )
        );
        complianceModule.blacklistFrom(users);
    }

    // Test double blacklistTo
    function test_DoubleBlacklistToFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistTo(users);
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedTo
                    .selector,
                user1
            )
        );
        complianceModule.blacklistTo(users);
    }

    // Test unblacklist non-blacklisted address
    function test_UnblacklistFromNonBlacklistedFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance.AddressNotBlacklistedFrom.selector,
                user1
            )
        );
        complianceModule.unblacklistFrom(users);
    }

    // Test unblacklistTo non-blacklisted address
    function test_UnblacklistToNonBlacklistedFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance.AddressNotBlacklistedTo.selector,
                user1
            )
        );
        complianceModule.unblacklistTo(users);
    }

    // Test blacklist with non-admin
    function test_BlacklistFromNonAdminFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        vm.startPrank(nonAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                nonAdmin,
                BLACKLIST_ADMIN_ROLE
            )
        );
        complianceModule.blacklistFrom(users);
        vm.stopPrank();
    }

    // Test partial blacklist/unblacklist transitions
    function test_PartialBlacklistStates() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // Blacklist from only
        complianceModule.blacklistFrom(users);
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));

        // Add blacklistTo - should be fully blacklisted
        complianceModule.blacklistTo(users);
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklisted(user1));

        // Remove blacklistFrom - should still be blacklistedTo
        complianceModule.unblacklistFrom(users);
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));
    }

    // Test transfer restrictions with various blacklist states
    function test_CanTransferRestrictionsComprehensive() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // Test when neither is blacklisted
        assertTrue(complianceModule.canTransfer(user1, user2, 100));
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertTrue(success);
        assertEq(reason, "");

        // Test with sender blacklisted
        complianceModule.blacklistFrom(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
        (success, reason) = complianceModule.canTransferWithFailureReason(
            user1,
            user2,
            100
        );
        assertFalse(success);
        assertEq(reason, "Sender address is blacklisted from sending");

        // Test with both sender and receiver blacklisted
        users[0] = user2;
        complianceModule.blacklistTo(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
        (success, reason) = complianceModule.canTransferWithFailureReason(
            user1,
            user2,
            100
        );
        assertFalse(success);
        assertEq(reason, "Sender address is blacklisted from sending");
    }

    function test_DoubleBlacklistFromToCase() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // First blacklistFrom and then blacklistTo
        complianceModule.blacklistFrom(users);
        complianceModule.blacklistTo(users);

        // Should be fully blacklisted
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklisted(user1));

        // Now unblacklist only one direction
        complianceModule.unblacklistFrom(users);

        // Should still be blacklisted in one direction but not the other
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));
    }

    function test_BlacklistAlreadyBlacklistedBoth() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // First blacklist in one direction
        complianceModule.blacklistFrom(users);

        // Then blacklist in both directions
        complianceModule.blacklist(users);

        // Should be fully blacklisted
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklisted(user1));

        // Then unblacklist both directions
        complianceModule.unblacklist(users);

        // Should be fully unblacklisted
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));
    }

    function test_UnblacklistAlreadyUnblacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // First blacklist both directions
        complianceModule.blacklist(users);

        // Unblacklist from
        complianceModule.unblacklistFrom(users);

        // Try to unblacklist from again
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance.AddressNotBlacklistedFrom.selector,
                user1
            )
        );
        complianceModule.unblacklistFrom(users);
    }

    function test_EmptyArrayCases() public {
        address[] memory emptyArray = new address[](0);

        // Test all functions with empty array
        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklistFrom(emptyArray);

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklistTo(emptyArray);

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklist(emptyArray);

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.unblacklistFrom(emptyArray);

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.unblacklistTo(emptyArray);

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.unblacklist(emptyArray);
    }

    function test_CanTransferWithFailureReasonComprehensive() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // Test blacklistFrom then check transfer
        complianceModule.blacklistFrom(users);
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(reason, "Sender address is blacklisted from sending");

        // Unblacklist from and blacklist to, then check again
        complianceModule.unblacklistFrom(users);
        users[0] = user2;
        complianceModule.blacklistTo(users);
        (success, reason) = complianceModule.canTransferWithFailureReason(
            user1,
            user2,
            100
        );
        assertFalse(success);
        assertEq(reason, "Recipient address is blacklisted from receiving");
    }
}
