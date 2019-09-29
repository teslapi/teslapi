package main

import (
	"log"
	"net/http"

	"github.com/teslapi/teslapi/internal/handlers"
	"gocloud.dev/server"
)

func main() {
	srv := server.New(http.DefaultServeMux, nil)

	http.HandleFunc("/api/login", handlers.Login())
	http.HandleFunc("/api/recordings", handlers.Recordings())

	if err := srv.ListenAndServe(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}
