-- name: CreateFeedFollow :one
WITH inserted_feed AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed
INNER JOIN feeds ON inserted_feed.feed_id = feeds.id
INNER JOIN users ON inserted_feed.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follows WHERE user_id = $1;
   
-- name: DeleteFeedFollow :exec 
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;   
