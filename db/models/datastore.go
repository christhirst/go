package models

import (
	"context"
	dbc "go-postgres/db/sqlc"
	"log"

	_ "github.com/jackc/pgx/stdlib"
)

func (d *DbCon) SearchAccount(username string) {
	testQueries := dbc.New(d.DB)

	arg := dbc.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	e, _ := (testQueries.GetAccount(context.Background(), "cshzxd"))
	println(e.Username)
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	if err != nil {
		log.Panic(err)
	}

	for _, account := range accounts {
		if account.Username == username {
			println(account.Username)
		}

	}

}
