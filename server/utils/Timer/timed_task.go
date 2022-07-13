package Timer

import (
	"github.com/robfig/cron/v3"
	"sync"
)
/**
作为一个任务管理器：主要实现了任务列表的增删改查
 */

// Timer 时间任务管理器的接口
type Timer interface {
	AddTaskByFunc(taskName string, spec string, task func()) (cron.EntryID, error)
	AddTaskByJob(taskName string, spec string, job interface{ Run() }) (cron.EntryID, error)
	FindCron(taskName string) (*cron.Cron, bool)
	StartTask(taskName string)
	StopTask(taskName string)
	Remove(taskName string, id int)
	Clear(taskName string)
	Close()
}


//timer 定时任务管理器
type timer struct {
	taskList map[string]*cron.Cron 	//存储任务列表
	sync.Mutex		//互斥锁,方便对资源进行锁控制
}

// NewTimerTask 构造函数
//通过了构造函数和被实现的接口(声明返回接口但实际返回的是结构体)，实现私有结构体对外的沟通
func NewTimerTask() Timer {
	return &timer{taskList: make(map[string]*cron.Cron)}
}

// AddTaskByFunc 通过函数的方法添加任务
func (t *timer)AddTaskByFunc(taskName string,spec string,task func())(cron.EntryID, error)  {
	//常规操作，每次对timer进行操作时，必须要先将资源锁住，避免脏读数据
	t.Lock()
	defer t.Unlock()
	//判断存储列表中是否有该方法名，若没有则填充一个cron实例，准备添加任务
	if _,ok := t.taskList[taskName]; !ok{
		t.taskList[taskName] = cron.New(cron.WithSeconds())
	}
	//添加回调任务，再开启
	id,err := t.taskList[taskName].AddFunc(spec,task)
	t.taskList[taskName].Start()
	return id, err
}

// AddTaskByJob 通过接口的方法添加任务
func (t *timer)AddTaskByJob(taskName string,spec string ,job interface{Run()})(cron.EntryID,error)  {
	t.Lock()
	defer t.Unlock()
	if _,ok := t.taskList[taskName]; !ok {
		t.taskList[taskName] = cron.New()
	}
	id,err := t.taskList[taskName].AddJob(spec,job)
	t.taskList[taskName].Start()
	return id, err
}

// FindCron 获取对应taskName的cron 可能会为空
func (t *timer)FindCron(taskName string)(*cron.Cron, bool)  {
	t.Lock()
	defer t.Unlock()
	v,ok := t.taskList[taskName]
	return v,ok
}

// StartTask 开始任务
func (t *timer) StartTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v,ok := t.taskList[taskName]; ok {
		v.Start()
	}
}

// StopTask 停止任务
func (t *timer) StopTask(taskName string) {
	t.Lock()
	defer t.Unlock()
	if v,ok := t.taskList[taskName]; ok {
		v.Stop()
	}
}

// Remove 从taskName 删除指定任务(只删除回调任务，不map存储删除任务名和cron)
func (t *timer) Remove(taskName string, id int) {
	t.Lock()
	defer t.Unlock()
	if v,ok:= t.taskList[taskName]; ok  {
		v.Remove(cron.EntryID(id))
	}
}

// Clear 清除任务（从mao存储中删掉这个cron和任务）
func (t *timer)Clear(taskName string)  {
	t.Lock()
	defer t.Unlock()
	if v,ok := t.taskList[taskName]; ok{
		v.Stop()
		delete(t.taskList,taskName)
	}
}

// Close 释放资源（停掉所有任务）
func (t *timer) Close() {
	t.Lock()
	defer t.Unlock()
	for _, v := range t.taskList {
		v.Stop()
	}
}

