package repository

import (
	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type ChatRepository interface {
	FindByID(id string) (*models.Chat, error)
	FindByUserID(userID string) ([]*models.Chat, error)
	Create(chat *models.Chat) error
	Update(chat *models.Chat) error
	Delete(id string) error
}

type chatRepository struct{}

func NewChatRepository() ChatRepository {
	return &chatRepository{}
}

func (r *chatRepository) FindByID(id string) (*models.Chat, error) {
	var chat models.Chat
	err := database.DB.Where("id = ?", id).First(&chat).Error
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *chatRepository) FindByUserID(userID string) ([]*models.Chat, error) {
	var chats []*models.Chat
	err := database.DB.Where("user_id = ?", userID).Order("updated_at DESC").Find(&chats).Error
	return chats, err
}

func (r *chatRepository) Create(chat *models.Chat) error {
	return database.DB.Create(chat).Error
}

func (r *chatRepository) Update(chat *models.Chat) error {
	return database.DB.Save(chat).Error
}

func (r *chatRepository) Delete(id string) error {
	return database.DB.Delete(&models.Chat{}, "id = ?", id).Error
}
