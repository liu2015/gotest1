package plugin

import "github.com/gin-gonic/gin"

const (
	OnlyFuncName = "Plugin"
)

// 插件接口化
type Plugin interface {

	// register 注册路由
	Register(group *gin.RouterGroup)

	RouterPath() string
}
