package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Video struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

var videos = []Video{
	{ID: 1, Title: "Go语言基础教程", URL: "https://example.com/videos/1"},
	{ID: 2, Title: "Go Web开发", URL: "https://example.com/videos/2"},
}

func GetVideos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": videos,
	})
}

func GetVideo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	for _, video := range videos {
		if video.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": video,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
}
