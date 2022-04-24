package main

import (
	"gitee.com/hixnav/hixnav.git/cmd"
	"gitee.com/hixnav/hixnav.git/internal/e"
	"gitee.com/hixnav/hixnav.git/middleware"
	"gitee.com/wennmu/pkg.git/doorm"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	cmd.Init()
	r := gin.Default()
	//r.Static("/assets", "./web/dist/assets")
	//r.LoadHTMLGlob("web/dist/index.html")

	fs := assetfs.AssetFS{
		Asset:     cmd.Asset,
		AssetDir:  cmd.AssetDir,
		AssetInfo: nil,
		Prefix:    "web/dist/assets", //一定要加前缀
	}
	r.StaticFS("/assets", &fs)

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)

	r.Use(middleware.Request())
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "海芯导航",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/login", e.ErrorWrapper(cmd.Login))
		api.POST("/logout", e.ErrorWrapper(cmd.Logout))

		api.POST("/home", new(Nav).home)
		api.POST("/cates", new(Cate).list)
	}
	api.Use(middleware.Check())
	{
		//导航
		api.POST("/addLink", new(Nav).addLink)
		api.POST("/editLink", new(Nav).editLink)
		api.POST("/delLink", new(Nav).delLink)
		// 文链
		api.POST("/article", new(Article).list)
		api.POST("/addArticleLink", new(Article).addArticleLink)
		// 云存储
		api.POST("/uploadIO", new(cmd.Upload).UploadIO)
		api.POST("/listIO", new(cmd.Upload).GetFileIO)
		api.POST("/downIO", new(cmd.Upload).DownFileIO)
		api.POST("/delIO", new(cmd.Upload).DelFileIO)

		// 云图
		api.POST("/upload", new(cmd.Upload).UploadFile)
		//云账号
		api.POST("/account", new(cmd.Account).List)
		api.POST("/addAccount", new(cmd.Account).Add)
		api.POST("/delAccount", new(cmd.Account).Del)

	}
	//初始化数据库
	if err := doorm.InitMysqlDB(doorm.GlobalDBOptions{DSN: cmd.GlobalMysqlDNS}); err != nil {
		log.Fatal(err)
	}
	if err := r.Run("0.0.0.0:8543"); err != nil {
		log.Fatal(err)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "海芯导航",
	})
}

type Nav struct {
	Id       int64
	Cate     int64 `json:"Cate, int"`
	Catename string
	Name     string
	Logo     string
	Desc     string
	Url      string
	Uid      int64
}

func (s *Nav) home(c *gin.Context) {
	var cates []Cate
	uid := c.GetString("uid")
	var uids = []string{"1"}
	if uid != "" {
		uids = append(uids, uid)
	}
	_ = doorm.DB().Table("navs").Select("cate as cate, catename as catename").Where("uid IN ?", uids).Group("cate,catename").Scan(&cates)
	var data []Nav

	res := make(map[int64]interface{}, len(cates))
	for _, cate := range cates {
		_ = doorm.DB().Table("navs").Where("cate = ?", cate.Cate).Find(&data)
		res[cate.Cate] = data
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"navs": res,
	})
}

func (s *Nav) addLink(c *gin.Context) {
	var nav Nav
	if err := c.ShouldBindJSON(&nav); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	nav.Uid = c.GetInt64("uid")
	result := doorm.DB().Table("navs").Create(&nav)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

func (s *Nav) editLink(c *gin.Context) {
	var nav Nav
	if err := c.ShouldBindJSON(&nav); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	nav.Uid = c.GetInt64("uid")
	result := doorm.DB().Table("navs").Updates(&nav)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

func (s *Nav) delLink(c *gin.Context) {
	var nav Nav
	id := c.PostForm("id")
	uid := c.GetInt64("uid")
	result := doorm.DB().Table("navs").Where("id = ? and uid = ?", id, uid).Delete(&nav)
	if result.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

type Cate struct {
	Cate     int64 `json:"Cate, int"`
	Catename string
}

func (s *Cate) list(c *gin.Context) {
	var cates []Cate
	uid := c.GetString("uid")
	var uids = []string{"1"}
	if uid != "" {
		uids = append(uids, uid)
	}
	res := doorm.DB().Table("navs").Select("cate as cate, catename as catename").Where("uid IN ?", uids).Group("cate,catename").Scan(&cates)
	if res.Error != nil {
		log.Println(res.Error)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"cates": cates,
	})
}

// 文链操作
type Article struct {
	Id       int64
	Type     int64 `json:"Type, string"`
	Catename string
	Name     string
	Logo     string
	Url      string
	Uid      int64
}

func (s *Article) list(c *gin.Context) {
	var req Article
	uid := c.GetInt64("uid")
	log.Println("ids: ", uid)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	var articles []Article
	dbimpl := doorm.DB().Table("links").Where("type = ?", req.Type).Where("uid = ?", uid)
	if req.Catename != "" {
		dbimpl = dbimpl.Where("catename = ?", req.Catename)
	}
	dbimpl.Find(&articles)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"links": articles,
		},
	})
}

func (s *Article) addArticleLink(c *gin.Context) {
	var req Article
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	data := Article{
		Type:     req.Type,
		Catename: req.Catename,
		Name:     req.Name,
		Logo:     req.Logo,
		Url:      req.Url,
		Uid:      c.GetInt64("uid"),
	}
	result := doorm.DB().Table("links").Create(&data)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range cmd.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := cmd.Asset(name)
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
