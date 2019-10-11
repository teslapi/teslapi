package middleware

import (
	"errors"
	"net/http"
)

// Authorize takes a request and verifies the JWT against the signed token
// if there is no authorization header or the token does not have a valid
// signature it will return an error.
func Authorize(r *http.Request) error {
	header := r.Header.Get("Authorization")

	switch header {
	case "":
		return errors.New("Unauthorized request")
	}

	// get the bearer token

	// parse the bearer token

	// verify the token signature matches

	return nil
}
