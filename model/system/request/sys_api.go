package request

import (
	"ginserver/model/common/request"
	"ginserver/model/system"
)

type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
