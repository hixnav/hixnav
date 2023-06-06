package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
)

func (s *Nav) Home(c *gin.Context) (interface{}, error) {
	var cates []Cate
	uid := c.GetString("uid")
	var uids = []string{"1"}
	if uid != "" {
		uids = append(uids, uid)
	}
	_ = doorm.DB().Table("navs").Select("cate as cate, catename as catename").Where("uid IN ?", uids).Group("cate,catename").Scan(&cates)

	res := make(map[int64]interface{}, len(cates))
	for _, cate := range cates {
		var data []Nav
		_ = doorm.DB().Table("navs").Where("cate = ?", cate.Cate).Find(&data)
		res[cate.Cate] = data
	}
	return map[string]interface{}{
		"navs": res,
	}, nil
}

func (s *Nav) AddLink(c *gin.Context) (interface{}, error) {
	var nav Nav
	if err := c.ShouldBindJSON(&nav); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
	}
	nav.Uid = c.GetInt64("uid")
	result := doorm.DB().Table("navs").Create(&nav)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return nil, result.Error
	}
	return "", nil
}

func (s *Nav) EditLink(c *gin.Context) (interface{}, error) {
	var nav Nav
	if err := c.ShouldBindJSON(&nav); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
	}
	nav.Uid = c.GetInt64("uid")
	result := doorm.DB().Table("navs").Updates(&nav)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return nil, result.Error
	}
	return "", nil
}

func (s *Nav) DelLink(c *gin.Context) (interface{}, error) {
	var nav Nav
	id := c.PostForm("id")
	uid := c.GetInt64("uid")
	result := doorm.DB().Table("navs").Where("id = ? and uid = ?", id, uid).Delete(&nav)
	if result.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  result.Error.Error(),
		})
		return nil, result.Error
	}
	return "", nil
}

func (s *Nav) ExportLink(c *gin.Context) (interface{}, error) {
	uid := c.GetInt64("uid")
	// 执行查询
	var navs []Nav
	if err := doorm.DB().Model(&Nav{}).
		Where("uid = ?", uid).
		Select("catename, name, url, logo").
		Group("catename, name, url, logo").
		Find(&navs).Error; err != nil {
		return nil, err
	}

	// 按照 catename 归类，输出 JSON 结果
	type Row struct {
		Name string `json:"name"`
		Href string `json:"href"`
		Icon string `json:"icon"`
	}
	result := make(map[string][]Row)
	for _, nav := range navs {
		var row = Row{
			Name: nav.Name,
			Href: nav.Url,
			Icon: nav.Logo,
		}
		result[nav.Catename] = append(result[nav.Catename], row)
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil
}
