package api

import (
	"encoding/json"
	"net/http"
)

type loginResponse struct {
	Token string `json:"token,omitempty"`
}

func Routes() {
	http.HandleFunc("/api", handleAPIRequests)
	http.HandleFunc("/api/login", loginHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {

	}
}

func handleAPIRequests(w http.ResponseWriter, r *http.Request) {
	token := loginResponse{Token: "123456789"}

	js, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
