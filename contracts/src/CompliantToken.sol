// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ICompliantToken} from "./interfaces/ICompliantToken.sol";
import {ICompliance} from "./interfaces/ICompliance.sol";

/**
 * @title CompliantToken
 * @dev Implementation of a compliant ERC20 token that enforces regulatory compliance.
 * This token ensures that:
 * 1. All transfers must pass compliance checks (including entity verification)
 * 2. Token parameters can be managed by governance
 */
contract CompliantToken is ICompliantToken, ERC20, Ownable {
    // Immutable compliance reference
    ICompliance public immutable override compliance;

    // Custom errors
    error ComplianceCheckFailed(
        address from,
        address to,
        uint256 amount,
        string reason
    );
    error ZeroAddressCompliance();

    /**
     * @dev Constructor to initialize the token with basic parameters and required contract reference
     * @param name_ The name of the token
     * @param symbol_ The symbol of the token
     * @param compliance_ The address of the compliance contract
     */
    constructor(
        string memory name_,
        string memory symbol_,
        address compliance_
    ) ERC20(name_, symbol_) Ownable(msg.sender) {
        if (compliance_ == address(0)) revert ZeroAddressCompliance();
        compliance = ICompliance(compliance_);
    }

    /**
     * @dev Override of _update to enforce compliance checks
     * @param from The sender address
     * @param to The recipient address
     * @param value The amount of tokens to transfer
     */
    function _update(
        address from,
        address to,
        uint256 value
    ) internal virtual override {
        // Skip checks for minting and burning
        if (from != address(0) && to != address(0)) {
            // Check compliance with failure reason (includes entity verification)
            (bool isCompliant, string memory failureReason) = compliance
                .canTransferWithFailureReason(from, to, value);
            if (!isCompliant) {
                revert ComplianceCheckFailed(from, to, value, failureReason);
            }
        }

        super._update(from, to, value);
    }

    /**
     * @dev Mint new tokens. Only callable by the owner (governance)
     * @param to The address to mint tokens to
     * @param amount The amount of tokens to mint
     */
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    /**
     * @dev Burn tokens from an address. Only callable by the owner (governance)
     * @param from The address to burn tokens from
     * @param amount The amount of tokens to burn
     */
    function burn(address from, uint256 amount) external onlyOwner {
        _burn(from, amount);
    }
}
