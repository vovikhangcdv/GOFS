// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {SupplyCompliance} from "../../src/compliances/SupplyCompliance.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
contract MockToken is ERC20 {
    constructor(
        string memory name,
        string memory symbol,
        uint8 decimals
    ) ERC20(name, symbol) {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract SupplyComplianceTest is Test {
    SupplyCompliance public compliance;
    MockToken public token;
    address public admin;
    address public user1;
    address public user2;

    bytes32 public constant SUPPLY_ADMIN_ROLE = keccak256("SUPPLY_ADMIN_ROLE");
    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    event MaxSupplyUpdated(uint256 oldMax, uint256 newMax);

    function setUp() public {
        // Set up accounts
        admin = makeAddr("admin");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        vm.startPrank(admin);

        // Deploy token
        token = new MockToken("Test Token", "TEST", 0);

        // Deploy compliance module
        compliance = SupplyCompliance(
            address(
                new TransparentUpgradeableProxy(
                    address(new SupplyCompliance()),
                    address(new ProxyAdmin(admin)),
                    abi.encodeWithSelector(
                        SupplyCompliance.initialize.selector,
                        address(token)
                    )
                )
            )
        );

        // Setup roles
        compliance.grantRole(compliance.DEFAULT_ADMIN_ROLE(), admin);
        compliance.grantRole(compliance.SUPPLY_ADMIN_ROLE(), admin);

        vm.stopPrank();
    }
    function test_SetMaxSupply() public {
        uint256 newMax = 1000000;

        // Test event emission
        vm.expectEmit(true, true, true, true);
        emit MaxSupplyUpdated(0, newMax);
        vm.prank(admin);
        compliance.setMaxSupply(newMax);

        // Test max supply updated
        assertEq(compliance.maxSupply(), newMax);
    }

    function test_SetMaxSupply_Unauthorized() public {
        vm.startPrank(user1);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                user1,
                SUPPLY_ADMIN_ROLE
            )
        );
        compliance.setMaxSupply(1000);
        vm.stopPrank();
    }

    function test_CanTransfer_RegularTransfer() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);
        token.mint(user1, 100);

        // Regular transfer between users should always return true
        assertTrue(compliance.canTransfer(user1, user2, 50));
    }

    function test_CanTransfer_Mint_UnderLimit() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);
        token.mint(user1, 100);

        // Mint (transfer from address(0)) under limit should return true
        assertTrue(compliance.canTransfer(address(0), user1, 900));
    }

    function test_CanTransfer_Mint_OverLimit() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);
        token.mint(user1, 100);

        // Mint (transfer from address(0)) over limit should return false
        assertFalse(compliance.canTransfer(address(0), user1, 901));
    }

    function test_CanTransferWithFailureReason_RegularTransfer() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);
        token.mint(user1, 100);

        // Regular transfer between users
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(user1, user2, 50);
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithFailureReason_Mint_UnderLimit() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);
        token.mint(user1, 100);

        // Mint under limit
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(address(0), user1, 900);
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithFailureReason_Mint_OverLimit() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);
        token.mint(user1, 100);

        // Mint over limit
        (bool success, string memory reason) = compliance
            .canTransferWithFailureReason(address(0), user1, 901);
        assertFalse(success);
        assertEq(
            reason,
            string.concat(
                "Supply limit exceeded: requested=",
                Strings.toString(1001),
                ", max=",
                Strings.toString(1000)
            )
        );
    }

    function test_CurrentSupply() public {
        vm.prank(admin);
        compliance.setMaxSupply(1000);

        assertEq(compliance.currentSupply(), 0);
        token.mint(user1, 100);
        assertEq(compliance.currentSupply(), 100);
        token.mint(user2, 200);
        assertEq(compliance.currentSupply(), 300);
    }
}
