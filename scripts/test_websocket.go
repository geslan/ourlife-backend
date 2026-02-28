package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("🧪 Starting WebSocket Integration Test...")
	fmt.Println()

	// 1. 获取 JWT Token
	fmt.Println("1️⃣ Getting JWT Token...")
	token, err := getJWTToken()
	if err != nil {
		fmt.Printf("❌ Failed to get token: %v\n", err)
		return
	}
	fmt.Printf("✅ Token obtained: %s...\n", token[:20])
	fmt.Println()

	// 2. 建立 WebSocket 连接
	fmt.Println("2️⃣ Connecting to WebSocket...")
	conn, err := connectWebSocket(token, "test-chat-1")
	if err != nil {
		fmt.Printf("❌ Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("✅ Connected to WebSocket")
	fmt.Println()

	// 3. 设置在线状态
	fmt.Println("3️⃣ Setting user online...")
	err = setOnlineStatus(token)
	if err != nil {
		fmt.Printf("❌ Failed to set online: %v\n", err)
	} else {
		fmt.Println("✅ User set as online")
	}
	fmt.Println()

	// 4. 发送测试消息
	fmt.Println("4️⃣ Sending test message...")
	err = sendTestMessage(conn)
	if err != nil {
		fmt.Printf("❌ Failed to send message: %v\n", err)
	} else {
		fmt.Println("✅ Test message sent")
	}
	fmt.Println()

	// 5. 接收消息
	fmt.Println("5️⃣ Waiting for messages (5 seconds)...")
	done := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("❌ Failed to read message: %v\n", err)
				break
			}
			fmt.Printf("📩 Received message: %s\n", string(message))
		}
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("✅ Message reception test complete")
	case <-time.After(5 * time.Second):
		fmt.Println("⏱️ Timeout waiting for messages")
	}
	fmt.Println()

	// 6. 设置离线状态
	fmt.Println("6️⃣ Setting user offline...")
	err = setOfflineStatus(token)
	if err != nil {
		fmt.Printf("❌ Failed to set offline: %v\n", err)
	} else {
		fmt.Println("✅ User set as offline")
	}
	fmt.Println()

	fmt.Println("🎉 WebSocket Integration Test Complete!")
}

func getJWTToken() (string, error) {
	payload := map[string]interface{}{
		"telegramId": int64(123456789),
		"username":   "testuser",
		"name":       "Test User",
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:8080/api/auth/telegram-webapp",
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	token, ok := result["token"].(string)
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}

	return token, nil
}

func connectWebSocket(token, chatID string) (*websocket.Conn, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws/chat",
	}

	q := u.Query()
	q.Set("token", token)
	q.Set("chatId", chatID)
	u.RawQuery = q.Encode()

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func setOnlineStatus(token string) error {
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/online/set-online", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}

func setOfflineStatus(token string) error {
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/online/set-offline", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return nil
}

func sendTestMessage(conn *websocket.Conn) error {
	message := map[string]interface{}{
		"type":    "test",
		"content": "Hello from WebSocket test!",
	}

	bytes, _ := json.Marshal(message)
	return conn.WriteMessage(websocket.TextMessage, bytes)
}
