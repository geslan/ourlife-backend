package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/geslan/ourlife-backend/internal/repository"
	"github.com/geslan/ourlife-backend/internal/services"
	"github.com/geslan/ourlife-backend/internal/websocket"
)

// AIGenerate 生成 AI 回复
func AIGenerate(c *gin.Context) {
	userID := c.GetString("userId")
	chatID := c.Query("chatId")

	var req struct {
		Message     string `json:"message" binding:"required"`
		CharacterID string `json:"characterId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取角色配置
	character, err := repository.NewCharacterRepository().FindByID(req.CharacterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	// 构建请求
	aiReq := services.GenerateRequest{
		Message:     req.Message,
		CharacterID: req.CharacterID,
		UserID:      userID,
		Context: services.ConversationContext{
			Conversation: []string{}, // TODO: 从数据库加载历史对话
			CharacterConfig: services.CharacterConfig{
				Name:         character.Name,
				Personality:  character.Personality,
				Relationship: character.Relationship,
				Profession:   character.Profession,
				Interests:    character.Interests,
				Voice:       character.Voice,
				Bio:         character.Bio,
			},
		},
	}

	// 调用 AI 服务
	resp, err := services.NewAIService("http://localhost:8000").GenerateResponse(aiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 如果提供了 ChatID，通过 WebSocket 广播消息
	if chatID != "" {
		message := map[string]interface{}{
			"id":          uuid.New().String(),
			"chatId":      chatID,
			"senderId":    req.CharacterID,
			"senderType":  "character",
			"content":     resp.Content,
			"type":        resp.Type,
			"a2uiData":    resp.A2UIData,
			"tokenCost":   0,
			"createdAt":   time.Now().Format(time.RFC3339),
		}

		websocket.SendMessageToChat(chatID, message)
	}

	c.JSON(http.StatusOK, resp)
}

// MultiAgent 多智能体编排（待实现）
func MultiAgent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "多智能体编排功能将在 Phase 2 实现",
	})
}

// GenerateImage 图像生成（待实现）
func GenerateImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "图像生成功能将在 Phase 3 实现",
	})
}
