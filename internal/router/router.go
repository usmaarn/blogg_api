package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Initialize(addr string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//V1 Routes
	r.Route("/v1", V1Router)

	http.ListenAndServe(addr, r)
}
