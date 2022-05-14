package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/dologger"
)

//初始化系统所有子模块
func SysInit() error {
	//初始化配置
	if err := doconfig.InitConfigs(); err != nil {
		return err
	}
	//初始化日志
	dologger.InitLogger(&dologger.LoggerOptions{
		LogDebug:     true,
		LogPath:      doconfig.GetString("log_path"),
		LogErrorPath: doconfig.GetString("log_error_path"),
	})

	return nil
}
