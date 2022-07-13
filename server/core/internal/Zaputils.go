package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// TransportLevel 根据字符串转化为 zapcore.Level
//我们判断了当前的日志等级。最后写给某一等级的日志文件中
func (z *Zap_in)TransportLevel() zapcore.Level  {
	z.Zap.Level = strings.ToLower(z.Zap.Level)//所有字母小写
	switch z.Zap.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel		//默认
	}
}


// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
//我们若想把不同等级的日志分开记录，我们就需要动态的判断每一句的日志等级。这个方法就是用来判断日志等级的。
//我们封装了这个方法，并传入zapcore.newCore()方法的第三个参数里。不像一般的直接传入一个zap.Level。而是传入zap.LevelEnablerFunc即可实现
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap_in) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
