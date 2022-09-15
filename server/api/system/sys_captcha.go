package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/system/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type CaptchaApi struct {

}

/**
生产验证码
 */

//参数变量：catpcha所需参数为：store存储器，drvier驱动，存储器可自定义存储，驱动决定了验证码的格式
//

var store = base64Captcha.DefaultMemStore
//var store = NewDefaultRedisStore()	//自定义存储
var driver base64Captcha.Driver	//驱动

// InitCaptcha 初始化配置
func InitCaptcha() *base64Captcha.Captcha {
	//配置参数，其实可以放到配置文件中，假如所有驱动都有配置信息，我们只需要修改CaptchaType的值即可
	driver = internal.GetDriver("")	//使用默认配置
	return base64Captcha.NewCaptcha(driver, store)	//创建实例
}

// Captcha 生成验证码
func (*CaptchaApi)Captcha(c *gin.Context)  {
	//先初始化验证码,创建验证码实例
	cp := InitCaptcha()

	if id, b64s, err := cp.Generate(); err != nil {
		global.G_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		//若成功，则封装返回的json信息
		response.OkWithDetailed(systemRes.SysCaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: global.G_CONFIG.Captcha.Length,
		}, "验证码获取成功", c)
	}

}



// VerifyCaptcha 这里的封装是为了，若以后兼容自定义redis作为存储器时，方便调用
func VerifyCaptcha(id, VerifyValue string) bool {
	return store.Verify(id, VerifyValue, true)
}

