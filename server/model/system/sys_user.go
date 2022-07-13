package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.G_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`                                                           // 用户UUID
	Username    string         `json:"userName" gorm:"comment:用户登录名;unique;"`                                                        // 用户登录名
	NickName    string         `json:"nickName" gorm:"default:'系统用户';comment:用户昵称"`                                            // 用户昵称
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	//HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	RoleId 		uint         	`json:"roleId" gorm:"default:2;comment:用户角色ID"`                                        // 用户角色ID
	Role   		SysRole   		`json:"role" gorm:"comment:用户角色"`		//用于反向查询
}

func (SysUser) TableName() string {
	return "sys_users"
}
