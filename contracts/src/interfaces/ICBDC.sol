// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {IMonetaryPolicyRegistry} from "./IMonetaryPolicyRegistry.sol";

/**
 * @title ICBDC
 * @dev Interface for the  Bank Digital Currency token
 * that implements ERC20 standard with additional policy controls.
 */
interface ICBDC is IERC20, IERC20Metadata, IMonetaryPolicyRegistry {
    /**
     * @dev Events specific to CBDC operations
     */
    event TransferWithType(
        address indexed from,
        address indexed to,
        uint256 amount,
        TxType txType
    );
    event ApprovalWithType(
        address indexed owner,
        address indexed spender,
        uint256 amount,
        TxType txType
    );

    /**
     * @dev Enhanced transfer function with transaction type
     * @param to The recipient address
     * @param amount The amount to transfer
     * @param txType The type of transaction
     * @return success True if the transfer was successful
     */
    function transfer(
        address to,
        uint256 amount,
        TxType txType
    ) external returns (bool);

    /**
     * @dev Enhanced transferFrom function with transaction type
     * @param from The sender address
     * @param to The recipient address
     * @param amount The amount to transfer
     * @param txType The type of transaction
     * @return success True if the transfer was successful
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external returns (bool);

    /**
     * @dev Enhanced approve function with transaction type
     * @param spender The address to approve
     * @param amount The amount to approve
     * @param txType The type of transaction
     * @return success True if the approval was successful
     */
    function approve(
        address spender,
        uint256 amount,
        TxType txType
    ) external returns (bool);

    /**
     * @dev Enhanced allowance function with transaction type
     * @param owner The address of the token owner
     * @param spender The address of the spender
     * @param txType The type of transaction
     * @return remaining The remaining allowance
     */
    function allowance(
        address owner,
        address spender,
        TxType txType
    ) external view returns (uint256);
}
