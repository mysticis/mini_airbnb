-- name: CreateRoom :one
INSERT INTO rooms (
  owner_id, 
  home_type, 
  home_size, 
  furnished, 
  private_bathroom,
  balcony,
  garden,
  kitchen,
  pets_allowed,
  parking,
  wheelchair_accessible,
  basement,
  amenities,
  suitable_for,
  price,
  longitude,
  latitude 
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
)
RETURNING *;


-- name: GetRoom :one
SELECT * FROM rooms
WHERE id = $1 LIMIT 1;

-- name: ListRoomsByOwner :many
SELECT * FROM rooms
ORDER BY owner_id
LIMIT $1
OFFSET $2;

-- name: ListAllRooms :many
SELECT * FROM rooms
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRoom :one
UPDATE rooms
set home_type = $2,
 home_size = $3, 
  furnished = $4, 
  private_bathroom = $5,
  balcony = $6,
  garden = $7,
  kitchen = $8,
  pets_allowed = $9,
  parking = $10,
  wheelchair_accessible = $11,
  basement = $12,
  amenities = $13,
  suitable_for = $14,
  price = $15,
  longitude = $16,
  latitude = $17
WHERE owner_id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1;