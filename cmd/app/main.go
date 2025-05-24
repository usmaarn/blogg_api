package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/usmaarn/blogg_api/internals/middlewares"
	"github.com/usmaarn/blogg_api/internals/router"
)

func main() {

	r := router.NewRouter()

	r.Use(middlewares.Logger)

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Hello from homepage")
	})

	// r.HandleFunc("POST /uploade", func(w http.ResponseWriter, r *http.Request) {

	// })

	log.Fatal(r.Run(":8080"))
}
