package config

import (
	"fmt"
	"strings"
)

type Mongo struct {
	Coll     string `json:"coll" yaml:"coll" mapstructure:"coll"`
	Options  string `json:"options" yaml:"options" mapstructure:"options"`
	Database string `json:"database" yaml:"database" mapstructure:"database"`
	Username string `json:"username" yaml:"username" mapstructure:"username"`

	Password        string       `json:"password" yaml:"password" mapstructure:"password"`
	AuthSource      string       `json:"auth-source" yaml:"auth-source" mapstructure:"auth-source"`
	MinPoolSize     uint64       `json:"min-pool-size" yaml:"min-pool-size" mapstructure:"min-pool-size"`
	MaxPoolSize     uint64       `json:"max-pool-size" yaml:"max-pool-size" mapstructure:"max-pool-size"`
	SocketTimeoutMs int64        `json:"socket-timeout-ms" yaml:"socket-timeout-ms" mapstructure:"socket-timeout-ms"`
	IsZap           bool         `json:"is-zap" yaml:"is-zap" mapstructure:"is-zap"`
	Hosts           []*MongoHost `json:"hosts" yaml:"hosts" mapstructure:"hosts"`
}

type MongoHost struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port string `json:"port" yaml:"port" mapstructure:"port"`
}

func (x *Mongo) Uri() string {
	length := len(x.Hosts)
	hosts := make([]string, 0, length)
	for i := 0; i < length; i++ {
		if x.Hosts[i].Host != "" && x.Hosts[i].Port != "" {
			hosts = append(hosts, x.Hosts[i].Host+":"+x.Hosts[i].Port)
		}
	}
	if x.Options != "" {
		return fmt.Sprintf("mongodb://%s/%s?%s", strings.Join(hosts, ","), x.Database, x.Options)
	}
	return fmt.Sprintf("mongodb://%s/s", strings.Join(hosts, ","), x.Database)
}
