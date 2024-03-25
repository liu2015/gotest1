package core

import (
	"fmt"
	"ginserver/global"
	"ginserver/initialize"
	"ginserver/service/system"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		initialize.Redis()
	}
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	Router := initialize.Routers()
}
