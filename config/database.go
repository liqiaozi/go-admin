package config

type Database struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConnections  int    `mapstructure:"maxIdleConnections" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	MaxOpenConnections  int    `mapstructure:"maxOpenConnections" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	LogMode             string `mapstructure:"logMode" json:"logMode" yaml:"logMode"`
	EnableFileLogWriter bool   `mapstructure:"enableFileLogWriter" json:"enableFileLogWriter" yaml:"enableFileLogWriter"`
	LogFilename         string `mapstructure:"logFilename" json:"logFilename" yaml:"logFilename"`
}
