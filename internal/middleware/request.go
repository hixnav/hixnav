package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/hixnav/hixnav.git/internal/errcode"
	"net/http"
)

func Request() gin.HandlerFunc {
	errorToString := func(r interface{}) string {
		switch v := r.(type) {
		case error:
			return v.Error()
		default:
			return r.(string)
		}
	}
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		reqid := uuid.New()
		c.Set("requestId", reqid)

		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":      errcode.ERROR,
					"msg":       errorToString(r),
					"data":      &map[string]string{},
					"requestId": reqid,
				})
				c.Abort()
			}
		}()

		c.Next()
	}

}
