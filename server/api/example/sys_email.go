package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	resCom"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type EmailApi struct{}


// @Summary 发送邮件
// @Param data body email_response.Email true "发送邮件必须的参数"
// @Router /email/sendEmail [post]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email response.Email
	_ = c.ShouldBindJSON(&email)

	to := strings.Split(email.To, ",")
	if err := emailService.SendEmail(to, email.Subject, email.Body); err != nil {
		global.G_LOG.Error("发送失败!", zap.Error(err))
		resCom.FailWithMessage("发送失败", c)
	} else {
		resCom.OkWithMessage("发送成功", c)
	}
}


