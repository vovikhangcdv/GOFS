package models

import (
	"time"

	"gorm.io/gorm"
)

// SuspiciousTransfer represents a suspicious token transfer event
type SuspiciousTransfer struct {
	gorm.Model
	From          string `gorm:"index;column:from_address"`
	To            string `gorm:"index;column:to_address"`
	Amount        string
	TxHash        string `gorm:"uniqueIndex"`
	BlockNumber   uint64
	Timestamp     time.Time
	Reason        string        // Reason why the transfer is suspicious
	Severity      string        // "high", "medium", "low"
	Details       string        // JSON string containing additional details
	IsBlacklisted bool          `gorm:"default:false"` // Whether the address has been blacklisted
	RelatedTxs    []Transaction `gorm:"many2many:suspicious_transfer_related_txs;"`
}

// SuspiciousTransferRelatedTx represents the relationship between suspicious transfers and related transactions
type SuspiciousTransferRelatedTx struct {
	ID                   uint   `gorm:"primaryKey"`
	SuspiciousTransferID uint   `gorm:"not null;index"`
	TransactionHash      string `gorm:"not null;index"`
	RelationType         string `gorm:"not null"` // How this transaction is related (e.g., "same_sender", "same_receiver")
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

// TableName specifies the table name for SuspiciousTransferRelatedTx
func (SuspiciousTransferRelatedTx) TableName() string {
	return "suspicious_transfer_related_txs"
}
