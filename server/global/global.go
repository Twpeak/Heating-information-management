package global

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/Timer"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	G_CONFIG			config.Server
	G_DB				*gorm.DB
	G_LOG				*zap.Logger
	G_VIPER  			*viper.Viper
	G_REDIS         *redis.Client
	G_TimerTaskList Timer.Timer = Timer.NewTimerTask()
	//重点：我们直接写方法，是不会被接受的，因为这里存放的是变量，需要的是一个类型。要不就是接口，要不就是结构体。
	//重点：值得注意的是，我们的结构体(接口实现类)是小写私有的，通过了构造函数(声明返回接口但实际返回的是结构体)，实现了实现类对外的沟通
)