-- +goose Up
ALTER TABLE users ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (encode(digest(random()::text::bytea, 'sha256'), 'hex'));

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;