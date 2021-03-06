package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/dologger"
	"github.com/spf13/viper"
)

//初始化系统所有子模块
func SysInit() error {
	//初始化配置
	if err := doconfig.InitConfigs(); err != nil && err != err.(viper.ConfigFileNotFoundError) {
		return err
	}
	setConfig()
	dologger.InitLogger(&dologger.LoggerOptions{
		LogDebug:     true,
		LogPath:      doconfig.GetString("log_path"),
		LogErrorPath: doconfig.GetString("log_error_path"),
	})
	//初始化APP
	if err := InitApp(); err != nil {
		return err
	}
	//初始化数据库
	if err := InitDB(); err != nil {
		dologger.Error("connect db faild，err:", err.Error())
		//return err
	}
	//初始化存储
	if err := InitCos(); err != nil {
		return err
	}
	return nil
}
