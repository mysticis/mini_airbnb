-- name: CreateMedia :one
INSERT INTO room_media (
  room_id, room_images
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetRoomMedia :one
SELECT * FROM room_media
WHERE room_id = $1 LIMIT 1;

-- name: ListRoomsMedia :many
SELECT * FROM room_media
WHERE room_id = $1;

-- name: DeleteRoomMedia :exec
DELETE FROM room_media
WHERE room_id = $2 AND id = $1;

-- name: UpdateRoomMedia :one
UPDATE room_media
set room_images = $3
WHERE room_id = $2 AND id = $1
RETURNING *;