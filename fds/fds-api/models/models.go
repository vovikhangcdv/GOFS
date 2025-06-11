package models

import "time"

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
