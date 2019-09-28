package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Login handles login requests to the API
func Login() http.HandlerFunc {
	type loginRequest struct {
		Username string `json:"username"`
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

		if request.Username != os.Getenv("TESLAPI_USERNAME") || request.Password != os.Getenv("TESLAPI_PASSWORD") {
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
