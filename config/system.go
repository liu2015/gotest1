package config

type System struct {
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`

	Addr          int  `mapstructure:"addr" json:"addr" yaml:"addr"`
	LimitCountIP  int  `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int  `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`

	UseRedis bool `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	UseMongo bool `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`
}
