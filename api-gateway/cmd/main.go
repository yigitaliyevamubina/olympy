package main

import (
	"log"
	"olympy/api-gateway/api"
	"olympy/api-gateway/config"
	athleteservice "olympy/api-gateway/genproto/athlete_service"
	authservice "olympy/api-gateway/genproto/auth_service"
	countryservice "olympy/api-gateway/genproto/country_service"
	eventservice "olympy/api-gateway/genproto/event_service"
	medalservice "olympy/api-gateway/genproto/medal_service"
	"os"

	athletehandlers "olympy/api-gateway/api/handlers/athlete-handlers"
	authhandlers "olympy/api-gateway/api/handlers/auth-handlers"
	countryhandlers "olympy/api-gateway/api/handlers/country-handlers"
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers"
	medalhandlers "olympy/api-gateway/api/handlers/medal-handlers"

	"google.golang.org/grpc"
)

func main() {
	var cfg config.Config
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}

	// Connect to auth service
	connAuth, err := grpc.NewClient(cfg.AuthHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer connAuth.Close() // Ensuring connection is closed

	// Connect to event service
	connEvent, err := grpc.NewClient(cfg.EventHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to event service: %v", err)
	}
	defer connEvent.Close() // Ensuring connection is closed

	// Connect to medal service
	connMedal, err := grpc.NewClient(cfg.MedalHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to medal service: %v", err)
	}
	defer connMedal.Close() // Ensuring connection is closed

	// Connect to athlete service
	connAthlete, err := grpc.NewClient(cfg.AthleteHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to athlete service: %v", err)
	}
	defer connAthlete.Close() // Ensuring connection is closed

	// Creating clients for services
	authClient := authservice.NewAuthServiceClient(connAuth)
	eventClient := eventservice.NewEventServiceClient(connEvent)
	countryClient := countryservice.NewCountryServiceClient(connMedal)
	medalClient := medalservice.NewMedalServiceClient(connMedal)
	athleteClient := athleteservice.NewAthleteServiceClient(connAthlete)

	// Creating handler instances
	authHandlers := authhandlers.NewAuthHandlers(authClient, logger)
	eventHandlers := eventhandlers.NewEventHandlers(eventClient, logger)
	countryHandlers := countryhandlers.NewCountryHandlers(countryClient, logger)
	medalHandlers := medalhandlers.NewMedalHandlers(medalClient, logger)
	athleteHandlers := athletehandlers.NewAthleteHandlers(athleteClient, logger)

	// Creating API instance
	api := api.New(&cfg, logger, authHandlers, eventHandlers, countryHandlers, medalHandlers, athleteHandlers)
	logger.Fatal(api.RUN())
}
