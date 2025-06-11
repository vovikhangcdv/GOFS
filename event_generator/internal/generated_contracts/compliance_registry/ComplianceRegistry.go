// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package complianceRegistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ComplianceRegistryMetaData contains all meta data concerning the ComplianceRegistry contract.
var ComplianceRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"COMPLIANCE_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addModule\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canTransfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canTransferWithFailureReason\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getModules\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isModule\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRegisteredModule\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeModule\",\"inputs\":[{\"name\":\"_module\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ComplianceModuleAdded\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceModuleRemoved\",\"inputs\":[{\"name\":\"module\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidModuleAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidModuleInterface\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModuleAlreadyAdded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModuleNotFound\",\"inputs\":[]}]",
}

// ComplianceRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ComplianceRegistryMetaData.ABI instead.
var ComplianceRegistryABI = ComplianceRegistryMetaData.ABI

// ComplianceRegistry is an auto generated Go binding around an Ethereum contract.
type ComplianceRegistry struct {
	ComplianceRegistryCaller     // Read-only binding to the contract
	ComplianceRegistryTransactor // Write-only binding to the contract
	ComplianceRegistryFilterer   // Log filterer for contract events
}

// ComplianceRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComplianceRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplianceRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplianceRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplianceRegistrySession struct {
	Contract     *ComplianceRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ComplianceRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplianceRegistryCallerSession struct {
	Contract *ComplianceRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ComplianceRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplianceRegistryTransactorSession struct {
	Contract     *ComplianceRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ComplianceRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComplianceRegistryRaw struct {
	Contract *ComplianceRegistry // Generic contract binding to access the raw methods on
}

// ComplianceRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplianceRegistryCallerRaw struct {
	Contract *ComplianceRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ComplianceRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplianceRegistryTransactorRaw struct {
	Contract *ComplianceRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComplianceRegistry creates a new instance of ComplianceRegistry, bound to a specific deployed contract.
func NewComplianceRegistry(address common.Address, backend bind.ContractBackend) (*ComplianceRegistry, error) {
	contract, err := bindComplianceRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistry{ComplianceRegistryCaller: ComplianceRegistryCaller{contract: contract}, ComplianceRegistryTransactor: ComplianceRegistryTransactor{contract: contract}, ComplianceRegistryFilterer: ComplianceRegistryFilterer{contract: contract}}, nil
}

// NewComplianceRegistryCaller creates a new read-only instance of ComplianceRegistry, bound to a specific deployed contract.
func NewComplianceRegistryCaller(address common.Address, caller bind.ContractCaller) (*ComplianceRegistryCaller, error) {
	contract, err := bindComplianceRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryCaller{contract: contract}, nil
}

// NewComplianceRegistryTransactor creates a new write-only instance of ComplianceRegistry, bound to a specific deployed contract.
func NewComplianceRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ComplianceRegistryTransactor, error) {
	contract, err := bindComplianceRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryTransactor{contract: contract}, nil
}

// NewComplianceRegistryFilterer creates a new log filterer instance of ComplianceRegistry, bound to a specific deployed contract.
func NewComplianceRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ComplianceRegistryFilterer, error) {
	contract, err := bindComplianceRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryFilterer{contract: contract}, nil
}

// bindComplianceRegistry binds a generic wrapper to an already deployed contract.
func bindComplianceRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComplianceRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplianceRegistry *ComplianceRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplianceRegistry.Contract.ComplianceRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplianceRegistry *ComplianceRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.ComplianceRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplianceRegistry *ComplianceRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.ComplianceRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplianceRegistry *ComplianceRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplianceRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplianceRegistry *ComplianceRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplianceRegistry *ComplianceRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.contract.Transact(opts, method, params...)
}

// COMPLIANCEADMINROLE is a free data retrieval call binding the contract method 0xa11939bd.
//
// Solidity: function COMPLIANCE_ADMIN_ROLE() view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistryCaller) COMPLIANCEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "COMPLIANCE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// COMPLIANCEADMINROLE is a free data retrieval call binding the contract method 0xa11939bd.
//
// Solidity: function COMPLIANCE_ADMIN_ROLE() view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistrySession) COMPLIANCEADMINROLE() ([32]byte, error) {
	return _ComplianceRegistry.Contract.COMPLIANCEADMINROLE(&_ComplianceRegistry.CallOpts)
}

