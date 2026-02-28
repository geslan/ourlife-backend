package repository

import (
	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type UserRepository interface {
	FindByID(id string) (*models.User, error)
	FindByTelegramID(telegramID int64) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByTelegramID(telegramID int64) (*models.User, error) {
	var user models.User
	err := database.DB.Where("telegram_id = ?", telegramID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return database.DB.Save(user).Error
}
