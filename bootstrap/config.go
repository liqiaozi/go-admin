package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"lixuefei.com/go-admin/global"
	"os"
)

// 初始化加载配置文件
func initializeConfig() *viper.Viper {
	// 设置配置文件路径,生产环境可以通过设置环境变量来改变配置文件路径
	config := "./resources/application.json"
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}
	fmt.Println("[bootstrap] read config begin...")
	fmt.Println("[bootstrap] config file path: ", config)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("watch config changed:", in.Name)
		if err := v.Unmarshal(&global.App.Server); err != nil {
			fmt.Errorf("reload config error：%s \n", err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Server); err != nil {
		panic(fmt.Errorf("read config error: %s \n", err))
	}
	fmt.Println("[bootstrap] read config end...")
	return v
}
