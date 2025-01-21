package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	controller := NewArticleController(db)

	r.GET("/api/articles", controller.GetArticles)
	r.GET("/api/articles/:id", controller.GetArticle)
	r.GET("/api/videos", controller.GetVideos)
	r.GET("/api/videos/:id", controller.GetVideo)
	r.GET("/api/ebooks", controller.GetEbooks)
	r.GET("/api/ebooks/:id", controller.GetEbook)

	return r
}

func TestGetArticles(t *testing.T) {
	// 创建模拟数据库连接
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// 初始化测试路由
	r := setupTestRouter(db)

	// 创建测试请求
	req, _ := http.NewRequest("GET", "/api/articles?page=1&limit=10", nil)
	w := httptest.NewRecorder()

	// 执行请求
	r.ServeHTTP(w, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
}

func TestGetArticle(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := setupTestRouter(db)

	req, _ := http.NewRequest("GET", "/api/articles/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetVideos(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := setupTestRouter(db)

	req, _ := http.NewRequest("GET", "/api/videos", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(response))
}

func TestGetVideo(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := setupTestRouter(db)

	req, _ := http.NewRequest("GET", "/api/videos/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Go语言基础教程", response["title"])
}

func TestGetEbooks(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := setupTestRouter(db)

	req, _ := http.NewRequest("GET", "/api/ebooks", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(response))
}

func TestGetEbook(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := setupTestRouter(db)

	req, _ := http.NewRequest("GET", "/api/ebooks/1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Go语言编程", response["title"])
}
