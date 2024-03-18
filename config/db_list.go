package config

type DsnProvider interface {
	Dsn() string
}

type GeneralDB struct {
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	Dbname   string `mapstructure:"db-name" json:"db-anme" yaml:"db-name"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Engine   string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`
	LogMode  string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
}
