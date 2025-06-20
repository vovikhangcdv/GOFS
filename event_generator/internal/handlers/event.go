package handlers

import (
	"crypto/ecdsa"
	"context"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/config"
	"github.com/vovikhangcdv/GOFS/chainspammer/internal/utils"
)

func HandleLargeAmountTransfers(config *config.Config, sk *ecdsa.PrivateKey, largeAmount *big.Int) ([]common.Hash, error) {
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	balance, err := config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	if balance.Cmp(largeAmount) < 0 {
		log.Println("Insufficient balance for address ", addr.Hex(), " sending ", largeAmount.String(), " VND")
		log.Println("Airdropping ", largeAmount.String(), " VND to ", addr.Hex())
		_, err := AirdropVND(config, addr, largeAmount)
		if err != nil {
			return nil, fmt.Errorf("failed to airdrop: %w", err)
		}
		balance, err = config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, addr)
		if err != nil {
			return nil, fmt.Errorf("failed to get balance: %w", err)
		}
		if balance.Cmp(largeAmount) < 0 {
			return nil, fmt.Errorf("failed to airdrop: %w", err)
		}
	}
	txHashes := make([]common.Hash, 0)
	receiver := crypto.PubkeyToAddress(config.GetRandomKey().PublicKey)
	allowedTransactionTypes := GetAllowedTransactionTypes(config, addr, receiver)
	if len(allowedTransactionTypes) == 0 {
		return nil, fmt.Errorf("no allowed transaction types, skipping")
	}
	randomTransactionType := GetRandomAllowedTransactionType(allowedTransactionTypes)
	log.Println("Random transaction type: ", utils.GetTransactionTypeName(randomTransactionType))
	txHash, err := sendTransferEVNDTx(config, sk, receiver, largeAmount, randomTransactionType, true)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}
	log.Println("Transaction sent: ", txHash.Hex())
	txHashes = append(txHashes, txHash)
	return txHashes, nil
}

func HandleMultipleOutgoingTransfers(config *config.Config, sk *ecdsa.PrivateKey, blockDuration int64, numOfTxs int64) ([]common.Hash, error) {
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	balance, err := config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	if balance.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("insufficient balance for address %s", addr.Hex())
	}

	// Calculate delay between transactions in seconds
	// Assuming each block takes 3 seconds
	delayBetweenTxs := (blockDuration * 3) / numOfTxs

	log.Printf("MultipleOutgoingTransfers: Starting %d transactions over %d blocks (delay: %d seconds between txs)",
		numOfTxs, blockDuration, delayBetweenTxs)

	var txHashes []common.Hash
	for i := int64(0); i < numOfTxs; i++ {
		// Calculate amount for each transaction
		// Distribute balance evenly across transactions
		amount := new(big.Int).Div(balance, big.NewInt(numOfTxs))
		receiver := crypto.PubkeyToAddress(config.GetRandomKey().PublicKey)
		allowedTransactionTypes := GetAllowedTransactionTypes(config, addr, receiver)
		if len(allowedTransactionTypes) == 0 {
			return nil, fmt.Errorf("no allowed transaction types, skipping")
		}
		randomTransactionType := GetRandomAllowedTransactionType(allowedTransactionTypes)
		log.Println("Random transaction type: ", utils.GetTransactionTypeName(randomTransactionType))
		
		txHash, err := sendTransferEVNDTx(config, sk, receiver, amount, randomTransactionType, true)
		if err != nil {
			return txHashes, fmt.Errorf("failed to send transaction %d: %w", i+1, err)
		}

		txHashes = append(txHashes, txHash)
		log.Printf("Transaction %d/%d sent: %s", i+1, numOfTxs, txHash.Hex())

		// Wait for the calculated delay before next transaction
		if i < numOfTxs-1 { // Don't wait after the last transaction
			time.Sleep(time.Duration(delayBetweenTxs) * time.Second)
		}
	}
	return txHashes, nil
}

