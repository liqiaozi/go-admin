package config

// Service 服务信息
type Service struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"` // 服务名称
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 服务端口
}
