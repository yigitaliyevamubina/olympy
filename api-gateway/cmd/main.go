package main

import (
	"log"
	"olympy/api-gateway/api"
	"olympy/api-gateway/config"
	authservice "olympy/api-gateway/genproto/auth_service"
	countryservice "olympy/api-gateway/genproto/country_service" // Import for CountryService
	eventservice "olympy/api-gateway/genproto/event_service"
	medalservice "olympy/api-gateway/genproto/medal_service" // Import for MedalService
	"os"

	authhandlers "olympy/api-gateway/api/handlers/auth-handlers"
	countryhandlers "olympy/api-gateway/api/handlers/country-handlers" // Import for CountryHandlers
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers"
	medalhandlers "olympy/api-gateway/api/handlers/medal-handlers" // Import for MedalHandlers

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}

	// Connect to auth service
	connAuth, err := grpc.Dial(cfg.AuthHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer connAuth.Close() // Ensure connection is closed

	// Connect to event service
	connEvent, err := grpc.Dial(cfg.EventHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to event service: %v", err)
	}
	defer connEvent.Close() // Ensure connection is closed

	// Connect to medal && country service
	connMedal, err := grpc.Dial(cfg.MedalHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to medal service: %v", err)
	}
	defer connMedal.Close() // Ensure connection is closed

	// Create clients for services
	authClient := authservice.NewAuthServiceClient(connAuth)
	eventClient := eventservice.NewEventServiceClient(connEvent)
	countryClient := countryservice.NewCountryServiceClient(connMedal)
	medalClient := medalservice.NewMedalServiceClient(connMedal)

	// Create handler instances
	authHandlers := authhandlers.NewAuthHandlers(authClient, logger)
	eventHandlers := eventhandlers.NewEventHandlers(eventClient, logger)
	countryHandlers := countryhandlers.NewCountryHandlers(countryClient, logger)
	medalHandlers := medalhandlers.NewMedalHandlers(medalClient, logger)

	// Create API instance
	api := api.New(&cfg, logger, authHandlers, eventHandlers, countryHandlers, medalHandlers)
	logger.Fatal(api.RUN())
}
