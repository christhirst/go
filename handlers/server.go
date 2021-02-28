package handlers

import (
	db "go-postgres/db/sqlc"
)

type Server struct {
	store *db.Store
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	return server
}
