package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"slices"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
	entityRegistry "github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/entity_registry"
	localtypes "github.com/vovikhangcdv/GOFS/chainspammer/internal/types"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
)

var (
	ENTITY_TYPEHASH = crypto.Keccak256Hash([]byte("Entity(address entityAddress,uint8 entityType,bytes entityData,address verifier)"))
	typeBytes32, _  = abi.NewType("bytes32", "", nil)
	typeUint8, _    = abi.NewType("uint8", "", nil)
	typeAddress, _  = abi.NewType("address", "", nil)
	typeBytes, _    = abi.NewType("bytes", "", nil)
	typeString, _   = abi.NewType("string", "", nil)
)

func AirdropGas(config *config.Config, addr common.Address) (common.Hash, error) {
	airdropValue := big.NewInt(int64(float64(params.Ether) * 0.1))
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, config.Faucet, addr, airdropValue, 21_000, nil, false)
	if err != nil {
		return common.Hash{}, err
	}
	// Convert to float for better readability
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(airdropValue), new(big.Float).SetInt(big.NewInt(params.Ether)))
	log.Println("Airdropped ", ethValue.Text('f', 6), " ETH to ", addr.Hex(), " tx_hash: ", txHash.Hex())
	return txHash, nil
}

func AirdropGasIfNeeded(config *config.Config, addr common.Address) (common.Hash, error) {
	balance, err := config.Client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return common.Hash{}, err
	}
	if balance.Cmp(big.NewInt(int64(float64(params.Ether)*0.1))) == 0 {
		return AirdropGas(config, addr)
	}
	return common.Hash{}, nil
}

func AirdropVND(config *config.Config, receiver common.Address, airdropValue *big.Int) (common.Hash, error) {
	opts := &bind.TransactOpts{
		GasLimit: 1_000_000,
		NoSend:   true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	tx, err := config.SystemContracts.EVNDToken.Mint(opts, receiver, airdropValue)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, config.AdminKey, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), false)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Minted ", airdropValue.String(), " VND to ", receiver.Hex(), " tx hash: ", txHash.Hex())
	return txHash, nil
}

