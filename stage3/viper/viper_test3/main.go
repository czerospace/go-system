package main

import (
	"fmt"

	"github.com/spf13/viper"

	"go-system/stage3/viper/viper_test3/global"
)

// main 初始化配置文件
func main() {
	// new 得到一个 viper
	v := viper.New()
	// 设置配置文件路径
	configFileName := "config/config.yaml"
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 实例化一个 ServerConfig{} 结构体
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	fmt.Println(global.ServerConfig.ConsulInfo.Host)
	fmt.Println(global.ServerConfig.ConsulInfo.Port)
	fmt.Println(global.ServerConfig.NetInfo.Interface)
	fmt.Println(global.ServerConfig.NetInfo.Subnet)
}
