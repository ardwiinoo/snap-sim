-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS merchant.merchant_orders (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    snap_reference_no VARCHAR(255),
    client_id VARCHAR(255),
    amount BIGINT NOT NULL,
    status VARCHAR(20) DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


CREATE INDEX idx_merchant_orders_orderid ON merchant.merchant_orders(order_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS merchant.merchant_orders;
-- +goose StatementEnd