// COMPLIANCEADMINROLE is a free data retrieval call binding the contract method 0xa11939bd.
//
// Solidity: function COMPLIANCE_ADMIN_ROLE() view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) COMPLIANCEADMINROLE() ([32]byte, error) {
	return _ComplianceRegistry.Contract.COMPLIANCEADMINROLE(&_ComplianceRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ComplianceRegistry.Contract.DEFAULTADMINROLE(&_ComplianceRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ComplianceRegistry.Contract.DEFAULTADMINROLE(&_ComplianceRegistry.CallOpts)
}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address _from, address _to, uint256 _amount) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCaller) CanTransfer(opts *bind.CallOpts, _from common.Address, _to common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "canTransfer", _from, _to, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address _from, address _to, uint256 _amount) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistrySession) CanTransfer(_from common.Address, _to common.Address, _amount *big.Int) (bool, error) {
	return _ComplianceRegistry.Contract.CanTransfer(&_ComplianceRegistry.CallOpts, _from, _to, _amount)
}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address _from, address _to, uint256 _amount) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) CanTransfer(_from common.Address, _to common.Address, _amount *big.Int) (bool, error) {
	return _ComplianceRegistry.Contract.CanTransfer(&_ComplianceRegistry.CallOpts, _from, _to, _amount)
}

// CanTransferWithFailureReason is a free data retrieval call binding the contract method 0x78a53f5a.
//
// Solidity: function canTransferWithFailureReason(address _from, address _to, uint256 _amount) view returns(bool, string)
func (_ComplianceRegistry *ComplianceRegistryCaller) CanTransferWithFailureReason(opts *bind.CallOpts, _from common.Address, _to common.Address, _amount *big.Int) (bool, string, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "canTransferWithFailureReason", _from, _to, _amount)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// CanTransferWithFailureReason is a free data retrieval call binding the contract method 0x78a53f5a.
//
// Solidity: function canTransferWithFailureReason(address _from, address _to, uint256 _amount) view returns(bool, string)
func (_ComplianceRegistry *ComplianceRegistrySession) CanTransferWithFailureReason(_from common.Address, _to common.Address, _amount *big.Int) (bool, string, error) {
	return _ComplianceRegistry.Contract.CanTransferWithFailureReason(&_ComplianceRegistry.CallOpts, _from, _to, _amount)
}

// CanTransferWithFailureReason is a free data retrieval call binding the contract method 0x78a53f5a.
//
// Solidity: function canTransferWithFailureReason(address _from, address _to, uint256 _amount) view returns(bool, string)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) CanTransferWithFailureReason(_from common.Address, _to common.Address, _amount *big.Int) (bool, string, error) {
	return _ComplianceRegistry.Contract.CanTransferWithFailureReason(&_ComplianceRegistry.CallOpts, _from, _to, _amount)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ComplianceRegistry *ComplianceRegistryCaller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ComplianceRegistry *ComplianceRegistrySession) GetModules() ([]common.Address, error) {
	return _ComplianceRegistry.Contract.GetModules(&_ComplianceRegistry.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ComplianceRegistry *ComplianceRegistryCallerSession) GetModules() ([]common.Address, error) {
	return _ComplianceRegistry.Contract.GetModules(&_ComplianceRegistry.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ComplianceRegistry.Contract.GetRoleAdmin(&_ComplianceRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ComplianceRegistry.Contract.GetRoleAdmin(&_ComplianceRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ComplianceRegistry.Contract.HasRole(&_ComplianceRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ComplianceRegistry.Contract.HasRole(&_ComplianceRegistry.CallOpts, role, account)
}

// IsModule is a free data retrieval call binding the contract method 0x42f6e389.
//
// Solidity: function isModule(address _module) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCaller) IsModule(opts *bind.CallOpts, _module common.Address) (bool, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "isModule", _module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModule is a free data retrieval call binding the contract method 0x42f6e389.
//
// Solidity: function isModule(address _module) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistrySession) IsModule(_module common.Address) (bool, error) {
	return _ComplianceRegistry.Contract.IsModule(&_ComplianceRegistry.CallOpts, _module)
}

// IsModule is a free data retrieval call binding the contract method 0x42f6e389.
//
// Solidity: function isModule(address _module) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) IsModule(_module common.Address) (bool, error) {
	return _ComplianceRegistry.Contract.IsModule(&_ComplianceRegistry.CallOpts, _module)
}

// IsRegisteredModule is a free data retrieval call binding the contract method 0x0bcd4ebb.
//
// Solidity: function isRegisteredModule(address _module) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCaller) IsRegisteredModule(opts *bind.CallOpts, _module common.Address) (bool, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "isRegisteredModule", _module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredModule is a free data retrieval call binding the contract method 0x0bcd4ebb.
//
// Solidity: function isRegisteredModule(address _module) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistrySession) IsRegisteredModule(_module common.Address) (bool, error) {
	return _ComplianceRegistry.Contract.IsRegisteredModule(&_ComplianceRegistry.CallOpts, _module)
}

// IsRegisteredModule is a free data retrieval call binding the contract method 0x0bcd4ebb.
//
// Solidity: function isRegisteredModule(address _module) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) IsRegisteredModule(_module common.Address) (bool, error) {
	return _ComplianceRegistry.Contract.IsRegisteredModule(&_ComplianceRegistry.CallOpts, _module)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ComplianceRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ComplianceRegistry.Contract.SupportsInterface(&_ComplianceRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ComplianceRegistry *ComplianceRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ComplianceRegistry.Contract.SupportsInterface(&_ComplianceRegistry.CallOpts, interfaceId)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address _module) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactor) AddModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.contract.Transact(opts, "addModule", _module)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address _module) returns()
