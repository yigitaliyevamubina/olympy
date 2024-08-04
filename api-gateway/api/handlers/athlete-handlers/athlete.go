package athletehandlers

import (
	"log"
	athleteservice "olympy/api-gateway/genproto/athlete_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AthleteHandlers struct {
	client athleteservice.AthleteServiceClient
	logger *log.Logger
}

func NewAthleteHandlers(client athleteservice.AthleteServiceClient, logger *log.Logger) *AthleteHandlers {
	return &AthleteHandlers{
		client: client,
		logger: logger,
	}
}

// AddAthlete godoc
// @Summary Add an athlete
// @Description This endpoint adds a new athlete.
// @Accept json
// @Produce json
// @Param request body athleteservice.Athlete true "Athlete details to add"
// @Success 200 {object} athleteservice.Athlete
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/add [post]
func (a *AthleteHandlers) AddAthlete(ctx *gin.Context) {
	var req athleteservice.Athlete

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.AddAthlete(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// EditAthlete godoc
// @Summary Edit an athlete
// @Description This endpoint edits an existing athlete.
// @Accept json
// @Produce json
// @Param request body athleteservice.Athlete true "Athlete details to edit"
// @Success 200 {object} athleteservice.Athlete
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/edit [post]
func (a *AthleteHandlers) EditAthlete(ctx *gin.Context) {
	var req athleteservice.Athlete

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.EditAthlete(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// DeleteAthlete godoc
// @Summary Delete an athlete
// @Description This endpoint deletes an athlete by its ID.
// @Accept json
// @Produce json
// @Param id path string true "Athlete ID to delete"
// @Success 200 {object} athleteservice.Message
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/delete/{id} [delete]
func (a *AthleteHandlers) DeleteAthlete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &athleteservice.GetSingleRequest{Id: id}

	resp, err := a.client.DeleteAthlete(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetAthlete godoc
// @Summary Get an athlete
// @Description This endpoint retrieves an athlete by its ID.
// @Accept json
// @Produce json
// @Param id path string true "Athlete ID to retrieve"
// @Success 200 {object} athleteservice.Athlete
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/get/{id} [get]
func (a *AthleteHandlers) GetAthlete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &athleteservice.GetSingleRequest{Id: id}

	resp, err := a.client.GetAthlete(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// ListAthletes godoc
// @Summary List athletes
// @Description This endpoint retrieves all athletes with pagination and optional filters.
// @Accept json
// @Produce json
// @Param request body athleteservice.ListRequest true "Pagination and filter parameters"
// @Success 200 {object} athleteservice.ListResponse
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/getall [post]
func (a *AthleteHandlers) ListAthletes(ctx *gin.Context) {
	var req athleteservice.ListRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.ListAthletes(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
