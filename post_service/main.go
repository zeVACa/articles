package main

import (
	"articles/internal/config"
	"articles/internal/repository"
	"articles/internal/router"
	"articles/internal/service"
	authClient "articles/pkg/grpc/auth_client"
	"articles/storage"
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9" // ← ДОБАВИЛИ
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Error("Failed to connect to Redis", slog.String("error", err.Error()))
		os.Exit(1)
	}
	log.Info("Redis connected", slog.String("addr", cfg.Redis.Addr))

	cacheService := service.NewCacheService(redisClient, 5*time.Minute)

	db := storage.InitDatabase()
	postsService := service.NewPostsService(
		repository.NewPostsRepository(db),
		cacheService,
	)

	authGrpcClient, err := authClient.NewAuthClient("localhost:50051")
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
