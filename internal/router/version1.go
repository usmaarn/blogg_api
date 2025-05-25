package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/usmaarn/blogg_api/internal/handler"
)

func V1Router(r chi.Router) {
	//Authentication
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.RegisterUser)
	})

	//
}
