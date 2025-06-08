// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BLACKLIST_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blacklist\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blacklistFrom\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blacklistTo\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canTransfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canTransferWithFailureReason\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isBlacklisted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isBlacklistedFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isBlacklistedTo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unblacklist\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unblacklistFrom\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unblacklistTo\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddressBlacklisted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressBlacklistedFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressBlacklistedTo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressUnblacklisted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressUnblacklistedFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AddressUnblacklistedTo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressAlreadyBlacklistedFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressAlreadyBlacklistedTo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressNotBlacklistedFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressNotBlacklistedTo\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EmptyAddressList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BLACKLISTADMINROLE is a free data retrieval call binding the contract method 0x67ed99f4.
//
// Solidity: function BLACKLIST_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) BLACKLISTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "BLACKLIST_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLACKLISTADMINROLE is a free data retrieval call binding the contract method 0x67ed99f4.
//
// Solidity: function BLACKLIST_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) BLACKLISTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.BLACKLISTADMINROLE(&_Contracts.CallOpts)
}

// BLACKLISTADMINROLE is a free data retrieval call binding the contract method 0x67ed99f4.
//
// Solidity: function BLACKLIST_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) BLACKLISTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.BLACKLISTADMINROLE(&_Contracts.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.DEFAULTADMINROLE(&_Contracts.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Contracts.Contract.DEFAULTADMINROLE(&_Contracts.CallOpts)
}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address from, address to, uint256 amount) view returns(bool)
func (_Contracts *ContractsCaller) CanTransfer(opts *bind.CallOpts, from common.Address, to common.Address, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "canTransfer", from, to, amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address from, address to, uint256 amount) view returns(bool)
func (_Contracts *ContractsSession) CanTransfer(from common.Address, to common.Address, amount *big.Int) (bool, error) {
	return _Contracts.Contract.CanTransfer(&_Contracts.CallOpts, from, to, amount)
}

// CanTransfer is a free data retrieval call binding the contract method 0xe46638e6.
//
// Solidity: function canTransfer(address from, address to, uint256 amount) view returns(bool)
func (_Contracts *ContractsCallerSession) CanTransfer(from common.Address, to common.Address, amount *big.Int) (bool, error) {
	return _Contracts.Contract.CanTransfer(&_Contracts.CallOpts, from, to, amount)
}

