package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"token-monitor/contracts/restrict"
	"token-monitor/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// Analyzer provides comprehensive transaction and transfer analysis
type Analyzer struct {
	db              *gorm.DB
	largeAmount     *big.Float
	shortTimeBlocks int64
	suspiciousAddrs map[string]bool
	addressBalances map[string]*big.Float
	txChan          chan *models.Transaction
	stopChan        chan struct{}
	wg              sync.WaitGroup
	mu              sync.RWMutex
	maxBlocks       int
	interval        time.Duration
	rules           map[string]*models.Rule
}

// NewAnalyzer creates a new analyzer instance
func NewAnalyzer(db *gorm.DB, largeAmount float64, shortTimeBlocks int64, suspiciousAddrs []string, interval time.Duration) *Analyzer {
	suspiciousMap := make(map[string]bool)
	for _, addr := range suspiciousAddrs {
		suspiciousMap[addr] = true
	}

	analyzer := &Analyzer{
		db:              db,
		largeAmount:     new(big.Float).SetFloat64(largeAmount),
		shortTimeBlocks: shortTimeBlocks,
		suspiciousAddrs: suspiciousMap,
		addressBalances: make(map[string]*big.Float),
		txChan:          make(chan *models.Transaction, 1000),
		stopChan:        make(chan struct{}),
		maxBlocks:       20,
		interval:        interval,
		rules:           make(map[string]*models.Rule),
	}

	// Load rules from database
	analyzer.loadRules()

	return analyzer
}

// loadRules loads all active rules from the database
func (a *Analyzer) loadRules() {
	var rules []models.Rule
	if err := a.db.Where("status = ?", "active").Find(&rules).Error; err != nil {
		log.Printf("Error loading rules: %v", err)
		return
	}

	for _, rule := range rules {
		// Create a copy of the rule to avoid pointer issues
		ruleCopy := rule
		a.rules[rule.Name] = &ruleCopy
		log.Printf("Loaded rule: %s with parameters: %s", rule.Name, rule.Parameters)
	}
}

// getRuleParameter gets a parameter value from a rule's parameters JSON
func (a *Analyzer) getRuleParameter(ruleName, paramName string) (string, error) {
	rule, exists := a.rules[ruleName]
	if !exists {
		return "", fmt.Errorf("rule %s not found", ruleName)
	}

	var params map[string]interface{}
	if err := json.Unmarshal([]byte(rule.Parameters), &params); err != nil {
		return "", fmt.Errorf("error unmarshaling rule parameters: %w", err)
	}

	value, exists := params[paramName]
	if !exists {
		return "", fmt.Errorf("parameter %s not found in rule %s", paramName, ruleName)
	}

	return fmt.Sprintf("%v", value), nil
}

// Start begins processing transactions and periodic analysis
func (a *Analyzer) Start(ctx context.Context) {
	a.wg.Add(2)

	// Start transaction processing
	go func() {
		defer a.wg.Done()
		for {
			select {
			case tx := <-a.txChan:
				behaviors, err := a.AnalyzeTransaction(ctx, tx)
				if err != nil {
					log.Printf("Error analyzing transaction %s: %v", tx.Hash, err)
					continue
				}
				if len(behaviors) > 0 {
					a.handleSuspiciousBehaviors(tx, behaviors)
				}
			case <-ctx.Done():
				return
			case <-a.stopChan:
				return
			}
		}
	}()

	// Start periodic analysis
	go func() {
		defer a.wg.Done()
		ticker := time.NewTicker(a.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a.processUnanalyzedTransactions()
			case <-ctx.Done():
				return
			case <-a.stopChan:
				return
			}
		}
	}()
}

// Stop gracefully stops the analyzer
func (a *Analyzer) Stop() {
	close(a.stopChan)
	a.wg.Wait()
}

