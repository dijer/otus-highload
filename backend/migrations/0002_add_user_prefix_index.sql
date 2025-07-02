-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS user_prefix_idx
ON users (LOWER(first_name) text_pattern_ops, LOWER(last_name) text_pattern_ops);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS user_prefix_idx;
-- +goose StatementEnd
