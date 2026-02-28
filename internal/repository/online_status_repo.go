package repository

import (
	"time"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type OnlineStatusRepository interface {
	SetOnline(userID string) error
	SetOffline(userID string) error
	GetOnlineUserIDs() ([]string, error)
}

type onlineStatusRepository struct{}

func NewOnlineStatusRepository() OnlineStatusRepository {
	return &onlineStatusRepository{}
}

func (r *onlineStatusRepository) SetOnline(userID string) error {
	return database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("last_seen", time.Now()).
		Error
}

func (r *onlineStatusRepository) SetOffline(userID string) error {
	return database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("last_seen", time.Now().Add(-30*time.Minute)).
		Error
}

func (r *onlineStatusRepository) GetOnlineUserIDs() ([]string, error) {
	var userIDs []string
	err := database.DB.Model(&models.User{}).
		Where("last_seen > ?", time.Now().Add(-30*time.Minute)).
		Pluck("id", &userIDs).
		Error
	return userIDs, err
}
