package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"token-monitor/config"
	"token-monitor/models"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

// AnalyzerService defines the interface for transaction analysis
type AnalyzerService interface {
	QueueTransaction(tx *models.Transaction)
	Start(ctx context.Context)
	Stop()
}

// MonitorService defines the interface for token monitoring
type MonitorService interface {
	Start(ctx context.Context) error
	Stop() error
}

// monitor implements the MonitorService interface
type monitor struct {
	client         *ethclient.Client
	contractAddr   common.Address
	contractABI    abi.ABI
	db             *gorm.DB
	config         config.MonitorConfig
	analyzer       AnalyzerService
	subscriptions  []ethereum.Subscription
	transferEvents chan *models.Transfer
}

// NewMonitor creates a new instance of the monitor service
func NewMonitor(db *gorm.DB, config config.MonitorConfig, analyzer AnalyzerService) (MonitorService, error) {
	client, err := ethclient.Dial(config.EthereumWSURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum node: %w", err)
	}

	// Load contract ABI from file
	abiContent, err := config.GetContractABI()
	if err != nil {
		return nil, fmt.Errorf("failed to read contract ABI: %w", err)
	}

	// Parse contract ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return &monitor{
		client:         client,
		contractAddr:   common.HexToAddress(config.ContractAddress),
		contractABI:    parsedABI,
		db:             db,
		config:         config,
		analyzer:       analyzer,
		subscriptions:  make([]ethereum.Subscription, 0),
		transferEvents: make(chan *models.Transfer, 100),
	}, nil
}

