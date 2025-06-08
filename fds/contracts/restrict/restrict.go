package restrict

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Restrict represents the Restrict contract
type Restrict struct {
	RestrictCaller     // Read-only binding to the contract
	RestrictTransactor // Write-only binding to the contract
	RestrictFilterer   // Log filterer for contract events
	Client             *ethclient.Client
}

// RestrictCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictCaller struct {
	contract *bind.BoundContract
}

// RestrictTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictTransactor struct {
	contract *bind.BoundContract
}

// RestrictFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFilterer struct {
	contract *bind.BoundContract
}

// NewRestrict creates a new instance of Restrict, bound to a specific deployed contract.
func NewRestrict(address common.Address, backend bind.ContractBackend) (*Restrict, error) {
	contract, err := bindRestrict(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Restrict{
		RestrictCaller:     RestrictCaller{contract: contract},
		RestrictTransactor: RestrictTransactor{contract: contract},
		RestrictFilterer:   RestrictFilterer{contract: contract},
		Client:             backend.(*ethclient.Client),
	}, nil
}

// bindRestrict binds a generic wrapper to an already deployed contract.
func bindRestrict(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Blacklist is a paid mutator transaction binding the contract method.
func (r *RestrictTransactor) Blacklist(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return r.contract.Transact(opts, "blacklist", accounts)
}

// RestrictABI contains the ABI for the Restrict contract
const RestrictABI = `[
	{
		"inputs": [{"internalType": "address[]", "name": "accounts", "type": "address[]"}],
		"name": "blacklist",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [{"indexed": true, "internalType": "address", "name": "account", "type": "address"}],
		"name": "AddressBlacklisted",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [{"indexed": true, "internalType": "address", "name": "account", "type": "address"}],
		"name": "AddressBlacklistedFrom",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [{"indexed": true, "internalType": "address", "name": "account", "type": "address"}],
		"name": "AddressBlacklistedTo",
		"type": "event"
	}
]`
