-- name: CreateTenant :one
INSERT INTO tenant (
  first_name, last_name, email, phone, hashed_password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTenant :one
SELECT * FROM tenant
WHERE email = $1 LIMIT 1;

-- name: ListTenants :many
SELECT * FROM tenant
ORDER BY first_name;

-- name: DeleteTenant :exec
DELETE FROM tenant
WHERE id = $1;