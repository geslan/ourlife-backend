package main

import (
	"fmt"

	"github.com/geslan/ourlife-backend/internal/models"
	"github.com/geslan/ourlife-backend/internal/repository"
	"github.com/geslan/ourlife-backend/internal/services"
	"github.com/geslan/ourlife-backend/pkg/database"
)

func main() {
	// 初始化数据库
	fmt.Println("🔌 Connecting to database...")
	if err := database.Connect("host=localhost user=postgres password=postgres dbname=ourlife port=5432 sslmode=disable"); err != nil {
		fmt.Printf("❌ Failed to connect to database: %v\n", err)
		return
	}
	fmt.Println("✅ Database connected")
	fmt.Println()

	// 自动迁移
	fmt.Println("🔄 Running auto-migration...")
	if err := database.AutoMigrate(); err != nil {
		fmt.Printf("❌ Failed to migrate database: %v\n", err)
		return
	}
	fmt.Println("✅ Database migrated")
	fmt.Println()

	fmt.Println("🧪 Token Service Test")
	fmt.Println()

	// 初始化 Repositories
	userRepo := repository.NewUserRepository()
	transactionRepo := repository.NewTransactionRepository()

	// 初始化 TokenService
	tokenService := services.NewTokenService(userRepo, transactionRepo)

	// 使用现有的测试用户
	userID := "36ca43d4-e360-4bae-bd42-44b7ed19adb0"
	fmt.Printf("1️⃣ Using test user: %s\n", userID)

	_, err := userRepo.FindByID(userID)
	if err != nil {
		fmt.Printf("❌ Failed to find test user: %v\n", err)
		return
	}
	fmt.Println("✅ Test user found")
	fmt.Println()

	// 测试 1: 获取 Token 余额
	fmt.Println("2️⃣ Test 1: Get Token Balance")
	balance, err := tokenService.GetBalance(userID)
	if err != nil {
		fmt.Printf("❌ Failed to get balance: %v\n", err)
	} else {
		fmt.Printf("✅ Current balance: %d tokens\n", balance)
	}
	fmt.Println()

	// 测试 2: 添加 Tokens
	fmt.Println("3️⃣ Test 2: Add Tokens (+50)")
	err = tokenService.AddTokens(userID, 50, "Test topup")
	if err != nil {
		fmt.Printf("❌ Failed to add tokens: %v\n", err)
	} else {
		fmt.Println("✅ Tokens added successfully")
	}
	fmt.Println()

	// 验证余额
	balance, _ = tokenService.GetBalance(userID)
	fmt.Printf("✅ New balance: %d tokens\n", balance)
	fmt.Println()

	// 测试 3: 消耗 Tokens
	fmt.Println("4️⃣ Test 3: Consume Tokens (-30)")
	err = tokenService.ConsumeTokens(userID, 30)
	if err != nil {
		fmt.Printf("❌ Failed to consume tokens: %v\n", err)
	} else {
		fmt.Println("✅ Tokens consumed successfully")
	}
	fmt.Println()

	// 验证余额
	balance, _ = tokenService.GetBalance(userID)
	fmt.Printf("✅ New balance: %d tokens\n", balance)
	fmt.Println()

	// 测试 4: 尝试消耗不足的 Tokens
	fmt.Println("5️⃣ Test 4: Consume Insufficient Tokens (-200)")
	err = tokenService.ConsumeTokens(userID, 200)
	if err != nil {
		fmt.Printf("✅ Expected error: %v\n", err)
	} else {
		fmt.Println("❌ Should have failed but didn't")
	}
	fmt.Println()

	// 最终余额
	balance, _ = tokenService.GetBalance(userID)
	fmt.Printf("6️⃣ Final balance: %d tokens\n", balance)
	fmt.Println()

	// 查看交易记录
	fmt.Println("7️⃣ Test 5: Get Transaction History")
	transactions, err := transactionRepo.FindByUserID(userID, 10, 0)
	if err != nil {
		fmt.Printf("❌ Failed to get transactions: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d transactions\n", len(transactions))
		for i, tx := range transactions {
			fmt.Printf("   %d. Type: %s, Amount: %d, Description: %s\n",
				i+1, tx.Type, tx.Amount, tx.Description)
		}
	}
	fmt.Println()

	fmt.Println("🎉 Token Service Test Complete!")
}
