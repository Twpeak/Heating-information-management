package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

func Timer()  {
	//手动开启定时任务
	//自动删表定时任务
	if global.G_CONFIG.Timer.Detail[0].BaseTaskParameter.Start{
		ClearDB()
	}
	if global.G_CONFIG.Timer.EmailTask.BaseTaskParameter.Start{
		SendMessage()
	}

}


//定时删表任务
func ClearDB()  {
	for i := range global.G_CONFIG.Timer.Detail{	 // 需要清理的表名,需要比较时间的字段,时间间隔
		go func(detail config.ClearTables) {
			_ , err := global.G_TimerTaskList.AddTaskByFunc(global.G_CONFIG.Timer.Detail[0].BaseTaskParameter.TaskName,global.G_CONFIG.Timer.Detail[0].BaseTaskParameter.Spec, func() {	//在定时任务管理器中添加回调函数，名为ClearDB
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

//定时发送邮件任务
//分析：每周给医院管理员发送邮件
//推送内容：给医院管理员推送本周登记的记录数量，以及各个医生登记的患者数量
//查询数据库，本周，当前医院，医生(用户)名/id 分组 统计总数和每组数
func SendMessage()  {
	//并发执行定时任务
	go func(emailDate config.EmailTask) {
		_ , err := global.G_TimerTaskList.AddTaskByFunc(global.G_CONFIG.Timer.EmailTask.BaseTaskParameter.TaskName,global.G_CONFIG.Timer.EmailTask.BaseTaskParameter.Spec, func() {
			err := service.ServiceGroupApp.ExampleServiceGroup.EmailService.SendEamilToBoss()
			if err != nil {
				global.G_LOG.Error("发送邮件失败",zap.Error(err))
				return
			}
		})
		if err != nil {
			fmt.Println("add timer error:",err)
		}
	}(global.G_CONFIG.Timer.EmailTask)
}
