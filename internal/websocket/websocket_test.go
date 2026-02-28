package websocket_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWebSocketConnection(t *testing.T) {
	// 获取 JWT token
	resp, err := http.Post("http://localhost:8080/api/auth/telegram-webapp",
		"application/json",
		nil,
	)
	if err != nil {
		t.Fatalf("Failed to get token: %v", err)
	}
	defer resp.BodyClose()

	var tokenResp map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&tokenResp)

	token, ok := tokenResp["token"].(string)
	if !ok {
		t.Fatal("Failed to extract token from response")
	}

	// 建立 WebSocket 连接
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws/chat",
	}
	q := u.Query()
	q.Set("token", token)
	q.Set("chatId", "test-chat-1")
	u.RawQuery = q.Encode()

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer conn.Close()

	// 测试发送消息
	message := map[string]interface{}{
		"type":    "test",
		"content": "test message",
	}

	bytes, _ := json.Marshal(message)
	err = conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	// 测试接收消息
	_, p, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}

	if p != nil {
		t.Logf("Received message: %s", string(p))
	}
}
