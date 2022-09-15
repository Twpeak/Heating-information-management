package config

import "image/color"

type Captcha struct {
	StoreRedis	bool `mapstructure:"store-redis" json:"store-redis" yaml:"store-redis"`	//是否启用自定义redis当作存储器

	//驱动选项
	ImgWidth  	int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`    	// 验证码宽度
	ImgHeight 	int `mapstructure:"img-height" json:"img-height" yaml:"img-height"` 	// 验证码高度

	//DriverDigit
	MaxSkwe		float64	`mapstructure:"max-skwe" json:"max-skwe" yaml:"max-skwe"`			//单个数字的最大绝对偏差因子
	DotCount	int	`mapstructure:"dot-count" json:"dot-count" yaml:"dot-count"`		//背景噪音点的数量

	//DriverString
	NoiseCount	int	`mapstructure:"noise-count" json:"noise-count" yaml:"noise-count"`	//文本噪音数量
	Length		int `mapstructure:"length" json:"length" yaml:"length"`					//随机字符长度
	ShowLineOptions	int `mapstructure:"show-line-options" json:"show-line-options" yaml:"show-line-options"`	//背景线：OptionShowHollowLine | OptionShowSlimeLine | OptionShowSineLine . 2 | 4 | 8 ，曲线，直线，波浪线
	Source		string	`mapstructure:"source" json:"source" yaml:"source"`				//随机字符串的源（从中取字符）。
	BgColor		*color.RGBA	`mapstructure:"bgcolor" json:"bgcolor" yaml:"bgcolor"`		//背景颜色，透明度255为满
	Fonts		[]string	`mapstructure:"fonts" json:"fonts" yaml:"fonts"`			//从官方的fonts库中选择,

	//Audio
	Language 	[]string	`mapstructure:"language" json:"language" yaml:"language"`	//语音验证码选择语言
}

