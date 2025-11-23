-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snap.accounts (
    id SERIAL PRIMARY KEY,
    account_no VARCHAR(50) NOT NULL,
    bank_code VARCHAR(20) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_accounts_accno ON snap.accounts(account_no);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snap.accounts;
-- +goose StatementEnd
