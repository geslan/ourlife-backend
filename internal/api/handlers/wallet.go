package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBalance 获取 Token 余额
func GetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"balance": 0,
	})
}

// GetTransactions 获取交易记录
func GetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"transactions": []interface{}{},
	})
}

// Topup 充值（待实现）
func Topup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "not_implemented",
		"message": "充值功能将在 Phase 3 实现",
	})
}
