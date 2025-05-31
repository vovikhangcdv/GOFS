// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IExternalExchangePortal
 * @notice Interface for the External Exchange Portal that facilitates currency conversions
 * between CBDC-Token and other currencies/stablecoins
 */
interface IExchangePortal {
    /**
     * @notice Error thrown when exchange rate is invalid
     */
    error InvalidExchangeRate();

    /**
     * @notice Emitted when exchange rate is updated
     * @param token0 The address of the first token (usually CBDC)
     * @param token1 The address of the second token
     * @param newRate The new exchange rate (scaled by 1e18)
     */
    event ExchangeRateUpdated(
        address indexed token0,
        address indexed token1,
        uint256 newRate
    );

    /**
     * @notice Returns the address of the first token in the pair
     * @return Address of token0
     */
    function token0() external view returns (address);

    /**
     * @notice Returns the address of the second token in the pair
     * @return Address of token1
     */
    function token1() external view returns (address);

    /**
     * @notice Get the current exchange rate between token0 and token1
     * @return Current exchange rate scaled by 1e18 (1 token0 = rate * token1)
     */
    function getExchangeRate() external view returns (uint256);

    /**
     * @notice Set the exchange rate between token0 and token1
     * @param newRate New exchange rate scaled by 1e18
     * @dev Only callable by authorized monetary policy governors
     */
    function setExchangeRate(uint256 newRate) external;

    /**
     * @notice Exchange token0 for token1 or vice versa with slippage protection
     * @param fromToken Address of the token to exchange from
     * @param toToken Address of the token to exchange to
     * @param amountIn Amount of fromToken to exchange
     * @param minAmountOut Minimum amount of toToken that must be received (slippage protection)
     * @return amountOut Amount of toToken received after fees
     */
    function exchange(
        address fromToken,
        address toToken,
        uint256 amountIn,
        uint256 minAmountOut
    ) external returns (uint256 amountOut);

    /**
     * @notice Calculate the output amount for an exchange before fees
     * @param fromToken Address of the token to exchange from
     * @param toToken Address of the token to exchange to
     * @param amountIn Amount of fromToken to exchange
     * @return amountOut Amount of toToken that would be received before fees
     */
    function getExchangeAmount(
        address fromToken,
        address toToken,
        uint256 amountIn
    ) external view returns (uint256 amountOut);
}
