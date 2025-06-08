package main

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a token transaction
type Transaction struct {
	gorm.Model
	Hash        string `gorm:"uniqueIndex"`
	From        string `gorm:"column:from_address;index:idx_from_block"`
	To          string `gorm:"column:to_address;index:idx_to_block"`
	Value       string
	BlockNumber uint64 `gorm:"type:bigint;index:idx_from_block,idx_to_block"`
	Timestamp   time.Time
	IsAnalyzed  bool
	IsPending   bool   `gorm:"default:false"`
	Status      string `gorm:"default:'confirmed'"`
}

// SuspiciousTransfer represents a suspicious token transfer event
type SuspiciousTransfer struct {
	gorm.Model
	From          string `gorm:"index;column:from_address"`
	To            string `gorm:"index;column:to_address"`
	Amount        string
	TxHash        string `gorm:"uniqueIndex"`
	BlockNumber   uint64
	Timestamp     time.Time
	Reason        string
	Severity      string
	Details       string
	IsBlacklisted bool
}

// BlacklistedAddress represents a blacklisted address in the system
type BlacklistedAddress struct {
	gorm.Model
	Address     string `gorm:"uniqueIndex;not null"`
	TxHash      string `gorm:"index;not null"`
	BlockNumber uint64 `gorm:"index;not null"`
	Reason      string
	Severity    string
	Details     string
}
