package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/log"
	localtypes "github.com/vovikhangcdv/GOFS/chainspammer/internal/types"
)

func RandomTx() (*types.Transaction, error) {
	conf := initDefaultTxConf(nil, common.Address{}, 0, nil, nil)
	return types.NewTransaction(conf.Nonce, *conf.To, conf.Value, conf.GasLimit, conf.GasPrice, conf.Code), nil
}

func initDefaultTxConf(
	rpc *rpc.Client, 
	sender common.Address, 
	nonce uint64, 
	gasPrice, 
	chainID *big.Int) *localtypes.TxConf {
	// defaults
	gasCost := uint64(100000)
	to := RandomAddress()
	calldata := RandomCallData(128)
	value := big.NewInt(0)
	if len(calldata) > 128 {
		calldata = calldata[:128]
	}
	// Set fields if non-nil
	if rpc != nil {
		client := ethclient.NewClient(rpc)
		var err error
		if gasPrice == nil {
			gasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Warn("Error suggesting gas price: %v", err)
				gasPrice = big.NewInt(1)
			}
		}
		if chainID == nil {
			chainID, err = client.ChainID(context.Background())
			if err != nil {
				log.Warn("Error fetching chain id: %v", err)
				chainID = big.NewInt(1)
			}
		}
		// Try to estimate gas
		gas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:      sender,
			To:        &to,
			Gas:       30_000_000,
			GasPrice:  gasPrice,
			GasFeeCap: gasPrice,
			GasTipCap: gasPrice,
			Value:     value,
			Data:      calldata,
		})
		if err == nil {
			gasCost = gas
		} else {
			fmt.Printf("Error estimating gas: %v", err)
		}
	}

	return &localtypes.TxConf{
		Rpc:      rpc,
		Nonce:    nonce,
		Sender:   sender,
		To:       &to,
		Value:    value,
		GasLimit: gasCost,
		GasPrice: gasPrice,
		ChainID:  chainID,
		Code:     calldata,
	}
}

func getCaps(rpc *rpc.Client) (*big.Int, *big.Int, error) {
	if rpc == nil {
		return nil, nil, fmt.Errorf("rpc is nil")
	}
	client := ethclient.NewClient(rpc)
	tip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, err
	}
	feeCap, err := client.SuggestGasPrice(context.Background())
	return tip, feeCap, err
}

