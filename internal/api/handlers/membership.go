package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMembershipStatus 获取会员状态
func GetMembershipStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"membership": "free",
		"isPremium":  false,
	})
}

// GetMembershipPlans 获取会员方案
func GetMembershipPlans(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"plans": []interface{}{},
	})
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
