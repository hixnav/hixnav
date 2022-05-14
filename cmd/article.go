package cmd

import (
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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

func (s *Article) List(c *gin.Context) (interface{}, error) {
	var req Article
	uid := c.GetInt64("uid")
	log.Println("ids: ", uid)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
	}
	var articles []Article
	dbimpl := doorm.DB().Table("links").Where("type = ?", req.Type).Where("uid = ?", uid)
	if req.Catename != "" {
		dbimpl = dbimpl.Where("catename = ?", req.Catename)
	}
	dbimpl.Find(&articles)
	return map[string]interface{}{
		"links": articles,
	}, nil
}

func (s *Article) AddArticleLink(c *gin.Context) (interface{}, error) {
	var req Article
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
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
		return nil, result.Error
	}
	return "", nil
}
