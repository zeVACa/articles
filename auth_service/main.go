package main

import (
	"articles/internal/config"
	"articles/internal/router"
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

	storage.InitDatabase()
	router.ImplementRouter()
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger // ✅ Указатель!

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
