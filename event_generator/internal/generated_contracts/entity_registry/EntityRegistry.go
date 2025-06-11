// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package entityRegistry

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

// EntityRegistryMetaData contains all meta data concerning the EntityRegistry contract.
var EntityRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTITY_REGISTRY_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ENTITY_TYPE_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addVerifier\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityTypes\",\"type\":\"uint8[]\",\"internalType\":\"EntityType[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardRegister\",\"inputs\":[{\"name\":\"entity\",\"type\":\"tuple\",\"internalType\":\"structEntity\",\"components\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"EntityType\"},{\"name\":\"entityData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"verifierSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getEntity\",\"inputs\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEntity\",\"components\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"EntityType\"},{\"name\":\"entityData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isVerifiedEntity\",\"inputs\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"entity\",\"type\":\"tuple\",\"internalType\":\"structEntity\",\"components\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"EntityType\"},{\"name\":\"entityData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"verifierSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeVerifier\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateVerifier\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityTypes\",\"type\":\"uint8[]\",\"internalType\":\"EntityType[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyInfo\",\"inputs\":[{\"name\":\"entity\",\"type\":\"tuple\",\"internalType\":\"structEntity\",\"components\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"EntityType\"},{\"name\":\"entityData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"hashedInfo\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EntityRegistered\",\"inputs\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"entity\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structEntity\",\"components\":[{\"name\":\"entityAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"entityType\",\"type\":\"uint8\",\"internalType\":\"EntityType\"},{\"name\":\"entityData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierAdded\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"entityTypes\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"EntityType[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierRemoved\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierUpdated\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"entityTypes\",\"type\":\"uint8[]\",\"indexed\":false,\"internalType\":\"EntityType[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EmptyEntityTypes\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EntityAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVerifierSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySelfRegistrationAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedVerifier\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerifierAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerifierDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerifierNotAllowedForEntityType\",\"inputs\":[]}]",
}

// EntityRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EntityRegistryMetaData.ABI instead.
var EntityRegistryABI = EntityRegistryMetaData.ABI

// EntityRegistry is an auto generated Go binding around an Ethereum contract.
type EntityRegistry struct {
	EntityRegistryCaller     // Read-only binding to the contract
	EntityRegistryTransactor // Write-only binding to the contract
	EntityRegistryFilterer   // Log filterer for contract events
}

// EntityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntityRegistrySession struct {
	Contract     *EntityRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntityRegistryCallerSession struct {
	Contract *EntityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EntityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntityRegistryTransactorSession struct {
	Contract     *EntityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EntityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntityRegistryRaw struct {
	Contract *EntityRegistry // Generic contract binding to access the raw methods on
}

// EntityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntityRegistryCallerRaw struct {
	Contract *EntityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EntityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntityRegistryTransactorRaw struct {
	Contract *EntityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntityRegistry creates a new instance of EntityRegistry, bound to a specific deployed contract.
func NewEntityRegistry(address common.Address, backend bind.ContractBackend) (*EntityRegistry, error) {
	contract, err := bindEntityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntityRegistry{EntityRegistryCaller: EntityRegistryCaller{contract: contract}, EntityRegistryTransactor: EntityRegistryTransactor{contract: contract}, EntityRegistryFilterer: EntityRegistryFilterer{contract: contract}}, nil
}

// NewEntityRegistryCaller creates a new read-only instance of EntityRegistry, bound to a specific deployed contract.
func NewEntityRegistryCaller(address common.Address, caller bind.ContractCaller) (*EntityRegistryCaller, error) {
	contract, err := bindEntityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryCaller{contract: contract}, nil
}

// NewEntityRegistryTransactor creates a new write-only instance of EntityRegistry, bound to a specific deployed contract.
func NewEntityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EntityRegistryTransactor, error) {
	contract, err := bindEntityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryTransactor{contract: contract}, nil
}

// NewEntityRegistryFilterer creates a new log filterer instance of EntityRegistry, bound to a specific deployed contract.
func NewEntityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EntityRegistryFilterer, error) {
	contract, err := bindEntityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryFilterer{contract: contract}, nil
}

// bindEntityRegistry binds a generic wrapper to an already deployed contract.
func bindEntityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EntityRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityRegistry *EntityRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntityRegistry.Contract.EntityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityRegistry *EntityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityRegistry.Contract.EntityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityRegistry *EntityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityRegistry.Contract.EntityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityRegistry *EntityRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityRegistry *EntityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityRegistry *EntityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EntityRegistry.Contract.DEFAULTADMINROLE(&_EntityRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _EntityRegistry.Contract.DEFAULTADMINROLE(&_EntityRegistry.CallOpts)
}

// ENTITYREGISTRYADMINROLE is a free data retrieval call binding the contract method 0xd1b4cb72.
//
// Solidity: function ENTITY_REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCaller) ENTITYREGISTRYADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "ENTITY_REGISTRY_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTITYREGISTRYADMINROLE is a free data retrieval call binding the contract method 0xd1b4cb72.
//
// Solidity: function ENTITY_REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistrySession) ENTITYREGISTRYADMINROLE() ([32]byte, error) {
	return _EntityRegistry.Contract.ENTITYREGISTRYADMINROLE(&_EntityRegistry.CallOpts)
}

