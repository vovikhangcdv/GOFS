// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../interfaces/ICompliance.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import {TxType} from "../interfaces/ITypes.sol";

/**
 * @title AddressRestrictionCompliance
 * @dev Compliance module that implements address blacklisting functionality with reason tracking
 */
contract AddressRestrictionCompliance is ICompliance, AccessControl {
    // Role for managing blacklist
    bytes32 public constant BLACKLIST_ADMIN_ROLE =
        keccak256("BLACKLIST_ADMIN_ROLE");

    // Struct to store restriction information
    struct RestrictionInfo {
        bool isRestricted;
        string reason;
        uint256 timestamp;
        address restrictedBy;
    }

    // Mappings to track blacklisted addresses with detailed information
    mapping(address => RestrictionInfo) private _restrictedFrom; // restricted from sending
    mapping(address => RestrictionInfo) private _restrictedTo; // restricted from receiving

    // Events with reason tracking
    event AddressBlacklistedFrom(
        address indexed account,
        string reason,
        address indexed restrictedBy,
        uint256 timestamp
    );
    event AddressUnblacklistedFrom(
        address indexed account,
        string reason,
        address indexed unrestrictedBy,
        uint256 timestamp
    );
    event AddressBlacklistedTo(
        address indexed account,
        string reason,
        address indexed restrictedBy,
        uint256 timestamp
    );
    event AddressUnblacklistedTo(
        address indexed account,
        string reason,
        address indexed unrestrictedBy,
        uint256 timestamp
    );
    event AddressBlacklisted(
        address indexed account,
        string reason,
        address indexed restrictedBy,
        uint256 timestamp
    );
    event AddressUnblacklisted(
        address indexed account,
        string reason,
        address indexed unrestrictedBy,
        uint256 timestamp
    );

    // Custom errors
    error AddressAlreadyBlacklistedFrom(address account);
    error AddressNotBlacklistedFrom(address account);
    error AddressAlreadyBlacklistedTo(address account);
    error AddressNotBlacklistedTo(address account);
    error ZeroAddress();
    error EmptyAddressList();
    error EmptyReason();
    error ArrayLengthMismatch();

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(BLACKLIST_ADMIN_ROLE, msg.sender);
    }

    /**
     * @dev Implements IERC165 interface detection
     */
    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(AccessControl, IERC165) returns (bool) {
        return
            interfaceId == type(ICompliance).interfaceId ||
            interfaceId == type(AccessControl).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @dev Add addresses to the blacklist for sending tokens with reasons
     * @param accounts The addresses to blacklist
     * @param reasons The reasons for blacklisting (must match accounts length)
     */
    function blacklistFromWithReasons(
        address[] memory accounts,
        string[] memory reasons
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (accounts.length != reasons.length) revert ArrayLengthMismatch();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            string memory reason = reasons[i];

            if (account == address(0)) revert ZeroAddress();
            if (bytes(reason).length == 0) revert EmptyReason();
            if (_restrictedFrom[account].isRestricted)
                revert AddressAlreadyBlacklistedFrom(account);

            _restrictedFrom[account] = RestrictionInfo({
                isRestricted: true,
                reason: reason,
                timestamp: block.timestamp,
                restrictedBy: msg.sender
            });

            emit AddressBlacklistedFrom(
                account,
                reason,
                msg.sender,
                block.timestamp
            );
        }
    }

    /**
     * @dev Add addresses to the blacklist for sending tokens with a single reason
     * @param accounts The addresses to blacklist
     * @param reason The reason for blacklisting all addresses
     */
    function blacklistFrom(
        address[] memory accounts,
        string memory reason
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (bytes(reason).length == 0) revert EmptyReason();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();
            if (_restrictedFrom[account].isRestricted)
                revert AddressAlreadyBlacklistedFrom(account);

            _restrictedFrom[account] = RestrictionInfo({
                isRestricted: true,
                reason: reason,
                timestamp: block.timestamp,
                restrictedBy: msg.sender
            });

            emit AddressBlacklistedFrom(
                account,
                reason,
                msg.sender,
                block.timestamp
            );
        }
    }

    /**
     * @dev Add addresses to the blacklist for receiving tokens with reasons
     * @param accounts The addresses to blacklist
     * @param reasons The reasons for blacklisting (must match accounts length)
     */
    function blacklistToWithReasons(
        address[] memory accounts,
        string[] memory reasons
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (accounts.length != reasons.length) revert ArrayLengthMismatch();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            string memory reason = reasons[i];

            if (account == address(0)) revert ZeroAddress();
            if (bytes(reason).length == 0) revert EmptyReason();
            if (_restrictedTo[account].isRestricted)
                revert AddressAlreadyBlacklistedTo(account);

            _restrictedTo[account] = RestrictionInfo({
                isRestricted: true,
                reason: reason,
                timestamp: block.timestamp,
                restrictedBy: msg.sender
            });

            emit AddressBlacklistedTo(
                account,
                reason,
                msg.sender,
                block.timestamp
            );
        }
    }

    /**
     * @dev Add addresses to the blacklist for receiving tokens with a single reason
     * @param accounts The addresses to blacklist
     * @param reason The reason for blacklisting all addresses
     */
    function blacklistTo(
        address[] memory accounts,
        string memory reason
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (bytes(reason).length == 0) revert EmptyReason();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();
            if (_restrictedTo[account].isRestricted)
                revert AddressAlreadyBlacklistedTo(account);

            _restrictedTo[account] = RestrictionInfo({
                isRestricted: true,
                reason: reason,
                timestamp: block.timestamp,
                restrictedBy: msg.sender
            });

            emit AddressBlacklistedTo(
                account,
                reason,
                msg.sender,
                block.timestamp
            );
        }
    }

    /**
     * @dev Add addresses to both sending and receiving blacklists with reasons
     * @param accounts The addresses to blacklist
     * @param reasons The reasons for blacklisting (must match accounts length)
     */
    function blacklistWithReasons(
        address[] memory accounts,
        string[] memory reasons
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (accounts.length != reasons.length) revert ArrayLengthMismatch();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            string memory reason = reasons[i];

            if (account == address(0)) revert ZeroAddress();
            if (bytes(reason).length == 0) revert EmptyReason();

            RestrictionInfo memory restrictionInfo = RestrictionInfo({
                isRestricted: true,
                reason: reason,
                timestamp: block.timestamp,
                restrictedBy: msg.sender
            });

            if (!_restrictedFrom[account].isRestricted) {
                _restrictedFrom[account] = restrictionInfo;
                emit AddressBlacklistedFrom(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }
            if (!_restrictedTo[account].isRestricted) {
                _restrictedTo[account] = restrictionInfo;
                emit AddressBlacklistedTo(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }

            emit AddressBlacklisted(
                account,
                reason,
                msg.sender,
                block.timestamp
            );
        }
    }

    /**
     * @dev Add addresses to both sending and receiving blacklists with a single reason
     * @param accounts The addresses to blacklist
     * @param reason The reason for blacklisting all addresses
     */
    function blacklist(
        address[] memory accounts,
        string memory reason
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (bytes(reason).length == 0) revert EmptyReason();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();

            RestrictionInfo memory restrictionInfo = RestrictionInfo({
                isRestricted: true,
                reason: reason,
                timestamp: block.timestamp,
                restrictedBy: msg.sender
            });

            if (!_restrictedFrom[account].isRestricted) {
                _restrictedFrom[account] = restrictionInfo;
                emit AddressBlacklistedFrom(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }
            if (!_restrictedTo[account].isRestricted) {
                _restrictedTo[account] = restrictionInfo;
                emit AddressBlacklistedTo(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }

            emit AddressBlacklisted(
                account,
                reason,
                msg.sender,
                block.timestamp
            );
        }
    }

    /**
     * @dev Remove addresses from the sending blacklist with reason
     * @param accounts The addresses to remove from blacklist
     * @param reason The reason for unrestricting
     */
    function unblacklistFrom(
        address[] memory accounts,
        string memory reason
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (bytes(reason).length == 0) revert EmptyReason();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();

            if (_restrictedFrom[account].isRestricted) {
                delete _restrictedFrom[account];
                emit AddressUnblacklistedFrom(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }
        }
    }

    /**
     * @dev Remove addresses from the receiving blacklist with reason
     * @param accounts The addresses to remove from blacklist
     * @param reason The reason for unrestricting
     */
    function unblacklistTo(
        address[] memory accounts,
        string memory reason
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (bytes(reason).length == 0) revert EmptyReason();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();

            if (_restrictedTo[account].isRestricted) {
                delete _restrictedTo[account];
                emit AddressUnblacklistedTo(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }
        }
    }

    /**
     * @dev Remove addresses from both blacklists with reason
     * @param accounts The addresses to remove from blacklist
     * @param reason The reason for unrestricting
     */
    function unblacklist(
        address[] memory accounts,
        string memory reason
    ) external onlyRole(BLACKLIST_ADMIN_ROLE) {
        if (accounts.length == 0) revert EmptyAddressList();
        if (bytes(reason).length == 0) revert EmptyReason();

        for (uint256 i = 0; i < accounts.length; i++) {
            address account = accounts[i];
            if (account == address(0)) revert ZeroAddress();

            bool wasRestrictedFrom = _restrictedFrom[account].isRestricted;
            bool wasRestrictedTo = _restrictedTo[account].isRestricted;

            if (wasRestrictedFrom) {
                delete _restrictedFrom[account];
                emit AddressUnblacklistedFrom(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }
            if (wasRestrictedTo) {
                delete _restrictedTo[account];
                emit AddressUnblacklistedTo(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }

            if (wasRestrictedFrom || wasRestrictedTo) {
                emit AddressUnblacklisted(
                    account,
                    reason,
                    msg.sender,
                    block.timestamp
                );
            }
        }
    }

    /**
     * @dev Check if an address is blacklisted from sending
     * @param account The address to check
     * @return bool True if the address is blacklisted from sending
     */
    function isBlacklistedFrom(address account) public view returns (bool) {
        return _restrictedFrom[account].isRestricted;
    }

    /**
     * @dev Check if an address is blacklisted from receiving
     * @param account The address to check
     * @return bool True if the address is blacklisted from receiving
     */
    function isBlacklistedTo(address account) public view returns (bool) {
        return _restrictedTo[account].isRestricted;
    }

    /**
     * @dev Check if an address is completely blacklisted
     * @param account The address to check
     * @return bool True if the address is blacklisted from both sending and receiving
     */
    function isBlacklisted(address account) public view returns (bool) {
        return
            _restrictedFrom[account].isRestricted &&
            _restrictedTo[account].isRestricted;
    }

    /**
     * @dev Get restriction information for sending
     * @param account The address to check
     * @return RestrictionInfo Complete restriction information
     */
    function getRestrictionInfoFrom(
        address account
    ) external view returns (RestrictionInfo memory) {
        return _restrictedFrom[account];
    }

    /**
     * @dev Get restriction information for receiving
     * @param account The address to check
     * @return RestrictionInfo Complete restriction information
     */
    function getRestrictionInfoTo(
        address account
    ) external view returns (RestrictionInfo memory) {
        return _restrictedTo[account];
    }

    /**
     * @dev Get restriction reason for sending
     * @param account The address to check
     * @return string The reason for restriction (empty if not restricted)
     */
    function getRestrictionReasonFrom(
        address account
    ) external view returns (string memory) {
        return _restrictedFrom[account].reason;
    }

    /**
     * @dev Get restriction reason for receiving
     * @param account The address to check
     * @return string The reason for restriction (empty if not restricted)
     */
    function getRestrictionReasonTo(
        address account
    ) external view returns (string memory) {
        return _restrictedTo[account].reason;
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
        if (_restrictedFrom[from].isRestricted) {
            return (
                false,
                string(
                    abi.encodePacked(
                        "Sender address is blacklisted from sending. Reason: ",
                        _restrictedFrom[from].reason
                    )
                )
            );
        }
        if (_restrictedTo[to].isRestricted) {
            return (
                false,
                string(
                    abi.encodePacked(
                        "Recipient address is blacklisted from receiving. Reason: ",
                        _restrictedTo[to].reason
                    )
                )
            );
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

    /**
     * @dev Compliance check for typed transfers
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     * @return bool True if the transfer is allowed
     */
    function canTransferWithType(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external view override returns (bool) {
        (bool success, ) = this.canTransferWithTypeAndFailureReason(
            from,
            to,
            amount,
            txType
        );
        return success;
    }

    /**
     * @dev Compliance check for typed transfers with failure reason
     * @param from Address tokens are transferred from
     * @param to Address tokens are transferred to
     * @param amount Amount of tokens to transfer
     * @param txType Type of transaction
     * @return (bool, string) Tuple of (success, failure reason)
     */
    function canTransferWithTypeAndFailureReason(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external view override returns (bool, string memory) {
        // For this compliance module, transaction type doesn't affect the blacklist check
        // So we just delegate to the standard compliance check
        return this.canTransferWithFailureReason(from, to, amount);
    }
}
