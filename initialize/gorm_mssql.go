package initialize

import (
	"ginserver/global"
	"ginserver/initialize/internal"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GormMssql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 191,
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxIdleConns)
		return db
	}
}
