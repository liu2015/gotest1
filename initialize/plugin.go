package initialize

import (
	"fmt"
	"ginserver/global"
	"ginserver/middleware"
	"plugin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for v := range Plugin {
		// 用户注册路由
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

	// 添加跟角色挂钩的插件
	PluginInit(PrivateGroup, email.CreateEmailPlug{
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	})

}
