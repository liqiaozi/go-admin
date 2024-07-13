package config

// Log 配置
type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`                // 日志等级
	RootDir    string `mapstructure:"rootDir" json:"rootDir" yaml:"rootDir"`          // 日志目录
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`       // 文件名
	Format     string `mapstructure:"format" json:"format" yaml:"format"`             // 格式
	ShowLine   bool   `mapstructure:"showLine" json:"showLine" yaml:"showLine"`       // 是否显示调用行
	MaxBackups int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups"` // 旧文件的最大个数
	MaxSize    int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`          // MB
	MaxAge     int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`             // 保留天数day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`       // 是否压缩
}
