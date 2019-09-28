package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/teslapi/teslapi/internal/respond"
)

// HandleLogin handles the login requests for the API
func HandleLogin() http.HandlerFunc {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type LoginResponse struct {
		Token string `json:"token"`
	}

	var request = LoginRequest{}
	var response = LoginResponse{}

	return func(w http.ResponseWriter, r *http.Request) {
		// check if the method is post
		if notMethod(r, http.MethodPost) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			resp, err := json.Marshal(&respond.BadRequest{
				Error:  "Method not allowed",
				Status: http.StatusMethodNotAllowed,
			})

			if err != nil {
				log.Fatal(err)
			}

			w.Write(resp)

			return
		}

		// parse the form request
		r.ParseForm()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, err := json.Marshal(&respond.BadRequest{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			})

			if err != nil {
				log.Fatal(err)
			}

			w.Write(resp)

			return
		}

		// verify the email and password
		if request.Email != "jasonmccallister" && request.Password != "Password1!" {
			w.WriteHeader(http.StatusUnauthorized)
			resp, err := json.Marshal(&respond.BadRequest{
				Error:  "invalid credentials",
				Status: http.StatusUnauthorized,
			})

			if err != nil {
				log.Fatal(err)
			}

			w.Write(resp)

			return
		}

		response.Token = "1235678.23456.23456765876"
		body, mrshlErr := json.Marshal(&response)
		if mrshlErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, err := json.Marshal(&respond.BadRequest{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			})

			if err != nil {
				log.Fatal(err)
			}

			w.Write(resp)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func notMethod(r *http.Request, method string) bool {
	return r.Method == method
}
