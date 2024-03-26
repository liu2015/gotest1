package middleware

import (
	"ginserver/service"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

func CasbinHandler() gin.HandlerFunc {

	return
}
