-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snap.callback_logs (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER NOT NULL REFERENCES snap.transactions(id),
    request_payload TEXT,
    response_code INTEGER,
    response_body TEXT,
    delivered_at TIMESTAMP,
    attempts INTEGER DEFAULT 0,
    max_attempts INTEGER DEFAULT 5,
    status VARCHAR(20) DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_callback_logs_txn ON snap.callback_logs(transaction_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snap.callback_logs;
-- +goose StatementEnd
