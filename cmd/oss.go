package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"github.com/gin-gonic/gin"
)

func InitCos() error {
	//初始化存储
	GlobalCOS = doconfig.GetString("COS")
	GlobalCOSSecretID = doconfig.GetString("COSSecretID")
	GlobalCOSSecretKey = doconfig.GetString("COSSecretKey")
	return nil
}

//存储迁移
func MigrateCos(c *gin.Context) (interface{}, error) {
	//创建内置存储桶
	return "", nil
}
