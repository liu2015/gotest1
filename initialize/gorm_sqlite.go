package initialize

import (
	"ginserver/global"
	"ginserver/initialize/internal"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GormSqlite() *gorm.DB {
	s := global.GVA_CONFIG.Sqlite

	if s.Dbname == "" {
		return nil
	}
	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)

	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