// QueueTransaction adds a transaction to the analysis queue
func (a *Analyzer) QueueTransaction(tx *models.Transaction) {
	// If transaction is pending, analyze it directly
	if tx.IsPending {
		behaviors, err := a.AnalyzeTransaction(context.Background(), tx)
		if err != nil {
			log.Printf("Error analyzing pending transaction %s: %v", tx.Hash, err)
			return
		}

		// Run database update and suspicious behavior handling in parallel
		var wg sync.WaitGroup
		wg.Add(2)

		// Update database in one goroutine
		go func() {
			defer wg.Done()
			if err := a.db.Model(&models.PendingTransaction{}).
				Where("hash = ?", tx.Hash).
				Update("is_analyzed", true).Error; err != nil {
				log.Printf("Error marking pending transaction as analyzed: %v", err)
			}
		}()

		// Handle suspicious behaviors in another goroutine
		go func() {
			defer wg.Done()
			if len(behaviors) > 0 {
				a.handleSuspiciousBehaviors(tx, behaviors)
			}
		}()

		// Wait for both operations to complete
		wg.Wait()
		return
	}

	// For confirmed transactions, add to queue
	select {
	case a.txChan <- tx:
		// Transaction queued successfully
	default:
		log.Printf("Warning: Transaction analysis queue is full, dropping transaction %s", tx.Hash)
	}
}

