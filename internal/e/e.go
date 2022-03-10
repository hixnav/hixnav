package e

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err AppError) Error() string {
	return err.Msg
}

type WrapperHandle func(c *gin.Context) (interface{}, error)

func ErrorWrapper(handle WrapperHandle) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err != nil {
			appError := err.(AppError)
			c.JSON(http.StatusOK, gin.H{
				"code":    appError.Code,
				"message": appError.Msg,
				"data":    &map[string]string{},
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "",
			"data":    data,
		})
	}
}