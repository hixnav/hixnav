package cmd

import (
	"strconv"

	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/dologger"
	"github.com/hixnav/hixnav.git/internal/storage"
	"github.com/spf13/viper"
)

// GlobalSyncScheduler holds the reference to the COS sync scheduler
var GlobalSyncScheduler *storage.SyncScheduler

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
		dologger.Error("connect db failed, err:", err.Error())
		//return err
	}
	//初始化存储
	if err := InitCos(); err != nil {
		return err
	}
	//初始化COS定时同步 (仅当使用sqlite时)
	driverType := doconfig.GetString("storage_driver")
	if driverType == "sqlite" {
		syncEnabled := doconfig.GetString("cos_sync_enabled")
		if syncEnabled == "true" {
			intervalStr := doconfig.GetString("cos_sync_interval")
			if intervalStr == "" {
				intervalStr = "3600"
			}
			interval, err := strconv.Atoi(intervalStr)
			if err != nil {
				interval = 3600
			}

			sync := storage.NewCOSSync(
				GlobalCOS,
				GlobalCOSSecretID,
				GlobalCOSSecretKey,
				doconfig.GetString("cos_sync_bucket"),
				storage.GetDB(),
			)

			GlobalSyncScheduler = storage.NewSyncScheduler(sync, interval)
			GlobalSyncScheduler.Start()
			dologger.Info("COS sync scheduler started with interval:", intervalStr, "s")
		}
	}
	return nil
}
