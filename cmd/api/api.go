package api

import (
	"awesomeProject/service"
	"awesomeProject/storage"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Server — main server structure, contains address and database connection
type Server struct {
	addr string
	db   *sql.DB
}

// NewServer creates and returns a new instance of the server
func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

// Run starts the HTTP server and registers all routes
func (s *Server) Run() error {
	router := mux.NewRouter() // Create a new router using Gorilla Mux
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	requestStore := storage.NewStore(s.db)
	requestService := service.NewLayerService(requestStore)
	requestHandler := service.NewHandler(requestService)

	requestHandler.RegisterRoutes(subrouter) // Register handler routes in the router

	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router) // Start the HTTP server
}
