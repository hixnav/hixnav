package cmd

import (
	"errors"
	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
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

type AccountResponse struct {
	ID         int64
	Sitename   string
	Siteurl    string
	Name       string
	Password   string
	Createat   int64
	CreateTime string
	Uid        int64 `json:"Cate, int"`
}

func (a *Account) List(c *gin.Context) (interface{}, error) {
	var accounts []AccountResponse
	uid := c.GetInt64("uid")
	res := doorm.DB().Table("accounts").Select("id, sitename, siteurl, name, createat, uid").Where("uid = ?", uid).Find(&accounts)
	if res.Error != nil {
		log.Println(res.Error)
	}
	for i, _ := range accounts {
		accounts[i].CreateTime = time.Unix(accounts[i].Createat, 0).Format("2006-01-02")
	}
	return map[string]interface{}{
		"accounts": accounts,
	}, nil
}

func (s *Account) Edit(c *gin.Context) (interface{}, error) {
	var acc Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return nil, err
	}
	acc.Uid = c.GetInt64("uid")
	result := doorm.DB().Table("accounts").Where("uid = ?", acc.Uid).Where("id = ?", acc.ID).Updates(&acc)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed"})
		return nil, result.Error
	}
	return "", nil
}

func (a *Account) Add(c *gin.Context) (interface{}, error) {
	var req Account
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		return nil, errors.New("invalid params")
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
		return nil, e.AppError{Code: -1, Msg: result.Error.Error()}
	}
	return "", nil
}

func (a *Account) Del(c *gin.Context) (interface{}, error) {
	var account Account
	id := c.PostForm("id")
	uid := c.GetInt64("uid")
	result := doorm.DB().Table("accounts").Where("id = ? and uid = ?", id, uid).Delete(&account)
	if result.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  result.Error.Error(),
		})
		return nil, result.Error
	}
	return "", nil
}
