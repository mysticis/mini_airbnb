-- name: CreateReservation :one
INSERT INTO reservations (
 tenant_id, email_id, room_id, duration, price, total
) VALUES (
  $1, $2, $3, $4, $5, $6
)

RETURNING *;

-- name: GetReservations :many
SELECT * FROM reservations
WHERE tenant_id = $1;

-- name: ListReservations :many
SELECT * FROM reservations
ORDER BY id;

-- name: DeleteReservation :exec
DELETE FROM reservations
WHERE tenant_id = $2 AND id = $1;

-- name: UpdateReservation :one
UPDATE reservations
set room_id = $3,
duration = $4,
price = $5,
total = $6
WHERE tenant_id = $2 AND id = $1
RETURNING *;
