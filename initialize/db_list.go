package initialize

import (
	"ginserver/config"
	"ginserver/global"

	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GVA_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "mssql":

		}
	}
}
