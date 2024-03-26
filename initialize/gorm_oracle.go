package initialize

import (
	"ginserver/global"
	"ginserver/initialize/internal"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 191,
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db

	}

}
