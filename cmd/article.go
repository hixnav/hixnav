package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
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
	if !checkLink(req.Logo, "GET") {
		req.Logo = doconfig.GetString("default_article_icon")
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

func (s *Link) ExportArticleLink(c *gin.Context) (interface{}, error) {
	uid := c.GetInt64("uid")
	// 执行查询
	var links []Link
	if err := doorm.DB().Model(&Link{}).
		Where("uid = ?", uid).
		Select("catename, name, url, logo").
		Group("catename, name, url, logo").
		Find(&links).Error; err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  err.Error(),
		})
		return nil, err
	}

	// 按照 catename 归类，输出 JSON 结果
	type Row struct {
		Name string `json:"name"`
		Href string `json:"href"`
		Icon string `json:"icon"`
	}
	result := make(map[string][]Row)
	for _, link := range links {
		var row = Row{
			Name: link.Name,
			Href: link.Url,
			Icon: link.Logo,
		}
		result[link.Catename] = append(result[link.Catename], row)
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil
}
