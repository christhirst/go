package handlers

import (
	"context"
	"fmt"
	db "go-postgres/db/sqlc"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

// Products handler for getting and updating products
type Users struct {
	l hclog.Logger
	//	v         *data.Validation
	userDB db.DbCon
}

//, v *data.Validation, pdb *data.ProductsDB
//, v, pdb
func NewUsers(l hclog.Logger, userDB db.DbCon) *Users {
	return &Users{l, userDB}
}

func (u *Users) GetUsers(rw http.ResponseWriter, r *http.Request) {
	Query := db.New(u.userDB.DB)
	println(u.userDB.DB)
	list, _ := Query.ListAccounts(context.Background(), db.ListAccountsParams{20, 0})
	fmt.Printf("%v", list)
}

func (u *Users) GetUser(rw http.ResponseWriter, r *http.Request) {
	Query := db.New(u.userDB.DB)

	us, _ := Query.GetAccount(context.Background(), "cshzxd")
	println(us.Username)

}

func (u *Users) AddUser(rw http.ResponseWriter, r *http.Request) {
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
func (u *Users) deleteUser(rw http.ResponseWriter, r *http.Request) {
	Query := db.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}
func (u *Users) modifyUser(rw http.ResponseWriter, r *http.Request) {
	Query := db.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}
