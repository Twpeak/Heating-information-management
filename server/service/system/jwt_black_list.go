package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

/**
关于jwt的存储，黑名单操作服务
gva中，jwt黑名单都存储在了catch中，白名单在redis中
他约定的是，黑名单的jwt为key，值为空，白名单key为username，值为jwt
我们只用redis，且定下我们自己的约定
黑名单jwt，存储在redis set中，白名单jwt则直接用key-val存储
 */

type JwtService struct {}


// JsonInBlacklist
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService)JsonInBlacklist(jwtList system.JwtBlacklist)(err error) {
	//在jwt黑名单信息表中，添加进该jwt信息
	err = global.G_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	return
}


// IsBlacklist
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService)IsBlacklist(jwt string) bool  {
	 ok := global.G_REDIS.SIsMember(context.Background(),global.RedisC.JwtBlacklist,jwt).Val()
	return ok
}


// GetRedisJWT
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.G_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// SetRedisJWT
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.G_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.G_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

//@function: GetTokenByAuthHeader
//@description: 通过authHeader获取token
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService)GetTokenByAuthHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"msg":  "请求头中auth为空",
		})
		c.Abort()
		return ""
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code": 2004,
			"msg":  "请求头中auth格式有误",
		})
		c.Abort()
		return ""
	}
	return parts[1]
}

func LoadAll()  {
	var jwtBlackList [] string	//用来存储黑名单的切片
	//从jwtBlacklist表中查询所有jwt字段
	err  := global.G_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&jwtBlackList).Error
	if err != nil {
		global.G_LOG.Error("加载数据库jwt黑名单失败！",zap.Error(err))
		return
	}
	for i := 0; i < len(jwtBlackList); i++ {
		global.G_REDIS.SAdd(context.Background(),global.RedisC.JwtBlacklist)
	}                // jwt黑名单 加入 redisd缓存 中
}




