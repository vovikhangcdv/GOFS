// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {TxType} from "./ITypes.sol";

/**
 * @title ICompliance
 * @dev Interface for compliance checks in token transfers.
 * This interface defines the method to check if a transfer is allowed
 * based on compliance rules.
 */
interface ICompliance is IERC165 {
    function canTransfer(
        address _from,
        address _to,
        uint256 _amount
    ) external view returns (bool);

    /// @notice Checks if a transfer can be made with a reason for failure
    /// @dev Returns a boolean indicating if the transfer is allowed,
    /// and a string with the reason if it is not allowed.
    function canTransferWithFailureReason(
        address _from,
        address _to,
        uint256 _amount
    ) external view returns (bool, string memory);

    /// @notice Checks if a typed transfer can be made
    /// @dev Returns a boolean indicating if the transfer is allowed
    function canTransferWithType(
        address _from,
        address _to,
        uint256 _amount,
        TxType _txType
    ) external view returns (bool);

    /// @notice Checks if a typed transfer can be made with a reason for failure
    /// @dev Returns a boolean indicating if the transfer is allowed,
    /// and a string with the reason if it is not allowed.
    function canTransferWithTypeAndFailureReason(
        address _from,
        address _to,
        uint256 _amount,
        TxType _txType
    ) external view returns (bool, string memory);
}
