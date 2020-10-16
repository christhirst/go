package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
)

const (
	dbDriver = "pgx"
	dbSource = "postgresql://postgres:postmoskwadb@uxti.de:5432/userdb?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
