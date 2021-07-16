/**
* @Description: 获取项目配置信息
* @Author: jinyidong
* @Date: 2021/5/19
* @Version V1.0
 */
package util

import (
	"github.com/spf13/viper"
	"os"
)

var projConfig *viper.Viper

func init() {
	projConfig = viper.New()
	path, _ := os.Getwd()
	projConfig.AddConfigPath(path)
	projConfig.SetConfigName("config") //设置读取的文件名
	projConfig.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := projConfig.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetConfig(key string) interface{} {
	return projConfig.Get(key)
}

func GetString(key string) string {
	return projConfig.GetString(key)
}
