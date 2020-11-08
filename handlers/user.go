package handlers

import (
	"context"
	"go-postgres/db/models"
	dbc "go-postgres/db/sqlc"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

// Products handler for getting and updating products
type Users struct {
	l hclog.Logger
	//	v         *data.Validation
	userDB models.DbCon
}

//, v *data.Validation, pdb *data.ProductsDB
//, v, pdb
func NewUsers(l hclog.Logger, userDB models.DbCon) *Users {
	return &Users{l, userDB}
}

func (u *Users) getUsers(rw http.ResponseWriter, h *http.Request) {
	Query := dbc.New(u.userDB.DB)
	Query.ListAccounts(context.Background(), dbc.ListAccountsParams{1, 2})

}

func (u *Users) getUser(rw http.ResponseWriter, h *http.Request) {
	Query := dbc.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}

func (u *Users) addUser(rw http.ResponseWriter, h *http.Request) {
	u.l.Info("Handle Post User")

	account := &dbc.Account{}
	err := account.FromJson(r.Body)

}
func (u *Users) deleteUser(rw http.ResponseWriter, h *http.Request) {
	Query := dbc.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}
func (u *Users) modifyUser(rw http.ResponseWriter, h *http.Request) {
	Query := dbc.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}
