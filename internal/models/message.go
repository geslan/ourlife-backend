package models

import (
	"time"

	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB 用于存储 PostgreSQL JSONB 类型
type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = JSONB{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return errors.New("type not supported")
	}
	return json.Unmarshal(bytes, j)
}

type Message struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ChatID    string    `json:"chatId" gorm:"index"`
	SenderID  string    `json:"senderId"`
	SenderType string   `json:"senderType"` // user, character
	Content   string    `json:"content" gorm:"not null"`
	Type      string    `json:"type" gorm:"default:'text'"` // text, a2ui, image
	A2UIData  JSONB     `json:"a2uiData,omitempty" gorm:"type:jsonb"`
	ImageURL  *string   `json:"imageUrl,omitempty"`
	TokenCost int       `json:"tokenCost" gorm:"default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type MessageType string

const (
	MessageTypeText  MessageType = "text"
	MessageTypeA2UI  MessageType = "a2ui"
	MessageTypeImage MessageType = "image"
)

type SenderType string

const (
	SenderTypeUser      SenderType = "user"
	SenderTypeCharacter SenderType = "character"
)
