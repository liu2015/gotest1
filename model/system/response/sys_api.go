package response

import "ginserver/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"Api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
