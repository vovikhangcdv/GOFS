// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ICompliance} from "./ICompliance.sol";

/**
 * @title ICompliantToken
 * @dev Interface for a compliant ERC20 token that enforces regulatory compliance.
 * Any implementation of transfer-based functions (transfer, transferFrom) must:
 * 1. Pass compliance rules via ICompliance.canTransfer()
 */
interface ICompliantToken is IERC20 {
    function compliance() external view returns (ICompliance);
}
