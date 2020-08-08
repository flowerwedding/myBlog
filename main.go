/**
 * @Title  main
 * @Description  启动文件
 * @Author  沈来
 * @Update  2020/8/3 16:34
 **/
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"myBlog/global"
	"myBlog/internal/routers"
	"net/http"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.ertupTracer err: %v", err)
	}
}

// @title  博客后台
// @version  1.0
// @description 《Go语言编程之旅》项目练习
// @termsOfService  https://github.com/flowerwedding/myBlog
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()

	//global.Logger.Infof("%s: go-programming-tour-book/%s","eddycjy","blog-service")
}
