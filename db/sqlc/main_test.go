package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mysticis/airbnb_mini/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")

	if err != nil {
		log.Fatal("could not load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)

	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
