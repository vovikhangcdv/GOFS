package models

import (
	"gorm.io/gorm"
	"time"
)

// Rule represents a monitoring rule in the system
type Rule struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"` // Unique name/identifier for the rule
	Description string `gorm:"type:text;not null"`   // Detailed description of the rule
	Status      string `gorm:"not null"`             // Status: "active" or "inactive"
	Severity    string `gorm:"not null"`             // Severity level: "low", "medium", "high"
	Parameters  string `gorm:"type:json"`           // JSON string containing rule-specific parameters
	Actions     string `gorm:"type:json"`           // JSON string containing actions to take when rule is violated
	Violations  uint64 `gorm:"default:0"`           // Number of times the rule has been violated
	LastViolationAt time.Time // Timestamp of the last violation
}

// RuleHit represents a record of when a rule was triggered


// RuleViolation represents a record of when a rule was violated
type RuleViolation struct {
	gorm.Model
	RuleID      uint   `gorm:"index;not null"`      // Reference to the rule
	TxHash      string `gorm:"index;not null"`      // Transaction hash that violated the rule
	BlockNumber uint64 `gorm:"index;not null"`      // Block number when rule was violated
	Details     string `gorm:"type:json"`           // JSON string containing details about the violation
	ActionTaken string `gorm:"type:json"`           // JSON string containing actions taken in response
} 