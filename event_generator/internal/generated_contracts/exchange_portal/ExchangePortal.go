// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchangePortal

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

// Entity is an auto generated low-level Go binding around an user-defined struct.
type Entity struct {
	EntityAddress common.Address
	EntityType    uint8
	EntityData    []byte
	Verifier      common.Address
}

// ExchangePortalMetaData contains all meta data concerning the ExchangePortal contract.
var ExchangePortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_treasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_exchangeFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXCHANGE_FEE_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EXCHANGE_RATE_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FEE_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REGISTER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exchangeFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExchangeAmount\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExchangeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerWithRegistry\",\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entity\",\"type\":\"tuple\",\"internalType\":\"structEntity\",\"components\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"EntityType\"},{\"name\":\"entityData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"verifierSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExchangeFee\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExchangeRate\",\"inputs\":[{\"name\":\"newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTreasury\",\"inputs\":[{\"name\":\"newTreasury\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExchangeExecuted\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"toToken\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExchangeRateUpdated\",\"inputs\":[{\"name\":\"token0\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token1\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeUpdated\",\"inputs\":[{\"name\":\"newFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TreasuryUpdated\",\"inputs\":[{\"name\":\"newTreasury\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExcessiveSlippage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExchangeRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidFeeConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenPair\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokensMustBeDifferent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// ExchangePortalABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangePortalMetaData.ABI instead.
var ExchangePortalABI = ExchangePortalMetaData.ABI

// ExchangePortal is an auto generated Go binding around an Ethereum contract.
type ExchangePortal struct {
	ExchangePortalCaller     // Read-only binding to the contract
	ExchangePortalTransactor // Write-only binding to the contract
	ExchangePortalFilterer   // Log filterer for contract events
}

// ExchangePortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangePortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangePortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangePortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangePortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangePortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangePortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangePortalSession struct {
	Contract     *ExchangePortal   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangePortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangePortalCallerSession struct {
	Contract *ExchangePortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExchangePortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangePortalTransactorSession struct {
	Contract     *ExchangePortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExchangePortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangePortalRaw struct {
	Contract *ExchangePortal // Generic contract binding to access the raw methods on
}

// ExchangePortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangePortalCallerRaw struct {
	Contract *ExchangePortalCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangePortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangePortalTransactorRaw struct {
	Contract *ExchangePortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangePortal creates a new instance of ExchangePortal, bound to a specific deployed contract.
func NewExchangePortal(address common.Address, backend bind.ContractBackend) (*ExchangePortal, error) {
	contract, err := bindExchangePortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangePortal{ExchangePortalCaller: ExchangePortalCaller{contract: contract}, ExchangePortalTransactor: ExchangePortalTransactor{contract: contract}, ExchangePortalFilterer: ExchangePortalFilterer{contract: contract}}, nil
}

// NewExchangePortalCaller creates a new read-only instance of ExchangePortal, bound to a specific deployed contract.
func NewExchangePortalCaller(address common.Address, caller bind.ContractCaller) (*ExchangePortalCaller, error) {
	contract, err := bindExchangePortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalCaller{contract: contract}, nil
}

// NewExchangePortalTransactor creates a new write-only instance of ExchangePortal, bound to a specific deployed contract.
func NewExchangePortalTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangePortalTransactor, error) {
	contract, err := bindExchangePortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalTransactor{contract: contract}, nil
}

// NewExchangePortalFilterer creates a new log filterer instance of ExchangePortal, bound to a specific deployed contract.
func NewExchangePortalFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangePortalFilterer, error) {
	contract, err := bindExchangePortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalFilterer{contract: contract}, nil
}

// bindExchangePortal binds a generic wrapper to an already deployed contract.
func bindExchangePortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangePortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangePortal *ExchangePortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangePortal.Contract.ExchangePortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangePortal *ExchangePortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangePortal.Contract.ExchangePortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangePortal *ExchangePortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangePortal.Contract.ExchangePortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangePortal *ExchangePortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangePortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangePortal *ExchangePortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangePortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangePortal *ExchangePortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangePortal.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.DEFAULTADMINROLE(&_ExchangePortal.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.DEFAULTADMINROLE(&_ExchangePortal.CallOpts)
}

// EXCHANGEFEEADMINROLE is a free data retrieval call binding the contract method 0x9df6bcb3.
//
// Solidity: function EXCHANGE_FEE_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCaller) EXCHANGEFEEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "EXCHANGE_FEE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXCHANGEFEEADMINROLE is a free data retrieval call binding the contract method 0x9df6bcb3.
//
// Solidity: function EXCHANGE_FEE_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalSession) EXCHANGEFEEADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.EXCHANGEFEEADMINROLE(&_ExchangePortal.CallOpts)
}

