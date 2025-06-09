// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {ComplianceRegistry} from "../src/ComplianceRegistry.sol";
import {ICompliance} from "../src/interfaces/ICompliance.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {TxType} from "../src/interfaces/ITypes.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
// Mock compliance module for testing
contract MockCompliance is ICompliance {
    bool private _shouldAllow;
    string private _failureReason;

    constructor(bool shouldAllow_, string memory failureReason_) {
        _shouldAllow = shouldAllow_;
        _failureReason = failureReason_;
    }

    function supportsInterface(
        bytes4 interfaceId
    ) external pure returns (bool) {
        return
            interfaceId == type(ICompliance).interfaceId ||
            interfaceId == type(IERC165).interfaceId;
    }

    function canTransfer(
        address,
        address,
        uint256
    ) external view override returns (bool) {
        return _shouldAllow;
    }

    function canTransferWithFailureReason(
        address,
        address,
        uint256
    ) external view override returns (bool, string memory) {
        return (_shouldAllow, _failureReason);
    }

    function canTransferWithType(
        address,
        address,
        uint256,
        TxType
    ) external view override returns (bool) {
        return _shouldAllow;
    }

    function canTransferWithTypeAndFailureReason(
        address,
        address,
        uint256,
        TxType
    ) external view override returns (bool, string memory) {
        return (_shouldAllow, _failureReason);
    }
}

// Invalid mock that doesn't properly implement ICompliance
contract InvalidMockCompliance {
    function someOtherFunction() external pure returns (bool) {
        return true;
    }
}

