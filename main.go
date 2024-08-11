package main

import (
	"lixuefei.com/go-admin/bootstrap"
	"lixuefei.com/go-admin/global"
	"lixuefei.com/go-admin/router"
)

func main() {
	// 基础设施初始化
	bootstrap.Init()

	// 路由初始化
	Router := router.InitRouter()

	// 创建http服务器
	srv := bootstrap.InitServer(":"+global.App.Server.ServiceInfo.Port, Router)

	// 程序关闭前操作
	defer func() {
		// 释放数据库连接
		bootstrap.DestroyProcesses()
	}()

	// 等待中断信号以优雅地关闭服务器（设置 10 秒的超时时间）
	bootstrap.ServerShutdown(srv)
}
