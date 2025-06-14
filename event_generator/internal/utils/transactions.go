package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	// "log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetCaps(rpc *rpc.Client) (*big.Int, *big.Int, error) {
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

func GetNonce(backend *rpc.Client, addr common.Address) (uint64, error) {
	client := ethclient.NewClient(backend)
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func GetChainID(backend *rpc.Client) *big.Int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var chainIDHex string
	err := backend.CallContext(ctx, &chainIDHex, "eth_chainId")
	if err != nil {
		panic(err)
	}
	chainID, ok := new(big.Int).SetString(chainIDHex, 0)
	if !ok {
		panic(fmt.Errorf("invalid chain id: %s", chainIDHex))
	}
	return chainID
}

func SendNormalTx(backend *rpc.Client, chainId *big.Int, sk *ecdsa.PrivateKey, to common.Address, value *big.Int, gas uint64, data []byte, isUseRPC bool) (common.Hash, error) {
	if sk == nil {
		return common.Hash{}, fmt.Errorf("sk is nil")
	}

	nonce, err := GetNonce(backend, crypto.PubkeyToAddress(sk.PublicKey))
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      gas,
		GasPrice: big.NewInt(25_000_000_000),
		Data:     data,
	})
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(chainId), sk)
	if err != nil {
		return common.Hash{}, err
	}
	// double check the signature
	recovered, _ := types.Sender(types.NewCancunSigner(chainId), signedTx)
	if recovered != crypto.PubkeyToAddress(sk.PublicKey) {
		return common.Hash{}, fmt.Errorf("signature mismatch: recovered %s, expected %s", recovered.Hex(), crypto.PubkeyToAddress(sk.PublicKey).Hex())
	}

	if err := SendSignedTx(backend, signedTx, isUseRPC); err != nil {
		return common.Hash{}, err
	}
	return signedTx.Hash(), nil
}

func SendSignedTx(backend *rpc.Client, tx *types.Transaction, isUseRPC bool) error {
	rlpData, _ := tx.MarshalBinary()
	// log.Println("rlpData: ", hexutil.Encode(rlpData))

	if isUseRPC {
		if err := backend.CallContext(context.Background(), nil, "eth_sendRawTransaction", hexutil.Encode(rlpData)); err != nil {
			return err
		}
		return nil
	}

	client := ethclient.NewClient(backend)
	if err := client.SendTransaction(context.Background(), tx); err != nil {
		return err
	}
	// Wait for the transaction to be mined
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("Failed to wait for transaction confirmation: %v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction failed, tx_hash: %s", tx.Hash().Hex())
	}
	return nil
}
