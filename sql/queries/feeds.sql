-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, title, url, created_at, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;
