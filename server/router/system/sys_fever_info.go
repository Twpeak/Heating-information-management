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
	}
	return feverRouter
}