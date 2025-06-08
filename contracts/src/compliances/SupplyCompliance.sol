// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {ContextUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol";
import "../interfaces/ICompliance.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {TxType} from "../interfaces/ITypes.sol";

/**
 * @title SupplyRestrictionModuleCompliance
 * @dev A compliance module that enforces maximum supply restrictions for minting operations
 */
contract SupplyCompliance is ICompliance, AccessControlUpgradeable {
    /// @notice Role for managing supply limits
    bytes32 public constant SUPPLY_ADMIN_ROLE = keccak256("SUPPLY_ADMIN_ROLE");

    /// @notice Maximum allowed supply
    uint256 private _maxSupply;

    /// @notice Target token address (changed from immutable)
    address private _token;

    /// @notice Emitted when max supply is updated
    event MaxSupplyUpdated(uint256 oldMax, uint256 newMax);

    error TokenNotSet();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address token) public initializer {
        if (token == address(0)) revert TokenNotSet();

        __AccessControl_init();

        _token = token;

        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(SUPPLY_ADMIN_ROLE, _msgSender());
    }

    /**
     * @dev Implements IERC165 interface detection
     */
    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        virtual
        override(AccessControlUpgradeable, IERC165)
        returns (bool)
    {
        return
            interfaceId == type(ICompliance).interfaceId ||
            interfaceId == type(AccessControlUpgradeable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @notice Set the maximum supply limit
     * @param newMax The new maximum supply limit
     */
    function setMaxSupply(uint256 newMax) external onlyRole(SUPPLY_ADMIN_ROLE) {
        uint256 oldMax = _maxSupply;
        _maxSupply = newMax;
        emit MaxSupplyUpdated(oldMax, newMax);
    }

    /**
     * @notice Get the current maximum supply limit
     */
    function maxSupply() external view returns (uint256) {
        return _maxSupply;
    }

    /**
     * @notice Get the current total supply from the token
     */
    function currentSupply() public view returns (uint256) {
        return IERC20(_token).totalSupply();
    }

    /**
     * @inheritdoc ICompliance
     */
    function canTransfer(
        address from,
        address to,
        uint256 amount
    ) external view returns (bool) {
        // Only check mints (from = address(0))
        if (from == address(0)) {
            return currentSupply() + amount <= _maxSupply;
        }
        return true;
    }

    /**
     * @inheritdoc ICompliance
     */
    function canTransferWithFailureReason(
        address from,
        address to,
        uint256 amount
    ) external view returns (bool, string memory) {
        if (from == address(0)) {
            uint256 newSupply = currentSupply() + amount;
            if (newSupply > _maxSupply) {
                return (
                    false,
                    string.concat(
                        "Supply limit exceeded: requested=",
                        Strings.toString(newSupply),
                        ", max=",
                        Strings.toString(_maxSupply)
                    )
                );
            }
        }
        return (true, "");
    }

    /**
     * @dev Check if a typed transfer is allowed based on supply limits
     * @param from The sender address
     * @param to The recipient address
     * @param amount The amount to transfer
     * @param txType The type of transaction (unused in this module)
     * @return bool Whether the transfer is allowed
     */
    function canTransferWithType(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external view override returns (bool) {
        // For this compliance module, transaction type doesn't affect supply checks
        // So we just delegate to the standard compliance check
        return this.canTransfer(from, to, amount);
    }

    /**
     * @dev Check if a typed transfer is allowed based on supply limits with failure reason
     * @param from The sender address
     * @param to The recipient address
     * @param amount The amount to transfer
     * @param txType The type of transaction (unused in this module)
     * @return (bool, string) Whether the transfer is allowed and failure reason if not
     */
    function canTransferWithTypeAndFailureReason(
        address from,
        address to,
        uint256 amount,
        TxType txType
    ) external view override returns (bool, string memory) {
        // For this compliance module, transaction type doesn't affect supply checks
        // So we just delegate to the standard compliance check
        return this.canTransferWithFailureReason(from, to, amount);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
