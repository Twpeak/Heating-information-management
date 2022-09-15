package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {	//直接用角色去请求
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取用户的角色
		sub := waitUse.RoleId
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method

		e := casbinService.Casbin()

		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if  success {
			c.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
