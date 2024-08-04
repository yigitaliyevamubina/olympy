package main

import (
	"log"
	"olympy/api-gateway/api"
	"olympy/api-gateway/config"
	authservice "olympy/api-gateway/genproto/auth_service"
	eventservice "olympy/api-gateway/genproto/event_service"

	"os"

	authhandlers "olympy/api-gateway/api/handlers/auth-handlers"
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers"

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}

	conn, err := grpc.Dial(cfg.AuthHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to auth service: %v", err)
	}

	connProduct, err := grpc.Dial(cfg.EventHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to product service: %v", err)
	}

	// auth service
	authClient := authservice.NewAuthServiceClient(conn)

	authHandlers := authhandlers.NewAuthHandlers(authClient, logger)

	// event service

	eventClient := eventservice.NewEventServiceClient(connProduct)

	eventHandlers := eventhandlers.NewEventHandlers(eventClient, logger)

	// model && country service

	// athlete service

	api := api.New(&cfg, logger, authHandlers, eventHandlers)
	logger.Fatal(api.RUN())
}
