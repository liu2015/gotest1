package request

type SysAuthorityBthReq struct {
	MenuID      uint   `json:"menuID"`
	AuthorityId uint   `json:"authorityId"`
	Selectd     []uint `json:"selected"`
}
