package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {}

// InitBaseRouter 注册基础功能路由 不做鉴权,		//疑问：为什么要返回 gin.IRouters
func (s *BaseRouter)InitBaseRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	//子路由组
	baseRouter := Router.Group("base")
	captchaApi := api.ApiGroupApp.SystemApiGroup.CaptchaApi			//使用逻辑
	baseApi   := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("captcha",captchaApi.Captcha)
		baseRouter.POST("login",baseApi.Login)
		baseRouter.POST("register",baseApi.Register)
	}
	return baseRouter

}