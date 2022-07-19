package config

//定时任务基础参数
type baseTaskParameter struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` 							// 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    							// CRON表达式
	TaskName	 string	`mapstructure:"taskName" json:"taskName" yaml:"taskName"`			  // 定时任务名
}

//定时任务参数配置合集
type Timer struct {
	//数组类型是为了避免参数需要为数组的情况，例如删除多个表
	Detail []ClearTables `mapstructure:"detail" json:"detail" yaml:"detail"`	//信息集合，虽然我们在配置中只测试了一个表
	EmailTask EmailTask	 `mapstructure:"emailTask" json:"emailTask" yaml:"emailTask" `
}

//定时清除表数据，任务所需要的配置参数
type ClearTables struct {
	BaseTaskParameter baseTaskParameter `yaml:"baseTaskParameter"`
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}

//定时任务：发送邮件
type EmailTask struct {
	BaseTaskParameter baseTaskParameter `yaml:"baseTaskParameter"`
}