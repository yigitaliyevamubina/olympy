package eventhandlers

import (
	"log"
	medalservice "olympy/api-gateway/genproto/medal_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedalHandlers struct {
	client medalservice.MedalServiceClient
	logger *log.Logger
}

func NewMedalHandlers(client medalservice.MedalServiceClient, logger *log.Logger) *MedalHandlers {
	return &MedalHandlers{
		client: client,
		logger: logger,
	}
}

// AddMedal godoc
// @Summary Add a medal
// @Description This endpoint adds a new medal.
// @Tags Medal
// @Accept json
// @Produce json
// @Param request body medalservice.Medal true "Medal details to add"
// @Success 200 {object} medalservice.Medal
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/add [post]
func (m *MedalHandlers) AddMedal(ctx *gin.Context) {
	var req medalservice.Medal

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := m.client.AddMedal(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// EditMedal godoc
// @Summary Edit a medal
// @Description This endpoint edits an existing medal.
// @Tags Medal
// @Accept json
// @Produce json
// @Param request body medalservice.Medal true "Medal details to edit"
// @Success 200 {object} medalservice.Medal
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/edit [put]
func (m *MedalHandlers) EditMedal(ctx *gin.Context) {
	var req medalservice.Medal

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := m.client.EditMedal(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// DeleteMedal godoc
// @Summary Delete a medal
// @Description This endpoint deletes a medal by its ID.
// @Tags Medal
// @Accept json
// @Produce json
// @Param id query string true "Medal ID to delete"
// @Success 200 {object} medalservice.Message
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/delete [delete]
func (m *MedalHandlers) DeleteMedal(ctx *gin.Context) {
	idStr := ctx.Query("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &medalservice.GetSingleRequest{Id: id}

	resp, err := m.client.DeleteMedal(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetMedal godoc
// @Summary Get a medal
// @Description This endpoint retrieves a medal by its ID.
// @Tags Medal
// @Accept json
// @Produce json
// @Param id query string true "Medal ID to retrieve"
// @Success 200 {object} medalservice.Medal
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/get [get]
func (m *MedalHandlers) GetMedal(ctx *gin.Context) {
	idStr := ctx.Query("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &medalservice.GetSingleRequest{Id: id}

	resp, err := m.client.GetMedal(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// ListMedals godoc
// @Summary List medals
// @Description This endpoint retrieves all medals with pagination and optional filters.
// @Tags Medal
// @Accept json
// @Produce json
// @Param page query int32 false "Page number" default(1)
// @Param limit query int32 false "Number of items per page" default(10)
// @Param country query int64 false "Country ID"
// @Param event_id query int64 false "Event ID"
// @Param athlete_id query string false "Athlete ID"
// @Success 200 {object} medalservice.ListResponse
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/getall [get]
func (m *MedalHandlers) ListMedals(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	countryStr := ctx.Query("country")
	eventIdStr := ctx.Query("event_id")
	athleteId := ctx.Query("athlete_id")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid limit number"})
		return
	}

	var country int64
	if countryStr != "" {
		country, err = strconv.ParseInt(countryStr, 10, 64)
		if err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "Invalid country ID"})
			return
		}
	}

	var eventId int64
	if eventIdStr != "" {
		eventId, err = strconv.ParseInt(eventIdStr, 10, 64)
		if err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "Invalid event ID"})
			return
		}
	}

	req := &medalservice.ListRequest{
		Page:      int32(page),
		Limit:     int32(limit),
		Country:   country,
		EventId:   eventId,
		AthleteId: athleteId,
	}

	resp, err := m.client.ListMedals(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetMedalRanking godoc
// @Summary Get medal rankings
// @Description This endpoint retrieves the ranking of countries based on medals.
// @Tags Medal
// @Accept json
// @Produce json
// @Success 200 {object} medalservice.MedalRankingResponse
// @Failure 500 {object} medalservice.Message
// @Router /medals/ranking [get]
func (m *MedalHandlers) GetMedalRanking(ctx *gin.Context) {
	resp, err := m.client.GetMedalRanking(ctx, &medalservice.Empty{})
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
