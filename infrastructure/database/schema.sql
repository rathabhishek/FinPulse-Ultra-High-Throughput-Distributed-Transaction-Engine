CREATE TABLE high_value_alerts (
    alert_id UUID PRIMARY KEY,
    transaction_id VARCHAR(255) NOT NULL,
    account_id VARCHAR(50),
    amount DECIMAL(18, 2),
    detected_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);