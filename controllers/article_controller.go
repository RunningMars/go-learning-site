package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-learning-site/backend/models"
	"net/http"
	"strconv"
)

type ArticleController struct {
	repo *models.ArticleRepository
}

func NewArticleController(db *sql.DB) *ArticleController {
	return &ArticleController{
		repo: &models.ArticleRepository{DB: db},
	}
}

func (c *ArticleController) GetArticles(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	search := ctx.Query("search")

	articles, err := c.repo.GetArticles(page, limit, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

func (c *ArticleController) GetArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid article id"})
		return
	}

	article, err := c.repo.GetArticle(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if article == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func (c *ArticleController) GetVideos(ctx *gin.Context) {
	videos := []models.Video{
		{
			ID:         1,
			UserID:     1,
			Title:      "Go语言基础教程",
			URL:        "https://example.com/videos/1",
			CoverImage: sql.NullString{String: "cover1.jpg", Valid: true},
			Category:   "Education",
		},
		{
			ID:         2,
			UserID:     1,
			Title:      "Go Web开发",
			URL:        "https://example.com/videos/2",
			CoverImage: sql.NullString{String: "cover2.jpg", Valid: true},
			Category:   "Web Development",
		},
	}
	ctx.JSON(http.StatusOK, videos)
}

func (c *ArticleController) GetVideo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid video id"})
		return
	}

	// 模拟查找视频
	if id == 1 {
		ctx.JSON(http.StatusOK, models.Video{
			ID:         1,
			UserID:     1,
			Title:      "Go语言基础教程",
			URL:        "https://example.com/videos/1",
			CoverImage: sql.NullString{String: "cover1.jpg", Valid: true},
			Category:   "Education",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "video not found"})
}

func (c *ArticleController) GetEbooks(ctx *gin.Context) {
	ebooks := []models.Ebook{
		{
			ID:         1,
			UserID:     1,
			Title:      "Go语言编程",
			FileURL:    "/ebooks/go-programming.pdf",
			CoverImage: sql.NullString{String: "cover1.jpg", Valid: true},
			Category:   "Programming",
		},
		{
			ID:         2,
			UserID:     1,
			Title:      "Go Web编程",
			FileURL:    "/ebooks/go-web.pdf",
			CoverImage: sql.NullString{String: "cover2.jpg", Valid: true},
			Category:   "Web Development",
		},
	}
	ctx.JSON(http.StatusOK, ebooks)
}

func (c *ArticleController) GetEbook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ebook id"})
		return
	}

	// 模拟查找电子书
	if id == 1 {
		ctx.JSON(http.StatusOK, models.Ebook{
			ID:         1,
			UserID:     1,
			Title:      "Go语言编程",
			FileURL:    "/ebooks/go-programming.pdf",
			CoverImage: sql.NullString{String: "cover1.jpg", Valid: true},
			Category:   "Programming",
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "ebook not found"})
}
