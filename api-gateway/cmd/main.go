package main

import (
	"log"
	"olympy/api-gateway/api"
	"olympy/api-gateway/config"
	genprotos "olympy/api-gateway/genprotos"
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

	authClient := genprotos.NewAuthServiceClient(conn)

	authHandlers := authhandlers.NewAuthHandlers(authClient, logger)

	eventClient := genprotos.NewEventServiceClient(connProduct)

	eventHandlers := eventhandlers.NewEventHandlers(eventClient, logger)

	api := api.New(&cfg, logger, authHandlers, eventHandlers)
	logger.Fatal(api.RUN())
}
