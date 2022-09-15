package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

/**
我们假设我们设计用到了很多数据库，所以需要一个判断初始化
*/

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.G_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		//return GormPgSql()
	default:
		return GormMysql()
	}
	return GormMysql()
}

// RegisterTables 注册数据库表专用--初始化表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		//model
		//gorm当创建从表时，默认会创建依赖的主表,多对多关系的表，不能自动创建。
		system.FeverInfo{},

		system.District{},
		system.Hospital{},
		system.SysUser{},
		system.SysRole{},

		//时间管理任务的表
		system.JwtBlacklist{},
	)
	if err != nil {
		global.G_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.G_LOG.Info("register table success")
}
