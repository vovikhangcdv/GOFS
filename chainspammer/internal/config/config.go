package config

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/urfave/cli/v2"
	complianceRegistry "github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/compliance_registry"
	compliantToken "github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/compliant_token"
	entityRegistry "github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/entity_registry"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/erc20"
	exchangePortal "github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/exchange_portal"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/types"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
)

type Config struct {
	Client          *ethclient.Client
	Backend         *rpc.Client
	ChainID         *big.Int
	Faucet          *ecdsa.PrivateKey
	Verifiers       []*ecdsa.PrivateKey
	Blacklisters    []*ecdsa.PrivateKey
	Keys            []*ecdsa.PrivateKey
	MaxKeys         int
	DelayTime       int64
	Wallet          *hdwallet.Wallet
	Seed            int64
	AdminKey        *ecdsa.PrivateKey
	SystemContracts types.SystemContracts
	// Use to count number of private keys derived from Wallet
	AddressCounter int
	// Transaction type weights
	RegisterEntityWeight int
	SendEVNDWeight       int
	ExchangeVNDUSDWeight int
}

func (c *Config) GetRandomVerifier() *ecdsa.PrivateKey {
	return c.Verifiers[utils.RandomIdx(len(c.Verifiers))]
}

func (c *Config) GetRandomKey() *ecdsa.PrivateKey {
	return c.Keys[utils.RandomIdx(len(c.Keys))]
}

