-- name: CreateReservation :one
INSERT INTO reservations (
 tenant_id, room_id, start_date, end_date, price, total
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
start_date = $4,
end_date = $5,
price = $6,
total = $7
WHERE tenant_id = $2 AND id = $1
RETURNING *;
