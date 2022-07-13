package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis()  {
	//从配置中获取并初始化redis
	redisCfg := global.G_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: 		redisCfg.Addr,
		Password: 	redisCfg.Password,
		DB:			redisCfg.DB,
	})
	//测试redis是否可以连接成功。并打印日志
	pong ,err := client.Ping(context.Background()).Result()
	if err != nil {
		global.G_LOG.Error("redis connet ping failed ,err :",zap.Error(err))
		fmt.Printf("redis connet ping failed ,err :%v", zap.Error(err))
	}else {
		global.G_LOG.Info("redis connect ping response:" ,zap.String("pong",pong))
		global.G_REDIS = client
	}
}