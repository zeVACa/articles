package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello2"))
	if err != nil {

	}
}

func TestHandler2(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello2"))
	if err != nil {

	}
}

func TestHandler3(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello2"))
	if err != nil {

	}
}

func ImplementRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/hello", TestHandler)
	r.Get("/hello", TestHandler2)
	r.Get("/hello", TestHandler3)
}
