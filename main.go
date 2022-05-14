package main

import (
	"context"
	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/cmd"
	"github.com/hixnav/hixnav.git/internal/router"
	"log"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	cmd.Init()

	if err := cmd.SysInit(); err != nil {
		panic(err)
		return
	}

	r := gin.Default()

	router.Register(r)

	//初始化数据库
	if err := doorm.InitMysqlDB(doorm.GlobalDBOptions{DSN: cmd.GlobalMysqlDNS}); err != nil {
		log.Fatal(err)
	}
	if err := r.Run(doconfig.GetString("port")); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
}
