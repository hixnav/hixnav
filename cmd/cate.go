package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/storage"
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
	res := storage.GetDB().Table("navs").Select("cate as cate, catename as catename").Where("uid IN ?", uids).Group("cate,catename").Scan(&cates)
	if res.Error != nil {
		log.Println(res.Error)
	}
	return map[string]interface{}{
		"cates": cates,
	}, nil
}
