// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title ICompliance
 * @dev Interface for compliance checks in token transfers.
 * This interface defines the method to check if a transfer is allowed
 * based on compliance rules.
 */
interface ICompliance {
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
}
