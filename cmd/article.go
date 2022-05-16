package cmd

import (
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (s *Link) List(c *gin.Context) (interface{}, error) {
	var req Link
	uid := c.GetInt64("uid")
	log.Println("ids: ", uid)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
	}
	var articles []Link
	dbimpl := doorm.DB().Table("links").Where("type = ?", req.Type).Where("uid = ?", uid)
	if req.Catename != "" {
		dbimpl = dbimpl.Where("catename = ?", req.Catename)
	}
	dbimpl.Find(&articles)
	return map[string]interface{}{
		"links": articles,
	}, nil
}

func (s *Link) AddArticleLink(c *gin.Context) (interface{}, error) {
	var req Link
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
	}
	data := Link{
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
