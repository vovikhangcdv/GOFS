// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./interfaces/IExchange.sol";

/**
 * @title ExchangePortal
 * @dev Implementation of fixed rate IExchangePortal with role-based access control
 */
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

    bytes32 public constant RATE_ADMIN_ROLE = keccak256("RATE_ADMIN_ROLE");
    bytes32 public constant FEE_ADMIN_ROLE = keccak256("FEE_ADMIN_ROLE");

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

    /**
     * @dev Initializes the exchange portal with token addresses and sets up admin roles
     * @param _token0 Address of the first token (usually CBDC)
     * @param _token1 Address of the second token
     * @param initialRate Initial exchange rate scaled by 1e18
     * @param _treasury Address where fees will be collected
     * @param _exchangeFee Initial fee in basis points
     */
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
        _grantRole(RATE_ADMIN_ROLE, msg.sender);
        _grantRole(FEE_ADMIN_ROLE, msg.sender);

        emit ExchangeRateUpdated(_token0, _token1, initialRate);
        emit FeeUpdated(_exchangeFee);
        emit TreasuryUpdated(_treasury);
    }

    /**
     * @dev Updates the exchange fee
     * @param newFee Fee in basis points (1/100 of 1%)
     */
    function setExchangeFee(uint256 newFee) external onlyRole(FEE_ADMIN_ROLE) {
        if (newFee > MAX_FEE) revert FeeTooHigh();
        exchangeFee = newFee;
        emit FeeUpdated(newFee);
    }

    /**
     * @dev Updates the treasury address
     * @param newTreasury New treasury address
     */
    function setTreasury(
        address newTreasury
    ) external onlyRole(FEE_ADMIN_ROLE) {
        if (newTreasury == address(0)) revert ZeroAddress();
        treasury = newTreasury;
        emit TreasuryUpdated(newTreasury);
    }

    /**
     * @inheritdoc IExchangePortal
     */
    function getExchangeRate() external view override returns (uint256) {
        return _exchangeRate;
    }

    /**
     * @inheritdoc IExchangePortal
     */
    function setExchangeRate(
        uint256 newRate
    ) external override onlyRole(RATE_ADMIN_ROLE) {
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
