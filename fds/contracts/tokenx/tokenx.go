package tokenx

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenX represents the TokenX contract
type TokenX struct {
	TokenXCaller     // Read-only binding to the contract
	TokenXTransactor // Write-only binding to the contract
	TokenXFilterer   // Log filterer for contract events
}

// TokenXCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenXCaller struct {
	contract *bind.BoundContract
}

// TokenXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenXTransactor struct {
	contract *bind.BoundContract
}

// TokenXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenXFilterer struct {
	contract *bind.BoundContract
}

// NewTokenX creates a new instance of TokenX, bound to a specific deployed contract.
func NewTokenX(address common.Address, backend bind.ContractBackend) (*TokenX, error) {
	contract, err := bindTokenX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenX{TokenXCaller: TokenXCaller{contract: contract}, TokenXTransactor: TokenXTransactor{contract: contract}, TokenXFilterer: TokenXFilterer{contract: contract}}, nil
}

// bindTokenX binds a generic wrapper to an already deployed contract.
func bindTokenX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenXABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// AddToBlacklist is a paid mutator transaction binding the contract method.
func (t *TokenXTransactor) AddToBlacklist(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return t.contract.Transact(opts, "addToBlacklist", account)
}

// TokenXABI contains the ABI for the TokenX contract
const TokenXABI = `[
	{
		"inputs": [{"internalType": "address", "name": "account", "type": "address"}],
		"name": "addToBlacklist",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address[]", "name": "accounts", "type": "address[]"}],
		"name": "blacklistFrom",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// BlacklistFrom is a paid mutator transaction binding the contract method 0x12345678.
//
// Solidity: function blacklistFrom(address[] memory accounts) external onlyRole(BLACKLIST_ADMIN_ROLE)
func (_TokenX *TokenXTransactor) BlacklistFrom(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _TokenX.contract.Transact(opts, "blacklistFrom", accounts)
}