// ENTITYREGISTRYADMINROLE is a free data retrieval call binding the contract method 0xd1b4cb72.
//
// Solidity: function ENTITY_REGISTRY_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCallerSession) ENTITYREGISTRYADMINROLE() ([32]byte, error) {
	return _EntityRegistry.Contract.ENTITYREGISTRYADMINROLE(&_EntityRegistry.CallOpts)
}

// ENTITYTYPEHASH is a free data retrieval call binding the contract method 0xd066f2c9.
//
// Solidity: function ENTITY_TYPE_HASH() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCaller) ENTITYTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "ENTITY_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ENTITYTYPEHASH is a free data retrieval call binding the contract method 0xd066f2c9.
//
// Solidity: function ENTITY_TYPE_HASH() view returns(bytes32)
func (_EntityRegistry *EntityRegistrySession) ENTITYTYPEHASH() ([32]byte, error) {
	return _EntityRegistry.Contract.ENTITYTYPEHASH(&_EntityRegistry.CallOpts)
}

// ENTITYTYPEHASH is a free data retrieval call binding the contract method 0xd066f2c9.
//
// Solidity: function ENTITY_TYPE_HASH() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCallerSession) ENTITYTYPEHASH() ([32]byte, error) {
	return _EntityRegistry.Contract.ENTITYTYPEHASH(&_EntityRegistry.CallOpts)
}

// PROACTIVEREGISTRYFOWARDERADMINROLE is a free data retrieval call binding the contract method 0x437e6236.
//
// Solidity: function PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCaller) PROACTIVEREGISTRYFOWARDERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROACTIVEREGISTRYFOWARDERADMINROLE is a free data retrieval call binding the contract method 0x437e6236.
//
// Solidity: function PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistrySession) PROACTIVEREGISTRYFOWARDERADMINROLE() ([32]byte, error) {
	return _EntityRegistry.Contract.PROACTIVEREGISTRYFOWARDERADMINROLE(&_EntityRegistry.CallOpts)
}

