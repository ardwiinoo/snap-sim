-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snap.clients (
    id SERIAL PRIMARY KEY,
    client_key VARCHAR(255) UNIQUE NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    client_name VARCHAR(255),
    callback_url TEXT,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snap.clients;
-- +goose StatementEnd