// CanTransferWithFailureReason is a free data retrieval call binding the contract method 0x78a53f5a.
//
// Solidity: function canTransferWithFailureReason(address from, address to, uint256 ) view returns(bool, string)
func (_Contracts *ContractsCaller) CanTransferWithFailureReason(opts *bind.CallOpts, from common.Address, to common.Address, arg2 *big.Int) (bool, string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "canTransferWithFailureReason", from, to, arg2)

	if err != nil {
		return *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// CanTransferWithFailureReason is a free data retrieval call binding the contract method 0x78a53f5a.
//
// Solidity: function canTransferWithFailureReason(address from, address to, uint256 ) view returns(bool, string)
func (_Contracts *ContractsSession) CanTransferWithFailureReason(from common.Address, to common.Address, arg2 *big.Int) (bool, string, error) {
	return _Contracts.Contract.CanTransferWithFailureReason(&_Contracts.CallOpts, from, to, arg2)
}

// CanTransferWithFailureReason is a free data retrieval call binding the contract method 0x78a53f5a.
//
// Solidity: function canTransferWithFailureReason(address from, address to, uint256 ) view returns(bool, string)
func (_Contracts *ContractsCallerSession) CanTransferWithFailureReason(from common.Address, to common.Address, arg2 *big.Int) (bool, string, error) {
	return _Contracts.Contract.CanTransferWithFailureReason(&_Contracts.CallOpts, from, to, arg2)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRoleAdmin(&_Contracts.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Contracts.Contract.GetRoleAdmin(&_Contracts.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.HasRole(&_Contracts.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Contracts *ContractsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Contracts.Contract.HasRole(&_Contracts.CallOpts, role, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Contracts *ContractsCaller) IsBlacklisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isBlacklisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Contracts *ContractsSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Contracts.Contract.IsBlacklisted(&_Contracts.CallOpts, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Contracts *ContractsCallerSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Contracts.Contract.IsBlacklisted(&_Contracts.CallOpts, account)
}

// IsBlacklistedFrom is a free data retrieval call binding the contract method 0xfcbab1e3.
//
// Solidity: function isBlacklistedFrom(address account) view returns(bool)
func (_Contracts *ContractsCaller) IsBlacklistedFrom(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isBlacklistedFrom", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklistedFrom is a free data retrieval call binding the contract method 0xfcbab1e3.
//
// Solidity: function isBlacklistedFrom(address account) view returns(bool)
func (_Contracts *ContractsSession) IsBlacklistedFrom(account common.Address) (bool, error) {
	return _Contracts.Contract.IsBlacklistedFrom(&_Contracts.CallOpts, account)
}

// IsBlacklistedFrom is a free data retrieval call binding the contract method 0xfcbab1e3.
//
// Solidity: function isBlacklistedFrom(address account) view returns(bool)
func (_Contracts *ContractsCallerSession) IsBlacklistedFrom(account common.Address) (bool, error) {
	return _Contracts.Contract.IsBlacklistedFrom(&_Contracts.CallOpts, account)
}

// IsBlacklistedTo is a free data retrieval call binding the contract method 0x934187b2.
//
// Solidity: function isBlacklistedTo(address account) view returns(bool)
func (_Contracts *ContractsCaller) IsBlacklistedTo(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isBlacklistedTo", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklistedTo is a free data retrieval call binding the contract method 0x934187b2.
//
// Solidity: function isBlacklistedTo(address account) view returns(bool)
func (_Contracts *ContractsSession) IsBlacklistedTo(account common.Address) (bool, error) {
	return _Contracts.Contract.IsBlacklistedTo(&_Contracts.CallOpts, account)
}

// IsBlacklistedTo is a free data retrieval call binding the contract method 0x934187b2.
//
// Solidity: function isBlacklistedTo(address account) view returns(bool)
func (_Contracts *ContractsCallerSession) IsBlacklistedTo(account common.Address) (bool, error) {
	return _Contracts.Contract.IsBlacklistedTo(&_Contracts.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// Blacklist is a paid mutator transaction binding the contract method 0x041f173f.
//
// Solidity: function blacklist(address[] accounts) returns()
func (_Contracts *ContractsTransactor) Blacklist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "blacklist", accounts)
}

// Blacklist is a paid mutator transaction binding the contract method 0x041f173f.
//
// Solidity: function blacklist(address[] accounts) returns()
func (_Contracts *ContractsSession) Blacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Blacklist(&_Contracts.TransactOpts, accounts)
}

// Blacklist is a paid mutator transaction binding the contract method 0x041f173f.
//
// Solidity: function blacklist(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) Blacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Blacklist(&_Contracts.TransactOpts, accounts)
}

// BlacklistFrom is a paid mutator transaction binding the contract method 0x343dbc01.
//
// Solidity: function blacklistFrom(address[] accounts) returns()
func (_Contracts *ContractsTransactor) BlacklistFrom(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "blacklistFrom", accounts)
}

// BlacklistFrom is a paid mutator transaction binding the contract method 0x343dbc01.
//
// Solidity: function blacklistFrom(address[] accounts) returns()
func (_Contracts *ContractsSession) BlacklistFrom(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BlacklistFrom(&_Contracts.TransactOpts, accounts)
}

// BlacklistFrom is a paid mutator transaction binding the contract method 0x343dbc01.
//
// Solidity: function blacklistFrom(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) BlacklistFrom(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BlacklistFrom(&_Contracts.TransactOpts, accounts)
}

// BlacklistTo is a paid mutator transaction binding the contract method 0x6d907c47.
//
// Solidity: function blacklistTo(address[] accounts) returns()
func (_Contracts *ContractsTransactor) BlacklistTo(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "blacklistTo", accounts)
}

// BlacklistTo is a paid mutator transaction binding the contract method 0x6d907c47.
//
// Solidity: function blacklistTo(address[] accounts) returns()
func (_Contracts *ContractsSession) BlacklistTo(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BlacklistTo(&_Contracts.TransactOpts, accounts)
}

// BlacklistTo is a paid mutator transaction binding the contract method 0x6d907c47.
//
// Solidity: function blacklistTo(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) BlacklistTo(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.BlacklistTo(&_Contracts.TransactOpts, accounts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantRole(&_Contracts.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.GrantRole(&_Contracts.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Contracts *ContractsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Contracts *ContractsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceRole(&_Contracts.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Contracts *ContractsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RenounceRole(&_Contracts.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeRole(&_Contracts.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Contracts *ContractsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RevokeRole(&_Contracts.TransactOpts, role, account)
}

// Unblacklist is a paid mutator transaction binding the contract method 0xbfab6535.
//
// Solidity: function unblacklist(address[] accounts) returns()
func (_Contracts *ContractsTransactor) Unblacklist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unblacklist", accounts)
}

