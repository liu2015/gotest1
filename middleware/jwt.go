package middleware

import (
	"ginserver/service"
	"ginserver/utils"

	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := utils.GetToken(ctx)
		if token == "" {

		}
	}
}
