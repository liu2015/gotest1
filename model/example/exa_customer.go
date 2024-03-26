package example

import (
	"ginserver/global"
	"ginserver/model/system"
)

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string         `json:"customerName" form:"customerName" gorm:"comment:客户名"`
	CustomerPhoneData  string         `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:客户手机号"`
	SysUserID          uint           `json:"sysUserId" from :"sysUserId" gorm:"comment:管理ID"`
	SysUserAuthorityID uint           `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:管理角色ID"`
	SysUser            system.SysUser `json:"sysUser" form:"sysUser" gorm:"comment:管理详情"`
}
