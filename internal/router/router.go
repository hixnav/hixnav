package router

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	RegisterFrontRouters(r)
	RegisterWebRouters(r)
	RegisterApiRouters(r)
}
