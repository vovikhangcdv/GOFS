// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test, console2} from "forge-std/Test.sol";
import {CompliantToken} from "../src/CompliantToken.sol";
import {ICompliance} from "../src/interfaces/ICompliance.sol";

// Mock compliance contract for testing
contract MockCompliance is ICompliance {
    bool private allowTransfer;
    string private failureMessage;
    bool private allowMinting;
    bool private allowBurning;

    constructor() {
        // By default allow minting and burning for most tests
        allowMinting = true;
        allowBurning = true;
    }

    // Set whether transfers should be allowed in tests
    function setTransferAllowed(bool allowed, string memory message) external {
        allowTransfer = allowed;
        failureMessage = message;
    }

    function setMintingAllowed(bool allowed) external {
        allowMinting = allowed;
    }

    function setBurningAllowed(bool allowed) external {
        allowBurning = allowed;
    }

    function canTransfer(
        address from,
        address to,
        uint256
    ) external view returns (bool) {
        if (from == address(0)) {
            return allowMinting;
        }
        if (to == address(0)) {
            return allowBurning;
        }
        return allowTransfer;
    }

    function canTransferWithFailureReason(
        address from,
        address to,
        uint256 amount
    ) external view returns (bool, string memory) {
        if (from == address(0)) {
            return (allowMinting, allowMinting ? "" : "Minting not allowed");
        }
        if (to == address(0)) {
            return (allowBurning, allowBurning ? "" : "Burning not allowed");
        }
        return (allowTransfer, failureMessage);
    }
}

