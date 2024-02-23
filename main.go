package main

import (
	"context"
	"embed"
	"gitee.com/wennmu/gopkg.git/doconfig"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/cmd"
	"github.com/hixnav/hixnav.git/internal/router"
	"io/fs"
	"log"
	"net/http"
)

//go:embed front/dist
var staticFS embed.FS

func main() {
	ctx, _ := context.WithCancel(context.Background())

	if err := cmd.SysInit(); err != nil {
		panic(err)
		return
	}

	r := gin.Default()

	subFS, err := fs.Sub(staticFS, "front/dist")
	if err != nil {
		panic(err)
		return
	}
	g := r.Group("/")
	{
		g.StaticFS("/", http.FS(subFS))
	}
	router.Register(r)

	if err := r.Run(doconfig.GetString("port")); err != nil {
		log.Fatal(err)
		panic(err)
	}

	<-ctx.Done()

}
