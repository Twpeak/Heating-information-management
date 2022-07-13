package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type JwtBlacklist struct {
	global.G_MODEL
	Jwt string `gorm:"type:text;comment:jwt" gorm:"comment:jwt黑名单;"`
}

func (*JwtBlacklist) TableName() string {
	return "jwt_blacklist"
}
