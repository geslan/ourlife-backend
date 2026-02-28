package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/geslan/ourlife-backend/internal/services"
)

var aiService *services.AIService

func InitAIHandlers(baseURL string) {
	aiService = services.NewAIService(baseURL)
}

// AIGenerate 生成 AI 回复
func AIGenerate(c *gin.Context) {
	userID := c.GetString("userId")

	var req struct {
		Message     string `json:"message" binding:"required"`
		CharacterID string `json:"characterId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取角色配置
	character, err := characterRepo.FindByID(req.CharacterID)
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
	resp, err := aiService.GenerateResponse(aiReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
