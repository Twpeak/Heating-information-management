package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

/**
Jwt登录验证中间件
 */

/***
中间件操作流程：
1.获取token
2.判断token是否为空，或是否在黑名单(多点登录，所以jwt被拉黑)，则返回相应错误
3.解析token
4.判断token是否过期，返回相应错误
5.判断token是否恰好在缓冲期，
	-若是：则续签：续签基本流程：	（新旧会话基本信息一样的，地址和过期时间不一样）
		-1.用旧会话信息去创建新会话（也可以：先用旧会话创建新token，取出新会话修改时间。放入上下文和头信息）
		-2.修改会话的过期时间
		-3.放入上下文
		-4.设置头信息				（若不安全也可以）
6.无论是不是都最后将会话放入上下文中
 */
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者
		//本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("Authorization")	//验证信息

		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()		//获取jwt对象方法，和盐
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		//判断当前请求的jwt是否过期或是否在缓冲期
		if claims.ExpiresAt - time.Now().Unix() < claims.BufferTime{	//预设过期时间减去当前时间 ，跟预设的缓冲时间做对比。看是直接过期还是还有救
			//若还在缓冲时间内，还有救，就废弃就jwt，启用新jwt
			newToken, _ := j.CreateTokenByOldToken(token, *claims)  					//新token，通过旧jwt中claim直接创建
			claims.ExpiresAt = time.Now().Unix() + global.G_CONFIG.JWT.ExpiresTime		//修改新过期时间
			newClaims,_ := j.ParseToken(newToken)										//取出新会话，准备修改（其实也可以先修改，然后再创建新token）
			//设置两个头部信息，给前端说明，方便更新token和更新cookie的过期时间
			c.Header("Authorization", "Bearer"+" "+token)							//设置头信息
			c.Header("new-expires-at",strconv.FormatInt(newClaims.ExpiresAt,10))

			//多点登录
			if global.G_CONFIG.System.UseMultipoint {
				//从缓存中获取旧jwt（注意，这里的旧jwt 应该和这次请求的jwt一致）拉黑旧jwt，记录新jwt
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.G_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

