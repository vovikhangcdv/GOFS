// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {ICompliance} from "./interfaces/ICompliance.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {TxType} from "./interfaces/ITypes.sol";

contract ComplianceRegistry is ICompliance, AccessControlUpgradeable {
    using EnumerableSet for EnumerableSet.AddressSet;

    bytes32 public constant COMPLIANCE_ADMIN_ROLE =
        keccak256("COMPLIANCE_ADMIN_ROLE");

    EnumerableSet.AddressSet private _modules;

    // Custom errors
    error InvalidModuleAddress();
    error ModuleAlreadyAdded();
    error InvalidModuleInterface();
    error ModuleNotFound();

    // Events
    event ComplianceModuleAdded(address indexed module);
    event ComplianceModuleRemoved(address indexed module);

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __AccessControl_init();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(COMPLIANCE_ADMIN_ROLE, msg.sender);
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

    function addModule(
        address _module
    ) external onlyRole(COMPLIANCE_ADMIN_ROLE) {
        if (_module == address(0)) revert InvalidModuleAddress();
        if (_modules.contains(_module)) revert ModuleAlreadyAdded();

        // Check if module implements ICompliance interface using ERC165
        try
            IERC165(_module).supportsInterface(type(ICompliance).interfaceId)
        returns (bool supported) {
            if (!supported) {
                revert InvalidModuleInterface();
            }
        } catch {
            revert InvalidModuleInterface();
        }

        bool added = _modules.add(_module);
        require(added, "Failed to add module");

        emit ComplianceModuleAdded(_module);
    }

    function removeModule(
        address _module
    ) external onlyRole(COMPLIANCE_ADMIN_ROLE) {
        if (!_modules.contains(_module)) revert ModuleNotFound();

        bool removed = _modules.remove(_module);
        require(removed, "Failed to remove module");

        emit ComplianceModuleRemoved(_module);
    }

    function isModule(address _module) external view returns (bool) {
        return _modules.contains(_module);
    }

    function canTransfer(
        address _from,
        address _to,
        uint256 _amount
    ) external view returns (bool) {
        // Check all modules
        for (uint256 i = 0; i < _modules.length(); i++) {
            address module = _modules.at(i);
            if (!ICompliance(module).canTransfer(_from, _to, _amount)) {
                return false;
            }
        }
        return true;
    }

    function canTransferWithFailureReason(
        address _from,
        address _to,
        uint256 _amount
    ) external view returns (bool, string memory) {
        // Check all modules
        for (uint256 i = 0; i < _modules.length(); i++) {
            address module = _modules.at(i);
            (bool success, string memory reason) = ICompliance(module)
                .canTransferWithFailureReason(_from, _to, _amount);
            if (!success) {
                return (false, reason);
            }
        }
        return (true, "");
    }

    function getModules() external view returns (address[] memory) {
        return _modules.values();
    }

    function isRegisteredModule(address _module) external view returns (bool) {
        return _modules.contains(_module);
    }

    function canTransferWithType(
        address _from,
        address _to,
        uint256 _amount,
        TxType _txType
    ) external view returns (bool) {
        // Check all modules
        for (uint256 i = 0; i < _modules.length(); i++) {
            address module = _modules.at(i);
            if (
                !ICompliance(module).canTransferWithType(
                    _from,
                    _to,
                    _amount,
                    _txType
                )
            ) {
                return false;
            }
        }
        return true;
    }

    function canTransferWithTypeAndFailureReason(
        address _from,
        address _to,
        uint256 _amount,
        TxType _txType
    ) external view returns (bool, string memory) {
        // Check all modules
        for (uint256 i = 0; i < _modules.length(); i++) {
            address module = _modules.at(i);
            (bool success, string memory reason) = ICompliance(module)
                .canTransferWithTypeAndFailureReason(
                    _from,
                    _to,
                    _amount,
                    _txType
                );
            if (!success) {
                return (false, reason);
            }
        }
        return (true, "");
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[49] private __gap;
}