func AirdropVNDIfNeeded(config *config.Config, receiver common.Address) (*big.Int, error) {
	balance, err := config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, receiver)
	if err != nil {
		return nil, err
	}
	if balance.Cmp(big.NewInt(0)) != 0 {
		return balance, nil
	}
	txHash, err := AirdropVND(config, receiver, big.NewInt(10_000_000))
	if err != nil {
		return nil, err
	}
	log.Println("Airdropped eVND to ", receiver.Hex(), " tx hash: ", txHash.Hex())
	balance, err = config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, receiver)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func SendRegisterEntityTx(config *config.Config, skUser *ecdsa.PrivateKey, isUseRPC bool) (common.Hash, error) {
	skVerifier := config.GetRandomVerifier()
	isVerified, err := config.SystemContracts.EntityRegistry.IsVerifiedEntity(&bind.CallOpts{}, crypto.PubkeyToAddress(skUser.PublicKey))
	if err != nil {
		return common.Hash{}, err
	}
	if isVerified {
		log.Println("✅ User already verified, skipping")
		return common.Hash{}, nil
	}

	if _, err := AirdropGas(config, crypto.PubkeyToAddress(skUser.PublicKey)); err != nil {
		return common.Hash{}, err
	}
	opts := &bind.TransactOpts{
		GasLimit: 500_000,
		NoSend:   true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	tx, entityData, entityType, err := createRegisterTx(config, opts, skVerifier, skUser)
	if err != nil {
		return common.Hash{}, err
	}

	msg := ethereum.CallMsg{
		From: crypto.PubkeyToAddress(skUser.PublicKey),
		To:   &config.SystemContracts.EntityRegistryAddress,
		Data: tx.Data(),
	}
	gas, err := config.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, skUser, *tx.To(), tx.Value(), gas, tx.Data(), isUseRPC)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Registered Address: ", crypto.PubkeyToAddress(skUser.PublicKey).Hex(), " entity_type: ", entityType, " tx_hash: ", txHash.Hex())

	log.Println("Entity Info: Name=", entityData.Name, " IDNumber=", entityData.IDNumber, " Birthday=", entityData.Birthday, " Gender=", entityData.Gender, " Email=", entityData.Email, " Phone=", entityData.Phone, " Address=", entityData.Address, " Nationality=", entityData.Nationality, " Others=", entityData.Others)

	if _, err := AirdropVND(config, crypto.PubkeyToAddress(skUser.PublicKey), big.NewInt(10_000_000)); err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func SendExchangeVNDToUSDTx(config *config.Config, skFrom *ecdsa.PrivateKey, isUseRPC bool) (common.Hash, error) {
	addr := crypto.PubkeyToAddress(skFrom.PublicKey)
	if _, err := AirdropGas(config, addr); err != nil {
		return common.Hash{}, err
	}
	if !IsTransferBetweenAddressesAllowed(config, addr, config.SystemContracts.ExchangePortalAddress) {
		return common.Hash{}, fmt.Errorf("transfer not allowed between addresses, skipping")
	}
	// Temporary check if UNKNOWN_TX is allowed
	allowedTransactionTypes := GetAllowedTransactionTypes(config, addr, config.SystemContracts.ExchangePortalAddress)
	isUnknownTxAllowed := slices.Contains(allowedTransactionTypes, utils.UNKNOWN_TX)
	if !isUnknownTxAllowed {
		return common.Hash{}, fmt.Errorf("UNKNOWN_TX is not allowed, skipping")
	}

	verifiered, err := config.SystemContracts.EntityRegistry.IsVerifiedEntity(&bind.CallOpts{}, addr)
	if err != nil {
		return common.Hash{}, err
	}
	if !verifiered {
		return common.Hash{}, fmt.Errorf("address is not verified, skipping")
	}
	if _, err := AirdropGasIfNeeded(config, addr); err != nil {
		return common.Hash{}, err
	}
	fromBalance, err := AirdropVNDIfNeeded(config, addr)
	if err != nil {
		log.Println("❌ Error while funding eVND: ", err)
		return common.Hash{}, err
	}
	randomValue := utils.RandomBigInt(fromBalance)
	for !isValidAmount(config,
		config.SystemContracts.EVNDTokenAddress,
		config.SystemContracts.USDTokenAddress,
		randomValue,
	) {
		randomValue = utils.RandomBigInt(fromBalance)
	}

	err = SendApproveTxIfNeeded(config,
		skFrom,
		config.SystemContracts.ExchangePortalAddress,
		config.SystemContracts.EVNDTokenAddress,
		randomValue,
	)
	if err != nil {
		return common.Hash{}, err
	}

	opts := &bind.TransactOpts{
		GasLimit: 500_000,
		NoSend:   true,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}

	tx, err := config.SystemContracts.ExchangePortal.Exchange(
		opts,
		config.SystemContracts.EVNDTokenAddress,
		config.SystemContracts.USDTokenAddress,
		randomValue,
		big.NewInt(0),
	)
	if err != nil {
		return common.Hash{}, err
	}
	msg := ethereum.CallMsg{
		From: crypto.PubkeyToAddress(skFrom.PublicKey),
		To:   &config.SystemContracts.ExchangePortalAddress,
		Data: tx.Data(),
	}
	gas, err := config.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, skFrom, *tx.To(), tx.Value(), gas, tx.Data(), isUseRPC)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Exchanged VND to USDT from: ", addr.Hex(), " value: ", randomValue.String(), " tx_hash: ", txHash.Hex())
	return txHash, nil
}

