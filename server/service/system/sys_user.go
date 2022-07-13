package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//关于用户的操作逻辑:登录，注册

type UserService struct {}

// Login
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService)Login(u *system.SysUser)(UserInter *system.SysUser,err error)  {
	if nil == global.G_DB{		//先判断是否连接数据库
		return nil,fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.G_DB.Where("username = ?",u.Username).Preload("Role").First(&user).Error
	if err != nil {
		//没有解密，只是封装了加密和对比两个操作而已
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user,err
}

//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.G_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.G_DB.Create(&u).Error
	return u, err
}

func (userService *UserService)InitUserRole()  {
	rdb := global.G_DB.Model(&system.SysRole{})
	udb := global.G_DB.Model(&system.SysUser{})
	RoleDate := []system.SysRole{
		//....角色信息初始化数据
		{
			Id: 1,
			RoleName: "系统管理员",
		},
		{
			Id: 2,
			RoleName: "县区管理员",
		},
		{
			Id: 3,
			RoleName: "医院管理员",
		},
		{
			Id: 4,
			RoleName: "医生",
		},
	}
	UserDate := []system.SysUser{
		{
			Username: "admin",
			Password: "123456",
			RoleId: 1,
			Name: "系统管理员",
			IdentityCard:"410703900003074014",
			Phone: "15516575533",
			HospitalId: 1,
		}, {
			Username: "dis_admin",
			Password: "123456",
			RoleId: 2,
			Name: "县区管理员",
			IdentityCard:"410703900003074014",
			Phone: "15516575533",
			HospitalId: 1,
		},{
			Username: "hos_admin",
			Password: "123456",
			RoleId: 3,
			Name: "医院管理员",
			IdentityCard:"410703900003074014",
			Phone: "15516575533",
			HospitalId: 1,
		},{
			Username: "docter",
			Password: "123456",
			RoleId: 4,
			Name: "医生",
			IdentityCard:"410703900003074014",
			Phone: "15516575533",
			HospitalId: 1,
		},

	}
	for _,date := range RoleDate{
		if err := rdb.FirstOrCreate(&system.SysRole{},&date).Error;err != nil{
			global.G_LOG.Error("角色数据初始化失败",zap.Error(err))
			return
		}
	}

	for _,udate := range UserDate{
		if err := udb.FirstOrCreate(&system.SysUser{},&udate).Error;err != nil{
			global.G_LOG.Error("管理员数据初始化失败",zap.Error(err))
			return
		}
	}

	return
}