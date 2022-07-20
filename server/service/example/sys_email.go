package example

import (
	"bytes"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/dto"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"html/template"
	"time"
)

type EmailService struct{}
var EmailServiceApp = new(EmailService)


func (e *EmailService) SendEmail(to []string, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
//给医院管理员推送本周登记的记录数量，以及各个医生登记的患者数量

//定时任务的发送邮箱逻辑：给所有管理者发送邮件--是并发发送呢还是线性发送？
func (*EmailService)SendEamilToBoss() (err error) {
	//获取模板信息
	emailList,err := EmailServiceApp.AdminQuery()
	if err != nil{
		global.G_LOG.Error("获取信息失败",zap.Error(err))
	}
	for i := 0; i < len(emailList); i++ {
		var body bytes.Buffer
		subject := emailList[i].HospitalName + "医院发热信息周报"
		t,_:= template.ParseFiles("./templates/email-template.html")
		err := t.Execute(&body, emailList[i])
		if err != nil {
			global.G_LOG.Error("解析模板失败！",zap.Error(err))
			return err
		}
		 err = utils.EmailSendHtml(emailList[i].Email, subject, body)
	}
	return err
}

//获取模板信息
func (u *EmailService) AdminQuery() ([]dto.AdminPush, error) {
	startDate := time.Unix(time.Now().Unix()-(86400*7)-60, 0)
	sendDate := time.Now()
	var adm []dto.AdminPush
	err := global.G_DB.Model(&system.Hospital{}).Select("hospitals.boos_id,hospitals.id,hospitals.hospital_name,sys_users.name,sys_users.email").
		Joins("left join sys_users on hospitals.boos_id = sys_users.id").Find(&adm).Error

	for i := 0; i < len(adm); i++ {
		err = global.G_DB.Model(&system.SysUser{}).Select("sys_users.id,sys_users.name").Where("hospital_id = ?", adm[i].Id).Find(&adm[i].DockerPush).Error
		var all uint = 0
		for j := 0; j < len(adm[i].DockerPush); j++ {
			var iii []dto.IdDto
			m := global.G_DB.Model(&system.FeverInfo{}).Where("(doctor_id = ?) AND (fever_infos.created_at BETWEEN ? AND ?)", adm[i].DockerPush[j].Id, startDate, sendDate).Find(&iii)
			affected := uint(m.RowsAffected)
			adm[i].DockerPush[j].Number = affected
			all += affected
		}
		adm[i].Total = all
	}

	return adm, err
}
