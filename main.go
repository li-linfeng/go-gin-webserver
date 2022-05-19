package main

import (
	"context"
	"fmt"
	"log"
	"meta/models"
	"meta/pkg/setting"
	"meta/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//初始化全局配置
	setting.Init()
	models.Setup()

	// 注册中间件，注册路由
	router := routers.InitRouter()

	// 获取 http 对象
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 开启协程启动 http 服务
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// 启用 crontab
	// c := cron.New()
	// c.AddFunc("*/3 * * * * *", func() {
	// 	log.Println("start cron")
	// })

	// c.Start()

	// 监听管道是否有终止信号进入，如果信号进入，则进行平滑关闭服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
