-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS citus;

CREATE TABLE IF NOT EXISTS messages (
    chat_id BIGINT NOT NULL,
    msg_id BIGSERIAL,
    sender_id BIGINT NOT NULL,
    recipient_id BIGINT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY(chat_id, msg_id)
);

SELECT create_distributed_table('messages', 'chat_id');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
-- +goose StatementEnd
