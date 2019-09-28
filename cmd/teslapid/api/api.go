package main

import (
	"encoding/json"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
	"gocloud.dev/server"
)

type Config struct {
	DB *bolt.DB
}

func main() {
	srv := server.New(http.DefaultServeMux, nil)

	// Register a route.
	http.HandleFunc("/", login())

	// Start the server. If ListenAndServe returns an error, print it and exit.
	if err := srv.ListenAndServe(":8080"); err != nil {
		log.Fatalf("%v", err)
	}
}

func badRequest(message string, status int) {

}

func login() http.HandlerFunc {
	username := "jasonmcccallister"
	password := "Password1!"

	type loginRequest struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}

	type loginResponse struct {
		Token string `json:"token,omitempty"`
	}

	request := loginRequest{}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if request.Username != username && request.Password != password {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO sign a JWT
		response := loginResponse{
			Token: "12345678",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
