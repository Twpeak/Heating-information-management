package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api"
	"github.com/gin-gonic/gin"
)

type FeverRouter struct {
}

func (s *BaseRouter)InitFeverRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	feverRouter := Router.Group("fever")
	feverApi := api.ApiGroupApp.SystemApiGroup.FeverApi
	{
		feverRouter.GET("test",feverApi.Test)
		feverRouter.POST("update",feverApi.UpdateFeverInfo)
		feverRouter.DELETE("del",feverApi.DelFeverInfo)
		feverRouter.POST("add",feverApi.AddFeverInfo)
	}
	return feverRouter
}