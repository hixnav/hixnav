package cmd

import (
	"encoding/base64"
	"fmt"
	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/dologger"
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/errcode"
	"log"
	"time"
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
type DbConfig struct {
	Addr      string
	Dbname    string
	Account   string
	Password  string
	IsMigrate bool
}

func MigrateDB(c *gin.Context) (interface{}, error) {
	var dbconf DbConfig
	if err := c.ShouldBindJSON(&dbconf); err != nil {
		log.Println(err.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "invalid params",
		}
	}
	doconfig.Set("DB", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconf.Account, dbconf.Password, dbconf.Addr, dbconf.Dbname))

	if err := createDir("./config"); err != nil {
		return "", e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}
	doconfig.WriteConfigAs("config/config.yaml")

	//初始化数据库
	if err := InitDB(); err != nil {
		return "", e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}

	if !dbconf.IsMigrate {
		return "", nil
	}
	//创建数据库表
	err := doorm.DB().Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Nav{}, &Link{}, &Account{})
	if err != nil {
		log.Println(err.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "init failed",
		}
	}
	//内置用户
	res := doorm.DB().Create(&User{
		Account:  defaultUsername,
		Passwd:   base64.StdEncoding.EncodeToString([]byte(defaultUserPasswd)),
		CreateAt: time.Now().Unix(),
	})
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  res.Error.Error(),
		}
	}
	return "", nil
}

//检查数据库可用性
func DbTest(c *gin.Context) (interface{}, error) {
	if err := doorm.InitMysqlDB(doorm.GlobalDBOptions{DSN: doconfig.GetString("db")}); err != nil {
		return "", e.AppError{Code: -1, Msg: "db invalid.Please modify the database configuration information."}
	}
	return "", nil
}
