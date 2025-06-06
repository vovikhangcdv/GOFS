package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/entity_registry"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/compliance_registry"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/compliant_token"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/exchange_portal"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/erc20"
)

type TxConf struct {
	Rpc      *rpc.Client
	Nonce    uint64
	Sender   common.Address
	To       *common.Address
	Value    *big.Int
	GasLimit uint64
	GasPrice *big.Int
	ChainID  *big.Int
	Code     []byte
}

type TxType struct {
	Type string
	Weight int
}

type SystemContracts struct {
	EntityRegistry *entityRegistry.EntityRegistry
	EntityRegistryAddress common.Address
	ComplianceRegistry *complianceRegistry.ComplianceRegistry
	ComplianceRegistryAddress common.Address
	EVNDToken *compliantToken.CompliantToken
	EVNDTokenAddress common.Address
	USDToken *erc20.ERC20
	USDTokenAddress common.Address
	ExchangePortal *exchangePortal.ExchangePortal
	ExchangePortalAddress common.Address
}