package cmd

import (
	"gitee.com/wennmu/haixinnav.git/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	uid := (model.Admin{}).UserInfo(req.Username, req.Password)
	if uid <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return

	}

	atCliams := jwt.MapClaims{}

	atCliams["authorized"] = true
	atCliams["uid"] = uid
	atCliams["exp"] = time.Now().Add(time.Hour * 24).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atCliams)
	token, err := at.SignedString([]byte(APP_SECRET))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sign error"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"expire": atCliams["exp"],
	})
		return
}

func Logout(c *gin.Context) (interface{}, error) {
	return "", nil
}