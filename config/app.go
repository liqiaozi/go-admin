package config

type App struct {
	AppName string `mapstructure:"appName" json:"appName" yaml:"appName"` // 应用名称
	Port    string `mapstructure:"port" json:"port" yaml:"port"`          // 启动端口
}
