package authhandlers

import (
	"encoding/json"
	"log"
	"net/http"
	genprotos "olympy/api-gateway/genproto/auth_service"
	"github.com/streadway/amqp"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	client          genprotos.AuthServiceClient
	logger          *log.Logger
	rabbitMQChannel *amqp.Channel
}

func NewAuthHandlers(client genprotos.AuthServiceClient, logger *log.Logger, ch *amqp.Channel) *AuthHandlers {
	return &AuthHandlers{
		client:          client,
		logger:          logger,
		rabbitMQChannel: ch,
	}
}

// Register godoc
// @Summary Register user
// @Description This endpoint for registering user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body genprotos.RegisterUserRequest true "User details to register"
// @Success 200 {object} genprotos.RegisterUserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/register [post]
func (a *AuthHandlers) Register(ctx *gin.Context) {
	var req genprotos.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = a.rabbitMQChannel.Publish(
		"",
		"register_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        reqBytes,
		},
	)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User registration data sent to the queue"})
}

// Login godoc
// @Summary Login user
// @Description This endpoint for logging in user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body genprotos.LoginUserRequest true "User login details"
// @Success 200 {object} genprotos.LoginUserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/login [post]
func (a *AuthHandlers) Login(ctx *gin.Context) {
	var req genprotos.LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.LoginUser(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description This endpoint for refreshing access token using refresh token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body genprotos.RefreshTokenRequest true "Refresh token details"
// @Success 200 {object} genprotos.RefreshTokenResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /auth/refresh [post]
func (a *AuthHandlers) RefreshToken(ctx *gin.Context) {
	var req genprotos.RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.RefreshToken(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}
