package api

import (
	"encoding/json"
	"log"
	"net/http"
)



// ListenAndServe takes the
func (s *Server) ListenAndServe(addr string, handler http.Handler) error {
	server := &http.Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
