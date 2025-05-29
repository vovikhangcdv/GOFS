// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./interfaces/ICompliance.sol";

/**
 * @title ComplianceRegistry
 * @dev Registry that manages multiple compliance modules and orchestrates compliance checks between them.
 * This contract acts as a central registry for all compliance rules in the system.
 */
contract ComplianceRegistry is ICompliance, AccessControl {
    using EnumerableSet for EnumerableSet.AddressSet;

    // Role for managing compliance modules
    bytes32 public constant COMPLIANCE_ADMIN_ROLE =
        keccak256("COMPLIANCE_ADMIN_ROLE");

    // Set to store all compliance module addresses
    EnumerableSet.AddressSet private _modules;

    // Custom errors
    error InvalidModuleAddress();
    error ModuleAlreadyAdded();
    error InvalidModuleInterface();
    error ModuleNotFound();

    // Events
    event ComplianceModuleAdded(address indexed module);
    event ComplianceModuleRemoved(address indexed module);

    /**
     * @dev Constructor that sets up the admin role
     */
    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(COMPLIANCE_ADMIN_ROLE, msg.sender);
    }

    /**
     * @dev Add a new compliance module
     * @param _module Address of the compliance module to add
     */
    function addModule(
        address _module
    ) external onlyRole(COMPLIANCE_ADMIN_ROLE) {
        if (_module == address(0)) revert InvalidModuleAddress();
        if (_modules.contains(_module)) revert ModuleAlreadyAdded();

        // Test if module implements the correct interface
        try
            ICompliance(_module).canTransfer(address(0), address(0), 0)
        returns (bool) {
            // Module implements the interface correctly, proceed
        } catch {
            revert InvalidModuleInterface();
        }

        bool added = _modules.add(_module);
        require(added, "Failed to add module"); // Should never happen since we checked contains

        emit ComplianceModuleAdded(_module);
    }

    /**
     * @dev Remove a compliance module
     * @param _module Address of the compliance module to remove
     */
    function removeModule(
        address _module
    ) external onlyRole(COMPLIANCE_ADMIN_ROLE) {
        if (!_modules.contains(_module)) revert ModuleNotFound();

        bool removed = _modules.remove(_module);
        require(removed, "Failed to remove module"); // Should never happen since we checked contains

        emit ComplianceModuleRemoved(_module);
    }

    /**
     * @dev Check if a transfer is allowed by all compliance modules
     * @param _from Address tokens are transferred from
     * @param _to Address tokens are transferred to
     * @param _amount Amount of tokens transferred
     * @return bool Whether the transfer is allowed
     */
    function canTransfer(
        address _from,
        address _to,
        uint256 _amount
    ) public view override returns (bool) {
        uint256 length = _modules.length();
        for (uint256 i = 0; i < length; i++) {
            if (!ICompliance(_modules.at(i)).canTransfer(_from, _to, _amount)) {
                return false;
            }
        }
        return true;
    }

    /**
     * @dev Check if a transfer is allowed with failure reason
     * @param _from Address tokens are transferred from
     * @param _to Address tokens are transferred to
     * @param _amount Amount of tokens transferred
     * @return bool Whether the transfer is allowed
     * @return string The reason for failure if not allowed
     */
    function canTransferWithFailureReason(
        address _from,
        address _to,
        uint256 _amount
    ) public view override returns (bool, string memory) {
        uint256 length = _modules.length();
        for (uint256 i = 0; i < length; i++) {
            (bool success, string memory reason) = ICompliance(_modules.at(i))
                .canTransferWithFailureReason(_from, _to, _amount);
            if (!success) {
                return (false, reason);
            }
        }
        return (true, "");
    }

    /**
     * @dev Get all compliance modules
     * @return Array of compliance module addresses
     */
    function getModules() external view returns (address[] memory) {
        uint256 length = _modules.length();
        address[] memory moduleArray = new address[](length);
        for (uint256 i = 0; i < length; i++) {
            moduleArray[i] = _modules.at(i);
        }
        return moduleArray;
    }

    /**
     * @dev Check if an address is a registered compliance module
     * @param _module Address to check
     * @return bool Whether the address is a registered compliance module
     */
    function isRegisteredModule(address _module) external view returns (bool) {
        return _modules.contains(_module);
    }
}