// EXCHANGEFEEADMINROLE is a free data retrieval call binding the contract method 0x9df6bcb3.
//
// Solidity: function EXCHANGE_FEE_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCallerSession) EXCHANGEFEEADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.EXCHANGEFEEADMINROLE(&_ExchangePortal.CallOpts)
}

// EXCHANGERATEADMINROLE is a free data retrieval call binding the contract method 0x38c86b22.
//
// Solidity: function EXCHANGE_RATE_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCaller) EXCHANGERATEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "EXCHANGE_RATE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXCHANGERATEADMINROLE is a free data retrieval call binding the contract method 0x38c86b22.
//
// Solidity: function EXCHANGE_RATE_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalSession) EXCHANGERATEADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.EXCHANGERATEADMINROLE(&_ExchangePortal.CallOpts)
}

// EXCHANGERATEADMINROLE is a free data retrieval call binding the contract method 0x38c86b22.
//
// Solidity: function EXCHANGE_RATE_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCallerSession) EXCHANGERATEADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.EXCHANGERATEADMINROLE(&_ExchangePortal.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_ExchangePortal *ExchangePortalCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_ExchangePortal *ExchangePortalSession) FEEDENOMINATOR() (*big.Int, error) {
	return _ExchangePortal.Contract.FEEDENOMINATOR(&_ExchangePortal.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_ExchangePortal *ExchangePortalCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _ExchangePortal.Contract.FEEDENOMINATOR(&_ExchangePortal.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_ExchangePortal *ExchangePortalCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_ExchangePortal *ExchangePortalSession) MAXFEE() (*big.Int, error) {
	return _ExchangePortal.Contract.MAXFEE(&_ExchangePortal.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_ExchangePortal *ExchangePortalCallerSession) MAXFEE() (*big.Int, error) {
	return _ExchangePortal.Contract.MAXFEE(&_ExchangePortal.CallOpts)
}

// REGISTERADMINROLE is a free data retrieval call binding the contract method 0xc3f4c2a7.
//
// Solidity: function REGISTER_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCaller) REGISTERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "REGISTER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGISTERADMINROLE is a free data retrieval call binding the contract method 0xc3f4c2a7.
//
// Solidity: function REGISTER_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalSession) REGISTERADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.REGISTERADMINROLE(&_ExchangePortal.CallOpts)
}

// REGISTERADMINROLE is a free data retrieval call binding the contract method 0xc3f4c2a7.
//
// Solidity: function REGISTER_ADMIN_ROLE() view returns(bytes32)
func (_ExchangePortal *ExchangePortalCallerSession) REGISTERADMINROLE() ([32]byte, error) {
	return _ExchangePortal.Contract.REGISTERADMINROLE(&_ExchangePortal.CallOpts)
}

// ExchangeFee is a free data retrieval call binding the contract method 0x2ecd3be4.
//
// Solidity: function exchangeFee() view returns(uint256)
func (_ExchangePortal *ExchangePortalCaller) ExchangeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "exchangeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeFee is a free data retrieval call binding the contract method 0x2ecd3be4.
//
// Solidity: function exchangeFee() view returns(uint256)
func (_ExchangePortal *ExchangePortalSession) ExchangeFee() (*big.Int, error) {
	return _ExchangePortal.Contract.ExchangeFee(&_ExchangePortal.CallOpts)
}

// ExchangeFee is a free data retrieval call binding the contract method 0x2ecd3be4.
//
// Solidity: function exchangeFee() view returns(uint256)
func (_ExchangePortal *ExchangePortalCallerSession) ExchangeFee() (*big.Int, error) {
	return _ExchangePortal.Contract.ExchangeFee(&_ExchangePortal.CallOpts)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x95d6fd67.
//
// Solidity: function getExchangeAmount(address fromToken, address toToken, uint256 amountIn) view returns(uint256)
func (_ExchangePortal *ExchangePortalCaller) GetExchangeAmount(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "getExchangeAmount", fromToken, toToken, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x95d6fd67.
//
// Solidity: function getExchangeAmount(address fromToken, address toToken, uint256 amountIn) view returns(uint256)
func (_ExchangePortal *ExchangePortalSession) GetExchangeAmount(fromToken common.Address, toToken common.Address, amountIn *big.Int) (*big.Int, error) {
	return _ExchangePortal.Contract.GetExchangeAmount(&_ExchangePortal.CallOpts, fromToken, toToken, amountIn)
}

// GetExchangeAmount is a free data retrieval call binding the contract method 0x95d6fd67.
//
// Solidity: function getExchangeAmount(address fromToken, address toToken, uint256 amountIn) view returns(uint256)
func (_ExchangePortal *ExchangePortalCallerSession) GetExchangeAmount(fromToken common.Address, toToken common.Address, amountIn *big.Int) (*big.Int, error) {
	return _ExchangePortal.Contract.GetExchangeAmount(&_ExchangePortal.CallOpts, fromToken, toToken, amountIn)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_ExchangePortal *ExchangePortalCaller) GetExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_ExchangePortal *ExchangePortalSession) GetExchangeRate() (*big.Int, error) {
	return _ExchangePortal.Contract.GetExchangeRate(&_ExchangePortal.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_ExchangePortal *ExchangePortalCallerSession) GetExchangeRate() (*big.Int, error) {
	return _ExchangePortal.Contract.GetExchangeRate(&_ExchangePortal.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExchangePortal *ExchangePortalCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExchangePortal *ExchangePortalSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExchangePortal.Contract.GetRoleAdmin(&_ExchangePortal.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExchangePortal *ExchangePortalCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExchangePortal.Contract.GetRoleAdmin(&_ExchangePortal.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExchangePortal *ExchangePortalCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExchangePortal *ExchangePortalSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExchangePortal.Contract.HasRole(&_ExchangePortal.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExchangePortal *ExchangePortalCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExchangePortal.Contract.HasRole(&_ExchangePortal.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExchangePortal *ExchangePortalCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExchangePortal *ExchangePortalSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExchangePortal.Contract.SupportsInterface(&_ExchangePortal.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExchangePortal *ExchangePortalCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExchangePortal.Contract.SupportsInterface(&_ExchangePortal.CallOpts, interfaceId)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ExchangePortal *ExchangePortalCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ExchangePortal *ExchangePortalSession) Token0() (common.Address, error) {
	return _ExchangePortal.Contract.Token0(&_ExchangePortal.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ExchangePortal *ExchangePortalCallerSession) Token0() (common.Address, error) {
	return _ExchangePortal.Contract.Token0(&_ExchangePortal.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ExchangePortal *ExchangePortalCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ExchangePortal *ExchangePortalSession) Token1() (common.Address, error) {
	return _ExchangePortal.Contract.Token1(&_ExchangePortal.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ExchangePortal *ExchangePortalCallerSession) Token1() (common.Address, error) {
	return _ExchangePortal.Contract.Token1(&_ExchangePortal.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ExchangePortal *ExchangePortalCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangePortal.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ExchangePortal *ExchangePortalSession) Treasury() (common.Address, error) {
	return _ExchangePortal.Contract.Treasury(&_ExchangePortal.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_ExchangePortal *ExchangePortalCallerSession) Treasury() (common.Address, error) {
	return _ExchangePortal.Contract.Treasury(&_ExchangePortal.CallOpts)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address fromToken, address toToken, uint256 amountIn, uint256 minAmountOut) returns(uint256 amountOut)
func (_ExchangePortal *ExchangePortalTransactor) Exchange(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "exchange", fromToken, toToken, amountIn, minAmountOut)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address fromToken, address toToken, uint256 amountIn, uint256 minAmountOut) returns(uint256 amountOut)
func (_ExchangePortal *ExchangePortalSession) Exchange(fromToken common.Address, toToken common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.Contract.Exchange(&_ExchangePortal.TransactOpts, fromToken, toToken, amountIn, minAmountOut)
}

// Exchange is a paid mutator transaction binding the contract method 0x0ed2fc95.
//
// Solidity: function exchange(address fromToken, address toToken, uint256 amountIn, uint256 minAmountOut) returns(uint256 amountOut)
func (_ExchangePortal *ExchangePortalTransactorSession) Exchange(fromToken common.Address, toToken common.Address, amountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.Contract.Exchange(&_ExchangePortal.TransactOpts, fromToken, toToken, amountIn, minAmountOut)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExchangePortal *ExchangePortalTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExchangePortal *ExchangePortalSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.GrantRole(&_ExchangePortal.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.GrantRole(&_ExchangePortal.TransactOpts, role, account)
}

// RegisterWithRegistry is a paid mutator transaction binding the contract method 0xec049345.
//
// Solidity: function registerWithRegistry(address registryAddress, (address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_ExchangePortal *ExchangePortalTransactor) RegisterWithRegistry(opts *bind.TransactOpts, registryAddress common.Address, entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "registerWithRegistry", registryAddress, entity, verifierSignature)
}

// RegisterWithRegistry is a paid mutator transaction binding the contract method 0xec049345.
//
// Solidity: function registerWithRegistry(address registryAddress, (address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_ExchangePortal *ExchangePortalSession) RegisterWithRegistry(registryAddress common.Address, entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _ExchangePortal.Contract.RegisterWithRegistry(&_ExchangePortal.TransactOpts, registryAddress, entity, verifierSignature)
}

// RegisterWithRegistry is a paid mutator transaction binding the contract method 0xec049345.
//
// Solidity: function registerWithRegistry(address registryAddress, (address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) RegisterWithRegistry(registryAddress common.Address, entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _ExchangePortal.Contract.RegisterWithRegistry(&_ExchangePortal.TransactOpts, registryAddress, entity, verifierSignature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ExchangePortal *ExchangePortalTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ExchangePortal *ExchangePortalSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.RenounceRole(&_ExchangePortal.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.RenounceRole(&_ExchangePortal.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExchangePortal *ExchangePortalTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExchangePortal *ExchangePortalSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.RevokeRole(&_ExchangePortal.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.RevokeRole(&_ExchangePortal.TransactOpts, role, account)
}

// SetExchangeFee is a paid mutator transaction binding the contract method 0x9df8cc7b.
//
// Solidity: function setExchangeFee(uint256 newFee) returns()
func (_ExchangePortal *ExchangePortalTransactor) SetExchangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "setExchangeFee", newFee)
}

// SetExchangeFee is a paid mutator transaction binding the contract method 0x9df8cc7b.
//
// Solidity: function setExchangeFee(uint256 newFee) returns()
func (_ExchangePortal *ExchangePortalSession) SetExchangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.Contract.SetExchangeFee(&_ExchangePortal.TransactOpts, newFee)
}

// SetExchangeFee is a paid mutator transaction binding the contract method 0x9df8cc7b.
//
// Solidity: function setExchangeFee(uint256 newFee) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) SetExchangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.Contract.SetExchangeFee(&_ExchangePortal.TransactOpts, newFee)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_ExchangePortal *ExchangePortalTransactor) SetExchangeRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "setExchangeRate", newRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_ExchangePortal *ExchangePortalSession) SetExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.Contract.SetExchangeRate(&_ExchangePortal.TransactOpts, newRate)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(uint256 newRate) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) SetExchangeRate(newRate *big.Int) (*types.Transaction, error) {
	return _ExchangePortal.Contract.SetExchangeRate(&_ExchangePortal.TransactOpts, newRate)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_ExchangePortal *ExchangePortalTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _ExchangePortal.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_ExchangePortal *ExchangePortalSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.SetTreasury(&_ExchangePortal.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_ExchangePortal *ExchangePortalTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _ExchangePortal.Contract.SetTreasury(&_ExchangePortal.TransactOpts, newTreasury)
}

// ExchangePortalExchangeExecutedIterator is returned from FilterExchangeExecuted and is used to iterate over the raw logs and unpacked data for ExchangeExecuted events raised by the ExchangePortal contract.
type ExchangePortalExchangeExecutedIterator struct {
	Event *ExchangePortalExchangeExecuted // Event containing the contract specifics and raw log

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
func (it *ExchangePortalExchangeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalExchangeExecuted)
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
		it.Event = new(ExchangePortalExchangeExecuted)
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
func (it *ExchangePortalExchangeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalExchangeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalExchangeExecuted represents a ExchangeExecuted event raised by the ExchangePortal contract.
type ExchangePortalExchangeExecuted struct {
	FromToken common.Address
	ToToken   common.Address
	User      common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExchangeExecuted is a free log retrieval operation binding the contract event 0x5069a13fea9cb6e4b74ddfe95f6761207cda3f88a9091cd96007ffc727866b85.
//
// Solidity: event ExchangeExecuted(address indexed fromToken, address indexed toToken, address indexed user, uint256 amountIn, uint256 amountOut, uint256 feeAmount)
func (_ExchangePortal *ExchangePortalFilterer) FilterExchangeExecuted(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, user []common.Address) (*ExchangePortalExchangeExecutedIterator, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "ExchangeExecuted", fromTokenRule, toTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalExchangeExecutedIterator{contract: _ExchangePortal.contract, event: "ExchangeExecuted", logs: logs, sub: sub}, nil
}

// WatchExchangeExecuted is a free log subscription operation binding the contract event 0x5069a13fea9cb6e4b74ddfe95f6761207cda3f88a9091cd96007ffc727866b85.
//
// Solidity: event ExchangeExecuted(address indexed fromToken, address indexed toToken, address indexed user, uint256 amountIn, uint256 amountOut, uint256 feeAmount)
func (_ExchangePortal *ExchangePortalFilterer) WatchExchangeExecuted(opts *bind.WatchOpts, sink chan<- *ExchangePortalExchangeExecuted, fromToken []common.Address, toToken []common.Address, user []common.Address) (event.Subscription, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "ExchangeExecuted", fromTokenRule, toTokenRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalExchangeExecuted)
				if err := _ExchangePortal.contract.UnpackLog(event, "ExchangeExecuted", log); err != nil {
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

// ParseExchangeExecuted is a log parse operation binding the contract event 0x5069a13fea9cb6e4b74ddfe95f6761207cda3f88a9091cd96007ffc727866b85.
//
// Solidity: event ExchangeExecuted(address indexed fromToken, address indexed toToken, address indexed user, uint256 amountIn, uint256 amountOut, uint256 feeAmount)
func (_ExchangePortal *ExchangePortalFilterer) ParseExchangeExecuted(log types.Log) (*ExchangePortalExchangeExecuted, error) {
	event := new(ExchangePortalExchangeExecuted)
	if err := _ExchangePortal.contract.UnpackLog(event, "ExchangeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePortalExchangeRateUpdatedIterator is returned from FilterExchangeRateUpdated and is used to iterate over the raw logs and unpacked data for ExchangeRateUpdated events raised by the ExchangePortal contract.
type ExchangePortalExchangeRateUpdatedIterator struct {
	Event *ExchangePortalExchangeRateUpdated // Event containing the contract specifics and raw log

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
func (it *ExchangePortalExchangeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalExchangeRateUpdated)
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
		it.Event = new(ExchangePortalExchangeRateUpdated)
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
func (it *ExchangePortalExchangeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalExchangeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalExchangeRateUpdated represents a ExchangeRateUpdated event raised by the ExchangePortal contract.
type ExchangePortalExchangeRateUpdated struct {
	Token0  common.Address
	Token1  common.Address
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateUpdated is a free log retrieval operation binding the contract event 0x59a196fe84b02ac66eaa374ec156916370441ddfce08830f5b33dfbdb057791a.
//
// Solidity: event ExchangeRateUpdated(address indexed token0, address indexed token1, uint256 newRate)
func (_ExchangePortal *ExchangePortalFilterer) FilterExchangeRateUpdated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*ExchangePortalExchangeRateUpdatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "ExchangeRateUpdated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalExchangeRateUpdatedIterator{contract: _ExchangePortal.contract, event: "ExchangeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchExchangeRateUpdated is a free log subscription operation binding the contract event 0x59a196fe84b02ac66eaa374ec156916370441ddfce08830f5b33dfbdb057791a.
//
// Solidity: event ExchangeRateUpdated(address indexed token0, address indexed token1, uint256 newRate)
func (_ExchangePortal *ExchangePortalFilterer) WatchExchangeRateUpdated(opts *bind.WatchOpts, sink chan<- *ExchangePortalExchangeRateUpdated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "ExchangeRateUpdated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalExchangeRateUpdated)
				if err := _ExchangePortal.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
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

// ParseExchangeRateUpdated is a log parse operation binding the contract event 0x59a196fe84b02ac66eaa374ec156916370441ddfce08830f5b33dfbdb057791a.
//
// Solidity: event ExchangeRateUpdated(address indexed token0, address indexed token1, uint256 newRate)
func (_ExchangePortal *ExchangePortalFilterer) ParseExchangeRateUpdated(log types.Log) (*ExchangePortalExchangeRateUpdated, error) {
	event := new(ExchangePortalExchangeRateUpdated)
	if err := _ExchangePortal.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePortalFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the ExchangePortal contract.
type ExchangePortalFeeUpdatedIterator struct {
	Event *ExchangePortalFeeUpdated // Event containing the contract specifics and raw log

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
func (it *ExchangePortalFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalFeeUpdated)
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
		it.Event = new(ExchangePortalFeeUpdated)
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
func (it *ExchangePortalFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalFeeUpdated represents a FeeUpdated event raised by the ExchangePortal contract.
type ExchangePortalFeeUpdated struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFee)
func (_ExchangePortal *ExchangePortalFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*ExchangePortalFeeUpdatedIterator, error) {

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &ExchangePortalFeeUpdatedIterator{contract: _ExchangePortal.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFee)
func (_ExchangePortal *ExchangePortalFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *ExchangePortalFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalFeeUpdated)
				if err := _ExchangePortal.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x8c4d35e54a3f2ef1134138fd8ea3daee6a3c89e10d2665996babdf70261e2c76.
//
// Solidity: event FeeUpdated(uint256 newFee)
func (_ExchangePortal *ExchangePortalFilterer) ParseFeeUpdated(log types.Log) (*ExchangePortalFeeUpdated, error) {
	event := new(ExchangePortalFeeUpdated)
	if err := _ExchangePortal.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePortalRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ExchangePortal contract.
type ExchangePortalRoleAdminChangedIterator struct {
	Event *ExchangePortalRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExchangePortalRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalRoleAdminChanged)
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
		it.Event = new(ExchangePortalRoleAdminChanged)
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
func (it *ExchangePortalRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalRoleAdminChanged represents a RoleAdminChanged event raised by the ExchangePortal contract.
type ExchangePortalRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExchangePortal *ExchangePortalFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ExchangePortalRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalRoleAdminChangedIterator{contract: _ExchangePortal.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExchangePortal *ExchangePortalFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ExchangePortalRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalRoleAdminChanged)
				if err := _ExchangePortal.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ExchangePortal *ExchangePortalFilterer) ParseRoleAdminChanged(log types.Log) (*ExchangePortalRoleAdminChanged, error) {
	event := new(ExchangePortalRoleAdminChanged)
	if err := _ExchangePortal.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePortalRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ExchangePortal contract.
type ExchangePortalRoleGrantedIterator struct {
	Event *ExchangePortalRoleGranted // Event containing the contract specifics and raw log

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
func (it *ExchangePortalRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalRoleGranted)
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
		it.Event = new(ExchangePortalRoleGranted)
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
func (it *ExchangePortalRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalRoleGranted represents a RoleGranted event raised by the ExchangePortal contract.
type ExchangePortalRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExchangePortal *ExchangePortalFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExchangePortalRoleGrantedIterator, error) {

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

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalRoleGrantedIterator{contract: _ExchangePortal.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExchangePortal *ExchangePortalFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ExchangePortalRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalRoleGranted)
				if err := _ExchangePortal.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ExchangePortal *ExchangePortalFilterer) ParseRoleGranted(log types.Log) (*ExchangePortalRoleGranted, error) {
	event := new(ExchangePortalRoleGranted)
	if err := _ExchangePortal.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePortalRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ExchangePortal contract.
type ExchangePortalRoleRevokedIterator struct {
	Event *ExchangePortalRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ExchangePortalRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalRoleRevoked)
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
		it.Event = new(ExchangePortalRoleRevoked)
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
func (it *ExchangePortalRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalRoleRevoked represents a RoleRevoked event raised by the ExchangePortal contract.
type ExchangePortalRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExchangePortal *ExchangePortalFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExchangePortalRoleRevokedIterator, error) {

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

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangePortalRoleRevokedIterator{contract: _ExchangePortal.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExchangePortal *ExchangePortalFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ExchangePortalRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalRoleRevoked)
				if err := _ExchangePortal.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ExchangePortal *ExchangePortalFilterer) ParseRoleRevoked(log types.Log) (*ExchangePortalRoleRevoked, error) {
	event := new(ExchangePortalRoleRevoked)
	if err := _ExchangePortal.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangePortalTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the ExchangePortal contract.
type ExchangePortalTreasuryUpdatedIterator struct {
	Event *ExchangePortalTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *ExchangePortalTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangePortalTreasuryUpdated)
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
		it.Event = new(ExchangePortalTreasuryUpdated)
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
func (it *ExchangePortalTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangePortalTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangePortalTreasuryUpdated represents a TreasuryUpdated event raised by the ExchangePortal contract.
type ExchangePortalTreasuryUpdated struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_ExchangePortal *ExchangePortalFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*ExchangePortalTreasuryUpdatedIterator, error) {

	logs, sub, err := _ExchangePortal.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &ExchangePortalTreasuryUpdatedIterator{contract: _ExchangePortal.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_ExchangePortal *ExchangePortalFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *ExchangePortalTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _ExchangePortal.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangePortalTreasuryUpdated)
				if err := _ExchangePortal.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_ExchangePortal *ExchangePortalFilterer) ParseTreasuryUpdated(log types.Log) (*ExchangePortalTreasuryUpdated, error) {
	event := new(ExchangePortalTreasuryUpdated)
	if err := _ExchangePortal.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
