// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ICompliantToken} from "./interfaces/ICompliantToken.sol";
import {ICompliance} from "./interfaces/ICompliance.sol";
import {TypedToken} from "./TypedToken.sol";
import {TxType} from "./interfaces/ITypes.sol";

/**
 * @title CompliantToken
 * @dev Implementation of a compliant ERC20 token that enforces regulatory compliance.
 * This token ensures that:
 * 1. All transfers must pass compliance checks (including entity verification)
 * 2. Token parameters can be managed by governance
 */
contract CompliantToken is
    ICompliantToken,
    TypedToken,
    AccessControlUpgradeable
{
    // Role definitions
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant BURNER_ROLE = keccak256("BURNER_ROLE");
    bytes32 public constant COMPLIANCE_ADMIN_ROLE =
        keccak256("COMPLIANCE_ADMIN_ROLE");

    // Compliance reference (changed from immutable to allow upgrades)
    ICompliance public override compliance;

    // Custom errors
    error ComplianceCheckFailed(
        address from,
        address to,
        uint256 amount,
        string reason
    );
    error ZeroAddressCompliance();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev Initialize the token with basic parameters and required contract reference
     * @param name_ The name of the token
     * @param symbol_ The symbol of the token
     * @param compliance_ The address of the compliance contract
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        address compliance_
    ) public initializer {
        if (compliance_ == address(0)) revert ZeroAddressCompliance();

        __TypedToken_init(name_, symbol_);
        __AccessControl_init();

        compliance = ICompliance(compliance_);

        // Setup roles
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
        _grantRole(BURNER_ROLE, msg.sender);
        _grantRole(COMPLIANCE_ADMIN_ROLE, msg.sender);
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
        // Check compliance with failure reason (includes entity verification)
        (bool isCompliant, string memory failureReason) = compliance
            .canTransferWithFailureReason(from, to, value);
        if (!isCompliant) {
            revert ComplianceCheckFailed(from, to, value, failureReason);
        }

        super._update(from, to, value);
    }

    /**
     * @dev Mint new tokens. Only callable by accounts with MINTER_ROLE
     * @param to The address to mint tokens to
     * @param amount The amount of tokens to mint
     */
    function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        _mint(to, amount);
    }

    /**
     * @dev Burn tokens from an address. Only callable by accounts with BURNER_ROLE
     * @param from The address to burn tokens from
     * @param amount The amount of tokens to burn
     */
    function burn(address from, uint256 amount) external onlyRole(BURNER_ROLE) {
        _burn(from, amount);
    }

    /**
     * @dev Change the compliance contract. Only callable by accounts with COMPLIANCE_ADMIN_ROLE
     * @param newCompliance The address of the new compliance contract
     */
    function setCompliance(
        address newCompliance
    ) external onlyRole(COMPLIANCE_ADMIN_ROLE) {
        if (newCompliance == address(0)) revert ZeroAddressCompliance();
        compliance = ICompliance(newCompliance);
    }

    /**
     * @dev Override of _transferWithType to enforce compliance checks for typed transfers
     * @param from The sender address
     * @param to The recipient address
     * @param amount The amount of tokens to transfer
     * @param txType The type of transaction
     */
    function _transferWithType(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) internal virtual override {
        // Check compliance with failure reason and transaction type
        (bool isCompliant, string memory failureReason) = compliance
            .canTransferWithTypeAndFailureReason(from, to, amount, txType);
        if (!isCompliant) {
            revert ComplianceCheckFailed(from, to, amount, failureReason);
        }

        super._transferWithType(from, to, amount, txType);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
