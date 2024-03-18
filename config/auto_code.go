package config

type Autocode struct {
	SModel  string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRouter string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
}
