package middlewares

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer fmt.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
