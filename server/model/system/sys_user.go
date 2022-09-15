package system

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`              // 用户UUID
	Username string    `json:"username" gorm:"comment:用户登录名;unique;"`   // 用户登录名
	Name     string    `json:"name" gorm:"default:'系统用户';comment:真实姓名"` // 真实姓名
	Password string    `json:"password"  gorm:"comment:用户登录密码"`         // 用户登录密码
	//HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	RoleId uint    `json:"roleId" gorm:"default:4;comment:用户角色ID"` // 用户角色ID
	Role   SysRole `json:"role" gorm:"comment:用户角色"`               //用于反向查询
	//身份证号
	IdentityCard string `json:"identity_card" gorm:"comment:身份证号;"` //想写一个，注册时默认登录密码为身份证号码后六位的接口
	//电话号码
	Phone string `json:"phone" gorm:"comment:电话号码;"`
	//电子邮箱
	Email string `json:"email" gorm:"comment:电子邮箱;"`
	//所属医院
	HospitalId uint `json:"hospital_id" gorm:"comment:所属医院;"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
