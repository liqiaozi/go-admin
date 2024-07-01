package config

type Application struct {
	AppConfig App `mapstructure:"app" json:"app" yaml:"app"`
	Log       Log `mapstructure:"log" json:"log" yaml:"log"`
}
