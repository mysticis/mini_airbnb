-- name: CreateReview :one
INSERT INTO reviews (
  user_id, room_id, comment, rating
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: ListRoomReviews :many
SELECT * FROM reviews
WHERE room_id = $1
ORDER BY id;

-- name: UpdateReview :one
UPDATE reviews
set comment = $3,
rating = $4
WHERE room_id = $2 AND user_id = $1
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE user_id = $2 AND room_id = $1;

-- name: GetRoomReview :one
SELECT * FROM reviews
WHERE room_id = $2  AND user_id = $1 LIMIT 1;