// PROACTIVEREGISTRYFOWARDERADMINROLE is a free data retrieval call binding the contract method 0x437e6236.
//
// Solidity: function PROACTIVE_REGISTRY_FOWARDER_ADMIN_ROLE() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCallerSession) PROACTIVEREGISTRYFOWARDERADMINROLE() ([32]byte, error) {
	return _EntityRegistry.Contract.PROACTIVEREGISTRYFOWARDERADMINROLE(&_EntityRegistry.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EntityRegistry *EntityRegistrySession) DomainSeparator() ([32]byte, error) {
	return _EntityRegistry.Contract.DomainSeparator(&_EntityRegistry.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EntityRegistry *EntityRegistryCallerSession) DomainSeparator() ([32]byte, error) {
	return _EntityRegistry.Contract.DomainSeparator(&_EntityRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EntityRegistry *EntityRegistryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EntityRegistry *EntityRegistrySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EntityRegistry.Contract.Eip712Domain(&_EntityRegistry.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_EntityRegistry *EntityRegistryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _EntityRegistry.Contract.Eip712Domain(&_EntityRegistry.CallOpts)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address entityAddress) view returns((address,uint8,bytes,address))
func (_EntityRegistry *EntityRegistryCaller) GetEntity(opts *bind.CallOpts, entityAddress common.Address) (Entity, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "getEntity", entityAddress)

	if err != nil {
		return *new(Entity), err
	}

	out0 := *abi.ConvertType(out[0], new(Entity)).(*Entity)

	return out0, err

}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address entityAddress) view returns((address,uint8,bytes,address))
func (_EntityRegistry *EntityRegistrySession) GetEntity(entityAddress common.Address) (Entity, error) {
	return _EntityRegistry.Contract.GetEntity(&_EntityRegistry.CallOpts, entityAddress)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address entityAddress) view returns((address,uint8,bytes,address))
func (_EntityRegistry *EntityRegistryCallerSession) GetEntity(entityAddress common.Address) (Entity, error) {
	return _EntityRegistry.Contract.GetEntity(&_EntityRegistry.CallOpts, entityAddress)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EntityRegistry *EntityRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EntityRegistry *EntityRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EntityRegistry.Contract.GetRoleAdmin(&_EntityRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_EntityRegistry *EntityRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _EntityRegistry.Contract.GetRoleAdmin(&_EntityRegistry.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EntityRegistry *EntityRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EntityRegistry *EntityRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EntityRegistry.Contract.HasRole(&_EntityRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_EntityRegistry *EntityRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _EntityRegistry.Contract.HasRole(&_EntityRegistry.CallOpts, role, account)
}

// IsVerifiedEntity is a free data retrieval call binding the contract method 0x455f1f36.
//
// Solidity: function isVerifiedEntity(address entityAddress) view returns(bool)
func (_EntityRegistry *EntityRegistryCaller) IsVerifiedEntity(opts *bind.CallOpts, entityAddress common.Address) (bool, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "isVerifiedEntity", entityAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifiedEntity is a free data retrieval call binding the contract method 0x455f1f36.
//
// Solidity: function isVerifiedEntity(address entityAddress) view returns(bool)
func (_EntityRegistry *EntityRegistrySession) IsVerifiedEntity(entityAddress common.Address) (bool, error) {
	return _EntityRegistry.Contract.IsVerifiedEntity(&_EntityRegistry.CallOpts, entityAddress)
}

// IsVerifiedEntity is a free data retrieval call binding the contract method 0x455f1f36.
//
// Solidity: function isVerifiedEntity(address entityAddress) view returns(bool)
func (_EntityRegistry *EntityRegistryCallerSession) IsVerifiedEntity(entityAddress common.Address) (bool, error) {
	return _EntityRegistry.Contract.IsVerifiedEntity(&_EntityRegistry.CallOpts, entityAddress)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntityRegistry *EntityRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntityRegistry *EntityRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EntityRegistry.Contract.SupportsInterface(&_EntityRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EntityRegistry *EntityRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EntityRegistry.Contract.SupportsInterface(&_EntityRegistry.CallOpts, interfaceId)
}

// VerifyInfo is a free data retrieval call binding the contract method 0xa1f06cf5.
//
// Solidity: function verifyInfo((address,uint8,bytes,address) entity, bytes32 hashedInfo, bytes32[] proof) view returns(bool)
func (_EntityRegistry *EntityRegistryCaller) VerifyInfo(opts *bind.CallOpts, entity Entity, hashedInfo [32]byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _EntityRegistry.contract.Call(opts, &out, "verifyInfo", entity, hashedInfo, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyInfo is a free data retrieval call binding the contract method 0xa1f06cf5.
//
// Solidity: function verifyInfo((address,uint8,bytes,address) entity, bytes32 hashedInfo, bytes32[] proof) view returns(bool)
func (_EntityRegistry *EntityRegistrySession) VerifyInfo(entity Entity, hashedInfo [32]byte, proof [][32]byte) (bool, error) {
	return _EntityRegistry.Contract.VerifyInfo(&_EntityRegistry.CallOpts, entity, hashedInfo, proof)
}

// VerifyInfo is a free data retrieval call binding the contract method 0xa1f06cf5.
//
// Solidity: function verifyInfo((address,uint8,bytes,address) entity, bytes32 hashedInfo, bytes32[] proof) view returns(bool)
func (_EntityRegistry *EntityRegistryCallerSession) VerifyInfo(entity Entity, hashedInfo [32]byte, proof [][32]byte) (bool, error) {
	return _EntityRegistry.Contract.VerifyInfo(&_EntityRegistry.CallOpts, entity, hashedInfo, proof)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x828586c8.
//
// Solidity: function addVerifier(address verifier, uint8[] entityTypes) returns()
func (_EntityRegistry *EntityRegistryTransactor) AddVerifier(opts *bind.TransactOpts, verifier common.Address, entityTypes []uint8) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "addVerifier", verifier, entityTypes)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x828586c8.
//
// Solidity: function addVerifier(address verifier, uint8[] entityTypes) returns()
func (_EntityRegistry *EntityRegistrySession) AddVerifier(verifier common.Address, entityTypes []uint8) (*types.Transaction, error) {
	return _EntityRegistry.Contract.AddVerifier(&_EntityRegistry.TransactOpts, verifier, entityTypes)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x828586c8.
//
// Solidity: function addVerifier(address verifier, uint8[] entityTypes) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) AddVerifier(verifier common.Address, entityTypes []uint8) (*types.Transaction, error) {
	return _EntityRegistry.Contract.AddVerifier(&_EntityRegistry.TransactOpts, verifier, entityTypes)
}

// ForwardRegister is a paid mutator transaction binding the contract method 0x2160ebde.
//
// Solidity: function forwardRegister((address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_EntityRegistry *EntityRegistryTransactor) ForwardRegister(opts *bind.TransactOpts, entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "forwardRegister", entity, verifierSignature)
}

// ForwardRegister is a paid mutator transaction binding the contract method 0x2160ebde.
//
// Solidity: function forwardRegister((address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_EntityRegistry *EntityRegistrySession) ForwardRegister(entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _EntityRegistry.Contract.ForwardRegister(&_EntityRegistry.TransactOpts, entity, verifierSignature)
}

// ForwardRegister is a paid mutator transaction binding the contract method 0x2160ebde.
//
// Solidity: function forwardRegister((address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) ForwardRegister(entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _EntityRegistry.Contract.ForwardRegister(&_EntityRegistry.TransactOpts, entity, verifierSignature)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EntityRegistry *EntityRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EntityRegistry *EntityRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.GrantRole(&_EntityRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.GrantRole(&_EntityRegistry.TransactOpts, role, account)
}

// Register is a paid mutator transaction binding the contract method 0x01296ecf.
//
// Solidity: function register((address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_EntityRegistry *EntityRegistryTransactor) Register(opts *bind.TransactOpts, entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "register", entity, verifierSignature)
}

// Register is a paid mutator transaction binding the contract method 0x01296ecf.
//
// Solidity: function register((address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_EntityRegistry *EntityRegistrySession) Register(entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _EntityRegistry.Contract.Register(&_EntityRegistry.TransactOpts, entity, verifierSignature)
}

// Register is a paid mutator transaction binding the contract method 0x01296ecf.
//
// Solidity: function register((address,uint8,bytes,address) entity, bytes verifierSignature) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) Register(entity Entity, verifierSignature []byte) (*types.Transaction, error) {
	return _EntityRegistry.Contract.Register(&_EntityRegistry.TransactOpts, entity, verifierSignature)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_EntityRegistry *EntityRegistryTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_EntityRegistry *EntityRegistrySession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.RemoveVerifier(&_EntityRegistry.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.RemoveVerifier(&_EntityRegistry.TransactOpts, verifier)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EntityRegistry *EntityRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EntityRegistry *EntityRegistrySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.RenounceRole(&_EntityRegistry.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.RenounceRole(&_EntityRegistry.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EntityRegistry *EntityRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EntityRegistry *EntityRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.RevokeRole(&_EntityRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _EntityRegistry.Contract.RevokeRole(&_EntityRegistry.TransactOpts, role, account)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x520c886c.
//
// Solidity: function updateVerifier(address verifier, uint8[] entityTypes) returns()
func (_EntityRegistry *EntityRegistryTransactor) UpdateVerifier(opts *bind.TransactOpts, verifier common.Address, entityTypes []uint8) (*types.Transaction, error) {
	return _EntityRegistry.contract.Transact(opts, "updateVerifier", verifier, entityTypes)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x520c886c.
//
// Solidity: function updateVerifier(address verifier, uint8[] entityTypes) returns()
func (_EntityRegistry *EntityRegistrySession) UpdateVerifier(verifier common.Address, entityTypes []uint8) (*types.Transaction, error) {
	return _EntityRegistry.Contract.UpdateVerifier(&_EntityRegistry.TransactOpts, verifier, entityTypes)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x520c886c.
//
// Solidity: function updateVerifier(address verifier, uint8[] entityTypes) returns()
func (_EntityRegistry *EntityRegistryTransactorSession) UpdateVerifier(verifier common.Address, entityTypes []uint8) (*types.Transaction, error) {
	return _EntityRegistry.Contract.UpdateVerifier(&_EntityRegistry.TransactOpts, verifier, entityTypes)
}

// EntityRegistryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the EntityRegistry contract.
type EntityRegistryEIP712DomainChangedIterator struct {
	Event *EntityRegistryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *EntityRegistryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryEIP712DomainChanged)
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
		it.Event = new(EntityRegistryEIP712DomainChanged)
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
func (it *EntityRegistryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the EntityRegistry contract.
type EntityRegistryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EntityRegistry *EntityRegistryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*EntityRegistryEIP712DomainChangedIterator, error) {

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &EntityRegistryEIP712DomainChangedIterator{contract: _EntityRegistry.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EntityRegistry *EntityRegistryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *EntityRegistryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryEIP712DomainChanged)
				if err := _EntityRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_EntityRegistry *EntityRegistryFilterer) ParseEIP712DomainChanged(log types.Log) (*EntityRegistryEIP712DomainChanged, error) {
	event := new(EntityRegistryEIP712DomainChanged)
	if err := _EntityRegistry.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryEntityRegisteredIterator is returned from FilterEntityRegistered and is used to iterate over the raw logs and unpacked data for EntityRegistered events raised by the EntityRegistry contract.
type EntityRegistryEntityRegisteredIterator struct {
	Event *EntityRegistryEntityRegistered // Event containing the contract specifics and raw log

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
func (it *EntityRegistryEntityRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryEntityRegistered)
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
		it.Event = new(EntityRegistryEntityRegistered)
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
func (it *EntityRegistryEntityRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryEntityRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryEntityRegistered represents a EntityRegistered event raised by the EntityRegistry contract.
type EntityRegistryEntityRegistered struct {
	EntityAddress common.Address
	Entity        Entity
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEntityRegistered is a free log retrieval operation binding the contract event 0x5cda68d609c87809a2b53cb1ae9549a108342cf2fcae801b8aa033e9ca49453b.
//
// Solidity: event EntityRegistered(address indexed entityAddress, (address,uint8,bytes,address) entity)
func (_EntityRegistry *EntityRegistryFilterer) FilterEntityRegistered(opts *bind.FilterOpts, entityAddress []common.Address) (*EntityRegistryEntityRegisteredIterator, error) {

	var entityAddressRule []interface{}
	for _, entityAddressItem := range entityAddress {
		entityAddressRule = append(entityAddressRule, entityAddressItem)
	}

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "EntityRegistered", entityAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryEntityRegisteredIterator{contract: _EntityRegistry.contract, event: "EntityRegistered", logs: logs, sub: sub}, nil
}

// WatchEntityRegistered is a free log subscription operation binding the contract event 0x5cda68d609c87809a2b53cb1ae9549a108342cf2fcae801b8aa033e9ca49453b.
//
// Solidity: event EntityRegistered(address indexed entityAddress, (address,uint8,bytes,address) entity)
func (_EntityRegistry *EntityRegistryFilterer) WatchEntityRegistered(opts *bind.WatchOpts, sink chan<- *EntityRegistryEntityRegistered, entityAddress []common.Address) (event.Subscription, error) {

	var entityAddressRule []interface{}
	for _, entityAddressItem := range entityAddress {
		entityAddressRule = append(entityAddressRule, entityAddressItem)
	}

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "EntityRegistered", entityAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryEntityRegistered)
				if err := _EntityRegistry.contract.UnpackLog(event, "EntityRegistered", log); err != nil {
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

// ParseEntityRegistered is a log parse operation binding the contract event 0x5cda68d609c87809a2b53cb1ae9549a108342cf2fcae801b8aa033e9ca49453b.
//
// Solidity: event EntityRegistered(address indexed entityAddress, (address,uint8,bytes,address) entity)
func (_EntityRegistry *EntityRegistryFilterer) ParseEntityRegistered(log types.Log) (*EntityRegistryEntityRegistered, error) {
	event := new(EntityRegistryEntityRegistered)
	if err := _EntityRegistry.contract.UnpackLog(event, "EntityRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the EntityRegistry contract.
type EntityRegistryRoleAdminChangedIterator struct {
	Event *EntityRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EntityRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryRoleAdminChanged)
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
		it.Event = new(EntityRegistryRoleAdminChanged)
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
func (it *EntityRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the EntityRegistry contract.
type EntityRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EntityRegistry *EntityRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EntityRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryRoleAdminChangedIterator{contract: _EntityRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_EntityRegistry *EntityRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EntityRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryRoleAdminChanged)
				if err := _EntityRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_EntityRegistry *EntityRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*EntityRegistryRoleAdminChanged, error) {
	event := new(EntityRegistryRoleAdminChanged)
	if err := _EntityRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the EntityRegistry contract.
type EntityRegistryRoleGrantedIterator struct {
	Event *EntityRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *EntityRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryRoleGranted)
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
		it.Event = new(EntityRegistryRoleGranted)
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
func (it *EntityRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryRoleGranted represents a RoleGranted event raised by the EntityRegistry contract.
type EntityRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EntityRegistry *EntityRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EntityRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryRoleGrantedIterator{contract: _EntityRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_EntityRegistry *EntityRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EntityRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryRoleGranted)
				if err := _EntityRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_EntityRegistry *EntityRegistryFilterer) ParseRoleGranted(log types.Log) (*EntityRegistryRoleGranted, error) {
	event := new(EntityRegistryRoleGranted)
	if err := _EntityRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the EntityRegistry contract.
type EntityRegistryRoleRevokedIterator struct {
	Event *EntityRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EntityRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryRoleRevoked)
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
		it.Event = new(EntityRegistryRoleRevoked)
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
func (it *EntityRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryRoleRevoked represents a RoleRevoked event raised by the EntityRegistry contract.
type EntityRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EntityRegistry *EntityRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EntityRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryRoleRevokedIterator{contract: _EntityRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_EntityRegistry *EntityRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EntityRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryRoleRevoked)
				if err := _EntityRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_EntityRegistry *EntityRegistryFilterer) ParseRoleRevoked(log types.Log) (*EntityRegistryRoleRevoked, error) {
	event := new(EntityRegistryRoleRevoked)
	if err := _EntityRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryVerifierAddedIterator is returned from FilterVerifierAdded and is used to iterate over the raw logs and unpacked data for VerifierAdded events raised by the EntityRegistry contract.
type EntityRegistryVerifierAddedIterator struct {
	Event *EntityRegistryVerifierAdded // Event containing the contract specifics and raw log

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
func (it *EntityRegistryVerifierAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryVerifierAdded)
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
		it.Event = new(EntityRegistryVerifierAdded)
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
func (it *EntityRegistryVerifierAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryVerifierAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryVerifierAdded represents a VerifierAdded event raised by the EntityRegistry contract.
type EntityRegistryVerifierAdded struct {
	Verifier    common.Address
	EntityTypes []uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerifierAdded is a free log retrieval operation binding the contract event 0x398678206f3910abe75c3ab585dd972e3197332691c7e71f2bca55c03aa01bbb.
//
// Solidity: event VerifierAdded(address indexed verifier, uint8[] entityTypes)
func (_EntityRegistry *EntityRegistryFilterer) FilterVerifierAdded(opts *bind.FilterOpts, verifier []common.Address) (*EntityRegistryVerifierAddedIterator, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "VerifierAdded", verifierRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryVerifierAddedIterator{contract: _EntityRegistry.contract, event: "VerifierAdded", logs: logs, sub: sub}, nil
}

// WatchVerifierAdded is a free log subscription operation binding the contract event 0x398678206f3910abe75c3ab585dd972e3197332691c7e71f2bca55c03aa01bbb.
//
// Solidity: event VerifierAdded(address indexed verifier, uint8[] entityTypes)
func (_EntityRegistry *EntityRegistryFilterer) WatchVerifierAdded(opts *bind.WatchOpts, sink chan<- *EntityRegistryVerifierAdded, verifier []common.Address) (event.Subscription, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "VerifierAdded", verifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryVerifierAdded)
				if err := _EntityRegistry.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
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

// ParseVerifierAdded is a log parse operation binding the contract event 0x398678206f3910abe75c3ab585dd972e3197332691c7e71f2bca55c03aa01bbb.
//
// Solidity: event VerifierAdded(address indexed verifier, uint8[] entityTypes)
func (_EntityRegistry *EntityRegistryFilterer) ParseVerifierAdded(log types.Log) (*EntityRegistryVerifierAdded, error) {
	event := new(EntityRegistryVerifierAdded)
	if err := _EntityRegistry.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryVerifierRemovedIterator is returned from FilterVerifierRemoved and is used to iterate over the raw logs and unpacked data for VerifierRemoved events raised by the EntityRegistry contract.
type EntityRegistryVerifierRemovedIterator struct {
	Event *EntityRegistryVerifierRemoved // Event containing the contract specifics and raw log

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
func (it *EntityRegistryVerifierRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryVerifierRemoved)
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
		it.Event = new(EntityRegistryVerifierRemoved)
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
func (it *EntityRegistryVerifierRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryVerifierRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryVerifierRemoved represents a VerifierRemoved event raised by the EntityRegistry contract.
type EntityRegistryVerifierRemoved struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierRemoved is a free log retrieval operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address indexed verifier)
func (_EntityRegistry *EntityRegistryFilterer) FilterVerifierRemoved(opts *bind.FilterOpts, verifier []common.Address) (*EntityRegistryVerifierRemovedIterator, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "VerifierRemoved", verifierRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryVerifierRemovedIterator{contract: _EntityRegistry.contract, event: "VerifierRemoved", logs: logs, sub: sub}, nil
}

// WatchVerifierRemoved is a free log subscription operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address indexed verifier)
func (_EntityRegistry *EntityRegistryFilterer) WatchVerifierRemoved(opts *bind.WatchOpts, sink chan<- *EntityRegistryVerifierRemoved, verifier []common.Address) (event.Subscription, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "VerifierRemoved", verifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryVerifierRemoved)
				if err := _EntityRegistry.contract.UnpackLog(event, "VerifierRemoved", log); err != nil {
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

// ParseVerifierRemoved is a log parse operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address indexed verifier)
func (_EntityRegistry *EntityRegistryFilterer) ParseVerifierRemoved(log types.Log) (*EntityRegistryVerifierRemoved, error) {
	event := new(EntityRegistryVerifierRemoved)
	if err := _EntityRegistry.contract.UnpackLog(event, "VerifierRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityRegistryVerifierUpdatedIterator is returned from FilterVerifierUpdated and is used to iterate over the raw logs and unpacked data for VerifierUpdated events raised by the EntityRegistry contract.
type EntityRegistryVerifierUpdatedIterator struct {
	Event *EntityRegistryVerifierUpdated // Event containing the contract specifics and raw log

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
func (it *EntityRegistryVerifierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityRegistryVerifierUpdated)
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
		it.Event = new(EntityRegistryVerifierUpdated)
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
func (it *EntityRegistryVerifierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityRegistryVerifierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityRegistryVerifierUpdated represents a VerifierUpdated event raised by the EntityRegistry contract.
type EntityRegistryVerifierUpdated struct {
	Verifier    common.Address
	EntityTypes []uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerifierUpdated is a free log retrieval operation binding the contract event 0x6c0fb7a51dc2766276e13dc1531bda525f4c58549cd04cb930545da96eb50b6b.
//
// Solidity: event VerifierUpdated(address indexed verifier, uint8[] entityTypes)
func (_EntityRegistry *EntityRegistryFilterer) FilterVerifierUpdated(opts *bind.FilterOpts, verifier []common.Address) (*EntityRegistryVerifierUpdatedIterator, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _EntityRegistry.contract.FilterLogs(opts, "VerifierUpdated", verifierRule)
	if err != nil {
		return nil, err
	}
	return &EntityRegistryVerifierUpdatedIterator{contract: _EntityRegistry.contract, event: "VerifierUpdated", logs: logs, sub: sub}, nil
}

// WatchVerifierUpdated is a free log subscription operation binding the contract event 0x6c0fb7a51dc2766276e13dc1531bda525f4c58549cd04cb930545da96eb50b6b.
//
// Solidity: event VerifierUpdated(address indexed verifier, uint8[] entityTypes)
func (_EntityRegistry *EntityRegistryFilterer) WatchVerifierUpdated(opts *bind.WatchOpts, sink chan<- *EntityRegistryVerifierUpdated, verifier []common.Address) (event.Subscription, error) {

	var verifierRule []interface{}
	for _, verifierItem := range verifier {
		verifierRule = append(verifierRule, verifierItem)
	}

	logs, sub, err := _EntityRegistry.contract.WatchLogs(opts, "VerifierUpdated", verifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityRegistryVerifierUpdated)
				if err := _EntityRegistry.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
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

// ParseVerifierUpdated is a log parse operation binding the contract event 0x6c0fb7a51dc2766276e13dc1531bda525f4c58549cd04cb930545da96eb50b6b.
//
// Solidity: event VerifierUpdated(address indexed verifier, uint8[] entityTypes)
func (_EntityRegistry *EntityRegistryFilterer) ParseVerifierUpdated(log types.Log) (*EntityRegistryVerifierUpdated, error) {
	event := new(EntityRegistryVerifierUpdated)
	if err := _EntityRegistry.contract.UnpackLog(event, "VerifierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
