package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/teslapi/teslapi/internal/handlers"
)

var requiredEnvs = []string{
	"TESLAPI_KEY",
	"TESLAPI_USERNAME",
	"TESLAPI_PASSWORD",
}

func init() {
	// check the environment variables that are required for the app
	for _, e := range requiredEnvs {
		if env := os.Getenv(e); env == "" {
			log.Println(env)
			log.Fatalf("expected the environment variable %v to be defined", e)
		}
	}
}

func main() {
	// scan the directory for files

	// open the database to store clips in
	db, err := badger.Open(badger.DefaultOptions("../../storage/db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", handlers.UI(db))
	http.HandleFunc("/login", handlers.Login())
	http.HandleFunc("/api/login", handlers.Login())
	http.HandleFunc("/api/recordings", handlers.Recordings())

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("%v", err)
	}
}
