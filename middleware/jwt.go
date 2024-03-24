package middleware

import (
	"errors"
	"ginserver/global"
	"ginserver/model/common/response"
	"ginserver/model/system"
	"ginserver/service"
	"ginserver/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := utils.GetToken(ctx)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", ctx)
			ctx.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{
				"reload": true,
			}, "您的账户异地登录或令牌失效", ctx)
			utils.ClearToken(ctx)
			ctx.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{
					"reload": true,
				}, "授权已过期", ctx)
				utils.ClearToken(ctx)
				ctx.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), ctx)
			utils.ClearToken(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			ctx.Header("new-token", newToken)
			ctx.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(ctx, newToken, int(dr.Seconds()))
			if global.GVA_CONFIG.System.UseMultipoint {
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else {
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
			}
		}

	}
}
