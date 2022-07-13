package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"fmt"
)

func Timer()  {
	if global.G_CONFIG.Timer.Start { // 若配置文件中选择了开启定时任务
		for i := range global.G_CONFIG.Timer.Detail{	 // 需要清理的表名,需要比较时间的字段,时间间隔

			go func(detail config.Detail) {
				_ , err := global.G_TimerTaskList.AddTaskByFunc("ClearDB",global.G_CONFIG.Timer.Spec, func() {	//在定时任务管理器中添加回调函数，名为ClearDB
					err := utils.ClearTable(global.G_DB,detail.TableName,detail.CompareField,detail.Interval)	//输入待删除表的信息 global.Detail中的信息都是其参数信息
					if err != nil {
						fmt.Println("timer error :",err)
					}
				})
				if err != nil {
					fmt.Println("add timer error:",err)
				}
			}(global.G_CONFIG.Timer.Detail[i])	//调用闭包
		}

	}
}