package Viper

import (
	"log"

	"demo/Global"

	"github.com/spf13/viper"
)

func Config() {
	// 文件路径
	viper.AddConfigPath("./Config")
	// 文件信息
	viper.SetConfigName("Config")
	viper.SetConfigType("toml")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}
	// 解析至结构体
	if err := viper.Unmarshal(&Global.Config); err != nil {
		log.Fatal("Unmarshal confiig faled: ", err)
	}
}
