-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snap.transactions (
    id SERIAL PRIMARY KEY,
    snap_reference_no VARCHAR(255) UNIQUE NOT NULL,
    external_id VARCHAR(255) NOT NULL,
    client_id INTEGER NOT NULL REFERENCES snap.clients(id),
    amount BIGINT NOT NULL,
    currency VARCHAR(10) NOT NULL,
    source_account_no VARCHAR(50) NOT NULL,
    source_bank_code VARCHAR(20) NOT NULL,
    dest_account_no VARCHAR(50) NOT NULL,
    dest_bank_code VARCHAR(20) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL,
    failure_reason TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_transactions_client_id ON snap.transactions(client_id);
CREATE INDEX idx_transactions_snap_ref ON snap.transactions(snap_reference_no);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snap.transactions;
-- +goose StatementEnd
