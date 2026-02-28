package repository

import (
	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type CharacterRepository interface {
	FindByID(id string) (*models.Character, error)
	FindByUserID(userID string) ([]*models.Character, error)
	List(limit, offset int) ([]*models.Character, error)
	Create(character *models.Character) error
	Update(character *models.Character) error
	Delete(id string) error
}

type characterRepository struct{}

func NewCharacterRepository() CharacterRepository {
	return &characterRepository{}
}

func (r *characterRepository) FindByID(id string) (*models.Character, error) {
	var character models.Character
	err := database.DB.Where("id = ?", id).First(&character).Error
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *characterRepository) FindByUserID(userID string) ([]*models.Character, error) {
	var characters []*models.Character
	err := database.DB.Where("user_id = ?", userID).Find(&characters).Error
	return characters, err
}

func (r *characterRepository) List(limit, offset int) ([]*models.Character, error) {
	var characters []*models.Character
	err := database.DB.Limit(limit).Offset(offset).Find(&characters).Error
	return characters, err
}

func (r *characterRepository) Create(character *models.Character) error {
	return database.DB.Create(character).Error
}

func (r *characterRepository) Update(character *models.Character) error {
	return database.DB.Save(character).Error
}

func (r *characterRepository) Delete(id string) error {
	return database.DB.Delete(&models.Character{}, "id = ?", id).Error
}
