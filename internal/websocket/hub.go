package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Client WebSocket 客户端
type Client struct {
	ID     string
	Conn   *websocket.Conn
	Send   chan []byte
	UserID string
	ChatID string
}

// Hub 维护活跃客户端集合
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.RWMutex
}

var hub *Hub

// InitWebSocket 初始化 WebSocket Hub
func InitWebSocket() {
	hub = &Hub{
		broadcast:  make(chan []byte),
		register:  make(chan *Client),
		unregister: make(chan *Client),
		clients:  make(map[*Client]bool),
	}
	go hub.Run()
}

// GetHub 获取 Hub 实例
func GetHub() *Hub {
	return hub
}

// Run 运行 Hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("Client registered: %s", client.ID)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
			h.mu.Unlock()
			log.Printf("Client unregistered: %s", client.ID)

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.Send <- message:
					default:
						close(client.Send)
						delete(h.clients, client)
					}
				}
			h.mu.RUnlock()
		}
	}
}

// BroadcastMessage 广播消息给所有客户端
func (h *Hub) BroadcastMessage(event string, data interface{}) {
	message := map[string]interface{}{
		"event": event,
		"data":  data,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}

	h.broadcast <- bytes
}

// BroadcastToChat 广播消息给特定聊天室的客户端
func (h *Hub) BroadcastToChat(chatID string, event string, data interface{}) {
	message := map[string]interface{}{
		"event": event,
		"data":  data,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	for client := range h.clients {
		if client.ChatID == chatID {
			select {
			case client.Send <- bytes:
			default:
				close(client.Send)
				delete(h.clients, client)
			}
		}
	}
}
