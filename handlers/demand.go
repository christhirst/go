package handlers

import (
	"context"
	db "go-postgres/db/sqlc"
	"net/http"
)

func (u *Users) AddDemand(rw http.ResponseWriter, r *http.Request) {
	u.l.Info("Handle Post User")

	account := &db.Account{}
	err := account.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	arg := db.CreateAccountParams{
		Username: account.Username,
		Password: account.Password,
	}

	Query := db.New(u.userDB.DB)
	Query.CreateAccount(context.Background(), arg)

}
