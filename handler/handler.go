package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/blogg/store"
)

type Handler struct {
	Store    *store.Store
	Validate *validator.Validate
}

func New(store *store.Store) *Handler {

	return &Handler{
		Store:    store,
		Validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}
