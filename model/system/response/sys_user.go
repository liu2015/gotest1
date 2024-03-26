package response

import "ginserver/model/system"

type SysUserResponse struct {
	User system.SysUser `json:"suer"`
}

type LoginResponse struct {
	User      system.SysUser `json:"suer"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
