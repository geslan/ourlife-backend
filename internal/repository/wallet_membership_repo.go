package repository

import (
	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/pkg/database"
)

type WalletRepository interface {
	GetBalance(userID string) (int, error)
	AddBalance(userID string, amount int) error
	DeductBalance(userID string, amount int) error
}

type walletRepository struct{}

func NewWalletRepository() WalletRepository {
	return &walletRepository{}
}

func (r *walletRepository) GetBalance(userID string) (int, error) {
	var user models.User
	err := database.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.TokenBalance, nil
}

func (r *walletRepository) AddBalance(userID string, amount int) error {
	return database.DB.Model(&models.User{}).
		Where("id = ?", userID).
		UpdateColumn("token_balance", database.DB.Raw("token_balance + ?", amount)).Error
}

func (r *walletRepository) DeductBalance(userID string, amount int) error {
	return database.DB.Model(&models.User{}).
		Where("id = ? AND token_balance >= ?", userID, amount).
		UpdateColumn("token_balance", database.DB.Raw("token_balance - ?", amount)).Error
}

type MembershipRepository interface {
	GetStatus(userID string) (map[string]interface{}, error)
}

type membershipRepository struct{}

func NewMembershipRepository() MembershipRepository {
	return &membershipRepository{}
}

func (r *membershipRepository) GetStatus(userID string) (map[string]interface{}, error) {
	var user models.User
	err := database.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	status := map[string]interface{}{
		"membership": user.Membership,
		"isPremium":  user.Membership == string(models.RolePremium),
	}

	return status, nil
}
