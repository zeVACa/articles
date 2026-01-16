package main

import (
	"articles/internal/config"
	"articles/internal/repository"
	"articles/internal/router"
	"articles/internal/service"
	authClient "articles/pkg/grpc/auth_client"
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
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("App started", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db := storage.InitDatabase()
	postsService := service.NewPostsService(repository.NewPostsRepository(db))

	authGrpcClient, err := authClient.NewAuthClient("localhost:50051") // ← ИСПРАВЛЕНО!
	if err != nil {
		log.Error("Failed to connect to auth service", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer authGrpcClient.Close()

	router.ImplementRouter(postsService, authGrpcClient)

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
