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
// @Accept json
// @Produce json
// @Param request body medalservice.Medal true "Medal details to edit"
// @Success 200 {object} medalservice.Medal
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/edit [post]
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
// @Accept json
// @Produce json
// @Param id path string true "Medal ID to delete"
// @Success 200 {object} medalservice.Message
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/delete/{id} [delete]
func (m *MedalHandlers) DeleteMedal(ctx *gin.Context) {
	idStr := ctx.Param("id")

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
// @Accept json
// @Produce json
// @Param id path string true "Medal ID to retrieve"
// @Success 200 {object} medalservice.Medal
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/get/{id} [get]
func (m *MedalHandlers) GetMedal(ctx *gin.Context) {
	idStr := ctx.Param("id")

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
// @Accept json
// @Produce json
// @Param request body medalservice.ListRequest true "Pagination and filter parameters"
// @Success 200 {object} medalservice.ListResponse
// @Failure 400 {object} medalservice.Message
// @Failure 500 {object} medalservice.Message
// @Router /medals/getall [post]
func (m *MedalHandlers) ListMedals(ctx *gin.Context) {
	var req medalservice.ListRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := m.client.ListMedals(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetMedalRanking godoc
// @Summary Get medal rankings
// @Description This endpoint retrieves the ranking of countries based on medals.
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
