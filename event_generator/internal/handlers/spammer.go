package handlers

import (
	"fmt"
	"log"
	"math/big"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
	localtypes "github.com/vovikhangcdv/GOFS/chainspammer/internal/types"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
)

func Spam(config *config.Config) (common.Hash, error) {
	txTypes := []localtypes.TxType{
		{Type: "register_entity", Weight: config.RegisterEntityWeight},
		{Type: "send_evnd", Weight: config.SendEVNDWeight},
		{Type: "exchange_vnd_usd", Weight: config.ExchangeVNDUSDWeight},
		{Type: "set_exchange_rate", Weight: 1},
	}
	txType := utils.SelectTxType(txTypes)

	var txHash common.Hash
	var err error
	switch txType {
	case "register_entity":
		log.Println("🆕👤 Registering entity")
		skUser, err := config.GetNewKey()
		if err != nil {
			return common.Hash{}, err
		}
		txHash, err = SendRegisterEntityTx(config, skUser, false)
		if err != nil {
			return common.Hash{}, err
		}
		config.Keys = append(config.Keys, skUser)
		log.Println("Total users: ", len(config.Keys))
	case "send_evnd":
		log.Println("💸 Sending EVND")
		if len(config.Keys) < 2 {
			log.Println("❌ Not enough users, len=", len(config.Keys))
			return common.Hash{}, nil
		}
		skFrom := config.GetRandomKey()
		to := crypto.PubkeyToAddress(config.GetRandomKey().PublicKey)
		txHash, err = SendTransferEVNDRandomAmountTx(config, skFrom, to, false)
		if err != nil {
			return common.Hash{}, err
		}
	case "exchange_vnd_usd":
		log.Println("🔄 Exchanging VND to USD")
		if len(config.Keys) == 0 {
			log.Println("❌ Not enough users, len=", len(config.Keys))
			return common.Hash{}, nil
		}
		skFrom := config.GetRandomKey()
		txHash, err = SendExchangeVNDToUSDTx(config, skFrom, false)
		if err != nil {
			return common.Hash{}, err
		}
	case "set_exchange_rate":
		log.Println("✏️ Setting exchange rate")
		minExchangeRate, _ := new(big.Int).SetString("23_000_000_000_000_000_000", 0)
		maxExchangeRate, _ := new(big.Int).SetString("27_000_000_000_000_000_000", 0)
		rangeExchangeRate := new(big.Int).Sub(maxExchangeRate, minExchangeRate)
		newExchangeRate, err := rand.Int(rand.Reader, rangeExchangeRate)
		newExchangeRate = newExchangeRate.Add(newExchangeRate, minExchangeRate)
		if err != nil {
			return common.Hash{}, err
		}
		txHash, err = SendSetExchangeRateTx(config, newExchangeRate, false)
		if err != nil {
			return common.Hash{}, err
		}
	default:
		log.Println("⚠️ Unknown transaction type")
	}
	return txHash, nil
}

func SpamEvent(config *config.Config) (bool, error) {
	event := utils.SelectEvent(config.EventsConfig)
	txHashes := make([]common.Hash, 0)

	if len(config.Keys) < 2 {
		return false, fmt.Errorf("too few users")
	}
	sk := config.GetRandomKey()

	switch e := event.(type) {
	case *localtypes.LargeAmountTransfersConfig:
		log.Println("🚨 EVENT: Large amount transfers")
		txHashesCp, err := HandleLargeAmountTransfers(config, sk, e.TotalAmount)
		if err != nil {
			return false, err
		}
		txHashes = append(txHashes, txHashesCp...)
	case *localtypes.MultipleOutgoingTransfersConfig:
		log.Println("🚨 EVENT: Multiple outgoing transfers")
		txHashesCp, err := HandleMultipleOutgoingTransfers(config, sk, e.BlockDuration, e.TotalTxs)
		if err != nil {
			return false, err
		}
		txHashes = append(txHashes, txHashesCp...)
	case *localtypes.MultipleIncomingTransfersConfig:
		log.Println("🚨 EVENT: Multiple incoming transfers")
		txHashesCp, err := HandleMultipleIncomingTransfers(config, crypto.PubkeyToAddress(sk.PublicKey), e.BlockDuration, e.TotalAmount)
		if err != nil {
			return false, err
		}
		txHashes = append(txHashes, txHashesCp...)
	case *localtypes.SuspiciousAddressInteractionsConfig:
		log.Println("🚨 EVENT: Suspicious address interactions")
		txHashesCp, err := HandleSuspiciousAddressInteractions(config, sk, e.BlacklistAddresses)
		if err != nil {
			return false, err
		}
		txHashes = append(txHashes, txHashesCp...)
	}

	errors := GetReceipts(config, txHashes)
	for _, err := range errors {
		if err != nil {
			return false, err
		}
	}
	return true, nil
}