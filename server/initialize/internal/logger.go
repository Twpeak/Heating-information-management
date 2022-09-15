package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"fmt"
	"gorm.io/gorm/logger"
)

//此方法主要用于重写logger.writer，用于数据库的日志输出

//继承gorm.logger.writer
//我们就可以在其内部写很多操作，主要是重写Printf方法
type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.G_CONFIG.System.DbType {
	case "mysql":
		logZap = global.G_CONFIG.Mysql.LogZap
	case "pgsql":
		//logZap = global.G_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.G_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}