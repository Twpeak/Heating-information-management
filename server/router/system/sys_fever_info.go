package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api"
	"github.com/gin-gonic/gin"
)

type FeverRouter struct {
}

func (s *BaseRouter) InitFeverRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	feverRouter := Router.Group("fever")
	feverApi := api.ApiGroupApp.SystemApiGroup.FeverApi
	{
		feverRouter.POST("update",feverApi.UpdateFeverInfo)
		feverRouter.DELETE("del",feverApi.DelFeverInfo)
		feverRouter.POST("add",feverApi.AddFeverInfo)
		feverRouter.GET("", feverApi.FeverTextLimit)
		feverRouter.GET("export", feverApi.Export)
	}
	return feverRouter
}