contract ComplianceRegistryTest is Test {
    ComplianceRegistry public registry;
    address public admin;
    MockCompliance public allowingModule;
    MockCompliance public denyingModule;
    InvalidMockCompliance public invalidModule;

    event ComplianceModuleAdded(address indexed module);
    event ComplianceModuleRemoved(address indexed module);

    bytes32 public constant COMPLIANCE_ADMIN_ROLE =
        keccak256("COMPLIANCE_ADMIN_ROLE");

    function setUp() public {
        admin = address(this);
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(new ComplianceRegistry()),
            address(new ProxyAdmin(admin)),
            abi.encodeWithSelector(
                ComplianceRegistry.initialize.selector,
                admin
            )
        );
        registry = ComplianceRegistry(address(proxy));
        allowingModule = new MockCompliance(true, "");
        denyingModule = new MockCompliance(false, "Transfer not allowed");
        invalidModule = new InvalidMockCompliance();
    }

    function test_InitialState() public {
        assertTrue(registry.hasRole(registry.DEFAULT_ADMIN_ROLE(), admin));
        assertTrue(registry.hasRole(registry.COMPLIANCE_ADMIN_ROLE(), admin));
        assertEq(registry.getModules().length, 0);
    }

    function test_AddModule() public {
        registry.addModule(address(allowingModule));
        assertTrue(registry.isRegisteredModule(address(allowingModule)));
        assertEq(registry.getModules().length, 1);
        assertEq(registry.getModules()[0], address(allowingModule));
    }

    function test_AddModule_RevertZeroAddress() public {
        vm.expectRevert(ComplianceRegistry.InvalidModuleAddress.selector);
        registry.addModule(address(0));
    }

    function test_AddModule_RevertDuplicate() public {
        registry.addModule(address(allowingModule));
        vm.expectRevert(ComplianceRegistry.ModuleAlreadyAdded.selector);
        registry.addModule(address(allowingModule));
    }

    function test_AddModule_RevertInvalidInterface_EOA() public {
        vm.expectRevert(ComplianceRegistry.InvalidModuleInterface.selector);
        registry.addModule(address(invalidModule));
    }

    function test_RemoveModule() public {
        registry.addModule(address(allowingModule));
        registry.removeModule(address(allowingModule));
        assertFalse(registry.isRegisteredModule(address(allowingModule)));
        assertEq(registry.getModules().length, 0);
    }

    function test_RemoveModule_RevertNotFound() public {
        vm.expectRevert(ComplianceRegistry.ModuleNotFound.selector);
        registry.removeModule(address(allowingModule));
    }

    function test_CanTransfer_SingleAllowingModule() public {
        registry.addModule(address(allowingModule));
        assertTrue(registry.canTransfer(address(1), address(2), 100));
    }

    function test_CanTransfer_SingleDenyingModule() public {
        registry.addModule(address(denyingModule));
        assertFalse(registry.canTransfer(address(1), address(2), 100));
    }

    function test_CanTransfer_MultipleModules() public {
        registry.addModule(address(allowingModule));
        registry.addModule(address(denyingModule));
        assertFalse(registry.canTransfer(address(1), address(2), 100));
    }

    function test_CanTransferWithFailureReason_SingleAllowingModule() public {
        registry.addModule(address(allowingModule));
        (bool success, string memory reason) = registry
            .canTransferWithFailureReason(address(1), address(2), 100);
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithFailureReason_SingleDenyingModule() public {
        registry.addModule(address(denyingModule));
        (bool success, string memory reason) = registry
            .canTransferWithFailureReason(address(1), address(2), 100);
        assertFalse(success);
        assertEq(reason, "Transfer not allowed");
    }

    function test_CanTransferWithFailureReason_MultipleModules() public {
        registry.addModule(address(allowingModule));
        registry.addModule(address(denyingModule));
        (bool success, string memory reason) = registry
            .canTransferWithFailureReason(address(1), address(2), 100);
        assertFalse(success);
        assertEq(reason, "Transfer not allowed");
    }

    function test_OnlyAdminCanAddModule() public {
        address nonAdmin = address(0x1);
        vm.startPrank(nonAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                nonAdmin,
                COMPLIANCE_ADMIN_ROLE
            )
        );
        registry.addModule(address(allowingModule));
        vm.stopPrank();
    }

    function test_OnlyAdminCanRemoveModule() public {
        registry.addModule(address(allowingModule));

        address nonAdmin = address(0x1);
        vm.startPrank(nonAdmin);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                nonAdmin,
                COMPLIANCE_ADMIN_ROLE
            )
        );
        registry.removeModule(address(allowingModule));
        vm.stopPrank();
    }

    function test_EmitEvents() public {
        vm.expectEmit(true, false, false, false);
        emit ComplianceModuleAdded(address(allowingModule));
        registry.addModule(address(allowingModule));

        vm.expectEmit(true, false, false, false);
        emit ComplianceModuleRemoved(address(allowingModule));
        registry.removeModule(address(allowingModule));
    }

    function test_SupportsInterface() public {
        // Test ICompliance interface
        assertTrue(registry.supportsInterface(type(ICompliance).interfaceId));

        // Test AccessControl interface
        assertTrue(
            registry.supportsInterface(type(IAccessControl).interfaceId)
        );

        // Test IERC165 interface
        assertTrue(registry.supportsInterface(type(IERC165).interfaceId));

        // Test invalid interface
        assertFalse(registry.supportsInterface(0x12345678));
    }

    function test_IsModule() public {
        // Test module not added
        assertFalse(registry.isModule(address(allowingModule)));

        // Add module and test
        registry.addModule(address(allowingModule));
        assertTrue(registry.isModule(address(allowingModule)));

        // Remove module and test
        registry.removeModule(address(allowingModule));
        assertFalse(registry.isModule(address(allowingModule)));
    }

    function test_CanTransferWithType_SingleAllowingModule() public {
        registry.addModule(address(allowingModule));
        assertTrue(
            registry.canTransferWithType(
                address(1),
                address(2),
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithType_SingleDenyingModule() public {
        registry.addModule(address(denyingModule));
        assertFalse(
            registry.canTransferWithType(
                address(1),
                address(2),
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithType_MultipleModules() public {
        registry.addModule(address(allowingModule));
        registry.addModule(address(denyingModule));
        assertFalse(
            registry.canTransferWithType(
                address(1),
                address(2),
                100,
                TxType.wrap(1)
            )
        );
    }

    function test_CanTransferWithTypeAndFailureReason_SingleAllowingModule()
        public
    {
        registry.addModule(address(allowingModule));
        (bool success, string memory reason) = registry
            .canTransferWithTypeAndFailureReason(
                address(1),
                address(2),
                100,
                TxType.wrap(1)
            );
        assertTrue(success);
        assertEq(reason, "");
    }

    function test_CanTransferWithTypeAndFailureReason_SingleDenyingModule()
        public
    {
        registry.addModule(address(denyingModule));
        (bool success, string memory reason) = registry
            .canTransferWithTypeAndFailureReason(
                address(1),
                address(2),
                100,
                TxType.wrap(1)
            );
        assertFalse(success);
        assertEq(reason, "Transfer not allowed");
    }

    function test_CanTransferWithTypeAndFailureReason_MultipleModules() public {
        registry.addModule(address(allowingModule));
        registry.addModule(address(denyingModule));
        (bool success, string memory reason) = registry
            .canTransferWithTypeAndFailureReason(
                address(1),
                address(2),
                100,
                TxType.wrap(1)
            );
        assertFalse(success);
        assertEq(reason, "Transfer not allowed");
    }

    function test_AddModule_RevertInvalidInterface_NonContract() public {
        // Test with EOA address - should revert because EOA cannot implement interface
        address eoa = makeAddr("eoa");
        vm.expectRevert(); // EOA will cause a generic revert when trying to call supportsInterface
        registry.addModule(eoa);
    }
}
