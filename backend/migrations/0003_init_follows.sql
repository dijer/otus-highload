-- +goose Up

-- +goose StatementBegin
CREATE TABLE follows (
    user_id BIGINT NOT NULL,
    friend_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (user_id, friend_id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_follows_followee ON follows(friend_id);
CREATE INDEX idx_follows_follower ON follows(user_id);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE follows
ADD CONSTRAINT follows_user_fk FOREIGN KEY (user_id)
REFERENCES users(id) NOT VALID;
ALTER TABLE follows
ADD CONSTRAINT follows_friend_fk FOREIGN KEY (friend_id)
REFERENCES users(id) NOT VALID;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE follows VALIDATE CONSTRAINT follows_user_fk;
ALTER TABLE follows VALIDATE CONSTRAINT follows_friend_fk;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE follows DROP CONSTRAINT IF EXISTS follows_user_fk;
ALTER TABLE follows DROP CONSTRAINT IF EXISTS follows_friend_fk;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX IF EXISTS idx_follows_follower;
DROP INDEX IF EXISTS idx_follows_followee;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS follows;
-- +goose StatementEnd
