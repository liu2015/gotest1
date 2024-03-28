package request

import "ginserver/model/common/request"

type SysAutoHistory struct {
	request.PageInfo
}

type RollBack struct {
	ID          int  `json:"id" form:"id"`
	DeleteTable bool `json:"deleteTable" form:"deleteTable"`
}
