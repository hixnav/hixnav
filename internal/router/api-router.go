package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/cmd"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/middleware"
)

func RegisterApiRouters(r *gin.Engine) {
	api := r.Group("/api")

	api.Use(middleware.Request())
	{
		api.POST("/login", e.ErrorWrapper(cmd.Login))
		api.POST("/logout", e.ErrorWrapper(cmd.Logout))

		api.POST("/home", e.ErrorWrapper(new(cmd.Nav).Home))
		api.POST("/cates", e.ErrorWrapper(new(cmd.Cate).List))
	}
	api.Use(middleware.Request(), middleware.Check())
	{
		//导航
		api.POST("/addLink", e.ErrorWrapper(new(cmd.Nav).AddLink))
		api.POST("/editLink", e.ErrorWrapper(new(cmd.Nav).EditLink))
		api.POST("/delLink", e.ErrorWrapper(new(cmd.Nav).DelLink))
		// 文链
		api.POST("/article", e.ErrorWrapper(new(cmd.Article).List))
		api.POST("/addArticleLink", e.ErrorWrapper(new(cmd.Article).AddArticleLink))
		// 云存储
		api.POST("/uploadIO", e.ErrorWrapper(new(cmd.Upload).UploadIO))
		api.POST("/listIO", e.ErrorWrapper(new(cmd.Upload).GetFileIO))
		api.POST("/downIO", e.ErrorWrapper(new(cmd.Upload).DownFileIO))
		api.POST("/delIO", e.ErrorWrapper(new(cmd.Upload).DelFileIO))

		// 云图
		api.POST("/upload", e.ErrorWrapper(new(cmd.Upload).UploadFile))
		//云账号
		api.POST("/account", e.ErrorWrapper(new(cmd.Account).List))
		api.POST("/addAccount", e.ErrorWrapper(new(cmd.Account).Add))
		api.POST("/editAccount", e.ErrorWrapper(new(cmd.Account).Edit))
		api.POST("/delAccount", e.ErrorWrapper(new(cmd.Account).Del))

	}
}
