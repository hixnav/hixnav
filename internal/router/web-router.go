package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/cmd"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/middleware"
)

func RegisterWebRouters(r *gin.Engine) {
	admin := r.Group("/ginf")

	admin.Use(middleware.Request())
	{
		admin.GET("/", e.ErrorWrapper(new(cmd.Hello).Get))
	}
}
