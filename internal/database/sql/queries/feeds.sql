-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, update_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
update_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1;

-- name: UpdateFeed :one
UPDATE feeds 
SET name = $1, url = $2
WHERE id = $3
RETURNING *;
