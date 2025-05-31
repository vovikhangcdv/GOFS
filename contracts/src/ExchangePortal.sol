// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./interfaces/IExchange.sol";

contract ExchangePortal is IExchangePortal, AccessControl {
    using SafeERC20 for IERC20;

    error ZeroAddress();
    error TokensMustBeDifferent();
    error InvalidInitialRate();
    error InvalidTokenPair();
    error InvalidAmount();
    error ExcessiveSlippage();
    error InvalidFeeConfig();
    error FeeTooHigh();

    bytes32 public constant EXCHANGE_RATE_ADMIN_ROLE =
        keccak256("EXCHANGE_RATE_ADMIN_ROLE");
    bytes32 public constant EXCHANGE_FEE_ADMIN_ROLE =
        keccak256("EXCHANGE_FEE_ADMIN_ROLE");

    uint256 public constant MAX_FEE = 100_00; // 10% max fee (scaled by 1e4)
    uint256 public constant FEE_DENOMINATOR = 10000;

    address public immutable override token0;
    address public immutable override token1;
    uint256 private _exchangeRate;

    uint256 public exchangeFee; // Fee in basis points (1/100 of 1%)
    address public treasury; // Address where fees are collected

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

    constructor(
        address _token0,
        address _token1,
        uint256 initialRate,
        address _treasury,
        uint256 _exchangeFee
    ) {
        if (_token0 == address(0)) revert ZeroAddress();
        if (_token1 == address(0)) revert ZeroAddress();
        if (_treasury == address(0)) revert ZeroAddress();
        if (_token0 == _token1) revert TokensMustBeDifferent();
        if (initialRate == 0) revert InvalidInitialRate();
        if (_exchangeFee > MAX_FEE) revert FeeTooHigh();

        token0 = _token0;
        token1 = _token1;
        _exchangeRate = initialRate;
        treasury = _treasury;
        exchangeFee = _exchangeFee;

        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(EXCHANGE_RATE_ADMIN_ROLE, msg.sender);
        _grantRole(EXCHANGE_FEE_ADMIN_ROLE, msg.sender);

        emit ExchangeRateUpdated(_token0, _token1, initialRate);
        emit FeeUpdated(_exchangeFee);
        emit TreasuryUpdated(_treasury);
    }

    function setExchangeFee(
        uint256 newFee
    ) external onlyRole(EXCHANGE_FEE_ADMIN_ROLE) {
        if (newFee > MAX_FEE) revert FeeTooHigh();
        exchangeFee = newFee;
        emit FeeUpdated(newFee);
    }

    function setTreasury(
        address newTreasury
    ) external onlyRole(EXCHANGE_FEE_ADMIN_ROLE) {
        if (newTreasury == address(0)) revert ZeroAddress();
        treasury = newTreasury;
        emit TreasuryUpdated(newTreasury);
    }

    function getExchangeRate() external view override returns (uint256) {
        return _exchangeRate;
    }

    function setExchangeRate(
        uint256 newRate
    ) external override onlyRole(EXCHANGE_RATE_ADMIN_ROLE) {
        if (newRate == 0) revert InvalidInitialRate();
        _exchangeRate = newRate;
        emit ExchangeRateUpdated(token0, token1, newRate);
    }

    /**
     * @inheritdoc IExchangePortal
     */
    function exchange(
        address fromToken,
        address toToken,
        uint256 amountIn,
        uint256 minAmountOut
    ) external override returns (uint256 amountOut) {
        if (
            !((fromToken == token0 && toToken == token1) ||
                (fromToken == token1 && toToken == token0))
        ) {
            revert InvalidTokenPair();
        }

        // Transfer tokens from sender to this contract
        IERC20(fromToken).safeTransferFrom(msg.sender, address(this), amountIn);

        // Calculate the output amount including fees
        amountOut = getExchangeAmount(fromToken, toToken, amountIn);
        if (amountOut == 0) revert InvalidAmount();

        // Calculate and deduct fee
        uint256 feeAmount = (amountOut * exchangeFee) / FEE_DENOMINATOR;
        uint256 finalAmount = amountOut - feeAmount;

        // Check slippage
        if (finalAmount < minAmountOut) revert ExcessiveSlippage();

        // Transfer fee to treasury if applicable
        if (feeAmount > 0) {
            IERC20(toToken).safeTransfer(treasury, feeAmount);
        }

        // Transfer remaining tokens to sender
        IERC20(toToken).safeTransfer(msg.sender, finalAmount);

        emit ExchangeExecuted(
            fromToken,
            toToken,
            msg.sender,
            amountIn,
            finalAmount,
            feeAmount
        );

        return finalAmount;
    }

    /**
     * @inheritdoc IExchangePortal
     */
    function getExchangeAmount(
        address fromToken,
        address toToken,
        uint256 amountIn
    ) public view override returns (uint256) {
        uint256 baseAmount;
        if (fromToken == token0 && toToken == token1) {
            baseAmount = (amountIn * _exchangeRate) / 1e18;
        } else if (fromToken == token1 && toToken == token0) {
            baseAmount = (amountIn * 1e18) / _exchangeRate;
        } else {
            revert InvalidTokenPair();
        }

        return baseAmount;
    }
}
