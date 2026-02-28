package models

import (
	"time"
)

type Transaction struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string    `json:"userId" gorm:"index"`
	Type        string    `json:"type"` // voice, image, call, topup, membership, reward
	Amount      int       `json:"amount"` // 正数增加，负数减少
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type TransactionType string

const (
	TransactionTypeVoice       TransactionType = "voice"
	TransactionTypeImage       TransactionType = "image"
	TransactionTypeCall        TransactionType = "call"
	TransactionTypeTopup       TransactionType = "topup"
	TransactionTypeMembership  TransactionType = "membership"
	TransactionTypeReward      TransactionType = "reward"
)
