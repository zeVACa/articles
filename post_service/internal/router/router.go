package router

import (
	"articles/internal/config"
	"articles/internal/handlers"
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

	handler := handlers.NewPostsHandler(s)

	r.Get("/posts", handler.GetAllPosts)
	r.Get("/posts/{id}", handler.GetPostById)
	r.Post("/posts", handler.CreatePost)

	cfg := config.MustLoad()
	srv := http.Server{
		Addr:         cfg.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("failed to start server: %v", err)
	}
}
