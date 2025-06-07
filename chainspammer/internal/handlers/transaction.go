package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
	entityRegistry "github.com/vovikhangcdv/GOFS/chainspammer/internal/generated_contracts/entity_registry"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
)

var (
	ENTITY_TYPEHASH = crypto.Keccak256Hash([]byte("Entity(address entityAddress,uint8 entityType,bytes entityData,address verifier)"))
	typeBytes32, _  = abi.NewType("bytes32", "", nil)
	typeUint8, _    = abi.NewType("uint8", "", nil)
	typeAddress, _  = abi.NewType("address", "", nil)
	typeBytes, _    = abi.NewType("bytes", "", nil)
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
	if balance.Cmp(big.NewInt(int64(float64(params.Ether) * 0.1))) == 0 {
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

func SendRegisterEntityTx(config *config.Config, skUser *ecdsa.PrivateKey) (common.Hash, error) {
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
		GasLimit:  500_000,
		NoSend:    true,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	tx, entityType, err := createRegisterTx(config, opts, skVerifier, skUser)
	if err != nil {
		return common.Hash{}, err
	}

	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, skUser, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), false)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Registered Address: ", crypto.PubkeyToAddress(skUser.PublicKey).Hex(), " entity_type: ", entityType, " tx_hash: ", txHash.Hex())

	if _, err := AirdropVND(config, crypto.PubkeyToAddress(skUser.PublicKey), big.NewInt(10_000_000)); err != nil {
		return common.Hash{}, err
	}
	return txHash, nil
}

func SendExchangeVNDToUSDTx(config *config.Config, skFrom *ecdsa.PrivateKey) (common.Hash, error) {
	addr := crypto.PubkeyToAddress(skFrom.PublicKey)
	if _, err := AirdropGas(config, addr); err != nil {
		return common.Hash{}, err
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
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, skFrom, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), false)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Exchanged VND to USDT from: ", addr.Hex(), " value: ", randomValue.String(), " tx_hash: ", txHash.Hex())
	return txHash, nil
}

func SendTransferEVNDRandomAmountTx(config *config.Config, skFrom *ecdsa.PrivateKey, to common.Address) (common.Hash, error) {
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
	txHash, err := sendTransferEVNDTx(config, skFrom, to, randomValue)
	if err != nil {
		return common.Hash{}, err
	}
	log.Println("Transferred eVND from: ", from.Hex(), " to: ", to.Hex(), " value: ", randomValue.String(), " tx hash: ", txHash.Hex())
	return txHash, nil
}

func sendTransferEVNDTx(config *config.Config, sk *ecdsa.PrivateKey, to common.Address, value *big.Int) (common.Hash, error) {
	opts := &bind.TransactOpts{
		GasLimit:  500_000,
		NoSend:    true,
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return tx, nil
		},
	}
	tx, err := config.SystemContracts.EVNDToken.Transfer(opts, to, value)
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
	txHash, err := utils.SendNormalTx(config.Backend, config.ChainID, sk, *tx.To(), tx.Value(), tx.Gas(), tx.Data(), false)
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

func createRegisterTx(config *config.Config, opts *bind.TransactOpts, skVerifier *ecdsa.PrivateKey, skUser *ecdsa.PrivateKey) (*types.Transaction, uint8, error) {
	entity := entityRegistry.Entity{
		EntityAddress: crypto.PubkeyToAddress(skUser.PublicKey),
		EntityType:    utils.GetRandomEntityType(),
		EntityData:    utils.RandomCallData(32),
		Verifier:      crypto.PubkeyToAddress(skVerifier.PublicKey),
	}

	domainSeparator, err := config.SystemContracts.EntityRegistry.DomainSeparator(&bind.CallOpts{})
	if err != nil {
		return nil, 0, err
	}

	hash := hashEntity(common.Bytes2Hex(domainSeparator[:]), entity)

	signature, err := crypto.Sign(hash.Bytes(), skVerifier)
	if err != nil {
		return nil, 0, err
	}
	signature[64] += 27

	tx, err := config.SystemContracts.EntityRegistry.Register(opts, entity, signature)

	if err != nil {
		return nil, 0, err
	}
	return tx, entity.EntityType, nil
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