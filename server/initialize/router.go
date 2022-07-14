package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Routers 初始化总路由
/**
一般简单的注册路由组，是根据不同的功能来分组。例如：用户操作组，注册相关操作组。。等
原本路由组里应该实现getpost等注册操作，但我们创建了两个父路由组Public。Private。并在其下方创建了多个嵌套子路由组，如jwt操作，注册相关功能操作等、
知识点：我们注册的是 router 目录下的路由组，其每个路由类就是一个路由组，其中调用api层的方法添加入路由组，而api又调用service层的逻辑方法，具体区别可以看gva框架源码和ass实例代码
*/

func Routers() *gin.Engine {
	Router := gin.Default()

	//解决跨域
	Router.Use(cors.Default())

	systemRouter := router.RouterGroupsApp.System
	//exampleRouter := router.RouterGroupsApp.Example

	Router.StaticFS(global.G_CONFIG.Local.Path, http.Dir(global.G_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	//若没有静态模板需要解析，则不需要开启
	if 1 == 0 {
		Router.Static(global.G_CONFIG.Local.Static, global.G_CONFIG.Local.StaticPath) // 静态页面资源
		Router.LoadHTMLGlob("templates/*")
	}

	// 方便统一添加路由组前缀 多服务器上线使用，————PublicGroup父路由组
	PublicGroup := Router.Group("")
	{
		//健康检测，直接注册
		PublicGroup.GET("/health", func(context *gin.Context) {
			context.JSON(200, "ok")
		})
		//给 PublicGroup 父路由组注册基础（注册登录）等功能路由组
		systemRouter.InitBaseRouter(PublicGroup) //注册基础功能路由 不做鉴权
	}

	PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//一般来说路由组里面都是post，get等操作，但我们将这里的每一句都嵌套了一个路由组
	{
		systemRouter.InitHospitalRouter(PrivateGroup)
	}
	UserRelevant := Router.Group("/user")
	{
		//首页所有信息
		UserRelevant.GET("/text", system.UserInformation)
		//添加用户
		UserRelevant.POST("/add", system.UserAdd)
		//修改用户前的信息回显
		UserRelevant.GET("/utext", system.UserUpdateDisplay)
		//修改用户
		UserRelevant.PUT("/put", system.UserUpdate)
		//删除用户
		UserRelevant.DELETE("/del", system.UserDelete)
	}

	//InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	return Router
}
