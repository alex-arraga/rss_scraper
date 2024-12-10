-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, update_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedsFollows :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows WHERE feed_id = $1 AND user_id = $2;