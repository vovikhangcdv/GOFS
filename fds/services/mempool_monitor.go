package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"token-monitor/models"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"gorm.io/gorm"
)

// EventInfo represents decoded event information
type EventInfo struct {
	Event string `json:"event"`
	From  string `json:"from,omitempty"`
	To    string `json:"to,omitempty"`
	Value string `json:"value"`
}

// TxResult represents the result of a transaction simulation
type TxResult struct {
	TxHash string      `json:"txHash"`
	Status string      `json:"status"`
	Events []EventInfo `json:"events,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// ForkProcess manages the Anvil fork process
type ForkProcess struct {
	cmd    *exec.Cmd
	ctx    context.Context
	cancel context.CancelFunc
}

// MempoolMonitor monitors pending transactions in the mempool
type MempoolMonitor struct {
	client      *ethclient.Client
	db          *gorm.DB
	contract    common.Address
	analyzer    AnalyzerService
	stopChan    chan struct{}
	wg          sync.WaitGroup
	interval    time.Duration
	rpcClient   *rpc.Client
	contractABI abi.ABI
}

// NewForkProcess creates a new fork process
func NewForkProcess(mainnetRpcURL string) *ForkProcess {
	ctx, cancel := context.WithCancel(context.Background())
	return &ForkProcess{
		ctx:    ctx,
		cancel: cancel,
	}
}

// Start starts the Anvil fork process
func (f *ForkProcess) Start() error {
	f.cmd = exec.CommandContext(f.ctx, "anvil", "--fork-url", os.Getenv("MAINNET_RPC_URL"), "--port", "9000")
	f.cmd.Stdout = nil
	f.cmd.Stderr = nil

	if err := f.cmd.Start(); err != nil {
		return err
	}

	// Wait for Anvil to be ready
	for i := 0; i < 15; i++ {
		time.Sleep(500 * time.Millisecond)
		if _, err := rpc.DialContext(context.Background(), "http://127.0.0.1:9000"); err == nil {
			return nil
		}
	}
	return fmt.Errorf("anvil failed to start")
}

// Stop stops the Anvil fork process
func (f *ForkProcess) Stop() {
	if f.cmd != nil && f.cmd.Process != nil {
		// Send SIGTERM first
		f.cmd.Process.Signal(syscall.SIGTERM)

		// Give it a moment to terminate gracefully
		done := make(chan error, 1)
		go func() {
			done <- f.cmd.Wait()
		}()

		select {
		case <-time.After(2 * time.Second):
			// If it hasn't terminated after 2 seconds, force kill
			f.cmd.Process.Kill()
		case <-done:
			// Process terminated successfully
		}
	}
	f.cancel()
}

// NewMempoolMonitor creates a new mempool monitor
func NewMempoolMonitor(client *ethclient.Client, rpcClient *rpc.Client, db *gorm.DB, contract common.Address, analyzer AnalyzerService, interval time.Duration) *MempoolMonitor {
	// Parse the ERC20 ABI
	contractABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	if err != nil {
		log.Printf("Error parsing ERC20 ABI: %v", err)
		return nil
	}

	return &MempoolMonitor{
		client:      client,
		db:          db,
		contract:    contract,
		analyzer:    analyzer,
		stopChan:    make(chan struct{}),
		interval:    interval,
		rpcClient:   rpcClient,
		contractABI: contractABI,
	}
}

// ERC20 ABI for event decoding
const erc20ABIJSON = `[
    { "anonymous": false,
      "inputs": [
        {"indexed": true, "internalType": "address", "name": "from", "type": "address"},
        {"indexed": true, "internalType": "address", "name": "to", "type": "address"},
        {"indexed": false,"internalType": "uint256","name": "value","type": "uint256"}
      ],
      "name": "Transfer",
      "type": "event"
    },
    { "anonymous": false,
      "inputs": [
        {"indexed": true, "internalType": "address", "name": "owner",   "type": "address"},
        {"indexed": true, "internalType": "address", "name": "spender", "type": "address"},
        {"indexed": false,"internalType": "uint256","name": "value",   "type": "uint256"}
      ],
      "name": "Approval",
      "type": "event"
    }
]`

// Start begins monitoring the mempool
func (m *MempoolMonitor) Start(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		ticker := time.NewTicker(m.interval)
		defer ticker.Stop()

		// Subscribe to pending transactions
		pendingTxCh := make(chan common.Hash)
		sub, err := m.rpcClient.EthSubscribe(ctx, pendingTxCh, "newPendingTransactions")
		if err != nil {
			log.Printf("Failed to subscribe to pending transactions: %v", err)
			return
		}
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				log.Printf("Mempool subscription error: %v", err)
				return

			case txHash := <-pendingTxCh:
				go m.processPendingTransaction(ctx, txHash)

			case <-ticker.C:
				// Periodically check for stale pending transactions
				m.cleanupStaleTransactions()

			case <-ctx.Done():
				return

			case <-m.stopChan:
				return
			}
		}
	}()
}

// Stop gracefully stops the monitor
func (m *MempoolMonitor) Stop() {
	close(m.stopChan)
	m.wg.Wait()
}

// simulateTransaction simulates a transaction using Anvil fork
func (m *MempoolMonitor) simulateTransaction(ctx context.Context, tx *types.Transaction, targetAddr common.Address, chainID *big.Int) TxResult {
	result := TxResult{TxHash: tx.Hash().Hex()}

	// Start Anvil fork with retries
	var fork *ForkProcess
	var err error
	for i := 0; i < 3; i++ {
		fork = NewForkProcess(os.Getenv("MAINNET_RPC_URL"))
		if err = fork.Start(); err == nil {
			break
		}
		log.Printf("Attempt %d to start Anvil failed: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil {
		result.Error = fmt.Sprintf("Failed to start anvil after retries: %v", err)
		return result
	}
	defer fork.Stop()

	// Create a timeout context for the entire simulation
	simCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Connect to Anvil with retries
	var anvilRpc *rpc.Client
	for i := 0; i < 5; i++ {
		anvilRpc, err = rpc.DialContext(simCtx, "http://127.0.0.1:9000")
		if err == nil {
			break
		}
		log.Printf("Attempt %d to connect to Anvil failed: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil {
		result.Error = fmt.Sprintf("Failed to connect to anvil after retries: %v", err)
		return result
	}
	defer anvilRpc.Close()

	anvilClient := ethclient.NewClient(anvilRpc)

	// Take a snapshot with retry
	var snapshotID string
	for i := 0; i < 3; i++ {
		err = anvilRpc.CallContext(simCtx, &snapshotID, "evm_snapshot")
		if err == nil {
			break
		}
		log.Printf("Attempt %d to take snapshot failed: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil {
		result.Error = fmt.Sprintf("Snapshot error after retries: %v", err)
		return result
	}

	// Execute the transaction
	rawTxBytes, err := tx.MarshalBinary()
	if err != nil {
		result.Error = fmt.Sprintf("Encoding tx error: %v", err)
		return result
	}

	var simTxHash common.Hash

	// Send transaction with retry
	for i := 0; i < 3; i++ {
		err = anvilRpc.CallContext(simCtx, &simTxHash, "eth_sendRawTransaction", "0x"+common.Bytes2Hex(rawTxBytes))
		if err == nil {
			break
		}
		log.Printf("Attempt %d to send transaction failed: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil {
		result.Error = fmt.Sprintf("Simulation send error after retries: %v", err)
		return result
	}

	// Mine block with retry
	for i := 0; i < 3; i++ {
		err = anvilRpc.CallContext(simCtx, nil, "evm_mine")
		if err == nil {
			break
		}
		log.Printf("Attempt %d to mine block failed: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil {
		result.Error = fmt.Sprintf("Mining error after retries: %v", err)
		return result
	}

	// Get receipt with retries
	var receipt *types.Receipt
	for i := 0; i < 5; i++ {
		receipt, err = anvilClient.TransactionReceipt(simCtx, simTxHash)
		if err == nil && receipt != nil {
			break
		}
		log.Printf("Attempt %d to get receipt failed: %v", i+1, err)
		time.Sleep(time.Second * time.Duration(i+1))
	}
	if err != nil || receipt == nil {
		result.Error = fmt.Sprintf("Receipt error after retries: %v", err)
		return result
	}

	// Determine status
	if receipt.Status == types.ReceiptStatusSuccessful {
		result.Status = "success"
	} else {
		result.Status = "revert"
	}

	// Handle ETH transfer
	if tx.Value() != nil && tx.Value().Cmp(big.NewInt(0)) > 0 {
		signer := types.LatestSignerForChainID(chainID)
		fromAddr, err := signer.Sender(tx)
		if err != nil {
			result.Error = fmt.Sprintf("Failed to get sender: %v", err)
			return result
		}

		ev := EventInfo{
			Event: "ETHTransfer",
			From:  fromAddr.Hex(),
			To:    targetAddr.Hex(),
			Value: tx.Value().String(),
		}
		result.Events = append(result.Events, ev)
	}

	// Decode events
	for _, logEntry := range receipt.Logs {
		if !strings.EqualFold(logEntry.Address.Hex(), targetAddr.Hex()) {
			continue
		}
		if len(logEntry.Topics) == 0 {
			continue
		}

		switch logEntry.Topics[0] {
		case m.contractABI.Events["Transfer"].ID:
			var transferEv struct {
				Value *big.Int
			}
			if err := m.contractABI.UnpackIntoInterface(&transferEv, "Transfer", logEntry.Data); err != nil {
				continue
			}
			var fromAddr, toAddr common.Address
			if len(logEntry.Topics) > 2 {
				fromAddr = common.HexToAddress(logEntry.Topics[1].Hex())
				toAddr = common.HexToAddress(logEntry.Topics[2].Hex())
			}
			ev := EventInfo{
				Event: "Transfer",
				From:  fromAddr.Hex(),
				To:    toAddr.Hex(),
				Value: transferEv.Value.String(),
			}
			result.Events = append(result.Events, ev)
		case m.contractABI.Events["Approval"].ID:
			var approvalEv struct {
				Value *big.Int
			}
			if err := m.contractABI.UnpackIntoInterface(&approvalEv, "Approval", logEntry.Data); err != nil {
				continue
			}
			var ownerAddr, spenderAddr common.Address
			if len(logEntry.Topics) > 2 {
				ownerAddr = common.HexToAddress(logEntry.Topics[1].Hex())
				spenderAddr = common.HexToAddress(logEntry.Topics[2].Hex())
			}
			ev := EventInfo{
				Event: "Approval",
				From:  ownerAddr.Hex(),
				To:    spenderAddr.Hex(),
				Value: approvalEv.Value.String(),
			}
			result.Events = append(result.Events, ev)
		}
	}

	return result
}

// processPendingTransaction processes a pending transaction
func (m *MempoolMonitor) processPendingTransaction(ctx context.Context, txHash common.Hash) {
	// Get transaction details
	tx, isPending, err := m.client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Printf("Error getting transaction %s: %v", txHash.Hex(), err)
		return
	}

	// Skip if transaction is no longer pending
	if !isPending {
		return
	}

	// Check if transaction is for our contract
	if tx.To() == nil || *tx.To() != m.contract {
		return
	}

	// Get current block number for simulation
	currentBlock, err := m.client.BlockNumber(ctx)
	if err != nil {
		log.Printf("Error getting current block number: %v", err)
		return
	}

	// Get sender address - handle pending transaction case
	var from common.Address
	if tx.Protected() {
		// For EIP-155 transactions, we can recover the sender from the signature
		chainID, err := m.client.ChainID(ctx)
		if err != nil {
			log.Printf("Error getting chain ID: %v", err)
			return
		}
		signer := types.LatestSignerForChainID(chainID)
		from, err = signer.Sender(tx)
		if err != nil {
			log.Printf("Error recovering sender from signature: %v", err)
			from = common.Address{} // Use zero address if we can't get the sender
		}
	} else {
		// For legacy transactions, try to get sender from the node
		from, err = m.client.TransactionSender(ctx, tx, txHash, 0)
		if err != nil {
			log.Printf("Error getting transaction sender: %v", err)
			from = common.Address{} // Use zero address if we can't get the sender
		}
	}

	// Get chain ID for simulation
	chainID, err := m.client.ChainID(ctx)
	if err != nil {
		log.Printf("Error getting chain ID: %v", err)
		return
	}

	// Simulate the transaction
	result := m.simulateTransaction(ctx, tx, m.contract, chainID)

	// Decode the result based on the method signature
	var to common.Address
	var value *big.Int
	if len(tx.Data()) >= 4 {
		methodID := tx.Data()[:4]
		if common.Bytes2Hex(methodID) == "a9059cbb" { // transfer(address,uint256)
			if len(tx.Data()) >= 68 {
				to = common.BytesToAddress(tx.Data()[16:36])
				value = new(big.Int).SetBytes(tx.Data()[36:68])
			}
		}
	}

	// Create pending transaction record
	pendingTx := &models.PendingTransaction{
		Hash:        txHash.Hex(),
		From:        from.Hex(),
		To:          to.Hex(),
		Value:       value.String(),
		BlockNumber: currentBlock, // Use current block number instead of 0
		Timestamp:   time.Now(),
		IsAnalyzed:  false,
		Status:      result.Status,
	}

	// Save to database using FirstOrCreate
	if err := m.db.Where("hash = ?", txHash.Hex()).FirstOrCreate(pendingTx).Error; err != nil {
		log.Printf("Error saving pending transaction %s: %v", txHash.Hex(), err)
		return
	}

	// Queue for analysis using the pending transaction data
	m.analyzer.QueueTransaction(&models.Transaction{
		Hash:        pendingTx.Hash,
		From:        pendingTx.From,
		To:          pendingTx.To,
		Value:       pendingTx.Value,
		BlockNumber: pendingTx.BlockNumber,
		Timestamp:   pendingTx.Timestamp,
		IsAnalyzed:  pendingTx.IsAnalyzed,
		IsPending:   true,
		Status:      pendingTx.Status,
	})

	// Log simulation results
	if result.Error != "" {
		log.Printf("Simulation error for tx %s: %s", txHash.Hex(), result.Error)
	} else {
		log.Printf("Simulation result for tx %s: %s", txHash.Hex(), result.Status)
		for _, ev := range result.Events {
			log.Printf("  Event: %s From: %s To: %s Value: %s", ev.Event, ev.From, ev.To, ev.Value)
		}
	}
}

// cleanupStaleTransactions removes transactions that have been pending for too long
func (m *MempoolMonitor) cleanupStaleTransactions() {
	// Remove transactions pending for more than 1 hour
	cutoff := time.Now().Add(-time.Hour)
	if err := m.db.Where("status = ? AND timestamp < ?", "pending", cutoff).
		Delete(&models.PendingTransaction{}).Error; err != nil {
		log.Printf("Error cleaning up stale transactions: %v", err)
	}
}
