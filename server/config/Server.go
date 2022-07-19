package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Casbin Casbin `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System System `mapstructure:"system" json:"system" yaml:"system"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	//oss
	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`		//本地资源



	//待扩展业务
	//Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`


}
