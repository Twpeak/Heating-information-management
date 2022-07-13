package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/mojocn/base64Captcha"
)



func GetDriver(Drivertype string) base64Captcha.Driver{
	switch  Drivertype{
	case "Digit":
		return &base64Captcha.DriverDigit{
	Height: global.G_CONFIG.Captcha.ImgHeight,
	Width: global.G_CONFIG.Captcha.ImgWidth,
	Length: global.G_CONFIG.Captcha.Length,
	MaxSkew: global.G_CONFIG.Captcha.MaxSkwe,
	DotCount: global.G_CONFIG.Captcha.DotCount,
	}
	case "String":
		return &base64Captcha.DriverString{
			Height: global.G_CONFIG.Captcha.ImgHeight,
			Width: global.G_CONFIG.Captcha.ImgWidth,
			NoiseCount: global.G_CONFIG.Captcha.NoiseCount,
			ShowLineOptions: global.G_CONFIG.Captcha.ShowLineOptions,
			Length: global.G_CONFIG.Captcha.Length,
			Source: global.G_CONFIG.Captcha.Source,
			//BgColor: global.G_CONFIG.Captcha.BgColor,
			//Fonts: global.G_CONFIG.Captcha.Fonts,
		}
	case "Math":
		return &base64Captcha.DriverMath{
			Height: global.G_CONFIG.Captcha.ImgHeight,
			Width: global.G_CONFIG.Captcha.ImgWidth,
			NoiseCount: global.G_CONFIG.Captcha.NoiseCount,
			ShowLineOptions: global.G_CONFIG.Captcha.ShowLineOptions,
			//BgColor: global.G_CONFIG.Captcha.BgColor,
		}
	case "Chinese":
		return &base64Captcha.DriverChinese{
			Height: global.G_CONFIG.Captcha.ImgHeight,
			Width: global.G_CONFIG.Captcha.ImgWidth,
			NoiseCount: global.G_CONFIG.Captcha.NoiseCount,
			ShowLineOptions: global.G_CONFIG.Captcha.ShowLineOptions,
			Length: global.G_CONFIG.Captcha.Length,
			Source: global.G_CONFIG.Captcha.Source,
			BgColor: global.G_CONFIG.Captcha.BgColor,
			Fonts: global.G_CONFIG.Captcha.Fonts,
		}
	case "Audio":
		return &base64Captcha.DriverAudio{
			Length: global.G_CONFIG.Captcha.Length,
			Language: global.G_CONFIG.Captcha.Language[0],	//选择语言
		}
	default:
		return &base64Captcha.DriverDigit{
			Height: global.G_CONFIG.Captcha.ImgHeight,
			Width: global.G_CONFIG.Captcha.ImgWidth,
			Length: global.G_CONFIG.Captcha.Length,
			MaxSkew: global.G_CONFIG.Captcha.MaxSkwe,
			DotCount: global.G_CONFIG.Captcha.DotCount,
		}
	}
}
