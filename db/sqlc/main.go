package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

type DbCon struct {
	DB *sql.DB
}

func (d *DbCon) InitDB() (*sql.DB, error) {

	const (
		host     = "uxti.de"
		port     = 5432
		user     = "postgres"
		password = "postmoskwadb"
		dbname   = "userdb"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	d.DB = db
	return db, err
}

func (d *DbCon) DBClose() {
	defer d.DB.Close()
}
