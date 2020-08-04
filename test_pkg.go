/**
 * @Title  test_pkg
 * @Description  公共组件测试
 * @Author  沈来
 * @Update  2020/8/3 21:47
 **/
package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"myBlog/global"
	"myBlog/internal/model"
	"myBlog/pkg/logger"
	"myBlog/pkg/setting"
	"time"
)

func setupSetting() error{
	setting,err := setting.NewSetting()
	if err != nil{
		return err
	}
	err = setting.ReadSection("Server",&global.ServerSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("App",&global.AppSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("Database",&global.DatabaseSetting)
	if err != nil{
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error{
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil{
		return err
	}
	return nil
}

func setupLogger() error{
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		//生成的日志文件目录
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	}, "",log.LstdFlags).WithCaller(2)

	return nil
}