package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)


//自定义邮件模板发送方法
func EmailSendHtml(To string,subject string, body bytes.Buffer) error {
	to := strings.Split(To, ",")
	return send(to, subject, body.String())
}


//@function: Email
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func Email(To []string, subject string, body string) error {
	//to := strings.Split(To, ",")
	return send(To, subject, body)
}

//@function: ErrorToEmail
//@description: 给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func ErrorToEmail(subject string, body string) error {
	to := strings.Split(global.G_CONFIG.Email.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

//@function: EmailTest
//@description: Email测试方法
//@param: subject string, body string
//@return: error

func EmailTest(subject string, body string) error {
	to := []string{global.G_CONFIG.Email.From}
	return send(to, subject, body)
}

//@function: send
//@description: Email发送方法
//@param: subject string, body string
//@return: error

func send(to []string, subject string, body string) error {
	from := global.G_CONFIG.Email.From
	nickname := global.G_CONFIG.Email.Nickname
	secret := global.G_CONFIG.Email.Secret
	host := global.G_CONFIG.Email.Host
	port := global.G_CONFIG.Email.Port
	isSSL := global.G_CONFIG.Email.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
