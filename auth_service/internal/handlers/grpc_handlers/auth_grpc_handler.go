package grpc_handlers

import (
	"articles/pkg/grpc/pb"
	"articles/pkg/jwtPkg"
	"context"
	"fmt"
	"log"
	"strconv"
)

type AuthGRPCHandler struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthGRPCHandler() *AuthGRPCHandler {
	return &AuthGRPCHandler{}
}

func (h *AuthGRPCHandler) VerifyToken(ctx context.Context, req *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	claims, err := jwtPkg.ValidateToken(req.GetToken())
	if err != nil {
		log.Printf("gRPC: Invalid token: %v", err)
		return &pb.VerifyTokenResponse{
			IsValid: false,
			UserId:  0,
		}, nil
	}

	userID := claims.ID
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println(err)
	}
	userIDInt64 := int64(userIDInt)

	log.Printf("gRPC: Token valid, user_id: %d", userID)

	return &pb.VerifyTokenResponse{
		IsValid: true,
		UserId:  userIDInt64,
	}, nil
}
