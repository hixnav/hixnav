package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/errcode"
	"log"
)

func InitCos() error {
	//初始化存储
	GlobalCOS = doconfig.GetString("cos")
	GlobalCOSSecretID = doconfig.GetString("cos_secret_id")
	GlobalCOSSecretKey = doconfig.GetString("cos_secret_key")
	return nil
}

type COSConfig struct {
	COS          string `json:"cos"`
	COSSecretID  string `json:"cosSecretID"`
	COSSecretKey string `json:"cosSecretKey"`
}

//存储迁移
func MigrateCos(c *gin.Context) (interface{}, error) {
	var cosconf COSConfig
	if err := c.ShouldBindJSON(&cosconf); err != nil {
		log.Println(err.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "invalid params",
		}
	}
	doconfig.Set("cos", cosconf.COS)
	doconfig.Set("cos_secret_id", cosconf.COSSecretID)
	doconfig.Set("cos_secret_key", cosconf.COSSecretKey)

	doconfig.WriteConfigAs("config/config.yaml")
	//TODO 创建内置存储桶

	//初始化存储
	if err := InitCos(); err != nil {
		return "", err
	}
	return "", nil
}
