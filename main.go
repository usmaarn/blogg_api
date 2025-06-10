package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/usmaarn/blogg/database"
	"github.com/usmaarn/blogg/handler"
	"github.com/usmaarn/blogg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := database.Init()
	repo := store.New(db)
	h := handler.New(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//Post Routes
	r.Get("/", h.PostsHandler)
	r.Get("/create-post", h.CreatePostHandler)
	r.Post("/create-post", h.StorePostHandler)

	//Assets Routes
	r.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/assets", http.FileServer(http.Dir("./public")))
		fs.ServeHTTP(w, r)
	})

	fmt.Println("app running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
