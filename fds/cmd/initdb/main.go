package main

import (
	"flag"
	"fmt"
	"log"

	"token-monitor/config"
	"token-monitor/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Parse command line flags
	dropTables := flag.Bool("drop", false, "Drop existing tables before creating new ones")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Configure GORM logger
	gormConfig := &gorm.Config{}
	if *verbose {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(cfg.Database.GetDSN()), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get underlying SQL DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	// Drop tables if requested
	if *dropTables {
		log.Println("Dropping existing tables...")
		// Drop tables in reverse order of dependencies
		if err := db.Migrator().DropTable(
			&models.SuspiciousTransferRelatedTx{},
			&models.SuspiciousTransfer{},
			&models.TokenTransfer{},
			&models.PendingTransaction{},
			&models.Transaction{},
			&models.BlacklistedAddress{},
			&models.RuleViolation{},
			&models.Rule{},
		); err != nil {
			log.Fatalf("Failed to drop tables: %v", err)
		}
		log.Println("Tables dropped successfully")
	}

	// Auto-migrate the schema in correct order
	log.Println("Creating database schema...")

	// First create base tables without foreign keys
	if err := db.AutoMigrate(
		&models.Transaction{},
		&models.PendingTransaction{},
		&models.BlacklistedAddress{},
		&models.Rule{},
	); err != nil {
		log.Fatalf("Failed to migrate base tables: %v", err)
	}

	// Then create tables with foreign keys
	if err := db.AutoMigrate(
		&models.TokenTransfer{},
		&models.SuspiciousTransfer{},
		&models.SuspiciousTransferRelatedTx{},
		&models.RuleViolation{},
	); err != nil {
		log.Fatalf("Failed to migrate tables with foreign keys: %v", err)
	}

	// Initialize default rules
	defaultRules := []models.Rule{
		{
			Name:        "large_transfer",
			Description: "Detects transfers exceeding a large amount threshold",
			Status:      "active",
			Severity:    "high",
			Parameters:  `{"threshold": "1000000000000000000000", "description": "Transfer amount threshold in wei"}`,
			Actions:     `{"action": "record_violation", "description": "Record violation when transfer amount exceeds threshold"}`,
		},
		{
			Name:        "multiple_transfers",
			Description: "Detects multiple transfers from the same address in a short time period",
			Status:      "active",
			Severity:    "medium",
			Parameters:  `{
				"min_transfers": 4,
				"block_range": 10,
				"description": "Minimum number of transfers and block range to check"
			}`,
			Actions:     `{"action": "record_violation", "description": "Record violation when address makes multiple transfers in short time"}`,
		},
		{
			Name:        "multiple_incoming_transfers",
			Description: "Detects multiple incoming transfers to the same address in a short time period",
			Status:      "active",
			Severity:    "high",
			Parameters:  `{
				"threshold": "1000000000000000000000",
				"block_range": 10,
				"description": "Total amount threshold in wei and block range to check"
			}`,
			Actions:     `{"action": "record_violation", "description": "Record violation when address receives multiple transfers exceeding threshold"}`,
		},
		{
			Name:        "suspicious_address",
			Description: "Detects transactions involving known suspicious addresses",
			Status:      "active",
			Severity:    "high",
			Parameters:  `{
				"addresses": [],
				"description": "List of known suspicious addresses to monitor"
			}`,
			Actions:     `{"action": "record_violation", "description": "Record violation when transaction involves suspicious address"}`,
		},
		{
			Name:        "insufficient_balance",
			Description: "Detects transfers where the sender's balance is less than the transfer amount",
			Status:      "active",
			Severity:    "high",
			Parameters:  `{
				"description": "Check if sender has sufficient balance before transfer",
				"check_blocks": 5,
				"description": "Number of blocks to check for previous balance"
			}`,
			Actions:     `{"action": "record_violation", "description": "Record violation when transfer amount exceeds sender's previous balance"}`,
		},
	}

	// Insert default rules
	for _, rule := range defaultRules {
		result := db.Where("name = ?", rule.Name).FirstOrCreate(&rule)
		if result.Error != nil {
			log.Printf("Warning: Failed to create rule %s: %v", rule.Name, result.Error)
		}
	}

	// Print schema information
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatalf("Failed to get tables: %v", err)
	}

	fmt.Println("\nDatabase schema initialized successfully!")
	fmt.Println("\nCreated tables:")
	for _, table := range tables {
		fmt.Printf("- %s\n", table)
	}

	// Print indexes
	fmt.Println("\nIndexes:")
	for _, table := range tables {
		indexes, err := db.Migrator().GetIndexes(table)
		if err != nil {
			log.Printf("Failed to get indexes for table %s: %v", table, err)
			continue
		}
		fmt.Printf("\n%s:\n", table)
		for _, index := range indexes {
			fmt.Printf("  - %s\n", index)
		}
	}
}
