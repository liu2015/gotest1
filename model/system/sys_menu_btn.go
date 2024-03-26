package system

import "ginserver/global"

type SysBaseMenuBth struct {
	global.GVA_MODEL
	Nmae          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"SysBaseMenuID" gorm:"comment:菜单ID"`
}
