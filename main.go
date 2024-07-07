package main

import (
	"lixuefei.com/go-admin/bootstrap"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/router"
)

func main() {
	// 初始化操作
	bootstrap.Init()
	// 初始化路由
	Router := router.InitializeRouter()
	// 创建http服务器
	srv := bootstrap.InitServer(":"+global.App.Server.ServiceInfo.Port, Router)

	// 程序关闭前，释放数据库连接
	defer func() {
		bootstrap.DestroyProcesses()
	}()

	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	bootstrap.ServerShutdown(srv)
}
