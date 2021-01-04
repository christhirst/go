package main

import (
	"net/http"
	"time"

	//"time"

	db "go-postgres/db/sqlc"
	"go-postgres/handlers"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func main() {
	l := hclog.Default()
	co := db.DbCon{}
	co.InitDB()
	// create the handlers
	ph := handlers.NewUsers(l, co)

	//create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API s
	//getR := sm.Methods(http.MethodGet).Subrouter()
	postR := sm.Methods(http.MethodPost).Subrouter()

	postR.HandleFunc("/adddemand", ph.GetUsers)

	postR.HandleFunc("/user", ph.AddUser)

	postR.HandleFunc("/users", ph.GetUser)
	//getR.HandleFunc("/products", ph.ListAll).Queries("currency", "{[A-Z]{3}}")

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:8000"}))

	// create a new server
	s := http.Server{
		Addr:         ":9090",                                          // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	}
	s.ListenAndServe()
	// start the server
	//log := hclog.Default()
	//gs := grpc.NewServer()
	//cs := server.NewUser(l)

	//protos.RegisterLoginUserServer(gs, cs)

	//reflection.Register(gs)

	//ll, err := net.Listen("tcp", ":9092")
	//if err != nil {
	//	log.Error("Port already in use", "error", err)
	//	os.Exit(1)
	//}
	//gs.Serve(ll)

	co.DBClose()

}

// handlers for API
