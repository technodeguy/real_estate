package controllers

import (
	"fmt"
	"net/http"
)

func (server *Server) HealthCheck(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)

	fmt.Fprint(w, map[string]string{"established": "true"})
}
