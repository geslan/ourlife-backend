package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/geslan/ourlife-backend/internal/repository"
	"github.com/geslan/ourlife-backend/internal/websocket"
)

var onlineStatusRepo repository.OnlineStatusRepository

func init() {
	onlineStatusRepo = repository.NewOnlineStatusRepository()
}

// SetOnline 设置用户在线
func SetOnline(c *gin.Context) {
	userID := c.GetString("userId")

	if err := onlineStatusRepo.SetOnline(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 广播在线状态
	websocket.SendOnlineStatus(userID, true)

	// 获取所有在线用户
	onlineUserIDs, _ := onlineStatusRepo.GetOnlineUserIDs()

	c.JSON(http.StatusOK, gin.H{
		"message":    "User set as online",
		"onlineCount": len(onlineUserIDs),
	})
}

// SetOffline 设置用户离线
func SetOffline(c *gin.Context) {
	userID := c.GetString("userId")

	if err := onlineStatusRepo.SetOffline(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 广播离线状态
	websocket.SendOnlineStatus(userID, false)

	c.JSON(http.StatusOK, gin.H{
		"message": "User set as offline",
	})
}

// GetOnlineUsers 获取在线用户列表
func GetOnlineUsers(c *gin.Context) {
	onlineUserIDs, err := onlineStatusRepo.GetOnlineUserIDs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"onlineUserIDs": onlineUserIDs,
		"count":          len(onlineUserIDs),
	})
}
