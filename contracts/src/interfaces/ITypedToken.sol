// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {TxType} from "./ITypes.sol";
interface ITypedERC20 is IERC20 {
    /**
     * @notice Transfers the specified amount of tokens to the given address
     * @dev This is a standard ERC20 transfer function that defaults to TxType(0)
     * @param to The address to transfer tokens to
     * @param amount The amount of tokens to transfer
     * @return bool Returns true if the transfer was successful
     */
    function transfer(
        address to,
        uint256 amount
    ) external virtual override returns (bool) {
        return transferWithType(to, amount, TxType(0));
    }

    /**
     * @notice Approves the specified amount of tokens for the given spender
     * @dev This is a standard ERC20 approve function that defaults to TxType(0)
     * @param spender The address to approve tokens for
     * @param amount The amount of tokens to approve
     * @return bool Returns true if the approval was successful
     */
    function approve(
        address spender,
        uint256 amount
    ) external virtual override returns (bool) {
        return approveWithType(spender, amount, TxType(0));
    }

    /**
     * @dev Function to transfer tokens with transaction type
     * @param to The recipient address
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     */
    function transferWithType(
        address to,
        uint256 amount,
        TxType txType
    ) external returns (bool);

    /**
     * @dev Function for approving spender with transaction type
     * @param spender Address approved to spend tokens
     * @param amount Amount of tokens approved
     * @param txType Type of transaction
     */
    function approveWithType(
        address spender,
        uint256 amount,
        TxType txType
    ) external returns (bool);
}
