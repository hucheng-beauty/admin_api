package main

import (
	"admin_api/global"
	"admin_api/initialize"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := initialize.Routers()

	go func() {
		if err := r.Run(fmt.Sprintf("%s:%s", global.ServerConfig.Host, global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()

	zap.S().Info("port:", global.ServerConfig.Port)

	// 优雅退出,接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("3s 后关闭服务。。。")
	time.Sleep(3 * time.Second)
}
