package service

import (
	"context"
	"fmt"
	"olympy/auth-service/internal/storage"

	genprotos "olympy/auth-service/genproto/auth_service"
)

type AuthServiceServer struct {
	genprotos.UnimplementedAuthServiceServer
	authStorage *storage.AuthService
}

func NewAuthServiceServer(authStorage *storage.AuthService) *AuthServiceServer {
	return &AuthServiceServer{
		authStorage: authStorage,
	}
}

// RegisterUser handles user registration
func (s *AuthServiceServer) RegisterUser(ctx context.Context, req *genprotos.RegisterUserRequest) (*genprotos.RegisterUserResponse, error) {
	resp, err := s.authStorage.RegisterUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error during user registration: %v", err)
	}
	return resp, nil
}

// LoginUser handles user login
func (s *AuthServiceServer) LoginUser(ctx context.Context, req *genprotos.LoginUserRequest) (*genprotos.LoginUserResponse, error) {
	resp, err := s.authStorage.LoginUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("error during user login: %v", err)
	}
	return resp, nil
}
