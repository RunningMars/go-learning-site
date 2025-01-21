package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Ebook struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	File  string `json:"file"`
}

var ebooks = []Ebook{
	{ID: 1, Title: "Go语言编程", File: "/ebooks/go-programming.pdf"},
	{ID: 2, Title: "Go Web编程", File: "/ebooks/go-web.pdf"},
}

func GetEbooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": ebooks,
	})
}

func GetEbook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ebook ID"})
		return
	}

	for _, ebook := range ebooks {
		if ebook.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data": ebook,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Ebook not found"})
}
