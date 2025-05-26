package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/usmaarn/blogg_api/handlers"
	"github.com/usmaarn/blogg_api/middleware"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("POST /v1/auth/register", handlers.RegisterHandler)
	r.HandleFunc("POST /v1/auth/login", handlers.LoginHandler)
	r.HandleFunc("POST /v1/auth/logout", middleware.AuthMiddleware(nil, handlers.LogoutHandler))
	r.HandleFunc("GET /v1/auth/me", handlers.GetAuthenticatedUserHandler)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	fmt.Printf("App running on http://%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
