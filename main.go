package main

import (
	//"fmt"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"golang.org/x/net/context"
)

func main() {
	r := chi.NewRouter()

	// use middleware
	r.Use(middleware.Logger)

	// add routing
	r.Get("/", apiIndex)

	r.Route("/todo", func(r chi.Router) {
		r.Get("/", getTodoList)
	})
	http.ListenAndServe(":3333", r)
}

func apiIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}

var tasks = make(map[string]string, 0)

func getTodoList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get todo list"))
}
