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
}
