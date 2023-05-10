package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/hixnav/hixnav.git/cmd"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//校验登录状态
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearToken := c.GetHeader("Authorization")
		tokenSlice := strings.Split(bearToken, " ")
		tokenStr = tokenSlice[1]
		// tokenStr = c.Query("token")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(cmd.GlobalAppSecret), nil
		})
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "Illegal token",
			})
			c.Abort()
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		var uid int64
		if ok && token.Valid {
			uid, err = strconv.ParseInt(fmt.Sprintf("%.f", claims["uid"]), 10, 64)
			if err != nil || uid <= 0 {
				c.JSON(http.StatusOK, map[string]interface{}{
					"code":    http.StatusUnauthorized,
					"message": "Illegal user",
				})
			}
		}
		c.Set("uid", uid)
		log.Println("uid:", uid)
		c.Next()
	}
}
