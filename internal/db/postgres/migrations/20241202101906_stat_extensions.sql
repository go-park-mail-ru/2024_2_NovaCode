-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;
LOAD 'auto_explain';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION IF EXISTS pg_stat_statements;
-- +goose StatementEnd
