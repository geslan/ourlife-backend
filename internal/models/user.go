package models

import (
	"time"
)

type User struct {
	ID           string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	TelegramID   int64      `json:"telegramId" gorm:"uniqueIndex;not null"`
	Username     string     `json:"username"`
	Name         string     `json:"name"`
	Avatar       string     `json:"avatar"`
	Membership   string     `json:"membership" gorm:"default:'free'"` // free, premium
	TokenBalance int        `json:"tokenBalance" gorm:"default:0"`
	LastSeen     *time.Time `json:"lastSeen"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
}

type UserRole string

const (
	RoleGuest   UserRole = "guest"
	RoleUser    UserRole = "user"
	RolePremium UserRole = "premium"
	RoleAdmin   UserRole = "admin"
)
