package eventhandlers

import (
	"log"
	eventservice "olympy/api-gateway/genproto/event_service"

	"github.com/gin-gonic/gin"
)

type EventHandlers struct {
	client eventservice.EventServiceClient
	logger *log.Logger
}

func NewEventHandlers(client eventservice.EventServiceClient, logger *log.Logger) *EventHandlers {
	return &EventHandlers{
		client: client,
		logger: logger,
	}
}

// AddEvent godoc
// @Summary Add an event
// @Description This endpoint adds a new event.
// @Accept json
// @Produce json
// @Param request body eventservice.AddEventRequest true "Event details to add"
// @Success 200 {object} eventservice.AddEventResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/add [post]
func (e *EventHandlers) AddEvent(ctx *gin.Context) {
	var req eventservice.AddEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.client.AddEvent(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// EditEvent godoc
// @Summary Edit an event
// @Description This endpoint edits an existing event.
// @Accept json
// @Produce json
// @Param request body eventservice.EditEventRequest true "Event details to edit"
// @Success 200 {object} eventservice.EditEventResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/edit [post]
func (e *EventHandlers) EditEvent(ctx *gin.Context) {
	var req eventservice.EditEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.client.EditEvent(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// DeleteEvent godoc
// @Summary Delete an event
// @Description This endpoint deletes an event by its ID.
// @Accept json
// @Produce json
// @Param request body eventservice.DeleteEventRequest true "Event ID to delete"
// @Success 200 {object} eventservice.Message
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/delete [post]
func (e *EventHandlers) DeleteEvent(ctx *gin.Context) {
	var req eventservice.DeleteEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.client.DeleteEvent(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetEvent godoc
// @Summary Get an event
// @Description This endpoint retrieves an event by its ID.
// @Accept json
// @Produce json
// @Param request body eventservice.GetEventRequest true "Event ID to retrieve"
// @Success 200 {object} eventservice.GetEventResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/get [post]
func (e *EventHandlers) GetEvent(ctx *gin.Context) {
	var req eventservice.GetEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.client.GetEvent(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetAllEvents godoc
// @Summary Get all events
// @Description This endpoint retrieves all events with pagination.
// @Accept json
// @Produce json
// @Param request body eventservice.GetAllEventsRequest true "Pagination parameters"
// @Success 200 {object} eventservice.GetAllEventsResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/getall [post]
func (e *EventHandlers) GetAllEvents(ctx *gin.Context) {
	var req eventservice.GetAllEventsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.client.GetAllEvents(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// SearchEvents godoc
// @Summary Search events
// @Description This endpoint searches events by query with pagination.
// @Accept json
// @Produce json
// @Param request body eventservice.SearchEventsRequest true "Search parameters"
// @Success 200 {object} eventservice.GetAllEventsResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/search [post]
func (e *EventHandlers) SearchEvents(ctx *gin.Context) {
	var req eventservice.SearchEventsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.client.SearchEvents(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
