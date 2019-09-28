package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	testCases := []struct {
		desc           string
		method         string
		username       string
		password       string
		request        string
		expectedStatus int
	}{
		{
			desc:           "only POST requests are allowed",
			method:         http.MethodGet,
			username:       "teslapi",
			password:       "Password1!",
			request:        `{"username": "%v", "password": "%v"}`,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			desc:           "valid credentials return a successful response",
			method:         http.MethodPost,
			username:       "teslapi",
			password:       "Password1!",
			request:        `{"username": "%v", "password": "%v"}`,
			expectedStatus: http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			os.Setenv("TESLAPI_USERNAME", tC.username)
			os.Setenv("TESLAPI_PASSWORD", tC.password)
			jsonRequest := strings.NewReader(fmt.Sprintf(tC.request, tC.username, tC.password))
			req, err := http.NewRequest(tC.method, "login", jsonRequest)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := Login()

			// Act
			handler.ServeHTTP(rr, req)

			// Assert
			if status := rr.Code; status != tC.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tC.expectedStatus)
			}
		})
	}
}
