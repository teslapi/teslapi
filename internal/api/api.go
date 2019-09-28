package api

import (
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
	"gocloud.dev/server"
)

// Server is the strcut that contains the API
type Server struct {
	DB  *bolt.DB
	Cfg Config
}

// Config is used to start the new API
type Config struct {
	Server      *server.Server
	DBPath      string
	StoragePath string
}

// New takes a configuration and returns a new API
func New(cfg Config) *Server {
	// open the database
	db, err := bolt.Open(cfg.DBPath, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{DB: db}
}

// Routes defines all of the routes
func (s *Server) Routes() {
	http.HandleFunc("login", s.HandlerLogin())
}
