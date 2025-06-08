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
	); err != nil {
		log.Fatalf("Failed to migrate base tables: %v", err)
	}

	// Then create tables with foreign keys
	if err := db.AutoMigrate(
		&models.TokenTransfer{},
		&models.SuspiciousTransfer{},
		&models.SuspiciousTransferRelatedTx{},
	); err != nil {
		log.Fatalf("Failed to migrate tables with foreign keys: %v", err)
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
