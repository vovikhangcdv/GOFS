// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {CompliantToken} from "../src/CompliantToken.sol";
import {ICompliance} from "../src/interfaces/ICompliance.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {TxType} from "../src/interfaces/ITypes.sol";

// Simple mock compliance for testing typed transfers
contract SimpleCompliance is ICompliance {
    function supportsInterface(
        bytes4 interfaceId
    ) external pure returns (bool) {
        return interfaceId == type(ICompliance).interfaceId;
    }

    function canTransfer(
        address,
        address,
        uint256
    ) external pure returns (bool) {
        return true;
    }

    function canTransferWithFailureReason(
        address,
        address,
        uint256
    ) external pure returns (bool, string memory) {
        return (true, "");
    }

    function canTransferWithType(
        address,
        address,
        uint256,
        TxType
    ) external pure returns (bool) {
        return true;
    }

    function canTransferWithTypeAndFailureReason(
        address,
        address,
        uint256,
        TxType
    ) external pure returns (bool, string memory) {
        return (true, "");
    }
}

contract TypedTokenTest is Test {
    CompliantToken public token;
    SimpleCompliance public compliance;
    address public owner;
    address public user1;
    address public user2;

    event TypedTransfer(
        address indexed from,
        address indexed to,
        uint256 value,
        TxType indexed txType
    );

    function setUp() public {
        owner = makeAddr("owner");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        vm.startPrank(owner);
        compliance = new SimpleCompliance();
        token = new CompliantToken();
        token.initialize("Typed Token", "TYPED", address(compliance));
        vm.stopPrank();
    }

    function test_TransferWithType() public {
        // Mint tokens to user1
        vm.prank(owner);
        token.mint(user1, 1000);

        // Test typed transfer
        TxType txType = TxType.wrap(1);

        vm.expectEmit(true, true, true, true);
        emit TypedTransfer(user1, user2, 500, txType);

        vm.prank(user1);
        bool success = token.transferWithType(user2, 500, txType);

        assertTrue(success);
        assertEq(token.balanceOf(user1), 500);
        assertEq(token.balanceOf(user2), 500);
    }

    function test_TransferFromWithType() public {
        // Mint tokens to user1
        vm.prank(owner);
        token.mint(user1, 1000);

        // User1 approves user2
        vm.prank(user1);
        token.approve(user2, 500);

        // Test typed transferFrom
        TxType txType = TxType.wrap(2);

        vm.expectEmit(true, true, true, true);
        emit TypedTransfer(user1, user2, 500, txType);

        vm.prank(user2);
        bool success = token.transferFromWithType(user1, user2, 500, txType);

        assertTrue(success);
        assertEq(token.balanceOf(user1), 500);
        assertEq(token.balanceOf(user2), 500);
        assertEq(token.allowance(user1, user2), 0);
    }

    function test_StandardTransferEmitsTypedEvent() public {
        // Mint tokens to user1
        vm.prank(owner);
        token.mint(user1, 1000);

        // Standard transfer should emit TypedTransfer with TxType(0)
        vm.expectEmit(true, true, true, true);
        emit TypedTransfer(user1, user2, 500, TxType.wrap(0));

        vm.prank(user1);
        bool success = token.transfer(user2, 500);

        assertTrue(success);
        assertEq(token.balanceOf(user1), 500);
        assertEq(token.balanceOf(user2), 500);
    }

    function test_TypedTransferWithCompliance() public {
        // Mint tokens to user1
        vm.prank(owner);
        token.mint(user1, 1000);

        // Test that compliance is called for typed transfers
        TxType txType = TxType.wrap(10);

        vm.prank(user1);
        bool success = token.transferWithType(user2, 500, txType);

        assertTrue(success);
        assertEq(token.balanceOf(user1), 500);
        assertEq(token.balanceOf(user2), 500);
    }
}
