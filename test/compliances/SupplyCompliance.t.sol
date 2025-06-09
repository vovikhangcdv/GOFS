// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {SupplyCompliance} from "../../src/compliances/SupplyCompliance.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

// Mock ERC20 token for testing
contract MockToken {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
    mapping(address => uint256) public balanceOf;

    constructor(string memory _name, string memory _symbol, uint8 _decimals) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
    }

    function mint(address to, uint256 amount) external {
        totalSupply += amount;
        balanceOf[to] += amount;
    }

    function burn(address from, uint256 amount) external {
        require(balanceOf[from] >= amount, "Insufficient balance");
        totalSupply -= amount;
        balanceOf[from] -= amount;
    }
}

contract SupplyComplianceTest is Test {
    SupplyCompliance public compliance;
    MockToken public token;
    address public admin;
    address public user1;
    address public user2;

    function setUp() public {
        // Set up accounts
        admin = makeAddr("admin");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        vm.startPrank(admin);

        // Deploy token
        token = new MockToken("Test Token", "TEST", 18);

        // Deploy compliance module
        compliance = new SupplyCompliance();
        compliance.initialize(address(token));

        // Setup roles
        compliance.grantRole(compliance.DEFAULT_ADMIN_ROLE(), admin);
        compliance.grantRole(compliance.SUPPLY_ADMIN_ROLE(), admin);

        vm.stopPrank();
    }

    function testConstructor() public {
        // Test valid initialization
        SupplyCompliance newCompliance = new SupplyCompliance();
        newCompliance.initialize(address(token));
        assertEq(address(newCompliance.token()), address(token));
    }

    function testConstructorWithZeroAddress() public {
        SupplyCompliance newCompliance = new SupplyCompliance();
        vm.expectRevert(SupplyCompliance.TokenNotSet.selector);
        newCompliance.initialize(address(0));
    }

    // ... existing code ...
}
