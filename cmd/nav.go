package cmd

import (
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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

func (s *Nav) Home(c *gin.Context) (interface{}, error) {
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
