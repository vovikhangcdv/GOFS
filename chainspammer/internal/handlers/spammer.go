package handlers

import (
	"log"

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
	default:
		log.Println("⚠️ Unknown transaction type")
	}
	return txHash, nil
}
