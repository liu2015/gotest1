package initialize

import (
	"ginserver/config"
	"ginserver/global"
	"ginserver/initialize/internal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormPaSql() *gorm.DB {
	p := global.GVA_CONFIG.Pgsql

	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

func GormPgSqlByConfig(p config.Pgsql) *gorm.DB {
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db

	}

}
