.PHONY: initdb initdb-drop initdb-verbose migrate migrate-down migrate-status setup build run clean

# Initialize database schema
initdb:
	go run cmd/initdb/main.go

# Initialize database schema with verbose logging
initdb-verbose:
	go run cmd/initdb/main.go -verbose

# Drop and reinitialize database schema
initdb-drop:
	go run cmd/initdb/main.go -drop

# Drop and reinitialize database schema with verbose logging
initdb-drop-verbose:
	go run cmd/initdb/main.go -drop -verbose 

# Database migration commands
migrate:
	@echo "Running database migrations..."
	@go run scripts/migrate.go

# Add more migration commands as needed
migrate-status:
	@echo "Checking migration status..."
	@psql $(DATABASE_URL) -c "SELECT version, applied_at FROM schema_versions ORDER BY version;"

# Development setup
setup:
	@echo "Setting up development environment..."
	@go mod download
	@go mod tidy

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/token-monitor

# Run the application
run: build
	@echo "Running application..."
	@./bin/token-monitor

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/ 