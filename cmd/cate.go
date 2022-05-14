package cmd

import (
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"log"
)

type Cate struct {
	Cate     int64 `json:"Cate, int"`
	Catename string
}

func (s *Cate) List(c *gin.Context) (interface{}, error) {
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
	return map[string]interface{}{
		"cates": cates,
	}, nil
}
