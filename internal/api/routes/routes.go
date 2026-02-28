package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/geslan/ourlife-backend/internal/api/handlers"
	"github.com/geslan/ourlife-backend/internal/api/middleware"
	"github.com/geslan/ourlife-backend/internal/websocket"
)

func SetupRoutes(r *gin.Engine) {
	// 初始化 WebSocket Hub
	websocket.InitWebSocket()

	// Public routes
	public := r.Group("/api")
	{
		// Auth routes
		auth := public.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
			auth.POST("/telegram-webapp", handlers.TelegramWebApp)
		}

		// Character routes (public for listing)
		characters := public.Group("/characters")
		{
			characters.GET("", handlers.ListCharacters)
			characters.GET("/:id", handlers.GetCharacter)
		}
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		user := protected.Group("/user")
		{
			user.GET("/me", handlers.GetCurrentUser)
		}

		// Character management
		characters := protected.Group("/characters")
		{
			characters.POST("", handlers.CreateCharacter)
			characters.PUT("/:id", handlers.UpdateCharacter)
			characters.DELETE("/:id", handlers.DeleteCharacter)
			characters.GET("/me", handlers.GetMyCharacters)
		}

		// Chat routes
		chats := protected.Group("/chats")
		{
			chats.GET("", handlers.ListChats)
			chats.GET("/:id/messages", handlers.GetMessages)
			chats.POST("/:id/messages", handlers.SendMessage)
		}

		// Wallet routes
		wallet := protected.Group("/wallet")
		{
			wallet.GET("/balance", handlers.GetBalance)
			wallet.GET("/transactions", handlers.GetTransactions)
			wallet.POST("/topup", handlers.Topup)
		}

		// Membership routes
		membership := protected.Group("/membership")
		{
			membership.GET("/status", handlers.GetMembershipStatus)
			membership.GET("/plans", handlers.GetMembershipPlans)
			membership.POST("/subscribe", handlers.Subscribe)
			membership.POST("/cancel", handlers.CancelSubscription)
		}

		// AI routes
		ai := protected.Group("/ai")
		{
			ai.POST("/generate", handlers.AIGenerate)
			ai.POST("/multi-agent", handlers.MultiAgent)
			ai.POST("/generate-image", handlers.GenerateImage)
		}
	}

	// WebSocket endpoint
	r.GET("/ws/chat", websocket.HandleWebSocket)
}
