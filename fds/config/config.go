package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Database DatabaseConfig
	Monitor  MonitorConfig
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// EventCondition defines a condition to check for an event
type EventCondition struct {
	Field    string      // Field to check (e.g., "amount", "from", "to")
	Operator string      // Operator (e.g., ">", "<", "==", "in", "not_in")
	Value    interface{} // Value to compare against
	Severity string      // Severity if condition is met ("high", "medium", "low")
}

// MonitorConfig holds monitor service configuration
type MonitorConfig struct {
	EthereumWSURL        string
	ContractAddress      string
	ContractABI          string // Path to the ABI file
	MaxRetryAttempts     int
	RetryBackoff         time.Duration
	LargeAmountThreshold float64
	SuspiciousAddresses  []string
	EventConditions      map[string][]EventCondition // Map of event name to conditions
	ExcludedEvents       []string                    // Events to exclude from monitoring
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "token_monitor"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Monitor: MonitorConfig{
			EthereumWSURL:        getEnv("MAINNET_WS_URL", "ws://localhost:8545"),
			ContractAddress:      getEnv("EVND_TOKEN_ADDRESS", ""),
			ContractABI:          getEnv("CONTRACT_ABI", "./contracts/TokenX.json"),
			MaxRetryAttempts:     10,
			RetryBackoff:         time.Second * 5,
			LargeAmountThreshold: getEnvAsFloat64("LARGE_AMOUNT_THRESHOLD", 1000.0),
			SuspiciousAddresses:  strings.Split(getEnv("SUSPICIOUS_ADDRESSES", ""), ","),
			EventConditions: map[string][]EventCondition{
				"Transfer": {
					{
						Field:    "amount",
						Operator: ">",
						Value:    getEnvAsFloat64("LARGE_AMOUNT_THRESHOLD", 1000.0),
						Severity: "high",
					},
					{
						Field:    "from",
						Operator: "in",
						Value:    "SUSPICIOUS_ADDRESSES",
						Severity: "high",
					},
					{
						Field:    "to",
						Operator: "in",
						Value:    "SUSPICIOUS_ADDRESSES",
						Severity: "high",
					},
				},
				"Blacklisted": {
					{
						Field:    "address",
						Operator: "in",
						Value:    "SUSPICIOUS_ADDRESSES",
						Severity: "medium",
					},
				},
			},
			ExcludedEvents: strings.Split(getEnv("EXCLUDED_EVENTS", ""), ","),
		},
	}

	// Remove empty addresses from the list
	var cleanAddresses []string
	for _, addr := range config.Monitor.SuspiciousAddresses {
		if addr != "" {
			cleanAddresses = append(cleanAddresses, addr)
		}
	}
	config.Monitor.SuspiciousAddresses = cleanAddresses

	// Remove empty events from excluded list
	var cleanExcluded []string
	for _, event := range config.Monitor.ExcludedEvents {
		if event != "" {
			cleanExcluded = append(cleanExcluded, event)
		}
	}
	config.Monitor.ExcludedEvents = cleanExcluded

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// GetDSN returns the database connection string
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

// GetContractABI reads and returns the contract ABI from the specified file
func (c *MonitorConfig) GetContractABI() (string, error) {
	abiBytes, err := os.ReadFile(c.ContractABI)
	if err != nil {
		return "", fmt.Errorf("failed to read contract ABI file: %w", err)
	}
	return string(abiBytes), nil
}

// validate checks if all required configuration values are set
func (c *Config) validate() error {

	if c.Monitor.ContractAddress == "" {
		return fmt.Errorf("CONTRACT_ADDRESS is required")
	}
	if c.Monitor.ContractABI == "" {
		return fmt.Errorf("CONTRACT_ABI is required")
	}
	if c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}
	return nil
}

// Helper functions for environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsFloat64(key string, defaultValue float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			return floatValue
		}
	}
	return defaultValue
}
