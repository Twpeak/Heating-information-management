package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

func main() {
	global.G_VIPER = core.Viper()	//初始化Viper,载入配置文件
	//fmt.Println(global.G_CONFIG.Captcha)
	global.G_LOG = core.Zap()		// 初始化zap日志库
	zap.ReplaceGlobals(global.G_LOG)	//将zap提供的Logger and SugaredLogger替换为Logger
	global.G_DB = initialize.Gorm()	//gorm连接数据库
	initialize.Timer()				//定时任务管理器是在全局变量配置中初始化的，所以这里仅仅是开启了删除表的定时任务

	if global.G_DB != nil{			//若成功连接数据库

		initialize.RegisterTables(global.G_DB)	//初始化表
		//程序结束前关闭数据库连接。所以我们需要DB()获取sql.DB
		sqldb,_ := global.G_DB.DB()
		defer sqldb.Close()
	}

	//初始化redis服务
	//使用redis或使用多点登录
	if global.G_CONFIG.System.UseRedis || global.G_CONFIG.System.UseMultipoint{
		initialize.Redis()
	}
	//加载黑名单
	if global.G_DB != nil {
		//若数据库已连接，则从db中加载jwt黑名单数据,并放入Redis缓存中，准备一会的登录操作
		system.LoadAll()
	}
	initialize.InitBaseMysqlDate()
	//初始化总路由
	Router := initialize.Routers()

	//加载静态资源
	//Router.Static("/", "./server/resource/Static/index.html")
	//启动服务
	Router.Run(":8888")

}


