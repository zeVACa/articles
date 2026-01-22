package main

import (
	"articles/internal/config"
	grpcHandler "articles/internal/handlers/grpc_handlers"
	"articles/internal/kafka"
	"articles/internal/repository"
	"articles/internal/router"
	"articles/internal/service"
	"articles/pkg/grpc/pb"
	"articles/storage"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	logger := setupLogger(cfg.Env)

	logger.Info("App started", slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	kafkaBrokers := []string{"localhost:9092"}
	kafkaProducer, err := kafka.NewProducer(kafkaBrokers)
	if err != nil {
		logger.Error("Failed to create Kafka producer", slog.String("error", err.Error()))
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()
	logger.Info("Kafka producer initialized successfully")

	db := storage.InitDatabase()
	authService := service.NewAuthService(repository.NewAuthRepository(db))

	go func() {
		logger.Info("Starting HTTP server", slog.String("port", ":8080"))
		router.ImplementRouter(authService, kafkaProducer)
	}()

	go func() {
		log.Println("Starting gRPC server...")

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen on :50051: %v", err)
		}

		grpcServer := grpc.NewServer()

		authHandler := grpcHandler.NewAuthGRPCHandler()
		pb.RegisterAuthServiceServer(grpcServer, authHandler)

		log.Println("gRPC server listening on :50051")

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down servers...")
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
