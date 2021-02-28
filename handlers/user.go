package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	db "go-postgres/db/sqlc"
	"io/ioutil"
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
	body, _ := ioutil.ReadAll(r.Body)
	Query := db.New(u.userDB.DB)
	print(string(body))
	us, _ := Query.GetAccount(context.Background(), "cshzxd")
	println(us.Username)
	rw.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal([]string{"chicago", "moscow", "munich", "milan"})
	rw.Write(js)
}

func (server *Server) AddUser(ctx context.Context, rw http.ResponseWriter, r *http.Request) {
	var req db.Account

	accounts := &db.Account{}
	err := accounts.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	arg := db.CreateAccountParams{
		Username: req.Username,
		Password: req.Password,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		return
	}

}

func (u *Users) deleteUser(rw http.ResponseWriter, r *http.Request) {
	Query := db.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}
func (u *Users) modifyUser(rw http.ResponseWriter, r *http.Request) {
	Query := db.New(u.userDB.DB)
	Query.GetAccount(context.Background(), "cshzxd")

}
