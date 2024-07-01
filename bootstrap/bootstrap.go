package bootstrap

func Init() {
	// 初始化配置
	initializeConfig()
	// 初始化日志
	initLog()
	// 初始化数据库
	initializeDB()
}