func (_ComplianceRegistry *ComplianceRegistrySession) AddModule(_module common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.AddModule(&_ComplianceRegistry.TransactOpts, _module)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address _module) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactorSession) AddModule(_module common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.AddModule(&_ComplianceRegistry.TransactOpts, _module)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ComplianceRegistry *ComplianceRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.GrantRole(&_ComplianceRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.GrantRole(&_ComplianceRegistry.TransactOpts, role, account)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address _module) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactor) RemoveModule(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.contract.Transact(opts, "removeModule", _module)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address _module) returns()
func (_ComplianceRegistry *ComplianceRegistrySession) RemoveModule(_module common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.RemoveModule(&_ComplianceRegistry.TransactOpts, _module)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address _module) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactorSession) RemoveModule(_module common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.RemoveModule(&_ComplianceRegistry.TransactOpts, _module)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ComplianceRegistry *ComplianceRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.RenounceRole(&_ComplianceRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.RenounceRole(&_ComplianceRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ComplianceRegistry *ComplianceRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.RevokeRole(&_ComplianceRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ComplianceRegistry *ComplianceRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ComplianceRegistry.Contract.RevokeRole(&_ComplianceRegistry.TransactOpts, role, account)
}

// ComplianceRegistryComplianceModuleAddedIterator is returned from FilterComplianceModuleAdded and is used to iterate over the raw logs and unpacked data for ComplianceModuleAdded events raised by the ComplianceRegistry contract.
type ComplianceRegistryComplianceModuleAddedIterator struct {
	Event *ComplianceRegistryComplianceModuleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComplianceRegistryComplianceModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceRegistryComplianceModuleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComplianceRegistryComplianceModuleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComplianceRegistryComplianceModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceRegistryComplianceModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceRegistryComplianceModuleAdded represents a ComplianceModuleAdded event raised by the ComplianceRegistry contract.
type ComplianceRegistryComplianceModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterComplianceModuleAdded is a free log retrieval operation binding the contract event 0xe1be16b136432bd24b5b1388a40ec7863174aff11cc151ab59749f87da9fdee1.
//
// Solidity: event ComplianceModuleAdded(address indexed module)
func (_ComplianceRegistry *ComplianceRegistryFilterer) FilterComplianceModuleAdded(opts *bind.FilterOpts, module []common.Address) (*ComplianceRegistryComplianceModuleAddedIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.FilterLogs(opts, "ComplianceModuleAdded", moduleRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryComplianceModuleAddedIterator{contract: _ComplianceRegistry.contract, event: "ComplianceModuleAdded", logs: logs, sub: sub}, nil
}

// WatchComplianceModuleAdded is a free log subscription operation binding the contract event 0xe1be16b136432bd24b5b1388a40ec7863174aff11cc151ab59749f87da9fdee1.
//
// Solidity: event ComplianceModuleAdded(address indexed module)
func (_ComplianceRegistry *ComplianceRegistryFilterer) WatchComplianceModuleAdded(opts *bind.WatchOpts, sink chan<- *ComplianceRegistryComplianceModuleAdded, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.WatchLogs(opts, "ComplianceModuleAdded", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceRegistryComplianceModuleAdded)
				if err := _ComplianceRegistry.contract.UnpackLog(event, "ComplianceModuleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComplianceModuleAdded is a log parse operation binding the contract event 0xe1be16b136432bd24b5b1388a40ec7863174aff11cc151ab59749f87da9fdee1.
//
// Solidity: event ComplianceModuleAdded(address indexed module)
func (_ComplianceRegistry *ComplianceRegistryFilterer) ParseComplianceModuleAdded(log types.Log) (*ComplianceRegistryComplianceModuleAdded, error) {
	event := new(ComplianceRegistryComplianceModuleAdded)
	if err := _ComplianceRegistry.contract.UnpackLog(event, "ComplianceModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComplianceRegistryComplianceModuleRemovedIterator is returned from FilterComplianceModuleRemoved and is used to iterate over the raw logs and unpacked data for ComplianceModuleRemoved events raised by the ComplianceRegistry contract.
type ComplianceRegistryComplianceModuleRemovedIterator struct {
	Event *ComplianceRegistryComplianceModuleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComplianceRegistryComplianceModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceRegistryComplianceModuleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComplianceRegistryComplianceModuleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComplianceRegistryComplianceModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceRegistryComplianceModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceRegistryComplianceModuleRemoved represents a ComplianceModuleRemoved event raised by the ComplianceRegistry contract.
type ComplianceRegistryComplianceModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterComplianceModuleRemoved is a free log retrieval operation binding the contract event 0xf3f11f4e6b87d4c3587a337c4533dab5c1dee243f72c0e77c1341a876157bbbe.
//
// Solidity: event ComplianceModuleRemoved(address indexed module)
func (_ComplianceRegistry *ComplianceRegistryFilterer) FilterComplianceModuleRemoved(opts *bind.FilterOpts, module []common.Address) (*ComplianceRegistryComplianceModuleRemovedIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.FilterLogs(opts, "ComplianceModuleRemoved", moduleRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryComplianceModuleRemovedIterator{contract: _ComplianceRegistry.contract, event: "ComplianceModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchComplianceModuleRemoved is a free log subscription operation binding the contract event 0xf3f11f4e6b87d4c3587a337c4533dab5c1dee243f72c0e77c1341a876157bbbe.
//
// Solidity: event ComplianceModuleRemoved(address indexed module)
func (_ComplianceRegistry *ComplianceRegistryFilterer) WatchComplianceModuleRemoved(opts *bind.WatchOpts, sink chan<- *ComplianceRegistryComplianceModuleRemoved, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.WatchLogs(opts, "ComplianceModuleRemoved", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceRegistryComplianceModuleRemoved)
				if err := _ComplianceRegistry.contract.UnpackLog(event, "ComplianceModuleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComplianceModuleRemoved is a log parse operation binding the contract event 0xf3f11f4e6b87d4c3587a337c4533dab5c1dee243f72c0e77c1341a876157bbbe.
//
// Solidity: event ComplianceModuleRemoved(address indexed module)
func (_ComplianceRegistry *ComplianceRegistryFilterer) ParseComplianceModuleRemoved(log types.Log) (*ComplianceRegistryComplianceModuleRemoved, error) {
	event := new(ComplianceRegistryComplianceModuleRemoved)
	if err := _ComplianceRegistry.contract.UnpackLog(event, "ComplianceModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComplianceRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ComplianceRegistry contract.
type ComplianceRegistryRoleAdminChangedIterator struct {
	Event *ComplianceRegistryRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComplianceRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceRegistryRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComplianceRegistryRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComplianceRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the ComplianceRegistry contract.
type ComplianceRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ComplianceRegistry *ComplianceRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ComplianceRegistryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryRoleAdminChangedIterator{contract: _ComplianceRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ComplianceRegistry *ComplianceRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ComplianceRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceRegistryRoleAdminChanged)
				if err := _ComplianceRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ComplianceRegistry *ComplianceRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*ComplianceRegistryRoleAdminChanged, error) {
	event := new(ComplianceRegistryRoleAdminChanged)
	if err := _ComplianceRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComplianceRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ComplianceRegistry contract.
type ComplianceRegistryRoleGrantedIterator struct {
	Event *ComplianceRegistryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComplianceRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceRegistryRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComplianceRegistryRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComplianceRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceRegistryRoleGranted represents a RoleGranted event raised by the ComplianceRegistry contract.
type ComplianceRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ComplianceRegistry *ComplianceRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ComplianceRegistryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryRoleGrantedIterator{contract: _ComplianceRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ComplianceRegistry *ComplianceRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ComplianceRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceRegistryRoleGranted)
				if err := _ComplianceRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ComplianceRegistry *ComplianceRegistryFilterer) ParseRoleGranted(log types.Log) (*ComplianceRegistryRoleGranted, error) {
	event := new(ComplianceRegistryRoleGranted)
	if err := _ComplianceRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComplianceRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ComplianceRegistry contract.
type ComplianceRegistryRoleRevokedIterator struct {
	Event *ComplianceRegistryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComplianceRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceRegistryRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComplianceRegistryRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComplianceRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceRegistryRoleRevoked represents a RoleRevoked event raised by the ComplianceRegistry contract.
type ComplianceRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ComplianceRegistry *ComplianceRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ComplianceRegistryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceRegistryRoleRevokedIterator{contract: _ComplianceRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ComplianceRegistry *ComplianceRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ComplianceRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ComplianceRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceRegistryRoleRevoked)
				if err := _ComplianceRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ComplianceRegistry *ComplianceRegistryFilterer) ParseRoleRevoked(log types.Log) (*ComplianceRegistryRoleRevoked, error) {
	event := new(ComplianceRegistryRoleRevoked)
	if err := _ComplianceRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
