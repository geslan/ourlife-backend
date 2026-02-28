package repository

import (
	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type TransactionRepository interface {
	FindByID(id string) (*models.Transaction, error)
	FindByUserID(userID string, limit, offset int) ([]*models.Transaction, error)
	Create(transaction *models.Transaction) error
}

type transactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) FindByID(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *transactionRepository) FindByUserID(userID string, limit, offset int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := database.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) Create(transaction *models.Transaction) error {
	return database.DB.Create(transaction).Error
}
