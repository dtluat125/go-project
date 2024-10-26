// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, user_id, feed_id, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow, arg.ID, arg.UserID, arg.FeedID)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}