// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFollow = `-- name: CreateFollow :one
insert into
    feed_follows(id, created_at, updated_at, user_id, feed_id)
values
    ($1, $2, $3, $4, $5) returning id, created_at, updated_at, user_id, feed_id
`

type CreateFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const deleteFollow = `-- name: DeleteFollow :exec
delete
from
    feed_follows
where
    id = $1
`

func (q *Queries) DeleteFollow(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFollow, id)
	return err
}

const getFollows = `-- name: GetFollows :many
select
    id, created_at, updated_at, user_id, feed_id
from
    feed_follows
where
    user_id = $1
`

func (q *Queries) GetFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}