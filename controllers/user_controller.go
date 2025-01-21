package controllers

import (
	"go-learning-site/backend/models"
	"go-learning-site/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证手机验证码
	if !c.service.VerifyCode(req.Phone, req.Code) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误或已过期"})
		return
	}

	user, err := c.service.Register(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(req.Phone, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// SendVerificationCode 发送验证码
func (c *UserController) SendVerificationCode(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "手机号不能为空"})
		return
	}

	err := c.service.SendVerificationCode(phone)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "验证码已发送"})
}
