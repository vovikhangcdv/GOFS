package models

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

// PendingTransaction represents a transaction in the mempool
type PendingTransaction struct {
	gorm.Model
	Hash        string `gorm:"uniqueIndex"`
	From        string `gorm:"column:from_address;index:idx_pending_from_block"`
	To          string `gorm:"column:to_address;index:idx_pending_to_block"`
	Value       string
	BlockNumber uint64 `gorm:"type:bigint;index:idx_pending_from_block,idx_pending_to_block"`
	Timestamp   time.Time
	IsAnalyzed  bool
	Status      string `gorm:"default:'pending'"`
}

type TokenTransfer struct {
	gorm.Model
	TransactionHash string `gorm:"index"`
	From            string `gorm:"index;column:from_address"`
	To              string `gorm:"index;column:to_address"`
	Amount          float64
	TokenAddress    string `gorm:"index"`
	BlockNumber     uint64 `gorm:"type:bigint"`
	Timestamp       time.Time
	IsAbnormal      bool `gorm:"default:false"`
	IsAnalyzed      bool `gorm:"default:false"`
	Reason          string
}
