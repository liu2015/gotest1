package system

import (
	"ginserver/global"

	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GVA_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username string    `json:"userName" gorm:"index;comment:用户登录名"`
	Password string    `json:"-" gorm:"comment:用户登录密码"`
	Nickname string    `json:"nickName" gorm:"default:系统用户;conmment:用户昵称"`

	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	SideMode string `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`

	Headerimg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	BaseColor string `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`

	ActiveColor string `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`
	AuthorityId uint   `json:"authorityId" gorm:"default:888;comment:用户校色ID"`

	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"email"  gorm:"comment:用户邮箱"` // 用
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
