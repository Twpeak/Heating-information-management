package core

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)




// Zap 获取zap.logger
func Zap()(logger *zap.Logger)  {
	//判断配置中存放log文件的目录地址是否存在，若不存在则创建
	if ok,_ := utils.PathExists(global.G_CONFIG.Zap.Director); !ok{
		fmt.Printf("create %v directory\n", global.G_CONFIG.Zap.Director)
		_ = os.Mkdir(global.G_CONFIG.Zap.Director, os.ModePerm)
	}
	//得到Zap_in
	z := internal.NewzapIn()
	//根据配置文件的Level获取 []zapcore.Core 切片
	cores:=z.GetZapCores()
	//zapcore.NewTee方法可以把多个core衔接在一起，对应logger的操作会同时操作这些core。
	logger = zap.New(zapcore.NewTee(cores...))		//zap.new返回一个Logger，

	if global.G_CONFIG.Zap.ShowLine {	//判断是否添加日志行号
		//Logger 结构体中会包含很多配置信息，我们在开发中可以通过 WithOptions 来添加相应的参数。如添加日志行号：
		//AddCaller 函数会创建一个回调钩子给 WithOptions 执行，这也是函数式编程的魅力所在：
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}



