package main

import "testing"

func TestLogin(t *testing.T) {
	// Arrange
	// setup temp boltdb
	// create dummy username and password

	// Act & Assert
	testCases := []struct {
		desc       string
		email      string
		password   string
		statusCode int
	}{
		{
			desc:       "valid logins return a valid JWT",
			email:      "jason@mccallister.io",
			password:   "somedumbpassword",
			statusCode: 200,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
