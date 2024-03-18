package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key"  yaml:"signing-key"`
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` //过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
	Issuer      string `mapstructure:"issuer" json:"jssuer" yaml:"issuer"`
}