func SendTransferEVNDRandomAmountTx(config *config.Config, skFrom *ecdsa.PrivateKey, to common.Address, isUseRPC bool) (common.Hash, error) {
	from := crypto.PubkeyToAddress(skFrom.PublicKey)
	// Check if `from` and `to` are verified
	fromVerified, err := config.SystemContracts.EntityRegistry.IsVerifiedEntity(&bind.CallOpts{}, from)
	if err != nil {
		return common.Hash{}, err
	}
	toVerified, err := config.SystemContracts.EntityRegistry.IsVerifiedEntity(&bind.CallOpts{}, to)
	if err != nil {
		return common.Hash{}, err
	}
	if !fromVerified {
		return common.Hash{}, fmt.Errorf("address is not verified, skipping")
	}
	if !toVerified {
		return common.Hash{}, fmt.Errorf("address is not verified, skipping")
	}

	if !IsTransferBetweenAddressesAllowed(config, from, to) {
		return common.Hash{}, fmt.Errorf("transfer not allowed between addresses, skipping")
	}

	allowedTransactionTypes := GetAllowedTransactionTypes(config, from, to)
	if len(allowedTransactionTypes) == 0 {
		return common.Hash{}, fmt.Errorf("no allowed transaction types, skipping")
	}
	randomTransactionType := GetRandomAllowedTransactionType(allowedTransactionTypes)
	log.Println("Random transaction type: ", utils.GetTransactionTypeName(randomTransactionType))

	if _, err := AirdropGas(config, from); err != nil {
		return common.Hash{}, err
	}
	if _, err := AirdropGasIfNeeded(config, from); err != nil {
		return common.Hash{}, err
	}
	fromBalance, err := AirdropVNDIfNeeded(config, from)
	if err != nil {
		log.Println("❌ Error while funding eVND: ", err)
		return common.Hash{}, err
	}
	randomValue := utils.RandomBigInt(fromBalance)
	txHash, err := sendTransferEVNDTx(config, skFrom, to, randomValue, randomTransactionType, isUseRPC)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Transferred eVND from: ", from.Hex(), " to: ", to.Hex(), " value: ", randomValue.String(), " tx hash: ", txHash.Hex())
	return txHash, nil
}

func sendTransferEVNDTx(config *config.Config, sk *ecdsa.PrivateKey, to common.Address, value *big.Int, txType utils.TransactionType, isUseRPC bool) (common.Hash, error) {
	opts := &bind.TransactOpts{
		GasLimit: 500_000,
		NoSend:   true,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	var tx *types.Transaction
	var err error
	if txType == utils.UNKNOWN_TX {
		tx, err = config.SystemContracts.EVNDToken.Transfer(opts, to, value)
	} else {
		tx, err = config.SystemContracts.EVNDToken.TransferWithType(opts, to, value, uint8(txType))
	}
	if err != nil {
		return common.Hash{}, err
	}
	// msg := ethereum.CallMsg{
	// 	From: crypto.PubkeyToAddress(sk.PublicKey),
	// 	To:   &to,
	// 	Data: tx.Data(),
	// }
	// gas, err := config.Client.EstimateGas(context.Background(), msg)
	// if err != nil {
	// 	return common.Hash{}, err
	// }
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, sk, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), isUseRPC)
	if err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func hashEntity(domainSeparator string, entity entityRegistry.Entity) common.Hash {
	arguments := abi.Arguments{
		{Type: typeBytes32},
		{Type: typeAddress},
		{Type: typeUint8},
		{Type: typeBytes},
		{Type: typeAddress},
	}
	structAbiEncoded, err := arguments.Pack(ENTITY_TYPEHASH, entity.EntityAddress, entity.EntityType, entity.EntityData, entity.Verifier)
	if err != nil {
		panic(err)
	}
	structHash := crypto.Keccak256Hash(structAbiEncoded).Bytes()
	var buf []byte
	buf = append(buf, 0x19)
	buf = append(buf, 0x01)
	buf = append(buf, common.FromHex(domainSeparator)...)
	buf = append(buf, structHash...)
	return crypto.Keccak256Hash(buf)
}

func createRegisterTx(config *config.Config, opts *bind.TransactOpts, skVerifier *ecdsa.PrivateKey, skUser *ecdsa.PrivateKey) (*types.Transaction, localtypes.EntityData, uint8, error) {
	entityData, err := utils.GetRandomEntityData(config.EntityDataPath)
	if err != nil {
		return nil, localtypes.EntityData{}, 0, err
	}
	structType, _ := abi.NewType("tuple(string name, bytes32 root)", "", []abi.ArgumentMarshaling{
		{Name: "name", Type: "string"},
		{Name: "root", Type: "bytes32"},
	})
	arguments := abi.Arguments{
		{Type: structType},
	}
	structValue := struct {
		Name string
		Root [32]byte
	}{
		Name: entityData.Name,
		Root: common.HexToHash(entityData.Root),
	}
	structAbiEncoded, err := arguments.Pack(structValue)
	if err != nil {
		return nil, localtypes.EntityData{}, 0, err
	}
	entity := entityRegistry.Entity{
		EntityAddress: crypto.PubkeyToAddress(skUser.PublicKey),
		EntityType:    utils.GetRandomEntityType(),
		EntityData:    structAbiEncoded,
		Verifier:      crypto.PubkeyToAddress(skVerifier.PublicKey),
	}

	domainSeparator, err := config.SystemContracts.EntityRegistry.DomainSeparator(&bind.CallOpts{})
	if err != nil {
		return nil, localtypes.EntityData{}, 0, err
	}

	hash := hashEntity(common.Bytes2Hex(domainSeparator[:]), entity)

	signature, err := crypto.Sign(hash.Bytes(), skVerifier)
	if err != nil {
		return nil, localtypes.EntityData{}, 0, err
	}
	signature[64] += 27

	tx, err := config.SystemContracts.EntityRegistry.Register(opts, entity, signature)

	if err != nil {
		return nil, localtypes.EntityData{}, 0, err
	}
	return tx, entityData, entity.EntityType, nil
}

