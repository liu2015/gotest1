package response

import "ginserver/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
