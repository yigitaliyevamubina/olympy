package eventhandlers

import (
	"log"
	countryservice "olympy/api-gateway/genproto/country_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CountryHandlers struct {
	client countryservice.CountryServiceClient
	logger *log.Logger
}

func NewCountryHandlers(client countryservice.CountryServiceClient, logger *log.Logger) *CountryHandlers {
	return &CountryHandlers{
		client: client,
		logger: logger,
	}
}

// AddCountry godoc
// @Summary Add a country
// @Description This endpoint adds a new country.
// @Tags Country
// @Accept json
// @Produce json
// @Param request body countryservice.Country true "Country details to add"
// @Success 200 {object} countryservice.Country
// @Failure 400 {object} countryservice.Message
// @Failure 500 {object} countryservice.Message
// @Router /countries/add [post]
func (c *CountryHandlers) AddCountry(ctx *gin.Context) {
	var req countryservice.Country

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.client.AddCountry(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// EditCountry godoc
// @Summary Edit a country
// @Description This endpoint edits an existing country.
// @Tags Country
// @Accept json
// @Produce json
// @Param request body countryservice.Country true "Country details to edit"
// @Success 200 {object} countryservice.Country
// @Failure 400 {object} countryservice.Message
// @Failure 500 {object} countryservice.Message
// @Router /countries/edit [put]
func (c *CountryHandlers) EditCountry(ctx *gin.Context) {
	var req countryservice.Country

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.client.EditCountry(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description This endpoint deletes a country by its ID.
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country ID to delete"
// @Success 200 {object} countryservice.Message
// @Failure 400 {object} countryservice.Message
// @Failure 500 {object} countryservice.Message
// @Router /countries/delete/{id} [delete]
func (c *CountryHandlers) DeleteCountry(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &countryservice.GetSingleRequest{Id: id}

	resp, err := c.client.DeleteCountry(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// GetCountry godoc
// @Summary Get a country
// @Description This endpoint retrieves a country by its ID.
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "Country ID to retrieve"
// @Success 200 {object} countryservice.Country
// @Failure 400 {object} countryservice.Message
// @Failure 500 {object} countryservice.Message
// @Router /countries/get/{id} [get]
func (c *CountryHandlers) GetCountry(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	req := &countryservice.GetSingleRequest{Id: id}

	resp, err := c.client.GetCountry(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// ListCountries godoc
// @Summary List countries
// @Description This endpoint retrieves all countries with pagination.
// @Tags Country
// @Accept json
// @Produce json
// @Param page query int32 false "Page number" default(1)
// @Param limit query int32 false "Number of items per page" default(10)
// @Success 200 {object} countryservice.ListResponse
// @Failure 400 {object} countryservice.Message
// @Failure 500 {object} countryservice.Message
// @Router /countries/getall [get]
func (c *CountryHandlers) ListCountries(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

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

	req := &countryservice.ListRequest{
		Page:  int32(page),
		Limit: int32(limit),
	}

	resp, err := c.client.ListCountries(ctx, req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
