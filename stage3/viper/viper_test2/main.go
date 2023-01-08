package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

/*
 	如何将生产环境和开发环境的配置隔离
 	需求是实现 不用改任何代码，实现获取不同的配置文件
	方法:在开发环境的电脑或服务器上设置一个环境变量
*/

// GetEnvInfo 获取系统环境变量
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
func main() {
	debug := GetEnvInfo("SHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-debug.yaml", configFilePrefix)
	}
	// new 得到一个 viper
	v := viper.New()
	// 设置配置文件路径
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 实例化一个 ServerConfig{} 结构体
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	// viper 动态监听配置文件的变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("port"))
		fmt.Println(v.Get("mysql"))
		fmt.Println(serverConfig.MysqlInfo.Host)
	})

	fmt.Println(serverConfig)
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("port"))
	fmt.Println(v.Get("mysql"))
	fmt.Println(serverConfig.MysqlInfo.Host)
	time.Sleep(time.Second * 300)
}
