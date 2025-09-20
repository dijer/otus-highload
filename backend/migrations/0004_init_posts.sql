-- +goose Up

-- +goose StatementBegin
CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_posts_author_created ON posts(user_id, created_at DESC);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_posts_created ON posts(created_at DESC);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE posts
ADD CONSTRAINT posts_user_fk FOREIGN KEY (user_id)
REFERENCES users(id) NOT VALID;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE posts VALIDATE CONSTRAINT posts_user_fk;
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
ALTER TABLE posts DROP CONSTRAINT IF EXISTS posts_user_fk;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX IF EXISTS idx_posts_created;
DROP INDEX IF EXISTS idx_posts_author_created;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
