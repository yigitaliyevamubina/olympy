package eventhandlers

import (
	"log"
	eventservice "olympy/api-gateway/genproto/event_service"
	"strconv"

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
// @Tags Event
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
// @Tags Event
// @Accept json
// @Produce json
// @Param request body eventservice.EditEventRequest true "Event details to edit"
// @Success 200 {object} eventservice.EditEventResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/edit [put]
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
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "Event ID to delete"
// @Success 200 {object} eventservice.Message
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/delete/{id} [delete]
func (e *EventHandlers) DeleteEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")

	req := &eventservice.DeleteEventRequest{Id: idStr}

	resp, err := e.client.DeleteEvent(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetEvent godoc
// @Summary Get an event
// @Description This endpoint retrieves an event by its ID.
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "Event ID to retrieve"
// @Success 200 {object} eventservice.GetEventResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/get/{id} [get]
func (e *EventHandlers) GetEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")

	req := &eventservice.GetEventRequest{Id: idStr}

	resp, err := e.client.GetEvent(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetAllEvents godoc
// @Summary Get all events
// @Description This endpoint retrieves all events with pagination.
// @Tags Event
// @Accept json
// @Produce json
// @Param page query int32 false "Page number" default(1)
// @Param page_size query int32 false "Number of items per page" default(10)
// @Success 200 {object} eventservice.GetAllEventsResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/getall [get]
func (e *EventHandlers) GetAllEvents(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page size"})
		return
	}

	req := &eventservice.GetAllEventsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

	resp, err := e.client.GetAllEvents(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// SearchEvents godoc
// @Summary Search events
// @Description This endpoint searches events by query with pagination.
// @Tags Event
// @Accept json
// @Produce json
// @Param query query string false "Search query"
// @Param page query int32 false "Page number" default(1)
// @Param page_size query int32 false "Number of items per page" default(10)
// @Success 200 {object} eventservice.GetAllEventsResponse
// @Failure 400 {object} eventservice.Message
// @Failure 500 {object} eventservice.Message
// @Router /events/search [get]
func (e *EventHandlers) SearchEvents(ctx *gin.Context) {
	query := ctx.Query("query")
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page size"})
		return
	}

	req := &eventservice.SearchEventsRequest{
		Query:    query,
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

	resp, err := e.client.SearchEvents(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
