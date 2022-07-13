package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing_key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires_time" yaml:"expires_time"`
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
}
