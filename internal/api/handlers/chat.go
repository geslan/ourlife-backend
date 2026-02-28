package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/internal/repository"
	"github.com/geslan/ourlife-backend/internal/websocket"
)

// ListChats 聊天列表
func ListChats(c *gin.Context) {
	userID := c.GetString("userId")

	chats, err := repository.NewChatRepository().FindByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}

// GetMessages 获取消息历史
func GetMessages(c *gin.Context) {
	chatID := c.Param("id")
	limit := 50
	offset := 0

	if l, ok := c.GetQuery("limit"); ok {
		if val, err := strconv.Atoi(l); err == nil {
			limit = val
		}
	}
	if o, ok := c.GetQuery("offset"); ok {
		if val, err := strconv.Atoi(o); err == nil {
			offset = val
		}
	}

	messages, err := repository.NewMessageRepository().FindByChatID(chatID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

// SendMessage 发送消息
func SendMessage(c *gin.Context) {
	userID := c.GetString("userId")
	chatID := c.Param("id")

	var req struct {
		Content string `json:"content" binding:"required"`
		Type    string `json:"type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建用户消息
	message := &models.Message{
		ID:        uuid.New().String(),
		ChatID:    chatID,
		SenderID:  userID,
		SenderType: string(models.SenderTypeUser),
		Content:   req.Content,
		Type:      req.Type,
		TokenCost: 0,
		CreatedAt: time.Now(),
	}

	if err := repository.NewMessageRepository().Create(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 通过 WebSocket 广播消息
	websocket.SendMessageToChat(chatID, gin.H{
		"id":         message.ID,
		"chatId":     chatID,
		"senderId":   userID,
		"senderType": "user",
		"content":    req.Content,
		"type":       req.Type,
		"createdAt":  time.Now().Format(time.RFC3339),
	})

	c.JSON(http.StatusCreated, message)
}

// WebSocketHandler WebSocket 端点
func WebSocketHandler(c *gin.Context) {
	websocket.HandleWebSocket(c)
}

// BroadcastTypingStatus 广播打字状态（供 Handlers 使用）
func BroadcastTypingStatus(chatID, userID string, isTyping bool) {
	websocket.SendTypingStatus(chatID, userID, isTyping)
}

// BroadcastMessage 广播消息（供 Handlers 使用）
func BroadcastMessage(event string, data interface{}) {
	websocket.BroadcastMessage(event, data)
}

// BroadcastToChat 广播消息到特定聊天室（供 Handlers 使用）
func BroadcastToChat(chatID string, event string, data interface{}) {
	websocket.BroadcastToChat(chatID, event, data)
}
