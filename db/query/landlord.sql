-- name: CreateLandlord :one
INSERT INTO landlord (
  first_name, last_name, email, phone, password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: ListLandlords :many
SELECT * FROM landlord
ORDER BY first_name;