func (c *Config) GetNewKey() (*ecdsa.PrivateKey, error) {
	if len(c.Keys) >= c.MaxKeys {
		return nil, errors.New("max users reached")
	}
	key, err := utils.GetNextPrivateKey(c.Wallet, &c.AddressCounter)
	if err != nil {
		return nil, err
	}
	c.Keys = append(c.Keys, key)
	return key, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// Use later
func NewConfigFromEnv() *Config {
	rpcUrl := getEnv("RPC_URL", "http://localhost:8545")
	backend, err := rpc.Dial(rpcUrl)
	client := ethclient.NewClient(backend)
	if err != nil {
		panic(err)
	}
	faucetSk, _ := crypto.HexToECDSA(getEnv("FAUCET_SK", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	wallet, _ := hdwallet.NewFromMnemonic(getEnv("WALLET_MEMO", "test test test test test test test test test test test junk"))
	cnt := 0
	adminKey, _ := utils.GetNextPrivateKey(wallet, &cnt)
	entityRegistryAddress := common.HexToAddress(getEnv("ENTITY_REGISTRY", "0xD8a5a9b31c3C0232E196d518E89Fd8bF83AcAd43"))
	complianceRegistryAddress := common.HexToAddress(getEnv("COMPLIANCE_REGISTRY", "0x36b58F5C1969B7b6591D752ea6F5486D069010AB"))
	eVNDTokenAddress := common.HexToAddress(getEnv("EVND_TOKEN", "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"))
	mUSDTokenAddress := common.HexToAddress(getEnv("MUSD_TOKEN", "0x9A676e781A523b5d0C0e43731313A708CB607508"))
	exchangePortalAddress := common.HexToAddress(getEnv("EXCHANGE_PORTAL", "0xC9a43158891282A2B1475592D5719c001986Aaec"))
	entityRegistry, _ := entityRegistry.NewEntityRegistry(entityRegistryAddress, client)
	complianceRegistry, _ := complianceRegistry.NewComplianceRegistry(complianceRegistryAddress, client)
	eVNDToken, _ := compliantToken.NewCompliantToken(eVNDTokenAddress, client)
	mUSDToken, _ := erc20.NewERC20(mUSDTokenAddress, client)
	exchangePortal, _ := exchangePortal.NewExchangePortal(exchangePortalAddress, client)
	systemContracts := types.SystemContracts{
		EntityRegistry:            entityRegistry,
		EntityRegistryAddress:     entityRegistryAddress,
		ComplianceRegistry:        complianceRegistry,
		ComplianceRegistryAddress: complianceRegistryAddress,
		EVNDToken:                 eVNDToken,
		EVNDTokenAddress:          eVNDTokenAddress,
		USDToken:                  mUSDToken,
		USDTokenAddress:           mUSDTokenAddress,
		ExchangePortal:            exchangePortal,
		ExchangePortalAddress:     exchangePortalAddress,
	}
	config := &Config{
		Client:               client,
		Backend:              backend,
		ChainID:              utils.GetChainID(backend),
		Faucet:               faucetSk,
		Seed:                 int64(getEnvAsInt("SEED", 0)),
		Keys:                 utils.RandomSks(getEnvAsInt("MAX_KEYS", 10)),
		MaxKeys:              getEnvAsInt("MAX_KEYS", 100),
		DelayTime:            int64(getEnvAsInt("DELAY_TIME", 3)),
		Wallet:               wallet,
		AdminKey:             adminKey,
		SystemContracts:      systemContracts,
		AddressCounter:       cnt,
		RegisterEntityWeight: 1,
		SendEVNDWeight:       1,
		ExchangeVNDUSDWeight: 1,
	}
	return config
}

func NewConfigFromContext(c *cli.Context) (*Config, error) {
	rpcUrl := c.String("rpc")
	backend, err := rpc.Dial(rpcUrl)
	client := ethclient.NewClient(backend)
	if err != nil {
		return nil, err
	}
	faucetSk, err := crypto.HexToECDSA(c.String("faucet-sk"))
	if err != nil {
		return nil, err
	}
	wallet, err := hdwallet.NewFromMnemonic(c.String("mnemonic"))
	if err != nil {
		return nil, err
	}
	cnt := 0
	adminKey, err := utils.GetNextPrivateKey(wallet, &cnt)
	if err != nil {
		return nil, err
	}
	verifiers := make([]*ecdsa.PrivateKey, 2)
	for i := 0; i < 2; i++ {
		verifier, err := utils.GetNextPrivateKey(wallet, &cnt)
		if err != nil {
			return nil, err
		}
		verifiers[i] = verifier
	}
	blacklisters := make([]*ecdsa.PrivateKey, 2)
	for i := 0; i < 2; i++ {
		blacklister, err := utils.GetNextPrivateKey(wallet, &cnt)
		if err != nil {
			return nil, err
		}
		blacklisters[i] = blacklister
	}

	keys := make([]*ecdsa.PrivateKey, 0, c.Int("max-keys"))

	entityRegistryAddress := common.HexToAddress(c.String("entity-registry"))
	complianceRegistryAddress := common.HexToAddress(c.String("compliance-registry"))
	eVNDTokenAddress := common.HexToAddress(c.String("evnd-token"))
	mUSDTokenAddress := common.HexToAddress(c.String("musd-token"))
	exchangePortalAddress := common.HexToAddress(c.String("exchange-portal"))

	entityRegistry, err := entityRegistry.NewEntityRegistry(entityRegistryAddress, client)
	if err != nil {
		return nil, err
	}
	complianceRegistry, err := complianceRegistry.NewComplianceRegistry(complianceRegistryAddress, client)
	if err != nil {
		return nil, err
	}
	eVNDToken, err := compliantToken.NewCompliantToken(eVNDTokenAddress, client)
	if err != nil {
		return nil, err
	}
	mUSDToken, err := erc20.NewERC20(mUSDTokenAddress, client)
	if err != nil {
		return nil, err
	}
	exchangePortal, err := exchangePortal.NewExchangePortal(exchangePortalAddress, client)
	if err != nil {
		return nil, err
	}

	systemContracts := types.SystemContracts{
		EntityRegistry:            entityRegistry,
		EntityRegistryAddress:     entityRegistryAddress,
		ComplianceRegistry:        complianceRegistry,
		ComplianceRegistryAddress: complianceRegistryAddress,
		EVNDToken:                 eVNDToken,
		EVNDTokenAddress:          eVNDTokenAddress,
		USDToken:                  mUSDToken,
		USDTokenAddress:           mUSDTokenAddress,
		ExchangePortal:            exchangePortal,
		ExchangePortalAddress:     exchangePortalAddress,
	}
	config := &Config{
		Client:               client,
		Backend:              backend,
		ChainID:              utils.GetChainID(backend),
		DelayTime:            c.Int64("delay-time"),
		Faucet:               faucetSk,
		Seed:                 c.Int64("seed"),
		Keys:                 keys,
		MaxKeys:              c.Int("max-keys"),
		Verifiers:            verifiers,
		Blacklisters:         blacklisters,
		Wallet:               wallet,
		AdminKey:             adminKey,
		SystemContracts:      systemContracts,
		AddressCounter:       cnt,
		RegisterEntityWeight: c.Int("register-entity-weight"),
		SendEVNDWeight:       c.Int("send-evnd-weight"),
		ExchangeVNDUSDWeight: c.Int("exchange-vnd-usd-weight"),
	}
	return config, nil
}
