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

// Rule represents a compliance rule
type Rule struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Name            string         `gorm:"uniqueIndex;not null" json:"name"`
	Description     string         `gorm:"type:text;not null" json:"description"`
	Status          string         `gorm:"type:varchar(16);not null;default:'active'" json:"status"`
	Severity        string         `gorm:"type:varchar(16);not null;default:'medium'" json:"severity"`
	Parameters      string         `gorm:"type:jsonb;not null;default:'{}'" json:"parameters"`
	Actions         string         `gorm:"type:jsonb;not null;default:'{}'" json:"actions"`
	Violations      uint64         `gorm:"default:0" json:"violations"`
	LastViolationAt *time.Time     `gorm:"type:timestamp with time zone" json:"last_violation_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// RuleViolation represents a record of when a rule was violated
type RuleViolation struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	RuleID      uint           `gorm:"index;not null" json:"rule_id"`
	TxHash      string         `gorm:"index;not null" json:"tx_hash"`
	BlockNumber uint64         `gorm:"index;not null" json:"block_number"`
	Details     string         `gorm:"type:jsonb;not null;default:'{}'" json:"details"`
	ActionTaken string         `gorm:"type:varchar(255);default:''" json:"action_taken"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// SuspiciousAddress represents an address flagged as suspicious
type SuspiciousAddress struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Address   string    `gorm:"uniqueIndex;not null" json:"address"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}

// WhitelistAddress represents an address that is whitelisted
type WhitelistAddress struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Address   string    `gorm:"uniqueIndex;not null" json:"address"`
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
}
