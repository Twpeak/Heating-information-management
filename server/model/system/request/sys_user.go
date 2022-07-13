package request

/**
关于user操作所用到的所有返回请求体数据封装
*/

// Register 用户注册结构
type Register struct {
	Username     string   	`json:"userName"`
	Password     string   	`json:"passWord"`
	Name     	 string   	`json:"name" `
	HeaderImg    string   	`json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	RoleId  	 uint   	`json:"roleId"    gorm:"default:2"`
}

// Login 用户登录结构
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}