func HandleMultipleIncomingTransfers(config *config.Config, addr common.Address, blockDuration int64, totalAmount *big.Int) ([]common.Hash, error) {
	txHashes := make([]common.Hash, 0)

	delayBetweenTxs := (blockDuration * 3) / 10
	log.Printf("MultipleIncomingTransfers: Sending %s VND over %d blocks (delay: %d seconds between txs)",
		totalAmount.String(), blockDuration, delayBetweenTxs)
	totalReceived := big.NewInt(0)
	for i := 0; i < 10; i++ { // Avoid infinite loop
		senderPrivateKey := config.GetRandomKey()
		sender := crypto.PubkeyToAddress(senderPrivateKey.PublicKey)
		balance, err := config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, sender)
		if err != nil {
			return nil, fmt.Errorf("failed to get balance: %w", err)
		}
		if balance.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		randomValue := utils.RandomBigInt(balance)
		allowedTransactionTypes := GetAllowedTransactionTypes(config, sender, addr)
		if len(allowedTransactionTypes) == 0 {
			return nil, fmt.Errorf("no allowed transaction types, skipping")
		}
		randomTransactionType := GetRandomAllowedTransactionType(allowedTransactionTypes)
		log.Println("Random transaction type: ", utils.GetTransactionTypeName(randomTransactionType))
		txHash, err := sendTransferEVNDTx(config, senderPrivateKey, addr, randomValue, randomTransactionType, true)
		if err != nil {
			return nil, fmt.Errorf("failed to send transaction: %w", err)
		}
		txHashes = append(txHashes, txHash)
		log.Printf("Transaction sent: %s, received: %s VND", txHash.Hex(), totalReceived.String())
		totalReceived.Add(totalReceived, balance)
		if totalReceived.Cmp(totalAmount) >= 0 {
			break
		}
		time.Sleep(time.Duration(delayBetweenTxs) * time.Second)
	}
	return txHashes, nil
}

func HandleSuspiciousAddressInteractions(config *config.Config, sk *ecdsa.PrivateKey, blacklistAddresses []common.Address) ([]common.Hash, error) {
	if len(blacklistAddresses) == 0 {
		return nil, fmt.Errorf("no blacklist addresses provided")
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	balance, err := config.SystemContracts.EVNDToken.BalanceOf(&bind.CallOpts{}, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}
	randomValue := big.NewInt(0)
	if balance.Cmp(big.NewInt(0)) != 0 {
		randomValue = utils.RandomBigInt(balance)
	}
	txHashes := make([]common.Hash, 0)
	randomBlacklistAddress := blacklistAddresses[rand.Intn(len(blacklistAddresses))]
	allowedTransactionTypes := GetAllowedTransactionTypes(config, addr, randomBlacklistAddress)
	if len(allowedTransactionTypes) == 0 {
		return nil, fmt.Errorf("no allowed transaction types, skipping")
	}
	randomTransactionType := GetRandomAllowedTransactionType(allowedTransactionTypes)
	log.Println("Random transaction type: ", utils.GetTransactionTypeName(randomTransactionType))
	txHash, err := sendTransferEVNDTx(config, sk, randomBlacklistAddress, randomValue, randomTransactionType, true)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}
	log.Println("Transaction sent: ", txHash.Hex())
	txHashes = append(txHashes, txHash)
	return txHashes, nil
}

func GetReceipts(config *config.Config, txHashes []common.Hash) []error {
	errors := make([]error, 0)
	// Check if the transaction is successful
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, txHash := range txHashes {
		tx, isPending, err := config.Client.TransactionByHash(ctx, txHash)
		if err != nil {
			errors = append(errors, fmt.Errorf("transaction %s failed", txHash.Hex()))
			continue
		}
		// if the transaction is pending, wait for it to be mined
		if isPending {
			receipt, err := bind.WaitMined(ctx, config.Client, tx)
			if err != nil {
				errors = append(errors, fmt.Errorf("transaction %s failed", txHash.Hex()))
				continue
			}
			if receipt.Status != 1 {
				errors = append(errors, fmt.Errorf("transaction %s failed", txHash.Hex()))
				continue
			}
			errors = append(errors, nil)
			continue
		}
		errors = append(errors, nil)
	}
	return errors
}