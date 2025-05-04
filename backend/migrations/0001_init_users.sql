-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  first_name TEXT,
  last_name TEXT,
  birthday DATE,
  gender gender,
  interests TEXT[],
  city TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users

DROP TYPE IF EXISTS gender;
-- +goose StatementEnd
