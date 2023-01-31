package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// Register 服务注册
func Register(address string, port int, name string, tags []string, id string) error {
	// 生成配置信息
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.137.134:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.137.1:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	// 注册
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

// AllServices 发现所有服务
func AllServices() {
	// 生成配置信息
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.137.134:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

// FilterService 过滤服务
func FilterService() {
	// 生成配置信息
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.137.134:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func main() {
	err := Register("192.168.137.1", 8021, "user-web", []string{"mxshop", "niko"}, "user-web")
	if err != nil {
		panic(err)
	}

	AllServices()
	FilterService()
}
