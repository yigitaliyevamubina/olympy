package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"olympy/auth-service/internal/storage"

	"github.com/streadway/amqp"

	"olympy/auth-service/genproto/auth_service"
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

func (s *AuthServiceServer) RefreshToken(ctx context.Context, req *genprotos.RefreshTokenRequest) (*genprotos.RefreshTokenResponse, error) {
	return s.authStorage.RefreshToken(ctx, req)
}

func (s *AuthServiceServer) StartRabbitMQConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"register_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var req auth_service.RegisterUserRequest
			if err := json.Unmarshal(d.Body, &req); err != nil {
				log.Printf("Error decoding JSON: %v", err)
				continue
			}

			resp, err := s.RegisterUser(context.Background(), &req)
			if err != nil {
				log.Printf("Error registering user: %v", err)
				continue
			}

			log.Printf("User registered successfully: %v", resp)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
