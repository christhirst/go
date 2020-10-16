package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
)

func InitDB() (*sql.DB, error) {

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
	testQueries := New(db)
	println(testQueries)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
