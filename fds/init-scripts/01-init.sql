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

-- Compliance rules table
CREATE TABLE IF NOT EXISTS rules (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(16) NOT NULL DEFAULT 'active',
    severity VARCHAR(16) NOT NULL DEFAULT 'medium',
    parameters JSONB NOT NULL DEFAULT '{}',
    actions JSONB NOT NULL DEFAULT '{}',
    violations BIGINT DEFAULT 0,
    last_violation_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Rule violations table
CREATE TABLE IF NOT EXISTS rule_violations (
    id SERIAL PRIMARY KEY,
    rule_id INTEGER NOT NULL REFERENCES rules(id) ON DELETE CASCADE,
    tx_hash VARCHAR(66) NOT NULL,
    block_number BIGINT NOT NULL,
    details JSONB NOT NULL DEFAULT '{}',
    action_taken VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Table for suspicious addresses
CREATE TABLE IF NOT EXISTS suspicious_addresses (
    id SERIAL PRIMARY KEY,
    address VARCHAR(42) UNIQUE NOT NULL,
    reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Table for whitelist addresses
CREATE TABLE IF NOT EXISTS whitelist_addresses (
    id SERIAL PRIMARY KEY,
    address VARCHAR(42) UNIQUE NOT NULL,
    reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default rules
INSERT INTO rules (name, description, status, severity, parameters, actions)
VALUES
    (
        'large_transfer',
        'Detects transfers exceeding a large amount threshold',
        'active',
        'high',
        '{"threshold": "1000000000000000000000", "description": "Transfer amount threshold in wei"}',
        '{"action": "record_violation", "description": "Record violation when transfer amount exceeds threshold"}'
    ),
    (
        'multiple_transfers',
        'Detects multiple transfers from the same address in a short time period',
        'active',
        'medium',
        '{"min_transfers": 4, "block_range": 10, "description": "Minimum number of transfers and block range to check"}',
        '{"action": "record_violation", "description": "Record violation when address makes multiple transfers in short time"}'
    ),
    (
        'multiple_incoming_transfers',
        'Detects multiple incoming transfers to the same address in a short time period',
        'active',
        'high',
        '{"threshold": "1000000000000000000000", "block_range": 10, "description": "Total amount threshold in wei and block range to check"}',
        '{"action": "record_violation", "description": "Record violation when address receives multiple transfers exceeding threshold"}'
    ),
    (
        'suspicious_address',
        'Detects transactions involving known suspicious addresses',
        'active',
        'high',
        '{"addresses": [], "description": "List of known suspicious addresses to monitor"}',
        '{"action": "record_violation", "description": "Record violation when transaction involves suspicious address"}'
    ),
    (
        'insufficient_balance',
        'Detects transfers where the sender''s balance is less than the transfer amount',
        'inactive',
        'high',
        '{"description": "Check if sender has sufficient balance before transfer", "check_blocks": 5}',
        '{"action": "record_violation", "description": "Record violation when transfer amount exceeds sender''s previous balance"}'
    )
ON CONFLICT (name) DO NOTHING;

-- Insert initial whitelist address
INSERT INTO whitelist_addresses (address, reason)
VALUES ('0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92261', 'deployer')
ON CONFLICT (address) DO NOTHING;

INSERT INTO whitelist_addresses (address, reason)
VALUES ('0x439Fd6e51aad88F6F4ce6aB8827279cffFb92261', 'reviewer')
ON CONFLICT (address) DO NOTHING;

-- Insert initial suspicious address
INSERT INTO suspicious_addresses (address, reason)
VALUES ('0xa0Ee7A142d267C1f36714E4a8F75612F20a79720', 'tonardo relate adress')
ON CONFLICT (address) DO NOTHING;
INSERT INTO suspicious_addresses (address, reason)
VALUES ('0xb1Ee7A142d267C1f36714E4a8F75612F20a79720', 'tonardo relate adress')
ON CONFLICT (address) DO NOTHING;

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