package main

import (
	"articles/internal/config"
	"articles/internal/repository"
	"articles/internal/router"
	"articles/internal/service"
	"articles/storage"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO config
	// TODO logger
	// TODO init storage
	// TODO init router
	// TODO run server
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("App started", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	//var pool pgxpool.Pool
	//r := repository.NewAuthRepository(&pool)
	//srvs := service.NewAuthService(r)
	//h := handlers.RegisterUserHandlerDI{}
	db := storage.InitDatabase()
	authService := service.NewAuthService(repository.NewAuthRepository(db))

	router.ImplementRouter(authService)

}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
