package repository

import (
	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type MessageRepository interface {
	FindByID(id string) (*models.Message, error)
	FindByChatID(chatID string, limit, offset int) ([]*models.Message, error)
	Create(message *models.Message) error
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) FindByID(id string) (*models.Message, error) {
	var message models.Message
	err := database.DB.Where("id = ?", id).First(&message).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *messageRepository) FindByChatID(chatID string, limit, offset int) ([]*models.Message, error) {
	var messages []*models.Message
	err := database.DB.Where("chat_id = ?", chatID).
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&messages).Error
	return messages, err
}

func (r *messageRepository) Create(message *models.Message) error {
	return database.DB.Create(message).Error
}
