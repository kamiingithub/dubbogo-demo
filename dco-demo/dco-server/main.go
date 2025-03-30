package main

import (
	"dubbo-demo/dco-server/api"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
)

func main() {
	logger.SetLoggerLevel("debug") // 设置日志级别为 debug
	//api.SetProviderService(&api.DcoCreativeServiceImpl{})
	//config.SetProviderService(&api.DcoCreativeServiceImpl{})
	config.SetProviderServiceWithInfo(&api.DcoCreativeServiceImpl{}, &api.DcoService_ServiceInfo)
	if err := config.Load(); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	fmt.Println("server started")
	select {}
}
