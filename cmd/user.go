package cmd

import (
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"gitee.com/wennmu/gopkg.git/doorm"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/errcode"
)

type userInfoResponse struct {
	Id       int64
	Account  string
	CreateAt string `json:"create_at, int"`
}

func (a User) UserInfo(username, password string) (uid int64) {
	var u userInfoResponse
	err := doorm.DB().Table("users").Where("account = ? ", username).Where("passwd = ?", base64.StdEncoding.EncodeToString([]byte(password))).First(&u)
	if err != nil {
		uid = 0
	}
	return u.Id
}

type userListResponse struct {
	Id         int64
	Account    string
	CreateAt   int64
	CreateTime string
}

func (a *User) List(c *gin.Context) (interface{}, error) {
	var users []userListResponse
	res := doorm.DB().Table("users").Select("id, account, create_at").Find(&users)
	if res.Error != nil {
		log.Println(res.Error)
	}
	for i, _ := range users {
		users[i].CreateTime = time.Unix(users[i].CreateAt, 0).Format("2006-01-02")
	}
	return map[string]interface{}{
		"users": users,
	}, nil
}

func (a *User) Add(c *gin.Context) (interface{}, error) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "invalid params",
		}
	}
	data := User{
		Account:  req.Account,
		Passwd:   base64.StdEncoding.EncodeToString([]byte(req.Passwd)),
		CreateAt: time.Now().Unix(),
	}
	result := doorm.DB().Table("users").Create(&data)
	if result.Error != nil {
		return nil, e.AppError{Code: -1, Msg: result.Error.Error()}
	}
	return "", nil
}

func (a *User) Del(c *gin.Context) (interface{}, error) {
	var u User
	id := c.PostForm("id")
	result := doorm.DB().Table("users").Where("id = ? ", id).Delete(&u)
	if result.Error != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": -1,
			"msg":  result.Error.Error(),
		})
		return nil, result.Error
	}
	return "", nil
}
