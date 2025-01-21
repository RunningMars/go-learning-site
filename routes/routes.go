package routes

import (
	"database/sql"

	"github.com/RunningMars/go-learning-site/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	// 初始化controller
	articleController := controllers.NewArticleController(db)

	// 文章路由
	api := r.Group("/api")
	{
		api.GET("/articles", articleController.GetArticles)
		api.GET("/articles/:id", articleController.GetArticle)

		// 视频路由
		api.GET("/videos", articleController.GetVideos)
		api.GET("/videos/:id", articleController.GetVideo)

		// 电子书路由
		api.GET("/ebooks", articleController.GetEbooks)
		api.GET("/ebooks/:id", articleController.GetEbook)
	}
}
