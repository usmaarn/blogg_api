package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}

	fmt.Printf("App running on http://%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
