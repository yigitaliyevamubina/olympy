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
	streamservice "olympy/api-gateway/genproto/stream_service"
	"os"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"

	athletehandlers "olympy/api-gateway/api/handlers/athlete-handlers"
	authhandlers "olympy/api-gateway/api/handlers/auth-handlers"
	countryhandlers "olympy/api-gateway/api/handlers/country-handlers"
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers"
	medalhandlers "olympy/api-gateway/api/handlers/medal-handlers"
	streamhandlers "olympy/api-gateway/api/handlers/stream-handlers"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := cfg.Load(); err != nil {
		logger.Fatal(err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
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

	// Connect to auth service
	connAuth, err := grpc.Dial(cfg.AuthHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to auth service: %v", err)
	}
	defer connAuth.Close()

	// Connect to event service
	connEvent, err := grpc.Dial(cfg.EventHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to event service: %v", err)
	}
	defer connEvent.Close()

	// Connect to medal service
	connMedal, err := grpc.Dial(cfg.MedalHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to medal service: %v", err)
	}
	defer connMedal.Close()

	// Connect to athlete service
	connAthlete, err := grpc.Dial(cfg.AthleteHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to athlete service: %v", err)
	}
	defer connAthlete.Close()

	// Connect to stream service
	connStream, err := grpc.Dial(cfg.StreamHost, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("Failed to connect to stream service: %v", err)
	}
	defer connStream.Close()

	// Creating clients for services
	authClient := authservice.NewAuthServiceClient(connAuth)
	eventClient := eventservice.NewEventServiceClient(connEvent)
	countryClient := countryservice.NewCountryServiceClient(connMedal)
	medalClient := medalservice.NewMedalServiceClient(connMedal)
	athleteClient := athleteservice.NewAthleteServiceClient(connAthlete)
	streamClient := streamservice.NewStreamingServiceClient(connStream)

	// Creating handler instances
	authHandlers := authhandlers.NewAuthHandlers(authClient, logger, ch)
	eventHandlers := eventhandlers.NewEventHandlers(eventClient, logger)
	countryHandlers := countryhandlers.NewCountryHandlers(countryClient, logger)
	medalHandlers := medalhandlers.NewMedalHandlers(medalClient, logger)
	athleteHandlers := athletehandlers.NewAthleteHandlers(athleteClient, logger)
	streamHandlers := streamhandlers.NewStreamHandlers(streamClient, logger)
	// Creating API instance
	api := api.New(cfg, logger, authHandlers, eventHandlers, countryHandlers, medalHandlers, athleteHandlers, streamHandlers)
	logger.Fatal(api.RUN())
}
