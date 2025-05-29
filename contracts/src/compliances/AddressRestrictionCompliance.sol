// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../interfaces/ICompliance.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

/**
 * @title AddressRestrictionCompliance
 * @dev Compliance module that implements address blacklisting functionality
 */
contract AddressRestrictionCompliance is ICompliance, AccessControl {
    // Role for managing blacklist
    bytes32 public constant BLACKLIST_ADMIN_ROLE =
        keccak256("BLACKLIST_ADMIN_ROLE");

    // Mappings to track blacklisted addresses
    mapping(address => bool) private _blacklistedFrom; // blacklisted from sending
    mapping(address => bool) private _blacklistedTo; // blacklisted from receiving

    // Events
    event AddressBlacklistedFrom(address indexed account);
    event AddressUnblacklistedFrom(address indexed account);
    event AddressBlacklistedTo(address indexed account);
    event AddressUnblacklistedTo(address indexed account);
    event AddressBlacklisted(address indexed account);
    event AddressUnblacklisted(address indexed account);

    // Custom errors
    error AddressAlreadyBlacklistedFrom(address account);
    error AddressNotBlacklistedFrom(address account);
    error AddressAlreadyBlacklistedTo(address account);
    error AddressNotBlacklistedTo(address account);
    error ZeroAddress();
    error EmptyAddressList();

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(BLACKLIST_ADMIN_ROLE, msg.sender);
    }

    /**
     * @dev Add addresses to the blacklist for sending tokens
     * @param accounts The addresses to blacklist
     */
    function blacklistFrom(
        address[] memory accounts
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();
            if (_blacklistedFrom[account])
                revert AddressAlreadyBlacklistedFrom(account);

            _blacklistedFrom[account] = true;
            emit AddressBlacklistedFrom(account);
        }
    }

    /**
     * @dev Add addresses to the blacklist for receiving tokens
     * @param accounts The addresses to blacklist
     */
    function blacklistTo(
        address[] memory accounts
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();
            if (_blacklistedTo[account])
                revert AddressAlreadyBlacklistedTo(account);

            _blacklistedTo[account] = true;
            emit AddressBlacklistedTo(account);
        }
    }

    /**
     * @dev Add addresses to both sending and receiving blacklists
     * @param accounts The addresses to blacklist
     */
    function blacklist(
        address[] memory accounts
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();

            if (!_blacklistedFrom[account]) {
                _blacklistedFrom[account] = true;
                emit AddressBlacklistedFrom(account);
            }
            if (!_blacklistedTo[account]) {
                _blacklistedTo[account] = true;
                emit AddressBlacklistedTo(account);
            }
            emit AddressBlacklisted(account);
        }
    }

    /**
     * @dev Remove addresses from the sending blacklist
     * @param accounts The addresses to remove from blacklist
     */
    function unblacklistFrom(
        address[] memory accounts
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();
            if (!_blacklistedFrom[account])
                revert AddressNotBlacklistedFrom(account);

            _blacklistedFrom[account] = false;
            emit AddressUnblacklistedFrom(account);
        }
    }

    /**
     * @dev Remove addresses from the receiving blacklist
     * @param accounts The addresses to remove from blacklist
     */
    function unblacklistTo(
        address[] memory accounts
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();
            if (!_blacklistedTo[account])
                revert AddressNotBlacklistedTo(account);

            _blacklistedTo[account] = false;
            emit AddressUnblacklistedTo(account);
        }
    }

    /**
     * @dev Remove addresses from both blacklists
     * @param accounts The addresses to remove from blacklist
     */
    function unblacklist(
        address[] memory accounts
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();

            if (_blacklistedFrom[account]) {
                _blacklistedFrom[account] = false;
                emit AddressUnblacklistedFrom(account);
            }
            if (_blacklistedTo[account]) {
                _blacklistedTo[account] = false;
                emit AddressUnblacklistedTo(account);
            }
            emit AddressUnblacklisted(account);
        }
    }

    /**
     * @dev Check if an address is blacklisted from sending
     * @param account The address to check
     * @return bool True if the address is blacklisted from sending
     */
    function isBlacklistedFrom(address account) public view returns (bool) {
        return _blacklistedFrom[account];
    }

    /**
     * @dev Check if an address is blacklisted from receiving
     * @param account The address to check
     * @return bool True if the address is blacklisted from receiving
     */
    function isBlacklistedTo(address account) public view returns (bool) {
        return _blacklistedTo[account];
    }

    /**
     * @dev Check if an address is completely blacklisted
     * @param account The address to check
     * @return bool True if the address is blacklisted from both sending and receiving
     */
    function isBlacklisted(address account) public view returns (bool) {
        return _blacklistedFrom[account] && _blacklistedTo[account];
    }

    /**
     * @dev Compliance check for transfers with failure reason
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @return (bool, string) Tuple of (success, failure reason)
     */
    function canTransferWithFailureReason(
        address from,
        address to,
        uint256
    ) external view override returns (bool, string memory) {
        if (_blacklistedFrom[from]) {
            return (false, "Sender address is blacklisted from sending");
        }
        if (_blacklistedTo[to]) {
            return (false, "Recipient address is blacklisted from receiving");
        }
        return (true, "");
    }

    /**
     * @dev Compliance check for transfers
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @return bool True if the transfer is allowed
     */
    function canTransfer(
        address from,
        address to,
        uint256 amount
    ) external view override returns (bool) {
        (bool success, ) = this.canTransferWithFailureReason(from, to, amount);
        return success;
    }
}
