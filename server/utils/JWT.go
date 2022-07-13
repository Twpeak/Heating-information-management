package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"time"
)

type JWT struct {
	SigningKey [] byte	//盐（密钥）
}

var (
	TokenExpired     = errors.New("Token is expired")				//token到期
	TokenNotValidYet = errors.New("Token not active yet")			//token尚未激活
	TokenMalformed   = errors.New("That's not even a token")		//非合法token
	TokenInvalid     = errors.New("Couldn't handle this token:")	//无法处理此token
)

func NewJWT() *JWT{
	return &JWT{
		[]byte(global.G_CONFIG.JWT.SigningKey),
	}
}

// CreateClaims 创建一个常用会话
func (j *JWT)CreateClaims(baseClaims request.BaseClaims)request.CustomClaims  {
	claims := request.CustomClaims{
		BaseClaims : baseClaims,
		BufferTime: global.G_CONFIG.JWT.BufferTime,	//// 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.G_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.G_CONFIG.JWT.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT)CreateToken(claims request.CustomClaims) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(j.SigningKey)		//经典的两句，生产token对象和加盐加密为jwt字符串
}


// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil		//传入密钥
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {	//若token对象不为空，则进行类型断言，并获取其CustomClaims信息
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}



// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	var GVA_Concurrency_Control = &singleflight.Group{}
	v, err, _ := GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}