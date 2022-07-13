package system

type SysRole struct {
	Id       uint   `json:"roleId" gorm:"comment:角色ID"`          // 角色ID	0为开发者，1为管理员,2为用户
	RoleName string `json:"roleName" gorm:"comment:角色名;unique;"` // 角色名
}

func (SysRole) TableName() string {
	return "sys_Roles"
}
