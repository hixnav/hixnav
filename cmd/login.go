package cmd

import (
	"log"
	"gitee.com/wennmu/haixinnav.git/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	"gitee.com/wennmu/haixinnav.git/internal/e"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) (interface{}, error) {
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}

	uid := (model.Admin{}).UserInfo(req.Username, req.Password)
	if uid <= 0 {
		return nil, e.AppError{Code: -1, Msg: "user not found"}

	}

	atCliams := jwt.MapClaims{}

	atCliams["authorized"] = true
	atCliams["uid"] = uid
	atCliams["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atCliams)
	token, err := at.SignedString([]byte(APP_SECRET))
	if err != nil {
		return nil, e.AppError{Code: -1, Msg: err.Error()}
	}


	return map[string]interface{}{
		"token":  token,
		"expire": atCliams["exp"],
	}, nil
}

func Logout(c *gin.Context) (interface{}, error) {
	return "", nil
}