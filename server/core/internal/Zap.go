package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Zap_in 内部逻辑的Zap内置对象，包含（继承）了Zap配置对象
type Zap_in struct {
	Zap *config.Zap
}

func NewzapIn() *Zap_in {
	return &Zap_in{Zap: &global.G_CONFIG.Zap}
}

/**

func New(core zapcore.Core, options ...Option) *Logger
zapcore.Core需要三个配置——Encoder，WriteSyncer，LogLevel。

这里的函数关系：GetZapCores-->调用GetEncoderCore,其需要两个参数，一个是当前等级字符串，一个是判断当前等级的封装方法（这个参数给newCore用的）
GetEncoderCore,通过NewCore获取Core ，其需要三个参数，分别调用——>GetEncoder,GetWriteSyncer,判断当前等级的封装方法
GetWriteSyncer配置日志切割和写到哪里，而GetEncoder返回编码器，所以稍微麻烦点。我们主要关注配置EncoderConfig


 */

// GetZapCores 根据配置文件的Level获取 []zapcore.Core 切片
func (z *Zap_in)GetZapCores() []zapcore.Core {
	//先初始化容器,zap日志等级共分为：DebugLevel，InfoLevel，WarnLevel，ErrorLevel，DPanicLevel，PanicLevel，FatalLevel
	//共七个等级，其值依次增加。所以下面循环获取当前等级并循环
	cores := make([]zapcore.Core,0,7)
	for level := z.TransportLevel(); level <= zapcore.FatalLevel; level++{
		cores = append(cores,z.GetEncoderCore(level,z.GetLevelPriority(level)))	//不断往cores中添加不同等级的cores
	}
	return cores
}

// GetEncoderCore 获取Encoder的 zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap_in) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {	//这里参数l仅仅是用来分割日志时命名用的
	return zapcore.NewCore(z.GetEncoder(), GetWriteSyncer(l.String()), level)
}

// GetWriteSyncer 获取 zapcore.WriteSyncer
// GVA中使用rotatelogs 来切割日志，我们这里尝试使用Lumberjack
//rotatelogs 是根据天数时间来切割，方便设置过期时间，而LumberJack是通过文档数量大小来切割
func GetWriteSyncer(level string) zapcore.WriteSyncer {
	/*fileWriter, err := rotatelogs.New(
		path.Join(global.G_CONFIG.Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.G_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err*/
	lumberJackLogger := &lumberjack.Logger{
		Filename:   global.G_CONFIG.Zap.Director+"/"+time.Now().Format("2006.01.02")+level+".log",
		MaxSize:    100,		//若不设置，默认100M
		MaxBackups: 5,			//保留旧文件的最大个数
		MaxAge:     30,			//保留旧文件的最大天数
		Compress:   false,		//是否压缩/归档旧文件
		LocalTime:   true,		//是否使用本地时间
	}

	//根据配置是否输出到控制台
	if global.G_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)

}

// GetEncoder 获取 zapcore.Encoder
//这里又需要补一下 zapcore.Encoder 编码器的知识了
/**
Encoder接口内嵌了ObjectEncoder，定义了Clone、EncodeEntry方法；
ObjectEncoder接口定义了各种类型的一系列Add方法；
MapObjectEncoder实现了ObjectEncoder接口，内部使用map[string]interface{}来存放数据。
且Encoder接口中包含了其配置对象EncoderConfig结构体，我们通过这个去配置Encoder
 */
func (z *Zap_in)GetEncoder() zapcore.Encoder {
	if global.G_CONFIG.Zap.Format == "json"{	//若日志的输出格式为：json
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())	//则通过json格式解析配置，创建编译器
	}
	return zapcore.NewJSONEncoder(z.GetEncoderConfig())
}






