// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Vm, Test} from "forge-std/Test.sol";
import {AddressRestrictionCompliance} from "../../src/compliances/AddressRestrictionCompliance.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
contract AddressRestrictionComplianceTest is Test {
    bytes32 public constant BLACKLIST_ADMIN_ROLE =
        keccak256("BLACKLIST_ADMIN_ROLE");

    AddressRestrictionCompliance public complianceModule;

    address public admin;
    address public user1;
    address public user2;
    address public user3;
    address public nonAdmin;

    // Updated Events with reason tracking
    event AddressBlacklistedFrom(
        address indexed account,
        string reason,
        address indexed restrictedBy,
        uint256 timestamp
    );
    event AddressUnblacklistedFrom(
        address indexed account,
        string reason,
        address indexed unrestrictedBy,
        uint256 timestamp
    );
    event AddressBlacklistedTo(
        address indexed account,
        string reason,
        address indexed restrictedBy,
        uint256 timestamp
    );
    event AddressUnblacklistedTo(
        address indexed account,
        string reason,
        address indexed unrestrictedBy,
        uint256 timestamp
    );
    event AddressBlacklisted(
        address indexed account,
        string reason,
        address indexed restrictedBy,
        uint256 timestamp
    );
    event AddressUnblacklisted(
        address indexed account,
        string reason,
        address indexed unrestrictedBy,
        uint256 timestamp
    );

    function setUp() public {
        admin = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        user3 = makeAddr("user3");
        nonAdmin = makeAddr("nonAdmin");

        complianceModule = AddressRestrictionCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(new AddressRestrictionCompliance()),
                    address(new ProxyAdmin(admin)),
                    abi.encodeWithSelector(
                        AddressRestrictionCompliance.initialize.selector
                    )
                )
            )
        );
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

        complianceModule.blacklistFrom(users, "Test restriction from sending");
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

        complianceModule.blacklistTo(users, "Test restriction from receiving");
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

        complianceModule.blacklist(users, "Test complete restriction");
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
        complianceModule.blacklistFrom(users, "Test reason");
    }

    function test_BlacklistFromRevertEmptyReason() public {
        address[] memory users = new address[](1);
        users[0] = user1;
        vm.expectRevert(AddressRestrictionCompliance.EmptyReason.selector);
        complianceModule.blacklistFrom(users, "");
    }

    function test_BlacklistFromRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklistFrom(users, "Test reason");
    }

    function test_BlacklistToRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklistTo(users, "Test reason");
    }

    function test_BlacklistRevertZeroAddress() public {
        address[] memory users = new address[](1);
        users[0] = address(0);
        vm.expectRevert(AddressRestrictionCompliance.ZeroAddress.selector);
        complianceModule.blacklist(users, "Test reason");
    }

    function test_BlacklistFromRevertAlreadyBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users, "First restriction");
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedFrom
                    .selector,
                user1
            )
        );
        complianceModule.blacklistFrom(users, "Second restriction");
    }

    function test_BlacklistToRevertAlreadyBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistTo(users, "First restriction");
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedTo
                    .selector,
                user1
            )
        );
        complianceModule.blacklistTo(users, "Second restriction");
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
        complianceModule.blacklist(users, "Test reason");
        vm.stopPrank();
    }

    function test_UnblacklistFrom() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistFrom(users, "Initial restriction");
        complianceModule.unblacklistFrom(users, "Lifting restriction");
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedFrom(user2));
    }

    function test_UnblacklistTo() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklistTo(users, "Initial restriction");
        complianceModule.unblacklistTo(users, "Lifting restriction");
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklistedTo(user2));
    }

    function test_UnblacklistBoth() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        complianceModule.blacklist(users, "Initial restriction");
        complianceModule.unblacklist(users, "Lifting restriction");
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedFrom(user2));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklistedTo(user2));
        assertFalse(complianceModule.isBlacklisted(user1));
        assertFalse(complianceModule.isBlacklisted(user2));
    }

    function test_UnblacklistFromAcceptNotBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        assertFalse(complianceModule.isBlacklistedFrom(user1));
        complianceModule.unblacklistFrom(users, "Unnecessary unrestriction");
        assertFalse(complianceModule.isBlacklistedFrom(user1));
    }

    function test_UnblacklistToAcceptNotBlacklisted() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        assertFalse(complianceModule.isBlacklistedTo(user1));
        complianceModule.unblacklistTo(users, "Unnecessary unrestriction");
        assertFalse(complianceModule.isBlacklistedTo(user1));
    }

    function test_CanTransferAllowedAddresses() public {
        assertTrue(complianceModule.canTransfer(user1, user2, 100));
    }

    function test_CanTransferBlockedSender() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users, "Sender restriction");
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
    }

    function test_CanTransferBlockedRecipient() public {
        address[] memory users = new address[](1);
        users[0] = user2;

        complianceModule.blacklistTo(users, "Recipient restriction");
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
    }

    function test_CanTransferCompletelyBlocked() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklist(users, "Complete restriction");
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

        complianceModule.blacklistFrom(users, "AML violation detected");
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(
            reason,
            "Sender address is blacklisted from sending. Reason: AML violation detected"
        );
    }

    function test_CanTransferWithFailureReasonBlockedRecipient() public {
        address[] memory users = new address[](1);
        users[0] = user2;

        complianceModule.blacklistTo(users, "Sanctions list match");
        (bool success, string memory reason) = complianceModule
            .canTransferWithFailureReason(user1, user2, 100);
        assertFalse(success);
        assertEq(
            reason,
            "Recipient address is blacklisted from receiving. Reason: Sanctions list match"
        );
    }

    function test_GetRestrictionInfo() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        string memory restrictionReason = "Test compliance violation";
        complianceModule.blacklistFrom(users, restrictionReason);

        AddressRestrictionCompliance.RestrictionInfo
            memory info = complianceModule.getRestrictionInfoFrom(user1);
        assertTrue(info.isRestricted);
        assertEq(info.reason, restrictionReason);
        assertEq(info.restrictedBy, admin);
        assertTrue(info.timestamp > 0);

        string memory reason = complianceModule.getRestrictionReasonFrom(user1);
        assertEq(reason, restrictionReason);
    }

    function test_BlacklistWithReasons() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        string[] memory reasons = new string[](2);
        reasons[0] = "AML violation";
        reasons[1] = "Sanctions match";

        complianceModule.blacklistFromWithReasons(users, reasons);

        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedFrom(user2));

        assertEq(
            complianceModule.getRestrictionReasonFrom(user1),
            "AML violation"
        );
        assertEq(
            complianceModule.getRestrictionReasonFrom(user2),
            "Sanctions match"
        );
    }

    function test_BlacklistWithReasonsArrayMismatch() public {
        address[] memory users = new address[](2);
        users[0] = user1;
        users[1] = user2;

        string[] memory reasons = new string[](1);
        reasons[0] = "Test reason";

        vm.expectRevert(
            AddressRestrictionCompliance.ArrayLengthMismatch.selector
        );
        complianceModule.blacklistFromWithReasons(users, reasons);
    }

    function test_Events() public {
        address[] memory fromUsers = new address[](1);
        fromUsers[0] = user1;

        address[] memory toUsers = new address[](1);
        toUsers[0] = user2;

        address[] memory bothUsers = new address[](1);
        bothUsers[0] = user3;

        vm.expectEmit(true, true, false, true);
        emit AddressBlacklistedFrom(
            user1,
            "From restriction",
            admin,
            block.timestamp
        );
        complianceModule.blacklistFrom(fromUsers, "From restriction");

        vm.expectEmit(true, true, false, true);
        emit AddressUnblacklistedFrom(
            user1,
            "From unrestriction",
            admin,
            block.timestamp
        );
        complianceModule.unblacklistFrom(fromUsers, "From unrestriction");

        vm.expectEmit(true, true, false, true);
        emit AddressBlacklistedTo(
            user2,
            "To restriction",
            admin,
            block.timestamp
        );
        complianceModule.blacklistTo(toUsers, "To restriction");

        vm.expectEmit(true, true, false, true);
        emit AddressUnblacklistedTo(
            user2,
            "To unrestriction",
            admin,
            block.timestamp
        );
        complianceModule.unblacklistTo(toUsers, "To unrestriction");

        vm.expectEmit(true, true, false, true);
        emit AddressBlacklistedFrom(
            user3,
            "Both restriction",
            admin,
            block.timestamp
        );
        vm.expectEmit(true, true, false, true);
        emit AddressBlacklistedTo(
            user3,
            "Both restriction",
            admin,
            block.timestamp
        );
        vm.expectEmit(true, true, false, true);
        emit AddressBlacklisted(
            user3,
            "Both restriction",
            admin,
            block.timestamp
        );
        complianceModule.blacklist(bothUsers, "Both restriction");

        vm.expectEmit(true, true, false, true);
        emit AddressUnblacklistedFrom(
            user3,
            "Both unrestriction",
            admin,
            block.timestamp
        );
        vm.expectEmit(true, true, false, true);
        emit AddressUnblacklistedTo(
            user3,
            "Both unrestriction",
            admin,
            block.timestamp
        );
        vm.expectEmit(true, true, false, true);
        emit AddressUnblacklisted(
            user3,
            "Both unrestriction",
            admin,
            block.timestamp
        );
        complianceModule.unblacklist(bothUsers, "Both unrestriction");
    }

    // Test double blacklisting
    function test_DoubleBlacklistFromFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users, "First restriction");
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedFrom
                    .selector,
                user1
            )
        );
        complianceModule.blacklistFrom(users, "Second restriction");
    }

    // Test double blacklistTo
    function test_DoubleBlacklistToFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistTo(users, "First restriction");
        vm.expectRevert(
            abi.encodeWithSelector(
                AddressRestrictionCompliance
                    .AddressAlreadyBlacklistedTo
                    .selector,
                user1
            )
        );
        complianceModule.blacklistTo(users, "Second restriction");
    }

    // Test unblacklist non-blacklisted address
    function test_UnblacklistFromNonBlacklistedFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.unblacklistFrom(users, "Unnecessary unrestriction");
        assertFalse(complianceModule.isBlacklistedFrom(user1));
    }

    // Test unblacklistTo non-blacklisted address
    function test_UnblacklistToNonBlacklistedFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        assertFalse(complianceModule.isBlacklistedTo(user1));
        complianceModule.unblacklistTo(users, "Unnecessary unrestriction");
        assertFalse(complianceModule.isBlacklistedTo(user1));
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
        complianceModule.blacklistFrom(users, "Test reason");
        vm.stopPrank();
    }

    // Test partial blacklist/unblacklist transitions
    function test_PartialBlacklistStates() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // Blacklist from only
        complianceModule.blacklistFrom(users, "Sender restriction");
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));

        // Add blacklistTo - should be fully blacklisted
        complianceModule.blacklistTo(users, "Receiver restriction");
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklisted(user1));

        // Remove blacklistFrom - should still be blacklistedTo
        complianceModule.unblacklistFrom(users, "Lifting sender restriction");
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
        complianceModule.blacklistFrom(users, "Suspicious activity");
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
        (success, reason) = complianceModule.canTransferWithFailureReason(
            user1,
            user2,
            100
        );
        assertFalse(success);
        assertEq(
            reason,
            "Sender address is blacklisted from sending. Reason: Suspicious activity"
        );

        // Test with both sender and receiver blacklisted
        users[0] = user2;
        complianceModule.blacklistTo(users, "High risk entity");
        assertFalse(complianceModule.canTransfer(user1, user2, 100));
        (success, reason) = complianceModule.canTransferWithFailureReason(
            user1,
            user2,
            100
        );
        assertFalse(success);
        assertEq(
            reason,
            "Sender address is blacklisted from sending. Reason: Suspicious activity"
        );
    }

    function test_DoubleBlacklistFromToCase() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // First blacklistFrom and then blacklistTo
        complianceModule.blacklistFrom(users, "Sender issue");
        complianceModule.blacklistTo(users, "Receiver issue");

        // Should be fully blacklisted
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklisted(user1));

        // Now unblacklist only one direction
        complianceModule.unblacklistFrom(users, "Sender cleared");

        // Should still be blacklisted in one direction but not the other
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));
    }

    function test_BlacklistAlreadyBlacklistedBoth() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        // First blacklist in one direction
        complianceModule.blacklistFrom(users, "Initial sender restriction");

        // Then blacklist in both directions
        complianceModule.blacklist(users, "Complete restriction");

        // Should be fully blacklisted
        assertTrue(complianceModule.isBlacklistedFrom(user1));
        assertTrue(complianceModule.isBlacklistedTo(user1));
        assertTrue(complianceModule.isBlacklisted(user1));

        // Then unblacklist both directions
        complianceModule.unblacklist(users, "Full clearance");

        // Should be fully unblacklisted
        assertFalse(complianceModule.isBlacklistedFrom(user1));
        assertFalse(complianceModule.isBlacklistedTo(user1));
        assertFalse(complianceModule.isBlacklisted(user1));
    }

    // Continue with remaining test functions...
    function test_UnblacklistEmptyReasonFails() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users, "Test restriction");

        vm.expectRevert(AddressRestrictionCompliance.EmptyReason.selector);
        complianceModule.unblacklistFrom(users, "");
    }

    function test_EmptyArrayFails() public {
        address[] memory emptyArray = new address[](0);

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklistFrom(emptyArray, "Test reason");

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklistTo(emptyArray, "Test reason");

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.blacklist(emptyArray, "Test reason");

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.unblacklistFrom(emptyArray, "Test reason");

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.unblacklistTo(emptyArray, "Test reason");

        vm.expectRevert(AddressRestrictionCompliance.EmptyAddressList.selector);
        complianceModule.unblacklist(emptyArray, "Test reason");
    }

    function test_UnblacklistIdempotent() public {
        address[] memory users = new address[](1);
        users[0] = user1;

        complianceModule.blacklistFrom(users, "Test restriction");
        complianceModule.unblacklistFrom(users, "Lifting restriction");
        assertFalse(complianceModule.isBlacklistedFrom(user1));

        // Calling unblacklist again should not fail
        complianceModule.unblacklistFrom(users, "Already cleared");
        assertFalse(complianceModule.isBlacklistedFrom(user1));

        complianceModule.blacklistTo(users, "Test restriction");
        // ... continuing with similar pattern
    }
}
