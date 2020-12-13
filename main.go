package main

import (
	"net"
	"net/http"
	"os"

	//"time"

	db "go-postgres/db/sqlc"
	"go-postgres/handlers"
	protos "go-postgres/proto/user"
	"go-postgres/server"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	l := hclog.Default()
	co := db.DbCon{}
	co.InitDB()
	// create the handlers
	ph := handlers.NewUsers(l, co)

	//create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()

	getR.HandleFunc("/products", ph.GetUsers)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/product", ph.AddUser)

	//getR.HandleFunc("/products", ph.ListAll)
	//getR.HandleFunc("/products", ph.ListAll).Queries("currency", "{[A-Z]{3}}")

	//ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// create a new server
	/* 	s := http.Server{
		Addr:         ":9090",                                          // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	} */
	//s.ListenAndServe()
	// start the server
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewUser(l)

	protos.RegisterLoginUserServer(gs, cs)

	reflection.Register(gs)

	ll, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Port already in use", "error", err)
		os.Exit(1)
	}
	gs.Serve(ll)

	co.DBClose()

}

// handlers for API
