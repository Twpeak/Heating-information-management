package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type EmailService struct{}

func (e *EmailService) SendEmail(to []string, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
//给医院管理员推送本周登记的记录数量，以及各个医生登记的患者数量
var total = 5
var subject = "医院周报"
const body = "尊敬的 医院管理者 （xxx先生） ， 本周新增发热患者人数为: " + "" +  "人，其中各个医生登记患者数量如下 xxx"



//获取所有医院管理者邮箱信息
func getBossIds() (emailList []string) {
	if err := global.G_DB.Model(&system.SysUser{}).Where("role_id = ?",3).Select("email").Scan(&emailList).Error;err != nil{
		global.G_LOG.Error("查询医院管理者信息失败",zap.Error(err))
		return
	}
	return emailList
}

//定时任务的发送邮箱逻辑：给所有管理者发送邮件--是并发发送呢还是线性发送？
func (*EmailService)SendEamilToBoss() (err error) {
	emailList := getBossIds()
	for i := 0; i < len(emailList); i++ {
		 err = utils.Email(emailList, subject, body)
		 if err != nil{
			 global.G_LOG.Error("发送邮件失败",zap.Error(err))
		 }
	}
	return err
}