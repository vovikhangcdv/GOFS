// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ICompliance} from "./ICompliance.sol";
import {IEntityRegistry} from "./IEntityRegistry.sol";

/**
 * @title ICompliantToken
 * @dev Interface for a compliant ERC20 token that enforces regulatory compliance and entity verification.
 * Any implementation of transfer-based functions (transfer, transferFrom) must:
 * 1. Check compliance rules via ICompliance.canTransfer()
 * 2. Verify entities via IEntityRegistry.isVerifiedEntity()
 */
interface ICompliantToken is IERC20 {
    function entityRegistry() external view returns (IEntityRegistry);
    function compliance() external view returns (ICompliance);
}
