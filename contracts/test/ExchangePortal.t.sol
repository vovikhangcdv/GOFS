// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {ExchangePortal} from "../src/ExchangePortal.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {Entity, EntityType} from "../src/interfaces/ITypes.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
contract MockToken is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract MockRegistry {
    function register(
        Entity calldata entity,
        bytes memory verifierSignature
    ) external {
        // Mock implementation - just return without doing anything
    }
}

contract ExchangePortalTest is Test {
    ExchangePortal public portal;
    MockToken public token0;
    MockToken public token1;
    address public treasury;
    address public admin;
    address public user;

    uint256 public constant INITIAL_RATE = 2 * 1e18; // 1 token0 = 2 token1
    uint256 public constant INITIAL_FEE = 100; // 1% fee (100 basis points)

    event ExchangeRateUpdated(
        address indexed token0,
        address indexed token1,
        uint256 newRate
    );
    event FeeUpdated(uint256 newFee);
    event TreasuryUpdated(address newTreasury);
    event ExchangeExecuted(
        address indexed fromToken,
        address indexed toToken,
        address indexed user,
        uint256 amountIn,
        uint256 amountOut,
        uint256 feeAmount
    );

    function setUp() public {
        admin = makeAddr("admin");
        user = makeAddr("user");
        treasury = makeAddr("treasury");

        vm.startPrank(admin);

        // Deploy tokens
        token0 = new MockToken("Token0", "TK0");
        token1 = new MockToken("Token1", "TK1");

        // Deploy portal
        // portal = new ExchangePortal();
        portal = ExchangePortal(
            address(
                new TransparentUpgradeableProxy(
                    address(new ExchangePortal()),
                    address(new ProxyAdmin(admin)),
                    abi.encodeWithSelector(
                        ExchangePortal.initialize.selector,
                        address(token0),
                        address(token1),
                        INITIAL_RATE,
                        treasury,
                        INITIAL_FEE
                    )
                )
            )
        );

        // Fund user with tokens
        token0.transfer(user, 1000 * 1e18);
        token1.transfer(user, 2000 * 1e18);

        // Fund portal with tokens for liquidity
        token0.transfer(address(portal), 10000 * 1e18);
        token1.transfer(address(portal), 20000 * 1e18);

        vm.stopPrank();
    }

    function test_InitialState() public {
        assertEq(portal.token0(), address(token0));
        assertEq(portal.token1(), address(token1));
        assertEq(portal.getExchangeRate(), INITIAL_RATE);
        assertEq(portal.exchangeFee(), INITIAL_FEE);
        assertEq(portal.treasury(), treasury);

        // Check roles
        assertTrue(portal.hasRole(portal.DEFAULT_ADMIN_ROLE(), admin));
        assertTrue(portal.hasRole(portal.EXCHANGE_FEE_ADMIN_ROLE(), admin));
        assertTrue(portal.hasRole(portal.EXCHANGE_RATE_ADMIN_ROLE(), admin));
    }

    function test_Exchange_Token0ToToken1() public {
        uint256 amountIn = 100 * 1e18;
        uint256 expectedBaseAmount = (amountIn * INITIAL_RATE) / 1e18; // 200 token1
        uint256 expectedFee = (expectedBaseAmount * INITIAL_FEE) /
            portal.FEE_DENOMINATOR(); // 2 token1
        uint256 expectedAmountOut = expectedBaseAmount - expectedFee; // 198 token1

        vm.startPrank(user);
        token0.approve(address(portal), amountIn);

        uint256 amountOut = portal.exchange(
            address(token0),
            address(token1),
            amountIn,
            expectedAmountOut // Set minAmountOut to expected amount
        );

        vm.stopPrank();

        // Verify balances
        assertEq(token0.balanceOf(user), 900 * 1e18);
        assertEq(token1.balanceOf(user), 2198 * 1e18);
        assertEq(token1.balanceOf(treasury), expectedFee);
        assertEq(amountOut, expectedAmountOut);
    }

    function test_Exchange_Token1ToToken0() public {
        uint256 amountIn = 200 * 1e18;
        uint256 expectedBaseAmount = (amountIn * 1e18) / INITIAL_RATE; // 100 token0
        uint256 expectedFee = (expectedBaseAmount * INITIAL_FEE) /
            portal.FEE_DENOMINATOR(); // 1 token0
        uint256 expectedAmountOut = expectedBaseAmount - expectedFee; // 99 token0

        vm.startPrank(user);
        token1.approve(address(portal), amountIn);

        uint256 amountOut = portal.exchange(
            address(token1),
            address(token0),
            amountIn,
            expectedAmountOut
        );

        vm.stopPrank();

        assertEq(token1.balanceOf(user), 1800 * 1e18);
        assertEq(token0.balanceOf(user), 1099 * 1e18);
        assertEq(token0.balanceOf(treasury), expectedFee);
        assertEq(amountOut, expectedAmountOut);
    }

    function test_Exchange_RevertOnSlippage() public {
        uint256 amountIn = 100 * 1e18;
        uint256 expectedBaseAmount = (amountIn * INITIAL_RATE) / 1e18;
        uint256 expectedFee = (expectedBaseAmount * INITIAL_FEE) /
            portal.FEE_DENOMINATOR();
        uint256 expectedAmountOut = expectedBaseAmount - expectedFee;

        vm.startPrank(user);
        token0.approve(address(portal), amountIn);

        // Set minAmountOut higher than actual output
        uint256 minAmountOut = expectedAmountOut + 1;

        vm.expectRevert(ExchangePortal.ExcessiveSlippage.selector);
        portal.exchange(
            address(token0),
            address(token1),
            amountIn,
            minAmountOut
        );

        vm.stopPrank();
    }

    function test_UpdateExchangeRate() public {
        uint256 newRate = 3 * 1e18;

        vm.startPrank(admin);

        vm.expectEmit(true, true, true, true);
        emit ExchangeRateUpdated(address(token0), address(token1), newRate);

        portal.setExchangeRate(newRate);
        vm.stopPrank();

        assertEq(portal.getExchangeRate(), newRate);
    }

    function test_UpdateFee() public {
        uint256 newFee = 200; // 2%

        vm.startPrank(admin);

        vm.expectEmit(true, true, true, true);
        emit FeeUpdated(newFee);

        portal.setExchangeFee(newFee);
        vm.stopPrank();

        assertEq(portal.exchangeFee(), newFee);
    }

    function test_UpdateTreasury() public {
        address newTreasury = makeAddr("newTreasury");

        vm.startPrank(admin);

        vm.expectEmit(true, true, true, true);
        emit TreasuryUpdated(newTreasury);

        portal.setTreasury(newTreasury);
        vm.stopPrank();

        assertEq(portal.treasury(), newTreasury);
    }

    function test_RevertOnUnauthorizedRateUpdate() public {
        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                user,
                portal.EXCHANGE_RATE_ADMIN_ROLE()
            )
        );
        portal.setExchangeRate(3 * 1e18);
        vm.stopPrank();
    }

    function test_RevertOnUnauthorizedFeeUpdate() public {
        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                user,
                portal.EXCHANGE_FEE_ADMIN_ROLE()
            )
        );
        portal.setExchangeFee(200);
        vm.stopPrank();
    }

    function test_RevertOnInvalidTokenPair() public {
        address invalidToken = makeAddr("invalidToken");

        vm.startPrank(user);
        vm.expectRevert(ExchangePortal.InvalidTokenPair.selector);
        portal.exchange(invalidToken, address(token1), 100 * 1e18, 90 * 1e18);
        vm.stopPrank();
    }

    function test_RevertOnFeeTooHigh() public {
        vm.startPrank(admin);
        vm.expectRevert(ExchangePortal.FeeTooHigh.selector);
        portal.setExchangeFee(100_01); // Above MAX_FEE (100_00 = 10%)
        vm.stopPrank();
    }

    function test_RevertOnZeroExchangeAmount() public {
        vm.startPrank(user);
        token0.approve(address(portal), 0);
        vm.expectRevert(ExchangePortal.InvalidAmount.selector);
        portal.exchange(address(token0), address(token1), 0, 0);
        vm.stopPrank();
    }

    function test_GetExchangeAmount() public {
        uint256 amountIn = 100 * 1e18;

        // Test token0 to token1
        uint256 amountOut = portal.getExchangeAmount(
            address(token0),
            address(token1),
            amountIn
        );
        assertEq(amountOut, (amountIn * INITIAL_RATE) / 1e18);

        // Test token1 to token0
        amountOut = portal.getExchangeAmount(
            address(token1),
            address(token0),
            amountIn
        );
        assertEq(amountOut, (amountIn * 1e18) / INITIAL_RATE);
    }

    function test_RevertOnZeroTreasury() public {
        vm.startPrank(admin);
        vm.expectRevert(ExchangePortal.ZeroAddress.selector);
        portal.setTreasury(address(0));
        vm.stopPrank();
    }

    function test_RevertOnZeroExchangeRate() public {
        vm.startPrank(admin);
        vm.expectRevert(ExchangePortal.InvalidInitialRate.selector);
        portal.setExchangeRate(0);
        vm.stopPrank();
    }

    function test_GetExchangeAmountInvalidPair() public {
        address invalidToken = makeAddr("invalidToken");
        vm.expectRevert(ExchangePortal.InvalidTokenPair.selector);
        portal.getExchangeAmount(invalidToken, address(token1), 100 * 1e18);
    }

    function test_RegisterWithRegistry() public {
        address mockRegistry = makeAddr("mockRegistry");
        Entity memory entity = Entity({
            entityAddress: address(portal),
            entityType: EntityType.wrap(1),
            entityData: abi.encode("Exchange Portal", "Exchange entity"),
            verifier: makeAddr("verifier")
        });
        bytes memory signature = "mock signature";

        // Deploy a mock registry contract that implements the register function
        MockRegistry mockRegistryContract = new MockRegistry();

        // Admin can authorize registration
        vm.startPrank(admin);
        portal.registerWithRegistry(
            address(mockRegistryContract),
            entity,
            signature
        );
        vm.stopPrank();
    }

    function test_RegisterWithRegistry_Unauthorized() public {
        address mockRegistry = makeAddr("mockRegistry");
        Entity memory entity = Entity({
            entityAddress: address(portal),
            entityType: EntityType.wrap(1),
            entityData: abi.encode("Exchange Portal", "Exchange entity"),
            verifier: makeAddr("verifier")
        });
        bytes memory signature = "mock signature";

        // Non-admin cannot authorize registration
        vm.startPrank(user);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                user,
                portal.REGISTER_ADMIN_ROLE()
            )
        );
        portal.registerWithRegistry(mockRegistry, entity, signature);
        vm.stopPrank();
    }
}
