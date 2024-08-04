package authhandlers

import (
	"log"
	genprotos "olympy/api-gateway/genproto/auth_service"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	client genprotos.AuthServiceClient
	logger *log.Logger
}

func NewAuthHandlers(client genprotos.AuthServiceClient, logger *log.Logger) *AuthHandlers {
	return &AuthHandlers{
		client: client,
		logger: logger,
	}
}

// Register godoc
// @Summary Register user
// @Description This endpoint for registering user.
// @Accept json
// @Produce json
// @Param request body genprotos.RegisterUserRequest true "User details to register"
// @Success 200 {object} genprotos.RegisterUserResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
// @Router /auth/register [post]
func (a *AuthHandlers) Register(ctx *gin.Context) {
	var req genprotos.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := a.client.RegisterUser(ctx, &req)
	if err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(200, resp)
}

// Login godoc
// @Summary Login user
// @Description This endpoint for logging in user.
// @Accept json
// @Produce json
// @Param request body genprotos.LoginUserRequest true "User login details"
// @Success 200 {object} genprotos.LoginUserResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
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
// @Accept json
// @Produce json
// @Param request body genprotos.RefreshTokenRequest true "Refresh token details"
// @Success 200 {object} genprotos.RefreshTokenResponse
// @Failure 400 {object} gin.H{"error": string}
// @Failure 500 {object} gin.H{"error": string}
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
