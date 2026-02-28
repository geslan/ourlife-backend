package models

import (
	"time"
)

type Chat struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string    `json:"userId" gorm:"index"`
	CharacterID string    `json:"characterId" gorm:"index"`
	Type        string    `json:"type" gorm:"default:'direct'"` // direct, group
	IsPinned    bool      `json:"isPinned" gorm:"default:false"`
	UnreadCount int       `json:"unreadCount" gorm:"default:0"`
	LastMessage string    `json:"lastMessage"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ChatType string

const (
	ChatTypeDirect ChatType = "direct"
	ChatTypeGroup  ChatType = "group"
)
