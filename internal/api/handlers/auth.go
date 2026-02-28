package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/internal/repository"
	"github.com/geslan/ourlife-backend/pkg/jwt"
)

// Register 用户注册
func Register(c *gin.Context) {
	var req struct {
		Username   string `json:"username" binding:"required"`
		Name       string `json:"name" binding:"required"`
		TelegramID int64  `json:"telegramId" binding:"required"`
		Avatar     string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户是否已存在
	_, err := repository.NewUserRepository().FindByTelegramID(req.TelegramID)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// 创建用户
	user := &models.User{
		TelegramID:   req.TelegramID,
		Username:     req.Username,
		Name:         req.Name,
		Avatar:       req.Avatar,
		Membership:   string(models.RoleUser),
		TokenBalance: 0,
	}

	if err := repository.NewUserRepository().Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找用户（简化版本，实际应该验证密码）
	user, err := repository.NewUserRepository().FindByTelegramID(0)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

// TelegramWebApp Telegram WebApp 认证
func TelegramWebApp(c *gin.Context) {
	var req struct {
		TelegramID int64  `json:"telegramId" binding:"required"`
		Username   string `json:"username"`
		Name       string `json:"name"`
		Avatar     string `json:"avatar"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查找或创建用户
	user, err := repository.NewUserRepository().FindByTelegramID(req.TelegramID)
	if err != nil {
		// 创建新用户
		user = &models.User{
			TelegramID:   req.TelegramID,
			Username:     req.Username,
			Name:         req.Name,
			Avatar:       req.Avatar,
			Membership:   string(models.RoleUser),
			TokenBalance: 0,
		}
		if err := repository.NewUserRepository().Create(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) {
	userID := c.GetString("userId")

	user, err := repository.NewUserRepository().FindByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