// Unblacklist is a paid mutator transaction binding the contract method 0xbfab6535.
//
// Solidity: function unblacklist(address[] accounts) returns()
func (_Contracts *ContractsSession) Unblacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Unblacklist(&_Contracts.TransactOpts, accounts)
}

// Unblacklist is a paid mutator transaction binding the contract method 0xbfab6535.
//
// Solidity: function unblacklist(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) Unblacklist(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Unblacklist(&_Contracts.TransactOpts, accounts)
}

// UnblacklistFrom is a paid mutator transaction binding the contract method 0x7a5425be.
//
// Solidity: function unblacklistFrom(address[] accounts) returns()
func (_Contracts *ContractsTransactor) UnblacklistFrom(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unblacklistFrom", accounts)
}

// UnblacklistFrom is a paid mutator transaction binding the contract method 0x7a5425be.
//
// Solidity: function unblacklistFrom(address[] accounts) returns()
func (_Contracts *ContractsSession) UnblacklistFrom(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnblacklistFrom(&_Contracts.TransactOpts, accounts)
}

// UnblacklistFrom is a paid mutator transaction binding the contract method 0x7a5425be.
//
// Solidity: function unblacklistFrom(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) UnblacklistFrom(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnblacklistFrom(&_Contracts.TransactOpts, accounts)
}

// UnblacklistTo is a paid mutator transaction binding the contract method 0xc2a20abf.
//
// Solidity: function unblacklistTo(address[] accounts) returns()
func (_Contracts *ContractsTransactor) UnblacklistTo(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unblacklistTo", accounts)
}

// UnblacklistTo is a paid mutator transaction binding the contract method 0xc2a20abf.
//
// Solidity: function unblacklistTo(address[] accounts) returns()
func (_Contracts *ContractsSession) UnblacklistTo(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnblacklistTo(&_Contracts.TransactOpts, accounts)
}

// UnblacklistTo is a paid mutator transaction binding the contract method 0xc2a20abf.
//
// Solidity: function unblacklistTo(address[] accounts) returns()
func (_Contracts *ContractsTransactorSession) UnblacklistTo(accounts []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.UnblacklistTo(&_Contracts.TransactOpts, accounts)
}

