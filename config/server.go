package config

type Server struct {
	ServiceInfo Service  `mapstructure:"serviceInfo" json:"serviceInfo" yaml:"serviceInfo"` // 服务信息
	Log         Log      `mapstructure:"logger" json:"logger" yaml:"logger"`                // 日志配置
	Database    Database `mapstructure:"database" json:"database" yaml:"database"`
	Cors        CORS     `mapstructure:"cors" json:"cors" yaml:"cors"`
}
