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

    function testInitialState() public {
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

    function testBlacklistFrom() public {
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

    function testBlacklistTo() public {
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

    function testBlacklistBoth() public {
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

    function testBlacklistFromRevertEmptyList() public {
        address[] memory users = new address[](0);
        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklistFrom(users);
    }

    function testBlacklistFromRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklistFrom(users);
    }

    function testBlacklistToRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklistTo(users);
    }

    function testBlacklistRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklist(users);
    }

    function testBlacklistFromRevertAlreadyBlacklisted() public {
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

    function testBlacklistToRevertAlreadyBlacklisted() public {
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

    function testBlacklistRevertNonAdmin() public {
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

    function testUnblacklistFrom() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistFrom(users);
        complianceModule.unblacklistFrom(users);
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedFrom(user2));
    }

    function testUnblacklistTo() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistTo(users);
        complianceModule.unblacklistTo(users);
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklistedTo(user2));
    }

    function testUnblacklistBoth() public {
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

    function testUnblacklistFromRevertNotBlacklisted() public {
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

    function testUnblacklistToRevertNotBlacklisted() public {
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

    function testCanTransferAllowedAddresses() public {
        assertTrue(complianceModule.canTransfer(user1, user2, 100));
    }

    function testCanTransferBlockedSender() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
    }

    function testCanTransferBlockedRecipient() public {
        address[] memory users = new address[](1);
        users[0] = user2;

        complianceModule.blacklistTo(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
    }

    function testCanTransferCompletelyBlocked() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklist(users);
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
        assertFalse(complianceModule.canTransfer(user2, user1, 100));
    }

    function testCanTransferWithFailureReasonAllowedAddresses() public {
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertTrue(success);
        assertEq(reason, "");
    }

    function testCanTransferWithFailureReasonBlockedSender() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users);
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(reason, "Sender address is blacklisted from sending");
    }

    function testCanTransferWithFailureReasonBlockedRecipient() public {
        address[] memory users = new address[](1);
        users[0] = user2;

        complianceModule.blacklistTo(users);
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(reason, "Recipient address is blacklisted from receiving");
    }

    function testEvents() public {
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
    function testDoubleBlacklistFromFails() public {
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
    function testDoubleBlacklistToFails() public {
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
    function testUnblacklistFromNonBlacklistedFails() public {
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
    function testUnblacklistToNonBlacklistedFails() public {
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
    function testBlacklistFromNonAdminFails() public {
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
    function testPartialBlacklistStates() public {
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
    function testCanTransferRestrictionsComprehensive() public {
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

    function testDoubleBlacklistFromToCase() public {
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

    function testBlacklistAlreadyBlacklistedBoth() public {
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

    function testUnblacklistAlreadyUnblacklisted() public {
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

    function testEmptyArrayCases() public {
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

    function testCanTransferWithFailureReasonComprehensive() public {
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
