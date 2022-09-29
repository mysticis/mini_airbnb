package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mysticis/airbnb_mini/api"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
)

const (
	dbDriver      = "pgx"
	dbSource      = "postgresql://postgres:mercy@localhost:5432/airbnb?sslmode=disable"
	serverAddress = "0.0.0.0:9000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(serverAddress)
}
