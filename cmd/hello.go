package cmd

import "github.com/gin-gonic/gin"

type Hello struct {
}

func (h *Hello) Get(c *gin.Context) (interface{}, error) {
	return "hello", nil
}

func (h *Hello) Handle(c *gin.Context) {
}