func isValidAmount(config *config.Config, token0, token1 common.Address, amount *big.Int) bool {
	amountOut, err := config.SystemContracts.ExchangePortal.GetExchangeAmount(&bind.CallOpts{}, token0, token1, amount)
	if err != nil {
		return false
	}
	return amountOut.Cmp(big.NewInt(0)) > 0
}

func SendApproveTxIfNeeded(config *config.Config, sk *ecdsa.PrivateKey, exchangePortal common.Address, token common.Address, value *big.Int) error {
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	allowance, err := config.SystemContracts.EVNDToken.Allowance(&bind.CallOpts{}, addr, exchangePortal)
	if err != nil {
		return err
	}
	if allowance.Cmp(value) >= 0 {
		return nil
	}
	opts := &bind.TransactOpts{
		GasLimit: 100_000,
		NoSend:   true,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	tx, err := config.SystemContracts.EVNDToken.Approve(opts, exchangePortal, value)
	if err != nil {
		return err
	}
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, sk, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), false)
	if err != nil {
		return err
	}
	log.Println("Approved ", value.String(), " from ", addr.Hex(), " tx hash: ", txHash.Hex())
	return nil
}

func SendSetExchangeRateTx(config *config.Config, newExchangeRate *big.Int, isUseRPC bool) (common.Hash, error) {
	opts := &bind.TransactOpts{
		GasLimit: 100_000,
		NoSend:   true,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	tx, err := config.SystemContracts.ExchangePortal.SetExchangeRate(opts, newExchangeRate)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, config.AdminKey, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), isUseRPC)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Set exchange rate to ", newExchangeRate.String(), " tx hash: ", txHash.Hex())
	return txHash, nil
}

func IsTransferBetweenAddressesAllowed(config *config.Config, from, to common.Address) bool {
	fromEntity, err := config.SystemContracts.EntityRegistry.GetEntity(&bind.CallOpts{}, from)
	if err != nil {
		return false
	}
	toEntity, err := config.SystemContracts.EntityRegistry.GetEntity(&bind.CallOpts{}, to)
	if err != nil {
		return false
	}
	fromType := utils.EntityType(fromEntity.EntityType)
	toType := utils.EntityType(toEntity.EntityType)
	log.Println("From type: ", fromType, " To type: ", toType)
	return utils.IsTransferAllowed(fromType, toType)
}

func GetAllowedTransactionTypes(config *config.Config, from, to common.Address) []utils.TransactionType {
	fromEntity, err := config.SystemContracts.EntityRegistry.GetEntity(&bind.CallOpts{}, from)
	if err != nil {
		return []utils.TransactionType{}
	}
	toEntity, err := config.SystemContracts.EntityRegistry.GetEntity(&bind.CallOpts{}, to)
	if err != nil {
		return []utils.TransactionType{}
	}
	fromType := utils.EntityType(fromEntity.EntityType)
	toType := utils.EntityType(toEntity.EntityType)
	return utils.GetAllowedTransactionTypes(fromType, toType)
}

func GetRandomAllowedTransactionType(allowedTransactionTypes []utils.TransactionType) utils.TransactionType {
	if len(allowedTransactionTypes) == 0 {
		return utils.UNKNOWN_TX
	}
	if slices.Contains(allowedTransactionTypes, utils.UNKNOWN_TX) {
		return utils.UNKNOWN_TX
	}
	return allowedTransactionTypes[rand.Intn(len(allowedTransactionTypes))]
}