// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: tenant.sql

package db

import (
	"context"
)

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenant (
  first_name, last_name, email, phone, hashed_password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, first_name, last_name, email, phone, hashed_password
`

type CreateTenantParams struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, createTenant,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.HashedPassword,
	)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.HashedPassword,
	)
	return i, err
}

const deleteTenant = `-- name: DeleteTenant :exec
DELETE FROM tenant
WHERE id = $1
`

func (q *Queries) DeleteTenant(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTenant, id)
	return err
}

const getTenant = `-- name: GetTenant :one
SELECT id, first_name, last_name, email, phone, hashed_password FROM tenant
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetTenant(ctx context.Context, email string) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, getTenant, email)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.HashedPassword,
	)
	return i, err
}

const listTenants = `-- name: ListTenants :many
SELECT id, first_name, last_name, email, phone, hashed_password FROM tenant
ORDER BY first_name
`

func (q *Queries) ListTenants(ctx context.Context) ([]Tenant, error) {
	rows, err := q.db.QueryContext(ctx, listTenants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tenant{}
	for rows.Next() {
		var i Tenant
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.HashedPassword,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
