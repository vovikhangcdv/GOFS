package models

import (
	"time"
)

// Transfer represents a token transfer event
type Transfer struct {
	ID          uint   `gorm:"primaryKey"`
	From        string `gorm:"index"`
	To          string `gorm:"index"`
	Amount      string
	TxHash      string `gorm:"uniqueIndex"`
	BlockNumber uint64
	Timestamp   time.Time
}
