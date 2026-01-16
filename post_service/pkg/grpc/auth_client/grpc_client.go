package auth_client

import (
	"articles/pkg/grpc/pb"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	client pb.AuthServiceClient
	conn   *grpc.ClientConn
}

// NewAuthClient подключается к auth_service по gRPC
func NewAuthClient(authServiceAddr string) (*AuthClient, error) {
	log.Printf("Connecting to auth service at %s...", authServiceAddr)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		authServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // Ждем подключения
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	client := pb.NewAuthServiceClient(conn)
	log.Printf("✅ Connected to auth service")

	return &AuthClient{
		client: client,
		conn:   conn,
	}, nil
}

// VerifyToken проверяет токен через gRPC
func (c *AuthClient) VerifyToken(token string) (bool, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := c.client.VerifyToken(ctx, &pb.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		return false, 0, fmt.Errorf("grpc error: %w", err)
	}

	return resp.GetIsValid(), resp.GetUserId(), nil
}

// Close закрывает соединение
func (c *AuthClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
