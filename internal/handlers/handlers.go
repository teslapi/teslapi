package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 30).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
			Issuer:    "teslapi",
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("TESLAPI_KEY")))
		if err != nil {
			log.Fatal(err)
			return
		}

		response := loginResponse{Token: tokenString}
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
