package api

import (
	"log"

	athletehandlers "olympy/api-gateway/api/handlers/athlete-handlers" // Import path for AthleteHandlers
	authhandler "olympy/api-gateway/api/handlers/auth-handlers"        // Updated import path
	countryhandlers "olympy/api-gateway/api/handlers/country-handlers" // Import path for CountryHandlers
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers"     // Updated import path
	medalhandlers "olympy/api-gateway/api/handlers/medal-handlers"     // Import path for MedalHandlers
	streamhandlers "olympy/api-gateway/api/handlers/stream-handlers"
	"olympy/api-gateway/api/middleware/casbin"
	"olympy/api-gateway/config"
	_ "olympy/api-gateway/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	logger         *log.Logger
	cfg            *config.Config
	authhandler    *authhandler.AuthHandlers
	eventhandler   *eventhandlers.EventHandlers
	countryhandler *countryhandlers.CountryHandlers
	medalhandler   *medalhandlers.MedalHandlers
	athletehandler *athletehandlers.AthleteHandlers
	streamhandlers *streamhandlers.StreamHandlers
}

func New(
	cfg *config.Config,
	logger *log.Logger,
	authhandler *authhandler.AuthHandlers,
	eventhandler *eventhandlers.EventHandlers,
	countryhandler *countryhandlers.CountryHandlers,
	medalhandler *medalhandlers.MedalHandlers,
	athletehandler *athletehandlers.AthleteHandlers,
	streamhandler *streamhandlers.StreamHandlers,
) *API {
	return &API{
		logger:         logger,
		cfg:            cfg,
		authhandler:    authhandler,
		eventhandler:   eventhandler,
		countryhandler: countryhandler,
		medalhandler:   medalhandler,
		athletehandler: athletehandler,
		streamhandlers: streamhandler,
	}
}


// NewRoute
// @title API
// @description TEST
// @BasePath /api/v1
// @version 1.7
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func (a *API) RUN() error {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(casbin.NewAuthorizer())

	api := router.Group("/api/v1")
	{
		api.POST("/auth/register", a.authhandler.Register)    // Register user
		api.POST("/auth/login", a.authhandler.Login)          // Login user
		api.POST("/auth/refresh", a.authhandler.RefreshToken) // Refresh access token

		api.POST("/events/add", a.eventhandler.AddEvent)             // Add event
		api.PUT("/events/edit", a.eventhandler.EditEvent)            // Edit event
		api.DELETE("/events/delete", a.eventhandler.DeleteEvent) // Delete event by ID
		api.GET("/events/get", a.eventhandler.GetEvent)          // Get event by ID
		api.GET("/events/getall", a.eventhandler.GetAllEvents)       // Get all events
		api.GET("/events/search", a.eventhandler.SearchEvents)       // Search events

		api.POST("/countries/add", a.countryhandler.AddCountry)             // Add country
		api.PUT("/countries/edit", a.countryhandler.EditCountry)            // Edit country
		api.DELETE("/countries/delete", a.countryhandler.DeleteCountry) // Delete country by ID
		api.GET("/countries/get", a.countryhandler.GetCountry)          // Get country by ID
		api.GET("/countries/getall", a.countryhandler.ListCountries)        // List countries

		api.POST("/medals/add", a.medalhandler.AddMedal)             // Add medal
		api.PUT("/medals/edit", a.medalhandler.EditMedal)            // Edit medal
		api.DELETE("/medals/delete", a.medalhandler.DeleteMedal) // Delete medal by ID
		api.GET("/medals/get", a.medalhandler.GetMedal)          // Get medal by ID
		api.GET("/medals/getall", a.medalhandler.ListMedals)         // List medals
		api.GET("/medals/ranking", a.medalhandler.GetMedalRanking)   // Get country rankings sorted by the number of medals

		api.POST("/athletes/add", a.athletehandler.AddAthlete)             // Add athlete
		api.PUT("/athletes/edit", a.athletehandler.EditAthlete)            // Edit athlete
		api.DELETE("/athletes/delete", a.athletehandler.DeleteAthlete) // Delete athlete by ID
		api.GET("/athletes/get", a.athletehandler.GetAthlete)          // Get athlete by ID
		api.GET("/athletes/getall", a.athletehandler.ListAthletes)         // List athletes
		api.POST("/stream/send", a.streamhandlers.SendEvent)               // Send

	}

	return router.Run(a.cfg.ServerAddress)
}
