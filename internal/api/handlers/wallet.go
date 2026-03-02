package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/geslan/ourlife-backend/internal/repository"
	"github.com/geslan/ourlife-backend/internal/services"
)

var (
	tokenService      *services.TokenService
	userRepository     repository.UserRepository
	transactionRepository repository.TransactionRepository
)

func init() {
	userRepository = repository.NewUserRepository()
	transactionRepository = repository.NewTransactionRepository()
	tokenService = services.NewTokenService(userRepository, transactionRepository)
}

// GetBalance 获取 Token 余额
func GetBalance(c *gin.Context) {
	userID := c.GetString("userId")

	balance, err := tokenService.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
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

	transactions, err := transactionRepository.FindByUserID(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"limit":      limit,
		"offset":     offset,
	})
}

// Topup 充值 Tokens
func Topup(c *gin.Context) {
	userID := c.GetString("userId")

	var req struct {
		Amount int    `json:"amount" binding:"required,min=1"`
		Method string `json:"method"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tokenService.AddTokens(userID, req.Amount, "Topup via "+req.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	balance, _ := tokenService.GetBalance(userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tokens topped up successfully",
		"amount":  req.Amount,
		"balance": balance,
	})
}

// ConsumeTokens 消耗 Tokens（内部使用）
func ConsumeTokens(c *gin.Context) {
	userID := c.GetString("userId")

	var req struct {
		Amount int    `json:"amount" binding:"required,min=1"`
		Reason string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tokenService.ConsumeTokens(userID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	balance, _ := tokenService.GetBalance(userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tokens consumed successfully",
		"amount":  req.Amount,
		"balance": balance,
	})
}
