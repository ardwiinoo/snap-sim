-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS snap;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS snap CASCADE;
-- +goose StatementEnd
