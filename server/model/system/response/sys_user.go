package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

/**
用户返回体
 */

//登录时获取用户信息的结构体
type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

//用户登录成功的返回信息
type LoginResponse struct {
	User system.SysUser `json:"user"`
	Token string		`json:"token"`
	ExpiresAt int64		`json:"expiresAt"`
}
