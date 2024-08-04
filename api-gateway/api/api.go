package api

import (
	"log"

	authhandler "olympy/api-gateway/api/handlers/auth-handlers"    // Updated import path
	eventhandlers "olympy/api-gateway/api/handlers/event-handlers" // Updated import path
	"olympy/api-gateway/config"
	_ "olympy/api-gateway/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	logger       *log.Logger
	cfg          *config.Config
	authhandler  *authhandler.AuthHandlers
	eventhandler *eventhandlers.EventHandlers
}

func New(
	cfg *config.Config,
	logger *log.Logger,
	authhandler *authhandler.AuthHandlers,
	eventhandler *eventhandlers.EventHandlers) *API {
	return &API{
		logger:       logger,
		cfg:          cfg,
		authhandler:  authhandler,
		eventhandler: eventhandler,
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

		api.POST("/events/add", a.eventhandler.AddEvent)        // Add event
		api.POST("/events/edit", a.eventhandler.EditEvent)      // Edit event
		api.POST("/events/delete", a.eventhandler.DeleteEvent)  // Delete event
		api.POST("/events/get", a.eventhandler.GetEvent)        // Get event
		api.POST("/events/getall", a.eventhandler.GetAllEvents) // Get all events
		api.POST("/events/search", a.eventhandler.SearchEvents) // Search events
	}

	return router.Run(a.cfg.ServerAddress)
}
