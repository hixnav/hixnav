package cmd

import (
	"gitee.com/wennmu/pkg.git/doorm"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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

func (a *Account) Add(c *gin.Context) {
	var req Account
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}
	data := Account{
		Sitename: req.Sitename,
		Siteurl:  req.Siteurl,
		Name:     req.Name,
		Password: req.Password,
		Createat: time.Now().Unix(),
		Uid:      c.GetInt64("uid"),
	}
	result := doorm.DB().Table("accounts").Create(&data)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
	})
}

func (a *Account) Del(c *gin.Context) {
	var account Account
	id := c.PostForm("id")
	uid := c.GetInt64("uid")
	result := doorm.DB().Table("accounts").Where("id = ? and uid = ?", id, uid).Delete(&account)
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
