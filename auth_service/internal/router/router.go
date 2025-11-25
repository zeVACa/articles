package router

import (
	"articles/internal/config"
	"articles/internal/handlers"
	"articles/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func ImplementRouter(s service.Service) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := handlers.NewRegisterUserHandler(s)

	r.Post("/hello", handler.RegisterUser)

	cfg := config.MustLoad()
	srv := http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed to start server")
	}
}
