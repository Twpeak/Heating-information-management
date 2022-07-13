package utils

import (
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)
/**
关于会话的操作：JWT中有一个创建会话的函数，这里
 */

/**
从token中解析获取Claims的方法，跟/JWTAuth()中间件鉴权方法中的逻辑非常相似
 */

// GetClaims 获取会话
func GetClaims(c *gin.Context)(*systemReq.CustomClaims,error)  {
	token := c.Request.Header.Get("Authorization")		//获取token
	j := NewJWT()		//获取jwt对象方法，和盐
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		global.G_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {	//若上下文中不存在 “claims”
		if cl, err := GetClaims(c); err != nil {			//则从token中获取，还出错就返回0，不错就返回用户id
			return 0
		} else {
			return cl.ID
		}
	} else {												//若上下文中有"claims"则断言，直接获取用户id
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.ID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户 角色id
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleId
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.RoleId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色 会话信息
func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.UUID
	}
}