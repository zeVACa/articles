package router

import (
	"articles/internal/config"
	"articles/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func ImplementRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// handlers.NewRegisterUserHandler() - вот так тоже ошибки

	// если сделать вот тут "реализацию" хендлера, то опять лезут ошибки

	r.Get("/hello", handlers.)// Вот тут не понимаю как импортировать хендлер и использовать в роере


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
