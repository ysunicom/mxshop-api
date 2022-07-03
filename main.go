package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
)

func main(){
	//1.初始化logger
	initialize.InitLogger()
	//2.初始化配置文件
	initialize.InitConfig()
	//3.初始化routers
	Router := initialize.Routers()

	zap.S().Infof("启动服务器:%d",global.ServerConfig.Port)

	if err := Router.Run(fmt.Sprintf(":%d",global.ServerConfig.Port));err != nil{
		zap.S().Panic("启动失败:",err.Error())
	}
}