package models

import (
	"time"

	"database/sql/driver"
	"encoding/json"
	"errors"
)

// StringArray 用于存储 PostgreSQL 数组类型
type StringArray []string

func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return "[]", nil
	}
	return json.Marshal(a)
}

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = StringArray{}
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
	return json.Unmarshal(bytes, a)
}

type Character struct {
	ID          string      `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string      `json:"userId" gorm:"index"`
	Name        string      `json:"name" gorm:"not null"`
	Age         int         `json:"age"`
	Avatar      string      `json:"avatar"`
	Banner      string      `json:"banner"`
	Bio         string      `json:"bio"`
	Personality StringArray `json:"personality" gorm:"type:jsonb"`
	Relationship string     `json:"relationship"`
	Profession  string      `json:"profession"`
	Interests   StringArray `json:"interests" gorm:"type:jsonb"`
	Voice       string      `json:"voice"`
	Style       string      `json:"style"` // realistic, anime
	Gender      string      `json:"gender"` // female, trans
	IsOfficial  bool        `json:"isOfficial" gorm:"default:false"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
}