// AnalyzeTransaction analyzes a single transaction for suspicious behaviors
func (a *Analyzer) AnalyzeTransaction(ctx context.Context, tx *models.Transaction) ([]map[string]interface{}, error) {
	// Check whitelist: skip rule checks if from or to is whitelisted
	var whitelisted []models.WhitelistAddress
	if err := a.db.Model(&models.WhitelistAddress{}).Where("address IN (?, ?)", tx.From, tx.To).Find(&whitelisted).Error; err == nil && len(whitelisted) > 0 {
		log.Printf("Skipping analysis for whitelisted addresses in transaction %s", tx.Hash)
		return nil, nil
	}

	var behaviors []map[string]interface{}

	// Convert transaction value to big.Float
	value := new(big.Float)
	value.SetString(tx.Value)

	// 1. Check for large amount transfers
	if threshold, err := a.getRuleParameter("large_transfer", "threshold"); err == nil {
		thresholdFloat := new(big.Float)
		thresholdFloat.SetString(threshold)
		if value.Cmp(thresholdFloat) > 0 {
			log.Printf("Large transfer detected in transaction %s: %s > %s", tx.Hash, value.String(), threshold)
			details := map[string]interface{}{
				"from":      tx.From,
				"to":        tx.To,
				"amount":    tx.Value,
				"threshold": threshold,
			}
			behaviors = append(behaviors, map[string]interface{}{
				"type":        "large_transfer",
				"description": "Large amount transfer detected",
				"severity":    "high",
				"details":     details,
			})
			a.recordRuleViolation("large_transfer", tx, details)
		}
	} else {
		log.Printf("Error getting large_transfer threshold: %v", err)
	}

	// 2. Check for multiple transfers in short time
	if minTransfers, err := a.getRuleParameter("multiple_transfers", "min_transfers"); err == nil {
		if blockRange, err := a.getRuleParameter("multiple_transfers", "block_range"); err == nil {
			minTransfersInt, _ := strconv.Atoi(minTransfers)
			blockRangeInt, _ := strconv.Atoi(blockRange)
			if multipleBehaviors := a.checkMultipleTransfers(tx, minTransfersInt, blockRangeInt); len(multipleBehaviors) > 0 {
				log.Printf("Multiple transfers detected in transaction %s", tx.Hash)
				details := map[string]interface{}{
					"from":          tx.From,
					"count":         len(multipleBehaviors),
					"min_transfers": minTransfers,
					"block_range":   blockRange,
				}
				behaviors = append(behaviors, multipleBehaviors...)
				a.recordRuleViolation("multiple_transfers", tx, details)
			}
		} else {
			log.Printf("Error getting multiple_transfers block_range: %v", err)
		}
	} else {
		log.Printf("Error getting multiple_transfers min_transfers: %v", err)
	}

	// 3. Check for multiple incoming transfers in short time
	if threshold, err := a.getRuleParameter("multiple_incoming_transfers", "threshold"); err == nil {
		if blockRange, err := a.getRuleParameter("multiple_incoming_transfers", "block_range"); err == nil {
			thresholdFloat := new(big.Float)
			thresholdFloat.SetString(threshold)
			blockRangeInt, _ := strconv.Atoi(blockRange)
			if incomingBehaviors := a.checkMultipleIncomingTransfers(tx, thresholdFloat, blockRangeInt); len(incomingBehaviors) > 0 {
				log.Printf("Multiple incoming transfers detected in transaction %s -- number tx %d", tx.Hash, blockRangeInt)
				details := map[string]interface{}{
					"to":          tx.To,
					"count":       len(incomingBehaviors),
					"threshold":   threshold,
					"block_range": blockRange,
				}
				behaviors = append(behaviors, incomingBehaviors...)
				a.recordRuleViolation("multiple_incoming_transfers", tx, details)
			}
		} else {
			log.Printf("Error getting multiple_incoming_transfers block_range: %v", err)
		}
	} else {
		log.Printf("Error getting multiple_incoming_transfers threshold: %v", err)
	}

	// 4. Check for transfers to/from suspicious addresses
	var suspiciousAddrs []models.SuspiciousAddress
	if err := a.db.Model(&models.SuspiciousAddress{}).Find(&suspiciousAddrs).Error; err == nil {
		for _, addr := range suspiciousAddrs {
			if tx.From == addr.Address || tx.To == addr.Address {
				log.Printf("Suspicious address detected in transaction %s: %s", tx.Hash, addr.Address)
				details := map[string]interface{}{
					"from": tx.From,
					"to":   tx.To,
				}
				behaviors = append(behaviors, map[string]interface{}{
					"type":        "suspicious_address",
					"description": "Transaction involves suspicious address",
					"severity":    "high",
					"details":     details,
				})
				a.recordRuleViolation("suspicious_address", tx, details)
				break
			}
		}
	} else {
		log.Printf("Error checking suspicious addresses: %v", err)
	}

	// 5. Check for insufficient balance transfers
	/* 	if checkBlocks, err := a.getRuleParameter("insufficient_balance", "check_blocks"); err == nil {
	   		if insufficientBehaviors := a.checkInsufficientBalance(tx, value, checkBlocks); len(insufficientBehaviors) > 0 {
	   			log.Printf("Insufficient balance detected in transaction %s", tx.Hash)
	   			details := map[string]interface{}{
	   				"from":         tx.From,
	   				"to":           tx.To,
	   				"amount":       tx.Value,
	   				"check_blocks": checkBlocks,
	   			}
	   			behaviors = append(behaviors, insufficientBehaviors...)
	   			a.recordRuleViolation("insufficient_balance", tx, details)
	   		}
	   	} else {
	   		log.Printf("Error getting insufficient_balance check_blocks: %v", err)
	   	} */

	// Update recent transfers and balances
	a.updateState(tx)

	if len(behaviors) > 0 {
		log.Printf("Found %d suspicious behaviors in transaction %s", len(behaviors), tx.Hash)
	}

	return behaviors, nil
}

