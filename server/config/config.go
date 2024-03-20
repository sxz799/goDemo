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
var CapCWCat01 []string
var CapCWCat02 []string
var CapCWCat03 []string
var CapCWCat04 []string
var CapCWCat05 []string
var CapCWCat06 []string

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

	CapCWCat01=viper.GetStringSlice("dict.cwcat-01")
	CapCWCat02=viper.GetStringSlice("dict.cwcat-02")
	CapCWCat03=viper.GetStringSlice("dict.cwcat-03")
	CapCWCat04=viper.GetStringSlice("dict.cwcat-04")
	CapCWCat05=viper.GetStringSlice("dict.cwcat-05")
	CapCWCat06=viper.GetStringSlice("dict.cwcat-06")
}
