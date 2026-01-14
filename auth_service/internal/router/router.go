package router

import (
	"articles/internal/config"
	"articles/internal/handlers"
	mymiddleware "articles/internal/middleware"
	"articles/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ImplementRouter(s service.Service) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	handler := handlers.NewAuthHandler(s)

	r.Post("/register", handler.RegisterUser)
	r.Post("/login", handler.LoginUser)

	r.Group(func(r chi.Router) {
		r.Use(mymiddleware.AuthMiddleware)

		r.Post("/verify", handler.Verify)
	})

	cfg := config.MustLoad()
	srv := http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("failed to start serverrrrrrrrr %v", err)
	}
}
