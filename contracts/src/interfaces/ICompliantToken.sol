// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ICompliance} from "./ICompliance.sol";
import {IEntityRegistry} from "./IEntityRegistry.sol";

/**
 * @title ICompliantToken
 * @dev Interface for the compliant token that implements
 * ERC20 standard with additional policy controls.
 */
interface ICompliantToken is IERC20 {
    // compliance should be
    function compliance() external view returns (ICompliance);
    function entityRegistry() external view returns (IEntityRegistry);
}