// Start begins monitoring all contract events
func (m *monitor) Start(ctx context.Context) error {
	// Subscribe to all events from the contract
	query := ethereum.FilterQuery{
		Addresses: []common.Address{m.contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := m.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}

	m.subscriptions = append(m.subscriptions, sub)
	go m.processEvents(ctx, logs)
	return nil
}

// Stop gracefully stops the monitor service
func (m *monitor) Stop() error {
	for _, sub := range m.subscriptions {
		if sub != nil {
			sub.Unsubscribe()
		}
	}
	close(m.transferEvents)
	return nil
}

func (m *monitor) processEvents(ctx context.Context, logs chan types.Log) {
	for {
		select {
		case err := <-m.subscriptions[0].Err():
			if err != nil {
				log.Printf("Subscription error: %v", err)
				go m.reconnect(ctx)
			} else {
				log.Printf("Subscription closed gracefully")
			}
			return

		case eventLog := <-logs:
			// Get event name from the first topic
			eventName := m.getEventName(eventLog.Topics[0])
			if eventName == "" {
				continue
			}

			// Skip excluded events
			if m.isEventExcluded(eventName) {
				continue
			}

			// Parse event data
			eventData, err := m.parseEventData(eventLog, eventName)
			if err != nil {
				log.Printf("Error parsing event data: %v", err)
				continue
			}

			// Convert address types to strings
			from := ""
			to := ""
			amount := "0"

			if eventName == "Transfer" {
				if fromAddr, ok := eventData["from"].(common.Address); ok {
					from = fromAddr.Hex()
				}
				if toAddr, ok := eventData["to"].(common.Address); ok {
					to = toAddr.Hex()
				}
				if value, ok := eventData["value"].(*big.Int); ok {
					amount = value.String()
				}
			} else if eventName == "Blacklisted" || eventName == "RemovedFromBlacklist" {
				if addr, ok := eventData["account"].(common.Address); ok {
					from = addr.Hex()
				}
			}

			txHash := eventLog.TxHash.Hex()

			// Check if transaction already exists in transaction table
			/* var existingTx models.Transaction
			result := m.db.Where("hash = ?", txHash).First(&existingTx)
			if result.Error == nil {
				// Transaction already exists in transaction table, skip
				continue
			} else if result.Error != gorm.ErrRecordNotFound {
				// Only log if it's not a "record not found" error
				log.Printf("Error checking existing transaction: %v", result.Error)
				continue
			}
			*/
			// Check if this transaction was previously analyzed in pending state
			var pendingTx models.PendingTransaction
			pendingResult := m.db.Where("hash = ?", txHash).First(&pendingTx)
			isAnalyzed := false

			if pendingResult.Error == nil {
				// Transaction was in pending state, check if it was analyzed
				isAnalyzed = pendingTx.IsAnalyzed
				// Delete the pending transaction after getting its state
				if err := m.db.Delete(&pendingTx).Error; err != nil {
					log.Printf("Error deleting pending transaction: %v", err)
				}
			} else if pendingResult.Error != gorm.ErrRecordNotFound {
				// Only log if it's not a "record not found" error
				log.Printf("Error checking pending transaction: %v", pendingResult.Error)
				continue
			}

			// Create new confirmed transaction
			tx := &models.Transaction{
				Hash:        txHash,
				From:        from,
				To:          to,
				Value:       amount,
				BlockNumber: eventLog.BlockNumber,
				Timestamp:   time.Now(),
				IsAnalyzed:  isAnalyzed,
				IsPending:   false,
				Status:      "confirmed",
			}

			// Save to transaction table
			if err := m.db.Create(tx).Error; err != nil {
				log.Printf("Error saving new transaction: %v", err)
				continue
			}

			// Only queue for analysis if not already analyzed in pending state
			if !isAnalyzed {
				m.analyzer.QueueTransaction(tx)
			}

		case <-ctx.Done():
			log.Printf("Context cancelled, stopping event processing")
			return
		}
	}
}

// reconnect attempts to reestablish the subscription
func (m *monitor) reconnect(ctx context.Context) {
	log.Printf("Attempting to reconnect to event stream...")

	// Wait a bit before attempting to reconnect
	time.Sleep(5 * time.Second)

	// Create new subscription
	query := ethereum.FilterQuery{
		Addresses: []common.Address{m.contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := m.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Printf("Failed to reconnect: %v", err)
		// Try again after a delay
		go m.reconnect(ctx)
		return
	}

	// Replace the old subscription
	m.subscriptions[0] = sub

	// Start processing events again
	go m.processEvents(ctx, logs)
}

func (m *monitor) getEventName(topic common.Hash) string {
	for _, event := range m.contractABI.Events {
		eventSig := []byte(event.Sig)
		eventHash := common.BytesToHash(crypto.Keccak256(eventSig))
		if eventHash == topic {
			return event.Name
		}
	}
	return ""
}

func (m *monitor) isEventExcluded(eventName string) bool {
	for _, excluded := range m.config.ExcludedEvents {
		if excluded == eventName {
			return true
		}
	}
	return false
}

func (m *monitor) parseEventData(eventLog types.Log, eventName string) (map[string]interface{}, error) {
	event, exists := m.contractABI.Events[eventName]
	if !exists {
		return nil, fmt.Errorf("unknown event: %s", eventName)
	}

	// Create a map to store the parsed data
	data := make(map[string]interface{})

	// Parse indexed parameters (topics)
	topicIndex := 1 // Start from 1 as topic[0] is the event signature
	for _, input := range event.Inputs {
		if input.Indexed {
			if topicIndex >= len(eventLog.Topics) {
				return nil, fmt.Errorf("missing topic for indexed parameter %s", input.Name)
			}
			value, err := m.parseTopic(input.Type, eventLog.Topics[topicIndex])
			if err != nil {
				return nil, fmt.Errorf("failed to parse topic for %s: %w", input.Name, err)
			}
			data[input.Name] = value
			topicIndex++
		}
	}

	// Parse non-indexed parameters (data)
	if len(eventLog.Data) > 0 {
		values, err := event.Inputs.Unpack(eventLog.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack event data: %w", err)
		}
		nonIndexedIndex := 0
		for _, input := range event.Inputs {
			if !input.Indexed {
				if nonIndexedIndex >= len(values) {
					return nil, fmt.Errorf("missing data for non-indexed parameter %s", input.Name)
				}
				data[input.Name] = values[nonIndexedIndex]
				nonIndexedIndex++
			}
		}
	}

	return data, nil
}

func (m *monitor) parseTopic(t abi.Type, topic common.Hash) (interface{}, error) {
	switch t.T {
	case abi.AddressTy:
		return common.BytesToAddress(topic.Bytes()), nil
	case abi.IntTy, abi.UintTy:
		return new(big.Int).SetBytes(topic.Bytes()), nil
	case abi.BoolTy:
		return topic.Bytes()[0] == 1, nil
	case abi.StringTy, abi.BytesTy:
		return topic.Bytes(), nil
	default:
		return nil, fmt.Errorf("unsupported topic type: %v", t)
	}
}
