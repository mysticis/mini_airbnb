package db

import "database/sql"

// Store defines all functions to execute database queries
type Store interface {
	Querier
}

//SQLStore provides all functions to perform database queries

type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
