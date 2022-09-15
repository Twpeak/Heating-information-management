package config

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                   // 本地文件访问路径
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // 本地文件存储路径
	Static		string `mapstructure:"static" json:"static" yaml:"static"`
	StaticPath	string `mapstructure:"static-path" json:"static-path" yaml:"static-path"`
}
