/**
 * @Title  setting
 * @Description  配置管理
 * @Author  沈来
 * @Update  2020/8/3 20:22
 **/
package global

import (
	"myBlog/pkg/logger"
	"myBlog/pkg/setting"
)

var(
	ServerSetting *setting.ServerSettingS
	AppSetting    *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger  *logger.Logger
)