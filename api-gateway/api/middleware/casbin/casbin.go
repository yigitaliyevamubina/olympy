package casbin

import (
	"fmt"
	"log"
	"net/http"
	"olympy/api-gateway/api/models/model_common"
	"olympy/api-gateway/config"
	"olympy/api-gateway/internal/pkg/logger"
	jwt "olympy/api-gateway/internal/pkg/tokens"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type CasbinHandler struct {
	config     config.Config
	enforce    casbin.Enforcer
	jwthandler jwt.JWTHandler
}

func NewAuthorizer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token1 := ctx.GetHeader("Authorization")
		if token1 == "" {

			sub := "unauthorized"
			obj := ctx.Request.URL.Path
			etc := ctx.Request.Method
			e, _ := casbin.NewEnforcer(`auth.conf`, `auth.csv`)
			t, _ := e.Enforce(sub, obj, etc)
			if t {
				ctx.Next()
				return
			}
			fmt.Println(sub, obj, etc, t)
		}

		claims, err := jwt.ExtractClaim(token1)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				&model_common.ResponseError{
					Code:    http.StatusText(http.StatusUnauthorized),
					Message: "missing token in the header",
					Data:    err.Error(),
				})
			logger.Error(err)
			return
		}

		sub := claims["role"]
		obj := ctx.Request.URL.Path
		etc := ctx.Request.Method

		e, err := casbin.NewEnforcer(`auth.conf`, `auth.csv`)

		if err != nil {
			log.Fatal(err)
			return
		}
		t, err := e.Enforce(sub, obj, etc)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(sub, obj, etc)
		if t {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "permission denied",
		})
	}
}
