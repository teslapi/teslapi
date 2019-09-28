package main

import (
	"log"
	"net/http"

	"gocloud.dev/server"
)

func main() {
	srv := server.New(http.DefaultServeMux, nil)

	routes()

	if err := srv.ListenAndServe(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}

func routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
