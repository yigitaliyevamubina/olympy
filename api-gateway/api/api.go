package api

import (
	"log"

	authhandler "olympy/api-gateway/api/handlers/auth-handlers"        // Updated import path
	countryhandlers "olympy/api-gateway/api/handlers/country-handlers" // Import path for CountryHandlers
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers"     // Updated import path
	medalhandlers "olympy/api-gateway/api/handlers/medal-handlers"     // Import path for MedalHandlers
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
}

func New(
	cfg *config.Config,
	logger *log.Logger,
	authhandler *authhandler.AuthHandlers,
	eventhandler *eventhandlers.EventHandlers,
	countryhandler *countryhandlers.CountryHandlers,
	medalhandler *medalhandlers.MedalHandlers) *API {
	return &API{
		logger:         logger,
		cfg:            cfg,
		authhandler:    authhandler,
		eventhandler:   eventhandler,
		countryhandler: countryhandler,
		medalhandler:   medalhandler,
	}
}

// @title API
// @version 1.0
// @description TEST
// @host localhost:9090
// @BasePath /api/v1
func (a *API) RUN() error {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.POST("/auth/register", a.authhandler.Register)    // Register user
		api.POST("/auth/login", a.authhandler.Login)          // Login user
		api.POST("/auth/refresh", a.authhandler.RefreshToken) // Refresh access token

		api.POST("/events/add", a.eventhandler.AddEvent)             // Add event
		api.POST("/events/edit", a.eventhandler.EditEvent)           // Edit event
		api.DELETE("/events/delete/:id", a.eventhandler.DeleteEvent) // Delete event by ID
		api.GET("/events/get/:id", a.eventhandler.GetEvent)          // Get event by ID
		api.POST("/events/getall", a.eventhandler.GetAllEvents)      // Get all events
		api.POST("/events/search", a.eventhandler.SearchEvents)      // Search events

		api.POST("/countries/add", a.countryhandler.AddCountry)             // Add country
		api.POST("/countries/edit", a.countryhandler.EditCountry)           // Edit country
		api.DELETE("/countries/delete/:id", a.countryhandler.DeleteCountry) // Delete country by ID
		api.GET("/countries/get/:id", a.countryhandler.GetCountry)          // Get country by ID
		api.POST("/countries/getall", a.countryhandler.ListCountries)       // List countries

		api.POST("/medals/add", a.medalhandler.AddMedal)             // Add medal
		api.POST("/medals/edit", a.medalhandler.EditMedal)           // Edit medal
		api.DELETE("/medals/delete/:id", a.medalhandler.DeleteMedal) // Delete medal by ID
		api.GET("/medals/get/:id", a.medalhandler.GetMedal)          // Get medal by ID
		api.POST("/medals/getall", a.medalhandler.ListMedals)        // List medals
		api.POST("/medals/ranking", a.medalhandler.GetMedalRanking)  // Get country rankings sorted by the number of medals
	}

	return router.Run(a.cfg.ServerAddress)
}
