package cmd

import (
	"gitee.com/wennmu/pkg.git/doorm"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Account struct {
	ID       int64
	Sitename string
	Siteurl  string
	Name     string
	Password string
	Createat int64
	Uid      int64 `json:"Cate, int"`
}

func (a *Account) List(c *gin.Context) {
	var accounts []Account
	uid := c.GetInt64("uid")
	res := doorm.DB().Table("accounts").Select("id, sitename, siteurl, name, createat, uid").Where("uid = ?", uid).Find(&accounts)
	if res.Error != nil {
		log.Println(res.Error)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"data": map[string]interface{}{
			"accounts": accounts,
		},
	})
}
