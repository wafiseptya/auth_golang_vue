package controllers

import "github.com/wafiseptya/auth/api/middlewares"

func (s *Server) initializeRoutes() {

	// Login Route
	s.Router.HandleFunc("/api/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/api/auth/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")

}
