package services

import (
	"errors"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/internal/repository"
)

type TokenService struct {
	userRepo      repository.UserRepository
	transactionRepo repository.TransactionRepository
}

func NewTokenService(userRepo repository.UserRepository, transactionRepo repository.TransactionRepository) *TokenService {
	return &TokenService{
		userRepo:      userRepo,
		transactionRepo: transactionRepo,
	}
}

// ConsumeTokens 消耗 Tokens
func (s *TokenService) ConsumeTokens(userID string, amount int) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if user.TokenBalance < amount {
		return errors.New("insufficient token balance")
	}

	// 扣除 Tokens
	user.TokenBalance -= amount
	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	// 记录交易
	transaction := &models.Transaction{
		UserID:      userID,
		Type:        "purchase",
		Amount:      -amount,
		Description: "AI message generation",
	}
	return s.transactionRepo.Create(transaction)
}

// AddTokens 添加 Tokens
func (s *TokenService) AddTokens(userID string, amount int, description string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	// 添加 Tokens
	user.TokenBalance += amount
	if err := s.userRepo.Update(user); err != nil {
		return err
	}

	// 记录交易
	transaction := &models.Transaction{
		UserID:      userID,
		Type:        "topup",
		Amount:      amount,
		Description: description,
	}
	return s.transactionRepo.Create(transaction)
}

// GetBalance 获取 Token 余额
func (s *TokenService) GetBalance(userID string) (int, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return 0, err
	}

	return user.TokenBalance, nil
}
