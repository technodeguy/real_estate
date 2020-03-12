package controllers

import (
	"github.com/technodeguy/real-estate/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Home route
	s.router.HandleFunc("/", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.HealthCheck))).Methods("GET")

	//User routes
	s.router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.router.HandleFunc("/users/presigned_url", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.GetPresignedUrl))).Methods("POST")
	s.router.HandleFunc("/users/avatar", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuth(s.tokenService, s.SaveUserAvatar))).Methods("PUT")
	s.router.HandleFunc("/users/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

}
