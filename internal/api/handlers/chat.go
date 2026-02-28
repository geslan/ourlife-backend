package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/internal/repository"
)

var (
	chatRepo       repository.ChatRepository
	messageRepo    repository.MessageRepository
	transactionRepo repository.TransactionRepository
	walletRepo     repository.WalletRepository
	membershipRepo repository.MembershipRepository
)

func InitChatHandlers() {
	chatRepo = repository.NewChatRepository()
	messageRepo = repository.NewMessageRepository()
}

func InitWalletHandlers() {
	transactionRepo = repository.NewTransactionRepository()
	walletRepo = repository.NewWalletRepository()
}

func InitMembershipHandlers() {
	membershipRepo = repository.NewMembershipRepository()
}

// ListChats 聊天列表
func ListChats(c *gin.Context) {
	userID := c.GetString("userId")

	chats, err := chatRepo.FindByUserID(userID)
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

	messages, err := messageRepo.FindByChatID(chatID, limit, offset)
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

	if err := messageRepo.Create(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

// WebSocketHandler WebSocket 端点（待实现）
func WebSocketHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "WebSocket 功能将在 Phase 2 实现",
	})
}

// GetBalance 获取 Token 余额
func GetBalance(c *gin.Context) {
	userID := c.GetString("userId")

	balance, err := walletRepo.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// GetTransactions 获取交易记录
func GetTransactions(c *gin.Context) {
	userID := c.GetString("userId")
	limit := 20
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

	transactions, err := transactionRepo.FindByUserID(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// Topup 充值（待实现）
func Topup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "充值功能将在 Phase 3 实现",
	})
}

// GetMembershipStatus 获取会员状态
func GetMembershipStatus(c *gin.Context) {
	userID := c.GetString("userId")

	status, err := membershipRepo.GetStatus(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

// GetMembershipPlans 获取会员方案
func GetMembershipPlans(c *gin.Context) {
	plans := []gin.H{
		{
			"id":       "monthly",
			"name":     "Monthly",
			"price":    12,
			"currency": "USD",
			"savings":  0,
		},
		{
			"id":       "quarterly",
			"name":     "Quarterly",
			"price":    29,
			"currency": "USD",
			"savings":  19,
		},
		{
			"id":       "yearly",
			"name":     "Yearly",
			"price":    79,
			"currency": "USD",
			"savings":  45,
		},
	}

	c.JSON(http.StatusOK, gin.H{"plans": plans})
}

// Subscribe 开通会员（待实现）
func Subscribe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "会员订阅功能将在 Phase 3 实现",
	})
}

// CancelSubscription 取消订阅（待实现）
func CancelSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "取消订阅功能将在 Phase 3 实现",
	})
}
