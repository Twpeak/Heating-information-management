package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new (_gorm)

type _gorm struct {}

// Config gorm 自定义配置
func (g *_gorm)Config() *gorm.Config  {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,		//自动迁移时，禁用外键约束,不禁用
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	var logMode DBBASE
	switch global.G_CONFIG.System.DbType {
	case "mysql":
		logMode = &global.G_CONFIG.Mysql
		break
	case "pgsql":
		//logMode = &global.G_CONFIG.Pgsql
		break
	default:
		logMode = &global.G_CONFIG.Mysql
	}

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}


	return config


}