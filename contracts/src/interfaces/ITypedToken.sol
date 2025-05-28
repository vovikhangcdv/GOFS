/**
 * @title ITypedERC20
 * @dev Extension of ERC20 interface that adds transaction type support
 * @notice This interface extends the standard ERC20 interface by adding transaction type
 * functionality to transfers and approvals
 * @custom:txtype For standard ERC20 operations (transfer, transferFrom, approve), use TxType(0)
 */
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {TxType} from "./ITypes.sol";
interface ITypedERC20 is IERC20 {
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
     * @dev Function to transfer tokens from one address to another with transaction type
     * @param from The sender address
     * @param to The recipient address
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     */
    function transferFromWithType(
        address from,
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
