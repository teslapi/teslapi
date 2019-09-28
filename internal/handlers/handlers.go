package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Login handles login requests to the API
func Login() http.HandlerFunc {
	// TODO remove these and place into db
	email := "jason@mccallister.io"
	password := "Password1!"

	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type loginResponse struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()

		request := loginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		if request.Email != email || request.Password != password {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("invalid request")
			return
		}

		response := loginResponse{Token: "1234.56789.0987654321"}
		body, err := json.Marshal(&response)
		if err != nil {
			log.Fatal(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}
