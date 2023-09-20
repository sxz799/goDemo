package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var CapComeFrom []string
var CapManageType []string
var CapStatus []string
var CapZJSF []string

func init() {
	log.Println("正在应用配置文件...")

	if _, err := os.Stat("conf.yaml"); os.IsNotExist(err) {
		log.Panicln("配置文件不存在!")
		return
	}
	viper.SetConfigFile("conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln("配置文件读取失败...")
		return
	}

	CapComeFrom = viper.GetStringSlice("dict.from")
	CapManageType = viper.GetStringSlice("dict.manage")
	CapStatus = viper.GetStringSlice("dict.status")
	CapZJSF = viper.GetStringSlice("dict.zjsf")
}
