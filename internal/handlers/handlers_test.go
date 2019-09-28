package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	testCases := []struct {
		desc           string
		method         string
		request        string
		expectedStatus int
	}{
		{
			desc:           "only POST requests are allowed",
			method:         http.MethodGet,
			request:        `{"email": "jason@mccallister.io", "password": "Password1!"}`,
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			desc:           "valid credentials return a successful response",
			method:         http.MethodPost,
			request:        `{"email": "jason@mccallister.io", "password": "Password1!"}`,
			expectedStatus: http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest(tC.method, "login", strings.NewReader(tC.request))
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