// handleSuspiciousBehaviors processes suspicious behaviors and triggers appropriate actions
func (a *Analyzer) handleSuspiciousBehaviors(tx *models.Transaction, behaviors []map[string]interface{}) {
	log.Printf("Suspicious behaviors detected for transaction %s:", tx.Hash)

	// First check if any behavior has high severity
	highestSeverity := "low"
	for _, behavior := range behaviors {
		if behavior["severity"] == "high" {
			highestSeverity = "high"
			break
		} else if behavior["severity"] == "medium" && highestSeverity == "low" {
			highestSeverity = "medium"
		}
	}

	// If high severity, perform blacklist operation immediately
	isBlacklisted := false
	if highestSeverity == "high" {
		// Connect to mainnet for blacklist operations
		mainnetClient, err := ethclient.Dial(os.Getenv("MAINNET_RPC_URL"))
		if err != nil {
			log.Printf("Failed to connect to mainnet for blacklist operation: %v", err)
		} else {
			defer mainnetClient.Close()

			// Get the restrict contract address from environment
			restrictAddr := common.HexToAddress(os.Getenv("RESTRICT_CONTRACT_ADDRESS"))
			restrictContract, err := restrict.NewRestrict(restrictAddr, mainnetClient)
			if err != nil {
				log.Printf("Failed to create restrict contract instance: %v", err)
			} else {
				// Get auth for transaction
				auth, err := a.getAuth(mainnetClient)
				if err != nil {
					log.Printf("Failed to get auth for blacklist operation: %v", err)
				} else {
					// Prepare addresses to blacklist
					addressesToBlacklist := []common.Address{
						common.HexToAddress(tx.To),
					}

					// Call blacklist on the restrict contract
					blacklistTx, err := restrictContract.Blacklist(auth, addressesToBlacklist)
					if err != nil {
						log.Printf("Failed to add address to blacklist: %v", err)
					} else {
						log.Printf("Blacklist transaction sent: %s", blacklistTx.Hash().Hex())

						// Wait for transaction to be mined
						receipt, err := bind.WaitMined(context.Background(), mainnetClient, blacklistTx)
						if err != nil {
							log.Printf("Error waiting for blacklist transaction: %v", err)
						} else if receipt.Status == 1 {
							log.Printf("Successfully blacklisted address %s", tx.To)
							isBlacklisted = true

							// Store in blacklisted table
							blacklistedAddr := &models.BlacklistedAddress{
								Address:     tx.To,
								TxHash:      blacklistTx.Hash().Hex(),
								BlockNumber: receipt.BlockNumber.Uint64(),
								Reason:      "Multiple suspicious transfers",
								Severity:    "high",
								Details:     "Automatically blacklisted due to suspicious behavior",
							}

							if err := a.db.Create(blacklistedAddr).Error; err != nil {
								log.Printf("Error storing blacklisted address %s: %v", tx.To, err)
							} else {
								log.Printf("Added address %s to blacklist table", tx.To)
							}
						} else {
							log.Printf("Failed to blacklist address %s", tx.To)
						}
					}
				}
			}
		}
	}

	// Now process all behaviors and collect details
	allDetails := make(map[string]interface{})
	var relatedTxs []models.Transaction

	for _, behavior := range behaviors {
		log.Printf("- %s (Severity: %s): %s", behavior["type"], behavior["severity"], behavior["description"])
		log.Printf("  Details: %v", behavior["details"])

		// Add behavior details to combined details
		allDetails[behavior["type"].(string)] = behavior["details"]

		// Get related transactions based on behavior type
		switch behavior["type"] {
		case "multiple_transfers":
			// Get the specific transactions that were checked for multiple transfers
			var recentTxs []models.Transaction
			query := a.db.Model(&models.Transaction{}).
				Select("hash, from_address, to_address, value, block_number, timestamp")

			// Handle case where block number is less than shortTimeBlocks
			if tx.BlockNumber >= uint64(a.shortTimeBlocks) {
				query = query.Where("from_address = ? AND block_number >= ?", tx.From, tx.BlockNumber-uint64(a.shortTimeBlocks))
			} else {
				query = query.Where("from_address = ? AND block_number >= ?", tx.From, 0)
			}

			err := query.Find(&recentTxs).Error
			if err != nil {
				log.Printf("Error querying recent transactions: %v", err)
				continue
			}

			transferCount := len(recentTxs)
			if transferCount >= 6 {
				var oldestBlock, newestBlock uint64
				if len(recentTxs) > 0 {
					oldestBlock = recentTxs[len(recentTxs)-1].BlockNumber
					newestBlock = recentTxs[0].BlockNumber
				} else {
					oldestBlock = tx.BlockNumber
					newestBlock = tx.BlockNumber
				}

				blockRange := newestBlock - oldestBlock
				if blockRange > 0 {
					transfersPerBlock := float64(transferCount) / float64(blockRange)

					severity := "medium"
					if transfersPerBlock >= 2.0 {
						severity = "high"
					}

					behaviors = append(behaviors, map[string]interface{}{
						"type":        "multiple_transfers",
						"description": "Multiple transfers in short time",
						"severity":    severity,
						"details": map[string]interface{}{
							"address":             tx.From,
							"count":               transferCount,
							"block_range":         blockRange,
							"transfers_per_block": transfersPerBlock,
							"time_span":           time.Duration(int64(blockRange) * 12 * int64(time.Second)),
							"oldest_block":        oldestBlock,
							"newest_block":        newestBlock,
						},
					})
				}
			}

		case "multiple_incoming_transfers":
			// Get the specific transactions that were checked for multiple incoming transfers
			var incomingTxs []models.Transaction
			query := a.db.Where("to_address = ?", tx.To)
			if tx.BlockNumber > 10 {
				query = query.Where("block_number >= ?", tx.BlockNumber-10)
			}
			if err := query.Order("block_number DESC").Find(&incomingTxs).Error; err == nil {
				relatedTxs = append(relatedTxs, incomingTxs...)
			}
		}
	}

	// Convert combined details to JSON string
	detailsJSON, err := json.Marshal(allDetails)
	if err != nil {
		log.Printf("Error marshaling details: %v", err)
		return
	}
	detailsJSON = append(detailsJSON, []byte(" blocked by fds")...)

	// Set reason as the description(s) of the rules that marked the transaction as suspicious
	var reason string
	if len(behaviors) == 1 {
		desc, _ := behaviors[0]["description"].(string)
		reason = desc
	} else if len(behaviors) > 1 {
		var descs []string
		for _, b := range behaviors {
			if desc, ok := b["description"].(string); ok {
				descs = append(descs, desc)
			}
		}
		reason = strings.Join(descs, "; ")
	}

	// Use a database transaction to ensure atomicity
	err = a.db.Transaction(func(db *gorm.DB) error {
		// Create a single suspicious transfer record for all behaviors
		suspiciousTransfer := &models.SuspiciousTransfer{
			From:          tx.From,
			To:            tx.To,
			Amount:        tx.Value,
			TxHash:        tx.Hash,
			BlockNumber:   tx.BlockNumber,
			Timestamp:     time.Now(),
			Reason:        reason,
			Severity:      highestSeverity,
			Details:       string(detailsJSON),
			IsBlacklisted: isBlacklisted,
		}

		// Create new suspicious transfer record
		if err := db.Create(suspiciousTransfer).Error; err != nil {
			return fmt.Errorf("error creating suspicious transfer: %w", err)
		}

		// Create related transaction records
		for _, relatedTx := range relatedTxs {
			relatedTxRecord := &models.SuspiciousTransferRelatedTx{
				SuspiciousTransferID: suspiciousTransfer.ID,
				TransactionHash:      relatedTx.Hash,
				RelationType:         "related_transfer",
			}
			if err := db.Create(relatedTxRecord).Error; err != nil {
				return fmt.Errorf("error creating related transaction record: %w", err)
			}
		}

		// Update analyzed status
		if tx.IsPending {
			if err := db.Model(&models.PendingTransaction{}).
				Where("hash = ?", tx.Hash).
				Update("is_analyzed", true).Error; err != nil {
				return fmt.Errorf("error marking pending transaction as analyzed: %w", err)
			}
		} else {
			if err := db.Model(&models.Transaction{}).
				Where("hash = ?", tx.Hash).
				Update("is_analyzed", true).Error; err != nil {
				return fmt.Errorf("error marking transaction as analyzed: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Printf("Error in database transaction: %v", err)
		return
	}
}

// getAuth creates an authorized transactor for contract interactions
func (a *Analyzer) getAuth(client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("BLACKLIST_PRIVATE_KEY"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	// Get current gas price
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get gas price: %v", err)
	//}

	// Increase gas price by 50% to speed up transaction
	//gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(20))
	//gasPrice = new(big.Int).Div(gasPrice, big.NewInt(10))
	gasPrice := new(big.Int).Mul(big.NewInt(80), big.NewInt(1e9))
	auth.GasPrice = gasPrice

	return auth, nil
}

// processUnanalyzedTransactions processes any transactions that haven't been analyzed yet
func (a *Analyzer) processUnanalyzedTransactions() {
	// Process unanalyzed pending transactions

	// Process unanalyzed confirmed transactions
	var confirmedTxs []models.Transaction
	if err := a.db.Where("is_analyzed = ?", false).Find(&confirmedTxs).Error; err != nil {
		log.Printf("Error fetching unanalyzed confirmed transactions: %v", err)
	} else {
		for _, tx := range confirmedTxs {
			if err := a.db.Transaction(func(db *gorm.DB) error {
				// Analyze transaction
				behaviors, err := a.AnalyzeTransaction(context.Background(), &tx)
				if err != nil {
					return err
				}

				// Handle any suspicious behaviors
				if len(behaviors) > 0 {
					a.handleSuspiciousBehaviors(&tx, behaviors)
				}

				// Update confirmed transaction
				if err := db.Model(&models.Transaction{}).
					Where("hash = ?", tx.Hash).
					Update("is_analyzed", true).Error; err != nil {
					return fmt.Errorf("failed to mark transaction as analyzed: %w", err)
				}

				return nil
			}); err != nil {
				log.Printf("Error processing confirmed transaction %s: %v", tx.Hash, err)
				continue
			}
			log.Printf("Analyzed confirmed transaction: %s", tx.Hash)
		}
	}
}

// checkMultipleTransfers checks if an address has made multiple transfers in a short time
func (a *Analyzer) checkMultipleTransfers(tx *models.Transaction, minTransfers, blockRange int) []map[string]interface{} {
	var behaviors []map[string]interface{}

	var recentTxs []models.Transaction
	query := a.db.Model(&models.Transaction{}).
		Select("hash, from_address, to_address, value, block_number, timestamp")

	if tx.BlockNumber >= uint64(blockRange) {
		query = query.Where("from_address = ? AND block_number >= ?", tx.From, tx.BlockNumber-uint64(blockRange))
	} else {
		query = query.Where("from_address = ? AND block_number >= ?", tx.From, 0)
	}

	query = query.Order("block_number DESC")

	err := query.Find(&recentTxs).Error
	if err != nil {
		log.Printf("Error querying recent transactions: %v", err)
		return behaviors
	}

	transferCount := len(recentTxs)
	if transferCount >= minTransfers {
		var oldestBlock, newestBlock uint64
		if len(recentTxs) > 0 {
			oldestBlock = recentTxs[len(recentTxs)-1].BlockNumber
			newestBlock = recentTxs[0].BlockNumber
		} else {
			oldestBlock = tx.BlockNumber
			newestBlock = tx.BlockNumber
		}

		blockRange := newestBlock - oldestBlock
		if blockRange > 0 {
			transfersPerBlock := float64(transferCount) / float64(blockRange)

			severity := "medium"
			if transfersPerBlock >= 2.0 {
				severity = "high"
			}

			behaviors = append(behaviors, map[string]interface{}{
				"type":        "multiple_transfers",
				"description": "Multiple transfers in short time",
				"severity":    severity,
				"details": map[string]interface{}{
					"address":             tx.From,
					"count":               transferCount,
					"block_range":         blockRange,
					"transfers_per_block": transfersPerBlock,
					"time_span":           time.Duration(int64(blockRange) * 12 * int64(time.Second)),
					"oldest_block":        oldestBlock,
					"newest_block":        newestBlock,
				},
			})
		}
	}

	return behaviors
}

// checkMultipleIncomingTransfers checks if an address receives multiple transfers in a short time period
func (a *Analyzer) checkMultipleIncomingTransfers(tx *models.Transaction, threshold *big.Float, blockRange int) []map[string]interface{} {
	var behaviors []map[string]interface{}

	var recentTxs []models.Transaction
	query := a.db.Where("to_address = ?", tx.To)
	if tx.BlockNumber > uint64(blockRange) {
		query = query.Where("block_number >= ?", tx.BlockNumber-uint64(blockRange))
	}
	err := query.Order("block_number DESC").Find(&recentTxs).Error

	if err != nil {
		log.Printf("Error querying recent incoming transactions: %v", err)
		return behaviors
	}

	// Calculate total amount received
	totalAmount := new(big.Float)
	totalAmount.SetString(tx.Value)
	for _, recentTx := range recentTxs {
		amount := new(big.Float)
		amount.SetString(recentTx.Value)
		totalAmount.Add(totalAmount, amount)
	}

	// Check if total amount exceeds threshold
	if totalAmount.Cmp(threshold) > 0 {
		var blockRange uint64
		if len(recentTxs) > 0 {
			blockRange = tx.BlockNumber - recentTxs[len(recentTxs)-1].BlockNumber
		}

		behaviors = append(behaviors, map[string]interface{}{
			"type":        "multiple_incoming_transfers",
			"description": "Address received multiple transfers exceeding threshold in short time",
			"severity":    "high",
			"details": map[string]interface{}{
				"address":      tx.To,
				"total_amount": totalAmount.String(),
				"tx_count":     len(recentTxs) + 1,
				"block_range":  blockRange,
				"threshold":    threshold.String(),
			},
		})
	}

	return behaviors
}

// checkBalanceExceeded checks if a transfer amount exceeds the sender's previous balance
func (a *Analyzer) checkBalanceExceeded(tx *models.Transaction) map[string]interface{} {
	value := new(big.Float)
	value.SetString(tx.Value)

	a.mu.RLock()
	currentBalance := a.addressBalances[tx.From]
	a.mu.RUnlock()

	if currentBalance == nil {
		currentBalance = new(big.Float)
	}

	if value.Cmp(currentBalance) > 0 {
		return map[string]interface{}{
			"type":        "balance_exceeded",
			"description": "Transfer amount exceeds previous balance",
			"severity":    "medium",
			"details": map[string]interface{}{
				"address":          tx.From,
				"transfer_amount":  tx.Value,
				"previous_balance": currentBalance.String(),
			},
		}
	}

	return nil
}

// updateState updates the analyzer's state with the new transaction
func (a *Analyzer) updateState(tx *models.Transaction) {
	a.mu.Lock()
	defer a.mu.Unlock()

	value := new(big.Float)
	value.SetString(tx.Value)

	// Update sender's balance
	if balance, exists := a.addressBalances[tx.From]; exists {
		balance.Sub(balance, value)
	} else {
		a.addressBalances[tx.From] = new(big.Float).Neg(value)
	}

	// Update recipient's balance
	if balance, exists := a.addressBalances[tx.To]; exists {
		balance.Add(balance, value)
	} else {
		a.addressBalances[tx.To] = new(big.Float).Set(value)
	}

	// Cleanup old blocks
	a.cleanupOldBlocks(tx.BlockNumber)
}

// cleanupOldBlocks removes data for blocks older than maxBlocks
func (a *Analyzer) cleanupOldBlocks(currentBlock uint64) {
	if currentBlock > uint64(a.maxBlocks) {
		oldestBlock := currentBlock - uint64(a.maxBlocks)
		a.db.Where("block_number < ?", oldestBlock).Delete(&models.Transaction{})
	}
}

// checkRuleStatus checks if a rule is active
func (a *Analyzer) checkRuleStatus(ruleName string) bool {
	var rule models.Rule
	if err := a.db.Where("name = ? AND status = ?", ruleName, "active").First(&rule).Error; err != nil {
		return false
	}
	return true
}

// recordRuleViolation records a violation for an active rule
func (a *Analyzer) recordRuleViolation(ruleName string, tx *models.Transaction, details map[string]interface{}) {
	var rule models.Rule
	if err := a.db.Where("name = ? AND status = ?", ruleName, "active").First(&rule).Error; err != nil {
		return // Rule not found or not active
	}

	// Convert details to JSON
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		log.Printf("Error marshaling violation details: %v", err)
		return
	}

	// Create violation record
	violation := &models.RuleViolation{
		RuleID:      rule.ID,
		TxHash:      tx.Hash,
		BlockNumber: tx.BlockNumber,
		Details:     string(detailsJSON),
	}

	// Use transaction to ensure atomicity
	err = a.db.Transaction(func(db *gorm.DB) error {
		// Create violation record
		if err := db.Create(violation).Error; err != nil {
			return fmt.Errorf("error creating violation record: %w", err)
		}

		// Update rule violation count
		if err := db.Model(&rule).Updates(map[string]interface{}{
			"violations":        rule.Violations + 1,
			"last_violation_at": time.Now(),
		}).Error; err != nil {
			return fmt.Errorf("error updating rule violation count: %w", err)
		}

		return nil
	})

	if err != nil {
		log.Printf("Error recording %s violation: %v", ruleName, err)
	}
}

// getERC20BalanceAt returns the ERC20 token balance for a given address at a specific block (or nil for latest)
func getERC20BalanceAt(client *ethclient.Client, tokenAddress string, userAddress string, blockNumber *big.Int) (*big.Int, error) {
	balanceOfSignature := []byte("balanceOf(address)")
	balanceOfHash := crypto.Keccak256(balanceOfSignature)[:4]

	addr := common.HexToAddress(userAddress)
	paddedAddress := common.LeftPadBytes(addr.Bytes(), 32)
	data := append(balanceOfHash, paddedAddress...)

	tokenAddr := common.HexToAddress(tokenAddress)
	msg := ethereum.CallMsg{
		To:   &tokenAddr,
		Data: data,
	}

	ctx := context.Background()
	result, err := client.CallContract(ctx, msg, blockNumber)
	if err != nil {
		return nil, err
	}

	balance := new(big.Int).SetBytes(result)
	return balance, nil
}

// checkInsufficientBalance checks if the sender had sufficient token balance before the transfer
func (a *Analyzer) checkInsufficientBalance(tx *models.Transaction, transferAmount *big.Float, checkBlocks string) []map[string]interface{} {
	var behaviors []map[string]interface{}

	mainnetClient, err := ethclient.Dial(os.Getenv("MAINNET_RPC_URL"))
	if err != nil {
		log.Printf("Failed to connect to mainnet for balance check: %v", err)
		return behaviors
	}
	defer mainnetClient.Close()

	blockNumber := new(big.Int).SetUint64(tx.BlockNumber)
	blockNumber.Sub(blockNumber, big.NewInt(1))

	tokenAddress := tx.To    // string
	senderAddress := tx.From // string

	balance, err := getERC20BalanceAt(mainnetClient, tokenAddress, senderAddress, blockNumber)
	if err != nil {
		log.Printf("Error getting token balance for address %s at block %d: %v", tx.From, blockNumber, err)
		return behaviors
	}

	balanceFloat := new(big.Float).SetInt(balance)

	if transferAmount.Cmp(balanceFloat) > 0 {
		behaviors = append(behaviors, map[string]interface{}{
			"type":        "insufficient_balance",
			"description": "Token transfer amount exceeds sender's previous balance",
			"severity":    "high",
			"details": map[string]interface{}{
				"from":             tx.From,
				"to":               tx.To,
				"token_address":    tokenAddress,
				"transfer_amount":  transferAmount.String(),
				"previous_balance": balanceFloat.String(),
				"block_number":     blockNumber.String(),
			},
		})
	}

	return behaviors
}