// ContractsAddressBlacklistedIterator is returned from FilterAddressBlacklisted and is used to iterate over the raw logs and unpacked data for AddressBlacklisted events raised by the Contracts contract.
type ContractsAddressBlacklistedIterator struct {
	Event *ContractsAddressBlacklisted // Event containing the contract specifics and raw log

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
func (it *ContractsAddressBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressBlacklisted)
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
		it.Event = new(ContractsAddressBlacklisted)
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
func (it *ContractsAddressBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressBlacklisted represents a AddressBlacklisted event raised by the Contracts contract.
type ContractsAddressBlacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressBlacklisted is a free log retrieval operation binding the contract event 0xdaf49ab9345b6cb75bcb5a7f726bff9183c34dcf5c098c385730f9fd893765f6.
//
// Solidity: event AddressBlacklisted(address indexed account)
func (_Contracts *ContractsFilterer) FilterAddressBlacklisted(opts *bind.FilterOpts, account []common.Address) (*ContractsAddressBlacklistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressBlacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressBlacklistedIterator{contract: _Contracts.contract, event: "AddressBlacklisted", logs: logs, sub: sub}, nil
}

// WatchAddressBlacklisted is a free log subscription operation binding the contract event 0xdaf49ab9345b6cb75bcb5a7f726bff9183c34dcf5c098c385730f9fd893765f6.
//
// Solidity: event AddressBlacklisted(address indexed account)
func (_Contracts *ContractsFilterer) WatchAddressBlacklisted(opts *bind.WatchOpts, sink chan<- *ContractsAddressBlacklisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressBlacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressBlacklisted)
				if err := _Contracts.contract.UnpackLog(event, "AddressBlacklisted", log); err != nil {
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

// ParseAddressBlacklisted is a log parse operation binding the contract event 0xdaf49ab9345b6cb75bcb5a7f726bff9183c34dcf5c098c385730f9fd893765f6.
//
// Solidity: event AddressBlacklisted(address indexed account)
func (_Contracts *ContractsFilterer) ParseAddressBlacklisted(log types.Log) (*ContractsAddressBlacklisted, error) {
	event := new(ContractsAddressBlacklisted)
	if err := _Contracts.contract.UnpackLog(event, "AddressBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAddressBlacklistedFromIterator is returned from FilterAddressBlacklistedFrom and is used to iterate over the raw logs and unpacked data for AddressBlacklistedFrom events raised by the Contracts contract.
type ContractsAddressBlacklistedFromIterator struct {
	Event *ContractsAddressBlacklistedFrom // Event containing the contract specifics and raw log

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
func (it *ContractsAddressBlacklistedFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressBlacklistedFrom)
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
		it.Event = new(ContractsAddressBlacklistedFrom)
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
func (it *ContractsAddressBlacklistedFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressBlacklistedFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressBlacklistedFrom represents a AddressBlacklistedFrom event raised by the Contracts contract.
type ContractsAddressBlacklistedFrom struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressBlacklistedFrom is a free log retrieval operation binding the contract event 0x9dfdda91594e9fdc47d49655cb3e5be93cf9432fd57290eb605b597728b8b35f.
//
// Solidity: event AddressBlacklistedFrom(address indexed account)
func (_Contracts *ContractsFilterer) FilterAddressBlacklistedFrom(opts *bind.FilterOpts, account []common.Address) (*ContractsAddressBlacklistedFromIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressBlacklistedFrom", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressBlacklistedFromIterator{contract: _Contracts.contract, event: "AddressBlacklistedFrom", logs: logs, sub: sub}, nil
}

// WatchAddressBlacklistedFrom is a free log subscription operation binding the contract event 0x9dfdda91594e9fdc47d49655cb3e5be93cf9432fd57290eb605b597728b8b35f.
//
// Solidity: event AddressBlacklistedFrom(address indexed account)
func (_Contracts *ContractsFilterer) WatchAddressBlacklistedFrom(opts *bind.WatchOpts, sink chan<- *ContractsAddressBlacklistedFrom, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressBlacklistedFrom", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressBlacklistedFrom)
				if err := _Contracts.contract.UnpackLog(event, "AddressBlacklistedFrom", log); err != nil {
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

// ParseAddressBlacklistedFrom is a log parse operation binding the contract event 0x9dfdda91594e9fdc47d49655cb3e5be93cf9432fd57290eb605b597728b8b35f.
//
// Solidity: event AddressBlacklistedFrom(address indexed account)
func (_Contracts *ContractsFilterer) ParseAddressBlacklistedFrom(log types.Log) (*ContractsAddressBlacklistedFrom, error) {
	event := new(ContractsAddressBlacklistedFrom)
	if err := _Contracts.contract.UnpackLog(event, "AddressBlacklistedFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAddressBlacklistedToIterator is returned from FilterAddressBlacklistedTo and is used to iterate over the raw logs and unpacked data for AddressBlacklistedTo events raised by the Contracts contract.
type ContractsAddressBlacklistedToIterator struct {
	Event *ContractsAddressBlacklistedTo // Event containing the contract specifics and raw log

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
func (it *ContractsAddressBlacklistedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressBlacklistedTo)
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
		it.Event = new(ContractsAddressBlacklistedTo)
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
func (it *ContractsAddressBlacklistedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressBlacklistedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressBlacklistedTo represents a AddressBlacklistedTo event raised by the Contracts contract.
type ContractsAddressBlacklistedTo struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressBlacklistedTo is a free log retrieval operation binding the contract event 0x73a8a75339968fe050872be18a0d3551982e1a056b8c244c6d36de22b0ef571c.
//
// Solidity: event AddressBlacklistedTo(address indexed account)
func (_Contracts *ContractsFilterer) FilterAddressBlacklistedTo(opts *bind.FilterOpts, account []common.Address) (*ContractsAddressBlacklistedToIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressBlacklistedTo", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressBlacklistedToIterator{contract: _Contracts.contract, event: "AddressBlacklistedTo", logs: logs, sub: sub}, nil
}

// WatchAddressBlacklistedTo is a free log subscription operation binding the contract event 0x73a8a75339968fe050872be18a0d3551982e1a056b8c244c6d36de22b0ef571c.
//
// Solidity: event AddressBlacklistedTo(address indexed account)
func (_Contracts *ContractsFilterer) WatchAddressBlacklistedTo(opts *bind.WatchOpts, sink chan<- *ContractsAddressBlacklistedTo, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressBlacklistedTo", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressBlacklistedTo)
				if err := _Contracts.contract.UnpackLog(event, "AddressBlacklistedTo", log); err != nil {
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

// ParseAddressBlacklistedTo is a log parse operation binding the contract event 0x73a8a75339968fe050872be18a0d3551982e1a056b8c244c6d36de22b0ef571c.
//
// Solidity: event AddressBlacklistedTo(address indexed account)
func (_Contracts *ContractsFilterer) ParseAddressBlacklistedTo(log types.Log) (*ContractsAddressBlacklistedTo, error) {
	event := new(ContractsAddressBlacklistedTo)
	if err := _Contracts.contract.UnpackLog(event, "AddressBlacklistedTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAddressUnblacklistedIterator is returned from FilterAddressUnblacklisted and is used to iterate over the raw logs and unpacked data for AddressUnblacklisted events raised by the Contracts contract.
type ContractsAddressUnblacklistedIterator struct {
	Event *ContractsAddressUnblacklisted // Event containing the contract specifics and raw log

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
func (it *ContractsAddressUnblacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressUnblacklisted)
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
		it.Event = new(ContractsAddressUnblacklisted)
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
func (it *ContractsAddressUnblacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressUnblacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressUnblacklisted represents a AddressUnblacklisted event raised by the Contracts contract.
type ContractsAddressUnblacklisted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressUnblacklisted is a free log retrieval operation binding the contract event 0xb2aa8f2cce614e8fceaca560dbaf2a8ed3083e4ab371b10bf6d02e359216767a.
//
// Solidity: event AddressUnblacklisted(address indexed account)
func (_Contracts *ContractsFilterer) FilterAddressUnblacklisted(opts *bind.FilterOpts, account []common.Address) (*ContractsAddressUnblacklistedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressUnblacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressUnblacklistedIterator{contract: _Contracts.contract, event: "AddressUnblacklisted", logs: logs, sub: sub}, nil
}

// WatchAddressUnblacklisted is a free log subscription operation binding the contract event 0xb2aa8f2cce614e8fceaca560dbaf2a8ed3083e4ab371b10bf6d02e359216767a.
//
// Solidity: event AddressUnblacklisted(address indexed account)
func (_Contracts *ContractsFilterer) WatchAddressUnblacklisted(opts *bind.WatchOpts, sink chan<- *ContractsAddressUnblacklisted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressUnblacklisted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressUnblacklisted)
				if err := _Contracts.contract.UnpackLog(event, "AddressUnblacklisted", log); err != nil {
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

// ParseAddressUnblacklisted is a log parse operation binding the contract event 0xb2aa8f2cce614e8fceaca560dbaf2a8ed3083e4ab371b10bf6d02e359216767a.
//
// Solidity: event AddressUnblacklisted(address indexed account)
func (_Contracts *ContractsFilterer) ParseAddressUnblacklisted(log types.Log) (*ContractsAddressUnblacklisted, error) {
	event := new(ContractsAddressUnblacklisted)
	if err := _Contracts.contract.UnpackLog(event, "AddressUnblacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAddressUnblacklistedFromIterator is returned from FilterAddressUnblacklistedFrom and is used to iterate over the raw logs and unpacked data for AddressUnblacklistedFrom events raised by the Contracts contract.
type ContractsAddressUnblacklistedFromIterator struct {
	Event *ContractsAddressUnblacklistedFrom // Event containing the contract specifics and raw log

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
func (it *ContractsAddressUnblacklistedFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressUnblacklistedFrom)
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
		it.Event = new(ContractsAddressUnblacklistedFrom)
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
func (it *ContractsAddressUnblacklistedFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressUnblacklistedFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressUnblacklistedFrom represents a AddressUnblacklistedFrom event raised by the Contracts contract.
type ContractsAddressUnblacklistedFrom struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressUnblacklistedFrom is a free log retrieval operation binding the contract event 0x05ddbc1da122ce95119058bfb5e7156768d6a92bcf25629bc6608cc6d7627e15.
//
// Solidity: event AddressUnblacklistedFrom(address indexed account)
func (_Contracts *ContractsFilterer) FilterAddressUnblacklistedFrom(opts *bind.FilterOpts, account []common.Address) (*ContractsAddressUnblacklistedFromIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressUnblacklistedFrom", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressUnblacklistedFromIterator{contract: _Contracts.contract, event: "AddressUnblacklistedFrom", logs: logs, sub: sub}, nil
}

// WatchAddressUnblacklistedFrom is a free log subscription operation binding the contract event 0x05ddbc1da122ce95119058bfb5e7156768d6a92bcf25629bc6608cc6d7627e15.
//
// Solidity: event AddressUnblacklistedFrom(address indexed account)
func (_Contracts *ContractsFilterer) WatchAddressUnblacklistedFrom(opts *bind.WatchOpts, sink chan<- *ContractsAddressUnblacklistedFrom, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressUnblacklistedFrom", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressUnblacklistedFrom)
				if err := _Contracts.contract.UnpackLog(event, "AddressUnblacklistedFrom", log); err != nil {
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

// ParseAddressUnblacklistedFrom is a log parse operation binding the contract event 0x05ddbc1da122ce95119058bfb5e7156768d6a92bcf25629bc6608cc6d7627e15.
//
// Solidity: event AddressUnblacklistedFrom(address indexed account)
func (_Contracts *ContractsFilterer) ParseAddressUnblacklistedFrom(log types.Log) (*ContractsAddressUnblacklistedFrom, error) {
	event := new(ContractsAddressUnblacklistedFrom)
	if err := _Contracts.contract.UnpackLog(event, "AddressUnblacklistedFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsAddressUnblacklistedToIterator is returned from FilterAddressUnblacklistedTo and is used to iterate over the raw logs and unpacked data for AddressUnblacklistedTo events raised by the Contracts contract.
type ContractsAddressUnblacklistedToIterator struct {
	Event *ContractsAddressUnblacklistedTo // Event containing the contract specifics and raw log

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
func (it *ContractsAddressUnblacklistedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressUnblacklistedTo)
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
		it.Event = new(ContractsAddressUnblacklistedTo)
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
func (it *ContractsAddressUnblacklistedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressUnblacklistedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressUnblacklistedTo represents a AddressUnblacklistedTo event raised by the Contracts contract.
type ContractsAddressUnblacklistedTo struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddressUnblacklistedTo is a free log retrieval operation binding the contract event 0x0f8b068a900ac843db847c348ef052decfaa2ef3f6c4dbf70c55a354e5abc0e2.
//
// Solidity: event AddressUnblacklistedTo(address indexed account)
func (_Contracts *ContractsFilterer) FilterAddressUnblacklistedTo(opts *bind.FilterOpts, account []common.Address) (*ContractsAddressUnblacklistedToIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressUnblacklistedTo", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressUnblacklistedToIterator{contract: _Contracts.contract, event: "AddressUnblacklistedTo", logs: logs, sub: sub}, nil
}

// WatchAddressUnblacklistedTo is a free log subscription operation binding the contract event 0x0f8b068a900ac843db847c348ef052decfaa2ef3f6c4dbf70c55a354e5abc0e2.
//
// Solidity: event AddressUnblacklistedTo(address indexed account)
func (_Contracts *ContractsFilterer) WatchAddressUnblacklistedTo(opts *bind.WatchOpts, sink chan<- *ContractsAddressUnblacklistedTo, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressUnblacklistedTo", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressUnblacklistedTo)
				if err := _Contracts.contract.UnpackLog(event, "AddressUnblacklistedTo", log); err != nil {
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

// ParseAddressUnblacklistedTo is a log parse operation binding the contract event 0x0f8b068a900ac843db847c348ef052decfaa2ef3f6c4dbf70c55a354e5abc0e2.
//
// Solidity: event AddressUnblacklistedTo(address indexed account)
func (_Contracts *ContractsFilterer) ParseAddressUnblacklistedTo(log types.Log) (*ContractsAddressUnblacklistedTo, error) {
	event := new(ContractsAddressUnblacklistedTo)
	if err := _Contracts.contract.UnpackLog(event, "AddressUnblacklistedTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Contracts contract.
type ContractsRoleAdminChangedIterator struct {
	Event *ContractsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ContractsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleAdminChanged)
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
		it.Event = new(ContractsRoleAdminChanged)
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
func (it *ContractsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleAdminChanged represents a RoleAdminChanged event raised by the Contracts contract.
type ContractsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ContractsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleAdminChangedIterator{contract: _Contracts.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Contracts *ContractsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ContractsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleAdminChanged)
				if err := _Contracts.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseRoleAdminChanged(log types.Log) (*ContractsRoleAdminChanged, error) {
	event := new(ContractsRoleAdminChanged)
	if err := _Contracts.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Contracts contract.
type ContractsRoleGrantedIterator struct {
	Event *ContractsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ContractsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleGranted)
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
		it.Event = new(ContractsRoleGranted)
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
func (it *ContractsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleGranted represents a RoleGranted event raised by the Contracts contract.
type ContractsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractsRoleGrantedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleGrantedIterator{contract: _Contracts.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ContractsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleGranted)
				if err := _Contracts.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseRoleGranted(log types.Log) (*ContractsRoleGranted, error) {
	event := new(ContractsRoleGranted)
	if err := _Contracts.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Contracts contract.
type ContractsRoleRevokedIterator struct {
	Event *ContractsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ContractsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRoleRevoked)
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
		it.Event = new(ContractsRoleRevoked)
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
func (it *ContractsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRoleRevoked represents a RoleRevoked event raised by the Contracts contract.
type ContractsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ContractsRoleRevokedIterator, error) {

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

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsRoleRevokedIterator{contract: _Contracts.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Contracts *ContractsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ContractsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRoleRevoked)
				if err := _Contracts.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseRoleRevoked(log types.Log) (*ContractsRoleRevoked, error) {
	event := new(ContractsRoleRevoked)
	if err := _Contracts.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
