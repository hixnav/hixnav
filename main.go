package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gitee.com/wennmu/haixinnav.git/cmd"
	"gitee.com/wennmu/haixinnav.git/middleware"
)

var (
	db *gorm.DB
)

func main() {
	r := gin.Default()
	// r.Delims("{[{", "}]}")
	//  r.Static("/statics", "./statics")
	r.Static("/assets", "./hixnav/dist/assets")
	r.LoadHTMLGlob("hixnav/dist/index.html")
	// r.LoadHTMLGlob("views/*")
	r.Use(middleware.Request())
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "海芯导航",
		})
	})
	// r.GET("/add-link", func(c *gin.Context){
	// 	c.HTML(http.StatusOK, "add-link.html", gin.H{
	// 		"title": "海芯导航",
	// 	})
	// })
	// r.GET("/cloud-image", func(c *gin.Context){
	// 	c.HTML(http.StatusOK, "cloud-image.html", gin.H{
	// 		"title": "海芯云图",
	// 	})
	// })
	// r.GET("/cloud-HOS", func(c *gin.Context){
	// 	c.HTML(http.StatusOK, "cloud-HOS.html", gin.H{
	// 		"title": "海芯存储",
	// 	})
	// })
	// r.GET("/article", func(c *gin.Context){
	// 	c.HTML(http.StatusOK, "article.html", gin.H{
	// 		"title": "海芯文链",
	// 	})
	// })
	// r.GET("/docs", func(c *gin.Context){
	// 	c.HTML(http.StatusOK, "docs.html", gin.H{
	// 		"title": "海芯笔记",
	// 	})
	// })
	api := r.Group("/api")
	 {
		api.POST("/login", cmd.Login)

		api.POST("/home", new(Nav).home)
		api.POST("/cates", new(Cate).list)
		
		// 文链
		api.POST("/article", new(Article).list)
	}
	api.Use(middleware.Check())
	{
		api.POST("/addLink", new(Nav).addLink)
		api.POST("/addArticleLink", new(Article).addArticleLink)
		// 云存储
		api.POST("/uploadIO", cmd.UploadIO)
		api.POST("/listIO", cmd.GetFileIO)

		// 云图
		api.POST("/upload", cmd.UploadFile)

	}
	db, _ = gorm.Open(mysql.Open(cmd.DNS), &gorm.Config{})
	if err := r.Run("0.0.0.0:8543"); err != nil {
		log.Fatal(err)
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "海芯导航",
	})
}

type Nav struct{
		Id int64
		Cate int64 `json:"Cate, int"`
		Catename string
		Name string
		Logo string
		Desc string
		Url string
	}
	
func (s *Nav) home(c *gin.Context) {
		var cates []Cate
	_ = db.Table("navs").Select("cate as cate, catename as catename").Group("cate,catename").Scan(&cates)
	var data []Nav
	
	
	res := make(map[int64]interface{}, len(cates))
	for _, cate := range cates {
	_ = db.Table("navs").Where("cate = ?", cate.Cate).Find(&data)
		res[cate.Cate] = data
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"navs": res,
	})
}

func (s *Nav) addLink(c *gin.Context){
	var nav Nav
	if err := c.ShouldBindJSON(&nav); err != nil {
		log.Println(err.Error())
         c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
         return
      }
	
	 result := db.Table("navs").Create(&nav)
	 if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
         return
	 }
	 c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

type Cate struct{
		Cate int64 `json:"Cate, int"`
		Catename string
}

func (s *Cate) list(c *gin.Context){
	var cates []Cate
	res := db.Table("navs").Select("cate as cate, catename as catename").Group("cate,catename").Scan(&cates)
	if  res.Error != nil {
		log.Println(res.Error)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"cates": cates,
	})
}


// 文链操作
type Article struct{
	Id int64
	Type int64 `json:"Type, string"`
	Catename string
	Name string
	Logo string
	Url string
}

func (s *Article) list(c *gin.Context) {
	var req Article
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
	 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
	 	return
  }
	var articles []Article
	dbimpl := db.Table("links").Where("type = ?", req.Type).Find(&articles)
	if req.Catename != "" {
		dbimpl = dbimpl.Where("catename = ?", req.Catename)
	}
	dbimpl.Find(&articles)
	c.JSON(http.StatusOK, map[string]interface{}{
		"links": articles,
	})
}

func (s *Article) addArticleLink(c *gin.Context){
	var req Article
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	data := Article{
		Type : req.Type,
		Catename : req.Catename,
		Name :req.Name,
		Logo :req.Logo,
		Url: req.Url,
	}
	result := db.Table("links").Create(&data)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}