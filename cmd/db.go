package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/dologger"
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
)

func InitDB() error {
	//初始化数据库
	if err := doorm.InitMysqlDB(doorm.GlobalDBOptions{DSN: doconfig.GetString("DB")}); err != nil {
		dologger.Error(err)
		return err
	}
	return nil
}

//数据库迁移
func MigrateDB(c *gin.Context) (interface{}, error) {
	//创建数据库表
	//内置用户
	return "", nil
}
