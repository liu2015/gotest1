package system

import "ginserver/config"

type System struct {
	Config config.Server `json:"config"`
}
