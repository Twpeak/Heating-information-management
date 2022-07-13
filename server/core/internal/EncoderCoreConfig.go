package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap/zapcore"
	"time"
)

// GetEncoderConfig 获取zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap_in) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.G_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// CustomTimeEncoder 自定义日志输出时间格式
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap_in) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(global.G_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

// ZapEncodeLevel 因为存的是字符串而不是 zapcore.levelEncoder 对象，所以这个方法是转换获取对象用的
//根据 EncodeLevel 返回 zapcore.LevelEncoder
//这个地方返回的数据是 编码器中EncodeLevel这个字段的值
func (z *Zap_in)ZapEncodeLevel() zapcore.LevelEncoder  {
	switch  { //这里练习一下switch的用法，一个是判断，下面的方法是匹配，
	case z.Zap.EncodeLevel == "LowercaseLevelEncoder":		//小写编码器（默认）
		return zapcore.LowercaseLevelEncoder
	case z.Zap.EncodeLevel == "LowercaseColorLevelEncoder":	//小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.Zap.EncodeLevel == "CapitalLevelEncoder":	//小写编码器带颜色
		return zapcore.CapitalLevelEncoder
	case z.Zap.EncodeLevel == "CapitalColorLevelEncoder":	//小写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder			//默认
	}
}

