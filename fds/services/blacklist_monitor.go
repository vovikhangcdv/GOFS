package services

import (
	"context"
	"log"
	"sync"
	"time"

	"token-monitor/contracts/restrict"
	"token-monitor/models"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

// BlacklistMonitor monitors suspicious addresses and adds them to the blacklist
type BlacklistMonitor struct {
	db             *gorm.DB
	restrictClient *restrict.Restrict
	blacklistAddr  common.Address
	ownerKey       *bind.TransactOpts
	interval       time.Duration
	stopChan       chan struct{}
	wg             sync.WaitGroup
	batchSize      int // Number of addresses to blacklist in one transaction
}

// NewBlacklistMonitor creates a new blacklist monitor
func NewBlacklistMonitor(db *gorm.DB, restrictClient *restrict.Restrict, blacklistAddr common.Address, ownerKey *bind.TransactOpts, interval time.Duration) *BlacklistMonitor {
	return &BlacklistMonitor{
		db:             db,
		restrictClient: restrictClient,
		blacklistAddr:  blacklistAddr,
		ownerKey:       ownerKey,
		interval:       interval,
		stopChan:       make(chan struct{}),
		batchSize:      10, // Default batch size for blacklisting
	}
}

// Start begins monitoring suspicious addresses
func (m *BlacklistMonitor) Start(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		ticker := time.NewTicker(m.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				m.processNewSuspiciousAddresses()
			case <-ctx.Done():
				return
			case <-m.stopChan:
				return
			}
		}
	}()
}

// Stop gracefully stops the monitor
func (m *BlacklistMonitor) Stop() {
	close(m.stopChan)
	m.wg.Wait()
}

// isAddressBlacklisted checks if an address is already blacklisted
func (m *BlacklistMonitor) isAddressBlacklisted(address string) bool {
	var count int64
	m.db.Model(&models.BlacklistedAddress{}).Where("address = ?", address).Count(&count)
	return count > 0
}

// processNewSuspiciousAddresses checks for new suspicious addresses and blacklists them
func (m *BlacklistMonitor) processNewSuspiciousAddresses() {
	// Get all suspicious transfers that haven't been blacklisted yet
	var transfers []models.SuspiciousTransfer
	if err := m.db.Where("severity = ? AND is_blacklisted = ?", "high", false).Find(&transfers).Error; err != nil {
		log.Printf("Error querying suspicious transfers: %v", err)
		return
	}

	// Group addresses into batches
	var addressBatches [][]common.Address
	var currentBatch []common.Address

	for _, transfer := range transfers {
		// Skip if already blacklisted
		if m.isAddressBlacklisted(transfer.To) {
			continue
		}

		addr := common.HexToAddress(transfer.To)
		currentBatch = append(currentBatch, addr)

		// If batch is full, add it to batches and start a new one
		if len(currentBatch) >= m.batchSize {
			addressBatches = append(addressBatches, currentBatch)
			currentBatch = nil
		}
	}

	// Add remaining addresses as the last batch
	if len(currentBatch) > 0 {
		addressBatches = append(addressBatches, currentBatch)
	}

	// Process each batch
	for _, batch := range addressBatches {
		// Call blacklist on the restrict contract
		tx, err := m.restrictClient.Blacklist(m.ownerKey, batch)
		if err != nil {
			log.Printf("Error adding addresses to blacklist: %v", err)
			continue
		}

		// Store blacklisted addresses in database
		for _, addr := range batch {
			blacklistedAddr := &models.BlacklistedAddress{
				Address:     addr.Hex(),
				TxHash:      tx.Hash().Hex(),
				BlockNumber: 0, // Will be updated when transaction is mined
				Reason:      "Multiple suspicious transfers",
				Severity:    "high",
				Details:     "Automatically blacklisted due to suspicious behavior",
			}

			if err := m.db.Create(blacklistedAddr).Error; err != nil {
				log.Printf("Error storing blacklisted address %s: %v", addr.Hex(), err)
				continue
			}

			// Update suspicious transfer record
			if err := m.db.Model(&models.SuspiciousTransfer{}).
				Where("to_address = ?", addr.Hex()).
				Update("is_blacklisted", true).Error; err != nil {
				log.Printf("Error updating suspicious transfer for %s: %v", addr.Hex(), err)
			}

			log.Printf("Added address %s to blacklist. Transaction: %s", addr.Hex(), tx.Hash().Hex())
		}

		// Wait for transaction to be mined and update block number
		receipt, err := bind.WaitMined(context.Background(), m.restrictClient.Client, tx)
		if err != nil {
			log.Printf("Error waiting for blacklist transaction: %v", err)
			continue
		}

		// Update block numbers in database
		if err := m.db.Model(&models.BlacklistedAddress{}).
			Where("tx_hash = ?", tx.Hash().Hex()).
			Update("block_number", receipt.BlockNumber).Error; err != nil {
			log.Printf("Error updating block numbers: %v", err)
		}
	}
}
