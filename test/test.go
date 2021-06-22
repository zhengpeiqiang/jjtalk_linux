package main

import (
	"modules/test/router"
	"modules/test/task"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//服务路由注册、定义
	router.HttpServerRun()
	//
	task.InitTask()
	//做一个管道堵塞住服务，使服务可以一直运行，并且捕获信号，做退出操作
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
