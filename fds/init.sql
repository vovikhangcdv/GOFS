-- Create rules table if it doesn't exist
CREATE TABLE IF NOT EXISTS rules (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    status VARCHAR(50) DEFAULT 'active',
    severity VARCHAR(50) DEFAULT 'medium',
    parameters JSONB,
    actions JSONB,
    violations INTEGER DEFAULT 0,
    last_violation_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default rules if they don't exist
INSERT INTO rules (name, description, severity, parameters, actions)
VALUES 
    ('large_transfer', 'Detects transfers exceeding a large amount threshold', 'high', 
     '{"threshold": "1000000000000000000000"}', 
     '{"action": "record_violation"}')
ON CONFLICT (name) DO NOTHING;

INSERT INTO rules (name, description, severity, parameters, actions)
VALUES 
    ('multiple_transfers', 'Detects multiple transfers from the same address in a short time period', 'medium',
     '{"min_transfers": 4, "time_blocks": 10}',
     '{"action": "record_violation"}')
ON CONFLICT (name) DO NOTHING;

INSERT INTO rules (name, description, severity, parameters, actions)
VALUES 
    ('multiple_incoming_transfers', 'Detects multiple incoming transfers to the same address in a short time period', 'high',
     '{"min_transfers": 3, "time_blocks": 10}',
     '{"action": "record_violation"}')
ON CONFLICT (name) DO NOTHING;

INSERT INTO rules (name, description, severity, parameters, actions)
VALUES 
    ('suspicious_address', 'Detects transactions involving known suspicious addresses', 'high',
     '{"action": "record_violation"}',
     '{"action": "record_violation"}')
ON CONFLICT (name) DO NOTHING;

-- Create rule_violations table if it doesn't exist
CREATE TABLE IF NOT EXISTS rule_violations (
    id SERIAL PRIMARY KEY,
    rule_id INTEGER REFERENCES rules(id),
    tx_hash VARCHAR(255) NOT NULL,
    block_number BIGINT NOT NULL,
    details JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index on rule_violations for faster lookups
CREATE INDEX IF NOT EXISTS idx_rule_violations_tx_hash ON rule_violations(tx_hash);
CREATE INDEX IF NOT EXISTS idx_rule_violations_rule_id ON rule_violations(rule_id); 