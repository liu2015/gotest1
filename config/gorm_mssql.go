package config

// 这里使用的组合，GeneralDB 是在db_list 内

type Mssql struct {
	GeneralDB `"yaml:",inline" mapstructure:",squash"`
}

func (m *Mssql) Dsn() string {
	return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Path + "?database=" + m.Dbname + "&encrypt=disable"
}

// 所以，m.LogMode == m.Mssql.LogMode 调用
func (m *Mssql) GetLogMode() string {
	return m.LogMode
}
