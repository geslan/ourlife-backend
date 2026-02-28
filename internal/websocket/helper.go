package websocket

import (
	"github.com/gin-gonic/gin"
)

func BroadcastMessage(event string, data interface{}) {
	GetHub().BroadcastMessage(event, data)
}

func BroadcastToChat(chatID string, event string, data interface{}) {
	GetHub().BroadcastToChat(chatID, event, data)
}

func SendMessageToChat(chatID string, message map[string]interface{}) {
	GetHub().BroadcastToChat(chatID, "message", message)
}

func SendTypingStatus(chatID, userID string, isTyping bool) {
	GetHub().BroadcastToChat(chatID, "typing", gin.H{
		"userId":   userID,
		"isTyping": isTyping,
	})
}

func SendOnlineStatus(userID string, isOnline bool) {
	GetHub().BroadcastMessage("online", gin.H{
		"userId":    userID,
		"isOnline": isOnline,
	})
}
