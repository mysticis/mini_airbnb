-- name: CreateRoom :one
INSERT INTO rooms (
  owner_id, 
  email_id,
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
 $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
)

RETURNING *;


-- name: GetRoomByOwner :one
SELECT * FROM rooms
WHERE owner_id = $2 AND id = $1 LIMIT 1;

-- name: ListRoomsByOwner :many
SELECT * FROM rooms
WHERE owner_id = $1
ORDER BY id;

-- name: ListAllRooms :many
SELECT * FROM rooms
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRoom :one
UPDATE rooms
set home_type = $3,
 home_size = $4, 
  furnished = $5, 
  private_bathroom = $6,
  balcony = $7,
  garden = $8,
  kitchen = $9,
  pets_allowed = $10,
  parking = $11,
  wheelchair_accessible = $12,
  basement = $13,
  amenities = $14,
  suitable_for = $15,
  price = $16,
  longitude = $17,
  latitude = $18
WHERE owner_id = $2 AND id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE owner_id = $2 AND id = $1;