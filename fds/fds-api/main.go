package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// initializeDefaultRules creates default rules if they don't exist
func initializeDefaultRules() error {
	defaultRules := []Rule{
		{
			Name:        "large_transfer",
			Description: "Detect transfers above a certain threshold",
			Status:      "active",
			Parameters:  `{"threshold": "1000000"}`,
		},
		{
			Name:        "multiple_transfers",
			Description: "Detect multiple transfers from the same address in a short time",
			Status:      "active",
			Parameters:  `{"min_transfers": "3", "block_range": "10"}`,
		},
		{
			Name:        "multiple_incoming_transfers",
			Description: "Detect multiple incoming transfers to the same address in a short time",
			Status:      "active",
			Parameters:  `{"threshold": "100000", "block_range": "10"}`,
		},
		{
			Name:        "insufficient_balance",
			Description: "Detect transfers that would result in insufficient balance",
			Status:      "active",
			Parameters:  `{"check_blocks": "10"}`,
		},
	}

	for _, rule := range defaultRules {
		var existingRule Rule
		if err := db.Where("name = ?", rule.Name).First(&existingRule).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&rule).Error; err != nil {
					return fmt.Errorf("failed to create rule %s: %w", rule.Name, err)
				}
				log.Printf("Created default rule: %s", rule.Name)
			} else {
				return fmt.Errorf("error checking rule %s: %w", rule.Name, err)
			}
		}
	}
	return nil
}

// migrateDatabase handles database migrations
func migrateDatabase() error {
	// Auto migrate all models
	if err := db.AutoMigrate(
		&Transaction{},
		&SuspiciousTransfer{},
		&BlacklistedAddress{},
		&Rule{},
		&RuleViolation{},
		&SuspiciousAddress{},
		&WhitelistAddress{},
	); err != nil {
		return fmt.Errorf("failed to auto migrate database: %w", err)
	}

	// Add missing columns to rule_violations if they don't exist
	if !db.Migrator().HasColumn(&RuleViolation{}, "details") {
		if err := db.Migrator().AddColumn(&RuleViolation{}, "details"); err != nil {
			return fmt.Errorf("failed to add details column: %w", err)
		}
	}
	if !db.Migrator().HasColumn(&RuleViolation{}, "action_taken") {
		if err := db.Migrator().AddColumn(&RuleViolation{}, "action_taken"); err != nil {
			return fmt.Errorf("failed to add action_taken column: %w", err)
		}
	}

	return nil
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	// Set up database connection
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize default rules
	if err := initializeDefaultRules(); err != nil {
		log.Printf("Warning: Failed to initialize default rules: %v", err)
	}

	// Set up Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Disable caching
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Register API endpoints
	r.GET("/api/address/totals", getAddressTotals)
	r.GET("/api/suspicious", getSuspiciousTransactions)
	r.GET("/api/blacklist", getBlacklist)
	r.GET("/api/address/related", getRelatedAddresses)
	r.GET("/api/address/transactions", getTransactionsByAddress)
	r.GET("/api/suspicious/related", getRelatedTransactionsOfSuspicious)
	r.GET("/api/balance/eth", getETHBalance)
	r.GET("/api/balance/evnd", getEVNDBalance)
	r.POST("/api/blacklist", blacklistAddresses)
	r.POST("/api/unblacklist", unblacklistAddresses)
	r.DELETE("/api/blacklist", deleteBlacklistAddress)
	// New endpoints for rules
	r.GET("/api/rules", getRules)
	r.GET("/api/rules/violations", getRuleViolations)
	r.PUT("/api/rules", updateRule)
	// New endpoints for suspicious/whitelist/blacklist management
	r.GET("/api/suspicious-addresses", getSuspiciousAddresses)
	r.GET("/api/whitelist-addresses", getWhitelistAddresses)

	r.POST("/api/suspicious-addresses/add", addSuspiciousAddress)
	r.POST("/api/suspicious-addresses/remove", removeSuspiciousAddress)
	r.POST("/api/whitelist-addresses/add", addWhitelistAddress)
	r.POST("/api/whitelist-addresses/remove", removeWhitelistAddress)
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Printf("Starting API server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
