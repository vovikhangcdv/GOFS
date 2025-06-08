package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Get database connection string from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get SQL DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Create schema version table if it doesn't exist
	if err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_versions (
			id SERIAL PRIMARY KEY,
			version INTEGER NOT NULL,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`).Error; err != nil {
		log.Fatalf("Failed to create schema_versions table: %v", err)
	}

	// Get current schema version
	var currentVersion int
	if err := db.Raw("SELECT COALESCE(MAX(version), 0) FROM schema_versions").Scan(&currentVersion).Error; err != nil {
		log.Fatalf("Failed to get current schema version: %v", err)
	}

	// Define schema versions and their migrations
	migrations := []struct {
		version int
		up      func(*gorm.DB) error
	}{
		{
			version: 1,
			up: func(db *gorm.DB) error {
				// Create transactions table
				if err := db.Exec(`
					CREATE TABLE IF NOT EXISTS transactions (
						id SERIAL PRIMARY KEY,
						hash VARCHAR(66) UNIQUE NOT NULL,
						from_address VARCHAR(42) NOT NULL,
						to_address VARCHAR(42) NOT NULL,
						value NUMERIC NOT NULL,
						block_number BIGINT NOT NULL,
						timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
						is_analyzed BOOLEAN DEFAULT FALSE,
						is_pending BOOLEAN DEFAULT FALSE,
						status VARCHAR(20) NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
					);
				`).Error; err != nil {
					return err
				}

				// Create pending_transactions table
				if err := db.Exec(`
					CREATE TABLE IF NOT EXISTS pending_transactions (
						id SERIAL PRIMARY KEY,
						hash VARCHAR(66) UNIQUE NOT NULL,
						from_address VARCHAR(42) NOT NULL,
						to_address VARCHAR(42) NOT NULL,
						value NUMERIC NOT NULL,
						block_number BIGINT NOT NULL,
						timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
						is_analyzed BOOLEAN DEFAULT FALSE,
						status VARCHAR(20) NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
					);
				`).Error; err != nil {
					return err
				}

				// Create token_transfers table
				if err := db.Exec(`
					CREATE TABLE IF NOT EXISTS token_transfers (
						id SERIAL PRIMARY KEY,
						token_address VARCHAR(42) NOT NULL,
						from_address VARCHAR(42) NOT NULL,
						to_address VARCHAR(42) NOT NULL,
						amount NUMERIC NOT NULL,
						transaction_hash VARCHAR(66) NOT NULL REFERENCES transactions(hash),
						block_number BIGINT NOT NULL,
						timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
					);
				`).Error; err != nil {
					return err
				}

				// Create suspicious_transfers table
				if err := db.Exec(`
					CREATE TABLE IF NOT EXISTS suspicious_transfers (
						id SERIAL PRIMARY KEY,
						transaction_hash VARCHAR(66) NOT NULL REFERENCES transactions(hash),
						from_address VARCHAR(42) NOT NULL,
						to_address VARCHAR(42) NOT NULL,
						amount NUMERIC NOT NULL,
						severity VARCHAR(20) NOT NULL,
						reason TEXT NOT NULL,
						is_blacklisted BOOLEAN DEFAULT FALSE,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
					);
				`).Error; err != nil {
					return err
				}

				// Create suspicious_transfer_related_txs table
				if err := db.Exec(`
					CREATE TABLE IF NOT EXISTS suspicious_transfer_related_txs (
						id SERIAL PRIMARY KEY,
						suspicious_transfer_id INTEGER NOT NULL REFERENCES suspicious_transfers(id),
						transaction_hash VARCHAR(66) NOT NULL REFERENCES transactions(hash),
						relation_type VARCHAR(50) NOT NULL,
						created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
						updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
						UNIQUE(suspicious_transfer_id, transaction_hash)
					);
				`).Error; err != nil {
					return err
				}

				return nil
			},
		},
		{
			version: 2,
			up: func(db *gorm.DB) error {
				// Add indexes for performance
				return db.Exec(`
					CREATE INDEX IF NOT EXISTS idx_transactions_block_number ON transactions(block_number);
					CREATE INDEX IF NOT EXISTS idx_transactions_hash ON transactions(hash);
					CREATE INDEX IF NOT EXISTS idx_transactions_from ON transactions(from_address);
					CREATE INDEX IF NOT EXISTS idx_transactions_to ON transactions(to_address);
					CREATE INDEX IF NOT EXISTS idx_transactions_timestamp ON transactions(timestamp);
					
					CREATE INDEX IF NOT EXISTS idx_pending_transactions_hash ON pending_transactions(hash);
					CREATE INDEX IF NOT EXISTS idx_pending_transactions_timestamp ON pending_transactions(timestamp);
					CREATE INDEX IF NOT EXISTS idx_pending_transactions_from ON pending_transactions(from_address);
					CREATE INDEX IF NOT EXISTS idx_pending_transactions_to ON pending_transactions(to_address);
					
					CREATE INDEX IF NOT EXISTS idx_token_transfers_token_address ON token_transfers(token_address);
					CREATE INDEX IF NOT EXISTS idx_token_transfers_tx_hash ON token_transfers(transaction_hash);
					CREATE INDEX IF NOT EXISTS idx_token_transfers_from ON token_transfers(from_address);
					CREATE INDEX IF NOT EXISTS idx_token_transfers_to ON token_transfers(to_address);
					
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_severity ON suspicious_transfers(severity);
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_tx_hash ON suspicious_transfers(transaction_hash);
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_from ON suspicious_transfers(from_address);
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_to ON suspicious_transfers(to_address);
					
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfer_related_txs_suspicious_id ON suspicious_transfer_related_txs(suspicious_transfer_id);
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfer_related_txs_tx_hash ON suspicious_transfer_related_txs(transaction_hash);
				`).Error
			},
		},
		{
			version: 3,
			up: func(db *gorm.DB) error {
				// Add composite indexes for common queries
				return db.Exec(`
					CREATE INDEX IF NOT EXISTS idx_transactions_from_to ON transactions(from_address, to_address);
					CREATE INDEX IF NOT EXISTS idx_token_transfers_from_to ON token_transfers(from_address, to_address);
					CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_from_to ON suspicious_transfers(from_address, to_address);
					CREATE INDEX IF NOT EXISTS idx_pending_transactions_from_to ON pending_transactions(from_address, to_address);
				`).Error
			},
		},
	}

	// Apply pending migrations
	for _, migration := range migrations {
		if migration.version > currentVersion {
			log.Printf("Applying migration version %d...", migration.version)

			// Start transaction
			tx := db.Begin()
			if tx.Error != nil {
				log.Fatalf("Failed to start transaction: %v", tx.Error)
			}

			// Apply migration
			if err := migration.up(tx); err != nil {
				tx.Rollback()
				log.Fatalf("Failed to apply migration version %d: %v", migration.version, err)
			}

			// Record migration version
			if err := tx.Exec("INSERT INTO schema_versions (version) VALUES (?)", migration.version).Error; err != nil {
				tx.Rollback()
				log.Fatalf("Failed to record migration version %d: %v", migration.version, err)
			}

			// Commit transaction
			if err := tx.Commit().Error; err != nil {
				log.Fatalf("Failed to commit migration version %d: %v", migration.version, err)
			}

			log.Printf("Successfully applied migration version %d", migration.version)
		}
	}

	log.Println("All migrations completed successfully")
}
