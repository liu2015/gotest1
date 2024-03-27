package request

import (
	"fmt"
	"ginserver/config"
	"os"
)

type InitDB struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
	DBPath   string `json:"dbPath"`
}

// 空数据，初始化建库

func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"

	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)

}

// 空数据库，PGSQL建表连接
func (i *InitDB) PgsqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"

	}
	if i.Port == "" {
		i.Port = "5432"
	}
	return "host" + i.Host + "user=" + i.UserName + "password=" + i.Password + "port=" + i.Port + " dbname=" + "postgres" + " " + "sslmode=disable TimeZone=Asia/Shanghai"
}

// sqlite 空数据库 建库
func (i *InitDB) SqliteEmptyDsn() string {
	separator := string(os.PathSeparator)
	return i.DBPath + separator + i.DBName + ".db"
}

// mssql
func (i *InitDB) MssqlEmptyDsn() string {
	return "sqlserver://" + i.UserName + ":" + i.Password + "@" + i.Host + ":" + i.Port + "?database=" + i.DBName + "&encrypt=disable"
}

// 转换config。mysql

func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{

		GeneralDB: config.GeneralDB{
			Path:         i.Host,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
}

// topgsql 转换mysql

func (i *InitDB) ToPgsqlConfig() config.Pgsql {
	return config.Pgsql{
		GeneralDB: config.GeneralDB{
			Path:         i.Host,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "sslmode=disable TimeZone=Asia/Shanghai",
		},
	}
}

// tosql 转config。sqlite
func (i *InitDB) ToSqliteConfig() config.Sqlite {
	return config.Sqlite{
		GeneralDB: config.GeneralDB{
			Path:         i.DBPath,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "",
		},
	}
}

func (i *InitDB) ToMssqlConfig() config.Mssql {
	return config.Mssql{
		GeneralDB: config.GeneralDB{
			Path:         i.DBName,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "",
		},
	}
}
