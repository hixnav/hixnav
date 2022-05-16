package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"github.com/gin-gonic/gin"
)

func setConfig() {
	if doconfig.GetString("is_pack_front") == "" {
		doconfig.Set("is_pack_front", true)
	}
	if doconfig.GetString("port") == "" {
		doconfig.Set("port", "0.0.0.0:8543")
	}
	//初始化日志
	if doconfig.GetString("log_path") == "" {
		doconfig.Set("log_path", "log/hixnav-access.log")
	}
	if doconfig.GetString("log_error_path") == "" {
		doconfig.Set("log_error_path", "log/hixnav-error.log")
	}
}

//设置数据库配置
func SetDBConfig(c *gin.Context) (interface{}, error) {
	return "", nil
}

//设置存储配置
func SetOSSConfig(c *gin.Context) (interface{}, error) {
	return "", nil
}
