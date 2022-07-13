package system

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	
}

func ( *UserRouter)InitUserRouter(Router *gin.RouterGroup) *gin.RouterGroup {
	userRouter := Router.Group("user")
	//userService := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		//userRouter.GET("all",userService.)


	}
	return userRouter
}
