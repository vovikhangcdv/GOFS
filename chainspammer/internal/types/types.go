package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type SpammerConfig struct {
    RPCURL       string         `json:"rpcUrl"`
    PrivateKey   string         `json:"privateKey"`
    TokenAddress common.Address `json:"tokenAddress"`
    TPS          int           `json:"tps"`
    Duration     time.Duration `json:"duration"`
}

type SpammerStatus struct {
    IsRunning    bool      `json:"isRunning"`
    StartTime    time.Time `json:"startTime"`
    TPS          int       `json:"tps"`
    TotalTxs     int64     `json:"totalTxs"`
    SuccessTxs   int64     `json:"successTxs"`
    FailedTxs    int64     `json:"failedTxs"`
    CurrentNonce uint64    `json:"currentNonce"`
}

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