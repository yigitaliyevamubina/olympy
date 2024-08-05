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
// @Tags Athlete
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
// @Tags Athlete
// @Accept json
// @Produce json
// @Param request body athleteservice.Athlete true "Athlete details to edit"
// @Success 200 {object} athleteservice.Athlete
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/edit [put]
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
// @Tags Athlete
// @Accept json
// @Produce json
// @Param id query string true "Athlete ID to delete"
// @Success 200 {object} athleteservice.Message
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/delete [delete]
func (a *AthleteHandlers) DeleteAthlete(ctx *gin.Context) {
	idStr := ctx.Query("id")

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
// @Tags Athlete
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "Athlete ID to get"
// @Success 200 {object} athleteservice.Athlete
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/get [get]
func (a *AthleteHandlers) GetAthlete(ctx *gin.Context) {
	idStr := ctx.Query("id")

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
// @Tags Athlete
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int32 false "Page number" default(1)
// @Param limit query int32 false "Number of items per page" default(10)
// @Param country_id query int64 false "Country ID filter"
// @Param sport_type query string false "Sport type filter"
// @Success 200 {object} athleteservice.ListResponse
// @Failure 400 {object} athleteservice.Message
// @Failure 500 {object} athleteservice.Message
// @Router /athletes/getall [get]
func (a *AthleteHandlers) ListAthletes(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	countryIDStr := ctx.Query("country_id")
	sportType := ctx.Query("sport_type")

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

	var countryID int64
	if countryIDStr != "" {
		countryID, err = strconv.ParseInt(countryIDStr, 10, 64)
		if err != nil {
			ctx.IndentedJSON(400, gin.H{"error": "Invalid country ID"})
			return
		}
	}

	req := &athleteservice.ListRequest{
		Page:      int32(page),
		Limit:     int32(limit),
		CountryId: countryID,
		SportType: sportType,
	}

	resp, err := a.client.ListAthletes(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
