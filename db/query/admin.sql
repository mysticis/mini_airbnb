-- name: CreateAdmin :one
INSERT INTO admin (
  name, password
) VALUES (
  $1, $2
)
RETURNING *;