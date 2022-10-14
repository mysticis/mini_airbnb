package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/gomock"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mysticis/airbnb_mini/api"
	db "github.com/mysticis/airbnb_mini/db/sqlc"
	"github.com/mysticis/airbnb_mini/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not log environment variables", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	server.Start(config.ServerAddress)
}
