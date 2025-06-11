package models

import (
	"time"

	"gorm.io/gorm"
)

// BlacklistedAddress represents a blacklisted address in the system
type BlacklistedAddress struct {
	gorm.Model
	Address     string `gorm:"uniqueIndex;not null"`
	TxHash      string `gorm:"index;not null"` // Transaction hash that added this address to blacklist
	BlockNumber uint64 `gorm:"index;not null"` // Block number when address was blacklisted
	Reason      string // Reason for blacklisting
	Severity    string // Severity level of the suspicious behavior
	Details     string // Additional details about the blacklisting
}

// WhitelistAddress represents a whitelisted address
type WhitelistAddress struct {
	ID        uint   `gorm:"primaryKey"`
	Address   string `gorm:"uniqueIndex"`
	Reason    string
	CreatedAt time.Time
}

// SuspiciousAddress represents a suspicious address
type SuspiciousAddress struct {
	ID        uint   `gorm:"primaryKey"`
	Address   string `gorm:"uniqueIndex"`
	Reason    string
	CreatedAt time.Time
}
