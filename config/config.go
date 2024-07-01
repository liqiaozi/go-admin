package config

type Application struct {
	AppConfig App      `mapstructure:"app" json:"app" yaml:"app"`
	Log       Log      `mapstructure:"logger" json:"logger" yaml:"logger"`
	Database  Database `mapstructure:"database" json:"database" yaml:"database"`
	Cors      CORS     `mapstructure:"cors" json:"cors" yaml:"cors"`
}