contract CompliantTokenTest is Test {
    CompliantToken public token;
    MockCompliance public compliance;
    address public owner;
    address public user1;
    address public user2;

    // Events to test
    event Transfer(address indexed from, address indexed to, uint256 value);

    // Error type for Ownable unauthorized access
    error OwnableUnauthorizedAccount(address account);

    function setUp() public {
        owner = makeAddr("owner");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        vm.startPrank(owner);
        compliance = new MockCompliance();
        token = new CompliantToken("Test Token", "TEST", address(compliance));
        vm.stopPrank();
    }

    function test_Constructor() public {
        assertEq(token.name(), "Test Token");
        assertEq(token.symbol(), "TEST");
        assertEq(address(token.compliance()), address(compliance));
        assertEq(token.owner(), owner);
    }

    function test_ConstructorRevertsOnZeroComplianceAddress() public {
        vm.startPrank(owner);
        vm.expectRevert(CompliantToken.ZeroAddressCompliance.selector);
        new CompliantToken("Test Token", "TEST", address(0));
        vm.stopPrank();
    }

    function test_Mint() public {
        vm.startPrank(owner);
        token.mint(user1, 1000);
        vm.stopPrank();

        assertEq(token.balanceOf(user1), 1000);
        assertEq(token.totalSupply(), 1000);
    }

    function test_MintRevertsForNonOwner() public {
        vm.startPrank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, user1)
        );
        token.mint(user1, 1000);
        vm.stopPrank();
    }

    function test_Burn() public {
        // First mint some tokens
        vm.startPrank(owner);
        token.mint(user1, 1000);
        token.burn(user1, 500);
        vm.stopPrank();

        assertEq(token.balanceOf(user1), 500);
        assertEq(token.totalSupply(), 500);
    }

    function test_BurnRevertsForNonOwner() public {
        vm.startPrank(owner);
        token.mint(user1, 1000);
        vm.stopPrank();

        vm.startPrank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, user1)
        );
        token.burn(user1, 500);
        vm.stopPrank();
    }

    function test_TransferWithCompliance() public {
        // Setup compliance to allow transfers
        compliance.setTransferAllowed(true, "");

        // Mint tokens to user1
        vm.startPrank(owner);
        token.mint(user1, 1000);
        vm.stopPrank();

        // User1 transfers to user2
        vm.startPrank(user1);
        bool success = token.transfer(user2, 500);
        vm.stopPrank();

        assertTrue(success);
        assertEq(token.balanceOf(user1), 500);
        assertEq(token.balanceOf(user2), 500);
    }

    function test_TransferFailsWithoutCompliance() public {
        // Setup compliance to disallow transfers
        compliance.setTransferAllowed(false, "Transfer not allowed");

        // Mint tokens to user1
        vm.startPrank(owner);
        token.mint(user1, 1000);
        vm.stopPrank();

        // User1 attempts to transfer to user2
        vm.startPrank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(
                CompliantToken.ComplianceCheckFailed.selector,
                user1,
                user2,
                500,
                "Transfer not allowed"
            )
        );
        token.transfer(user2, 500);
        vm.stopPrank();

        // Verify balances remain unchanged
        assertEq(token.balanceOf(user1), 1000);
        assertEq(token.balanceOf(user2), 0);
    }

    function test_TransferFromWithCompliance() public {
        // Setup compliance to allow transfers
        compliance.setTransferAllowed(true, "");

        // Mint tokens to user1
        vm.startPrank(owner);
        token.mint(user1, 1000);
        vm.stopPrank();

        // User1 approves user2 to spend tokens
        vm.startPrank(user1);
        token.approve(user2, 500);
        vm.stopPrank();

        // User2 transfers from user1 to themselves
        vm.startPrank(user2);
        bool success = token.transferFrom(user1, user2, 500);
        vm.stopPrank();

        assertTrue(success);
        assertEq(token.balanceOf(user1), 500);
        assertEq(token.balanceOf(user2), 500);
        assertEq(token.allowance(user1, user2), 0);
    }

    function test_TransferFromFailsWithoutCompliance() public {
        // Setup compliance to disallow transfers
        compliance.setTransferAllowed(false, "Transfer not allowed");

        // Mint tokens to user1
        vm.startPrank(owner);
        token.mint(user1, 1000);
        vm.stopPrank();

        // User1 approves user2 to spend tokens
        vm.startPrank(user1);
        token.approve(user2, 500);
        vm.stopPrank();

        // User2 attempts to transfer from user1
        vm.startPrank(user2);
        vm.expectRevert(
            abi.encodeWithSelector(
                CompliantToken.ComplianceCheckFailed.selector,
                user1,
                user2,
                500,
                "Transfer not allowed"
            )
        );
        token.transferFrom(user1, user2, 500);
        vm.stopPrank();

        // Verify balances remain unchanged
        assertEq(token.balanceOf(user1), 1000);
        assertEq(token.balanceOf(user2), 0);
        assertEq(token.allowance(user1, user2), 500);
    }

    function test_MintingAndBurningWithCompliance() public {
        // Setup compliance to disallow transfers but allow minting/burning
        compliance.setTransferAllowed(false, "Transfer not allowed");
        compliance.setMintingAllowed(true);
        compliance.setBurningAllowed(true);

        // Minting should work when compliance allows it
        vm.startPrank(owner);
        token.mint(user1, 1000);
        assertEq(token.balanceOf(user1), 1000);

        // Burning should work when compliance allows it
        token.burn(user1, 500);
        assertEq(token.balanceOf(user1), 500);
        vm.stopPrank();
    }

    function test_MintingFailsWhenComplianceDisallows() public {
        compliance.setMintingAllowed(false);

        vm.startPrank(owner);
        vm.expectRevert(
            abi.encodeWithSelector(
                CompliantToken.ComplianceCheckFailed.selector,
                address(0),
                user1,
                1000,
                "Minting not allowed"
            )
        );
        token.mint(user1, 1000);
        vm.stopPrank();
    }

    function test_BurningFailsWhenComplianceDisallows() public {
        // First allow minting to setup the test
        compliance.setMintingAllowed(true);
        vm.prank(owner);
        token.mint(user1, 1000);

        // Now test burning compliance
        compliance.setBurningAllowed(false);

        vm.startPrank(owner);
        vm.expectRevert(
            abi.encodeWithSelector(
                CompliantToken.ComplianceCheckFailed.selector,
                user1,
                address(0),
                500,
                "Burning not allowed"
            )
        );
        token.burn(user1, 500);
        vm.stopPrank();
    }

    function test_setComplianceAddress() public {
        // Owner can set a new compliance address
        address newCompliance = makeAddr("newCompliance");
        vm.startPrank(owner);
        token.setCompliance(newCompliance);
        vm.stopPrank();

        assertEq(address(token.compliance()), newCompliance);
    }
    function test_setComplianceAddressRevertsForNonOwner() public {
        // Non-owner cannot set compliance address
        address newCompliance = makeAddr("newCompliance");
        vm.startPrank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(OwnableUnauthorizedAccount.selector, user1)
        );
        token.setCompliance(newCompliance);
        vm.stopPrank();
    }

    function test_setComplianceAddressRevertsOnZeroAddress() public {
        // Owner cannot set compliance to zero address
        vm.startPrank(owner);
        vm.expectRevert(CompliantToken.ZeroAddressCompliance.selector);
        token.setCompliance(address(0));
        vm.stopPrank();
        assertEq(address(token.compliance()), address(compliance));
    }
}
