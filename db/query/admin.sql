-- name: CreateAdmin :one
INSERT INTO admin (
  name, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAdmin :one
SELECT * FROM admin
WHERE name = $1 LIMIT 1;
