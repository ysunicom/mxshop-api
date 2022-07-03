package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
)

func GetEnvInfo(env string)bool{
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig(){
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-pro.yaml",configFilePrefix)
	if debug{
		configFileName = fmt.Sprintf("user-web/%s-debug.yaml",configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig();err != nil{
		zap.S().Panic("读取配置文件失败:",err.Error())
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil{
		zap.S().Panic("读取配置文件到结构体失败:",err.Error())
	}
	zap.S().Info("配置信息:%v",global.ServerConfig)
}
