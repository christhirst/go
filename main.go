package main

import (
	"go-postgres/db"
	"go-postgres/db/models"
	"log"
)

func main() {
	// create a new serve mux and register the handlers
	//sm := mux.NewRouter()

	// handlers for API
	//getR := sm.Methods(http.MethodGet).Subrouter()

	//ph := handlers.NewProducts(l, v, db)
	//getR.HandleFunc("/products", ph.ListAll)
	//getR.HandleFunc("/products", ph.ListAll).Queries("currency", "{[A-Z]{3}}")

	db, err := db.InitDB()
	if err != nil {
		log.Panic(err)
	}

	var p models.Credentials
	db.QueryRow("select Username, Password WHERE productCode = ?").Scan(&p.Username, &p.Password)

	print(p.Username)

}
