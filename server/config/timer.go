package config

type Timer struct {
	Start  bool     `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Spec   string   `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON表达式
	Detail []Detail `mapstructure:"detail" json:"detail" yaml:"detail"`	//待删除表的信息集合，虽然我们在配置中只测试了一个表
}

type Detail struct {	//定时清除表数据，任务所需要的配置参数
	TableName    string `mapstructure:"tableName" json:"tableName" yaml:"tableName"`          // 需要清理的表名
	CompareField string `mapstructure:"compareField" json:"compareField" yaml:"compareField"` // 需要比较时间的字段
	Interval     string `mapstructure:"interval" json:"interval" yaml:"interval"`             // 时间间隔
}
