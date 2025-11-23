-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS merchant;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS merchant CASCADE;
-- +goose StatementEnd
