package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

//关于用户的操作

type BaseApi struct {
	
}



/**
登录流程：
1.获取参数
2.检验参数是否合规
3.验证密码是否正确
4.通过角色和用户获取权限
5.做一些检查，例如检查当前用户角色是否为空
6.返回用户信息
7.签发jwt
 */

func (l *BaseApi)Login(c *gin.Context)  {
	//获取用户登录请求的封装参数
	var loginReq systemReq.Login
	//通过gin的方法，用请求体中的数据去填充数据
	_ = c.ShouldBind(&loginReq)
	fmt.Println(loginReq)
	if err:= utils.Verify(loginReq, utils.LoginVerify); err!=nil{
		response.FailWithMessage(err.Error(), c)
		return
	}
	if VerifyCaptcha(loginReq.CaptchaId, loginReq.Captcha) {	//若验证码校验成功
		//符合规定了，就存一下基本信息，去实现登录（去数据库换取完整的信息）
		u := &system.SysUser{Username: loginReq.Username,Password: loginReq.Password}
		if user,err := userService.Login(u); err!=nil {
			global.G_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {	//登录完成后，签发jwt
			l.tokenNext(c,*user)
		}
	} else {
		response.FailWithMessage("验证码错误",c)
	}
}


/**
签发jwt的操作：
1.创建会话
2.生产token
3.若是单点登录，则返回登录信息
4.若是多点登录：则判断缓存中是否还有jwt
	-若没有，设置缓存jwt，返回登录信息
	-若失败，则返回失败信息
	-若成功，获取缓存中的jwt，作废并设置新jwt，返回登录信息
 */
//登录后签发jwt
func (l *BaseApi)tokenNext(c *gin.Context,user system.SysUser)  {
	j := &utils.JWT{SigningKey: []byte(global.G_CONFIG.JWT.SigningKey)}// (盐)唯一签名
	//初始化并创建会话对象，我说参数类型怎么是baseclaims呢，自定义信息在参数中自己定义，基本信息是由user的信息确定
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Name:    user.Name,
		Username:    user.Username,
		RoleId: 	 user.RoleId,
	})
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		global.G_LOG.Error("获取token失败！",zap.Error(err))
		response.FailWithMessage("获取token失败！",c)
		return
	}

	//多点登录拦截，若是单点登录，则直接返回。无需修改登录状态
	if !global.G_CONFIG.System.UseMultipoint {	//返回信息登录成功
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}

	//先尝试用用户名获取jwt，若这次是新登录的，缓存中没有该用户的jwt信息，则去设置
	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil { //新登录，没有jwt在缓存中，则加入缓存设置登录状态
		if err := jwtService.SetRedisJWT(token,user.Username); err != nil{
			global.G_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		//使用自定义的登录响应体，封装登录信息并返回
		response.OkWithDetailed(systemRes.LoginResponse{
			User: user,
			Token: token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		},"登陆成功",c)
	} else if err!= nil {		//获取tokne失败
		global.G_LOG.Error("设置登录状态失败！",zap.Error(err))
		response.FailWithMessage("设置登录状态失败",c)
	} else {	//获取缓存token，缓存中还有token，没有过期,说明该用户曾经登录过，还未超过七天，则查看是否在黑名单中，再次尝试设置登录状态
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {	//因为是缓存中还有jwt，所以将旧jwt作废
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {	//设置新的jwt
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{					//返回响应信息
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/admin_register [post]
func (b *BaseApi) Register(c *gin.Context) {
	//1.获取参数信息，校验参数信息
	//2.封装到user对象并实现注册操作
	//		-查询数据库，用户名是否重复,重复则报错
	//		-生产uuid，加密密码信息。存储到数据库
	var r systemReq.Register
	c.ShouldBind(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{Username: r.Username, Name: r.Name, Password: r.Password,  RoleId: r.RoleId}
	//user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, RoleId: r.RoleId}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.G_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}
