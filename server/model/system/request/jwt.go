package request

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

/**
jwt 封装的返回方法
 */

// CustomClaims 常用会话信息，除了包含基本信息（我们自定义的额外信息）之外，还封装了jwt自带的主体信息
type CustomClaims struct {
	BaseClaims
	BufferTime int64		// 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
	jwt.StandardClaims		//这个类型中的字段都是基本类型，但不安全不推荐
	//jwt.RegisteredClaims		//这个类型中的字段大部分是复杂类型，需要使用jwt封装的方法获取值
}

// BaseClaims 基本会话信息
type BaseClaims struct {
	UUID        uuid.UUID		//用户uuid
	ID          uint			//用户id
	Username    string			//用户名
	Name   	    string				//真实姓名
	RoleId		uint			//角色id
}