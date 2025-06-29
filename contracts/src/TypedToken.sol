// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ITypedERC20} from "./interfaces/ITypedToken.sol";
import {TxType} from "./interfaces/ITypes.sol";

/**
 * @title TypedToken
 * @dev Extension of ERC20 that supports transaction types for transfers
 */
abstract contract TypedToken is ERC20Upgradeable, ITypedERC20 {
    // Events for typed operations
    event TypedTransfer(
        address indexed from,
        address indexed to,
        uint256 value,
        TxType indexed txType
    );

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev Initializer that sets up the ERC20 token
     * @param name_ The name of the token
     * @param symbol_ The symbol of the token
     */
    function __TypedToken_init(
        string memory name_,
        string memory symbol_
    ) internal onlyInitializing {
        __ERC20_init(name_, symbol_);
    }

    function __TypedToken_init_unchained(
        string memory name_,
        string memory symbol_
    ) internal onlyInitializing {
        // This function intentionally left empty
        // All initialization is done in __TypedToken_init
    }

    /**
     * @dev Transfer tokens with transaction type
     * @param to The recipient address
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     */
    function transferWithType(
        address to,
        uint256 amount,
        TxType txType
    ) external virtual override returns (bool) {
        address owner = _msgSender();
        _transferWithType(owner, to, amount, txType);
        return true;
    }

    /**
     * @dev Transfer tokens from one address to another with transaction type
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
    ) external virtual override returns (bool) {
        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        _transferWithType(from, to, amount, txType);
        return true;
    }

    /**
     * @dev Internal function to handle typed transfers
     * Can be overridden by derived contracts for additional logic
     * @param from The sender address
     * @param to The recipient address
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     */
    function _transferWithType(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) internal virtual {
        _transfer(from, to, amount);
        emit TypedTransfer(from, to, amount, txType);
    }

    /**
     * @dev Override the standard transfer to emit typed event with TxType(0)
     */
    function transfer(
        address to,
        uint256 amount
    ) public virtual override(ERC20Upgradeable, IERC20) returns (bool) {
        address owner = _msgSender();
        _transferWithType(owner, to, amount, TxType.wrap(0));
        return true;
    }

    /**
     * @dev Override the standard transferFrom to emit typed event with TxType(0)
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public virtual override(ERC20Upgradeable, IERC20) returns (bool) {
        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        _transferWithType(from, to, amount, TxType.wrap(0));
        return true;
    }

    /**
     * @dev Override the standard approve (no typed functionality for now)
     */
    function approve(
        address spender,
        uint256 amount
    ) public virtual override(ERC20Upgradeable, IERC20) returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, amount);
        return true;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[50] private __gap;
}
