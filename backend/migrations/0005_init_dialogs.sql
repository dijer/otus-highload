-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citus;

CREATE TABLE IF NOT EXISTS dialogs (
    chat_id BIGSERIAL PRIMARY KEY,
    user_a BIGINT NOT NULL,
    user_b BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

SELECT create_distributed_table('dialogs', 'chat_id');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS dialogs;
-- +goose StatementEnd
