-- name: CreateUser :one
INSERT INTO users (id, email, password, name, api_key) VALUES ($1, $2, $3, $4, encode(digest(random()::text::bytea, 'sha256'), 'hex'))
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;