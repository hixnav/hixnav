package router

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav-web/pkg"
	"html/template"
	"net/http"
	"strings"
)

func RegisterFrontRouters(r *gin.Engine) {
	if doconfig.GetBool("is_pack_front") {
		packFrontRouter(r)
	} else {
		frontRouter(r)
	}
}

//动态加载前端代码
func frontRouter(r *gin.Engine) {
	r.Delims("{[{", "}]}")
	r.Static("/assets", "./dist/assets")
	r.LoadHTMLGlob("./dist/*.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

//静态打包前端代码
func packFrontRouter(r *gin.Engine) {
	fs := assetfs.AssetFS{
		Asset:     pkg.Asset,
		AssetDir:  pkg.AssetDir,
		AssetInfo: nil,
		Prefix:    "dist/assets", //一定要加前缀
	}
	r.StaticFS("/assets", &fs)

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dist/index.html", gin.H{})
	})
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range pkg.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := pkg.Asset(name)
		if err != nil {
			continue
		}
		// name := strings.Replace(name, "templates/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

//加载静态文件资源
func loadDataFile(r *gin.Engine) {
	r.Static("/data/file", "./data/file")
}
