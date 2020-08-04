/**
 * @Title  setting
 * @Description  初始化配置信息
 * @Author  沈来
 * @Update  2020/8/3 19:48
 **/
package setting

import "github.com/spf13/viper"

type Setting struct{
	vp *viper.Viper
}

func NewSetting() (*Setting, error){
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil,err
	}

	return &Setting{vp},nil
}