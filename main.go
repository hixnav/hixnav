package main

import (
	"context"
	"gitee.com/wennmu/gopkg.git/doconfig"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/cmd"
	"github.com/hixnav/hixnav.git/internal/router"
	"log"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	if err := cmd.SysInit(); err != nil {
		panic(err)
		return
	}

	r := gin.Default()

	router.Register(r)

	if err := r.Run(doconfig.GetString("port")); err != nil {
		log.Fatal(err)
		panic(err)
	}

	<-ctx.Done()
}
