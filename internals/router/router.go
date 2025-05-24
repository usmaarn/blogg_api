package router

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Router struct {
	middlewares []Middleware
	handler     *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		middlewares: []Middleware{},
		handler:     http.NewServeMux(),
	}
}

func (r *Router) Get(path string, handlerFunc http.HandlerFunc) {
	r.HandleRequest(http.MethodGet, path, handlerFunc)
}

func (r *Router) Post(path string, handlerFunc http.HandlerFunc) {
	r.HandleRequest(http.MethodPost, path, handlerFunc)
}

func (r *Router) Put(path string, handlerFunc http.HandlerFunc) {
	r.HandleRequest(http.MethodPut, path, handlerFunc)
}

func (r *Router) Patch(path string, handlerFunc http.HandlerFunc) {
	r.HandleRequest(http.MethodPatch, path, handlerFunc)
}

func (r *Router) Delete(path string, handlerFunc http.HandlerFunc) {
	r.HandleRequest(http.MethodDelete, path, handlerFunc)
}

func (r *Router) HandleRequest(method string, path string, handler http.HandlerFunc) {
	for _, midleware := range r.middlewares {
		handler = midleware(handler)
	}
	r.handler.HandleFunc(fmt.Sprintf("%s %s", method, path), handler)
}

func (r *Router) Use(middleware Middleware) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) Run(addr string) error {
	server := http.Server{
		Addr:           addr,
		Handler:        r.handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server.ListenAndServe()
}
