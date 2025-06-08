-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create tables
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    hash VARCHAR(66) UNIQUE NOT NULL,
    from_address VARCHAR(42) NOT NULL,
    to_address VARCHAR(42) NOT NULL,
    value TEXT NOT NULL,
    block_number BIGINT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    is_analyzed BOOLEAN DEFAULT FALSE,
    is_pending BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) DEFAULT 'confirmed'
);

CREATE TABLE IF NOT EXISTS pending_transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    hash VARCHAR(66) UNIQUE NOT NULL,
    from_address VARCHAR(42) NOT NULL,
    to_address VARCHAR(42) NOT NULL,
    value TEXT NOT NULL,
    block_number BIGINT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    is_analyzed BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) DEFAULT 'pending'
);

CREATE TABLE IF NOT EXISTS token_transfers (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    transaction_hash VARCHAR(66) NOT NULL,
    from_address VARCHAR(42) NOT NULL,
    to_address VARCHAR(42) NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    token_address VARCHAR(42) NOT NULL,
    block_number BIGINT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    is_abnormal BOOLEAN DEFAULT FALSE,
    is_analyzed BOOLEAN DEFAULT FALSE,
    reason TEXT
);

CREATE TABLE IF NOT EXISTS blacklisted_addresses (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    address VARCHAR(42) UNIQUE NOT NULL,
    tx_hash VARCHAR(66) NOT NULL,
    block_number BIGINT NOT NULL,
    reason TEXT,
    severity VARCHAR(10),
    details TEXT
);

CREATE TABLE IF NOT EXISTS suspicious_transfers (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    from_address VARCHAR(42) NOT NULL,
    to_address VARCHAR(42) NOT NULL,
    amount TEXT NOT NULL,
    tx_hash VARCHAR(66) UNIQUE NOT NULL,
    block_number BIGINT NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    reason TEXT,
    severity VARCHAR(10),
    details TEXT,
    is_blacklisted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS suspicious_transfer_related_txs (
    id SERIAL PRIMARY KEY,
    suspicious_transfer_id INTEGER NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    relation_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_transactions_from_block ON transactions(from_address, block_number);
CREATE INDEX IF NOT EXISTS idx_transactions_to_block ON transactions(to_address, block_number);
CREATE INDEX IF NOT EXISTS idx_pending_transactions_from_block ON pending_transactions(from_address, block_number);
CREATE INDEX IF NOT EXISTS idx_pending_transactions_to_block ON pending_transactions(to_address, block_number);
CREATE INDEX IF NOT EXISTS idx_token_transfers_transaction_hash ON token_transfers(transaction_hash);
CREATE INDEX IF NOT EXISTS idx_token_transfers_from ON token_transfers(from_address);
CREATE INDEX IF NOT EXISTS idx_token_transfers_to ON token_transfers(to_address);
CREATE INDEX IF NOT EXISTS idx_token_transfers_token ON token_transfers(token_address);
CREATE INDEX IF NOT EXISTS idx_blacklisted_addresses_tx_hash ON blacklisted_addresses(tx_hash);
CREATE INDEX IF NOT EXISTS idx_blacklisted_addresses_block ON blacklisted_addresses(block_number);
CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_from ON suspicious_transfers(from_address);
CREATE INDEX IF NOT EXISTS idx_suspicious_transfers_to ON suspicious_transfers(to_address);
CREATE INDEX IF NOT EXISTS idx_suspicious_transfer_related_txs_transfer ON suspicious_transfer_related_txs(suspicious_transfer_id);
CREATE INDEX IF NOT EXISTS idx_suspicious_transfer_related_txs_tx ON suspicious_transfer_related_txs(transaction_hash); 