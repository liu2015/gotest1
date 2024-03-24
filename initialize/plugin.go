package initialize

import (
	"fmt"
	"ginserver/middleware"
	"ginserver/utils/plugin"

	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for v := range Plugin {
		// 用户反悔注册路由
		PluginGroup := group.Group(Plugin[v].RouterPath())
		Plugin[v].Register(PluginGroup)

	}

}

func InstallPlugin(Router *gin.Engine) {

	PublicGroup := Router.Group("")

	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("健权插件安装==", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

}
