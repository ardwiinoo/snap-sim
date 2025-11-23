-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snap.transaction_status_history (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER NOT NULL REFERENCES snap.transactions(id),
    old_status VARCHAR(20),
    new_status VARCHAR(20) NOT NULL,
    changed_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_status_history_txn ON snap.transaction_status_history(transaction_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snap.transaction_status_history;
-- +goose StatementEnd
