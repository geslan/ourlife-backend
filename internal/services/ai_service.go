package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AIService struct {
	baseURL    string
	httpClient *http.Client
}

func NewAIService(baseURL string) *AIService {
	return &AIService{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

type CharacterConfig struct {
	Name         string   `json:"name"`
	Personality  []string `json:"personality"`
	Relationship string   `json:"relationship"`
	Profession   string   `json:"profession"`
	Interests    []string `json:"interests"`
	Voice        string   `json:"voice"`
	Bio          string   `json:"bio"`
}

type ConversationContext struct {
	Conversation     []string        `json:"conversation"`
	CharacterConfig  CharacterConfig `json:"characterConfig"`
}

type GenerateRequest struct {
	Message      string             `json:"message"`
	CharacterID  string             `json:"characterId"`
	UserID       string             `json:"userId"`
	Context      ConversationContext `json:"context"`
}

type A2UIData struct {
	Surface string   `json:"surface"`
	Title   string   `json:"title"`
	Options []string `json:"options,omitempty"`
	Fields  []map[string]interface{} `json:"fields,omitempty"`
}

type GenerateResponse struct {
	Type     string    `json:"type"`
	Content  string    `json:"content"`
	A2UIData *A2UIData `json:"a2uiData,omitempty"`
}

func (s *AIService) GenerateResponse(req GenerateRequest) (*GenerateResponse, error) {
	url := fmt.Sprintf("%s/api/generate", s.baseURL)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("AI service returned status %d: %s", resp.StatusCode, string(body))
	}

	var response GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

func (s *AIService) HealthCheck() error {
	url := fmt.Sprintf("%s/health", s.baseURL)

	resp, err := s.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to check health: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("AI service unhealthy: status %d", resp.StatusCode)
	}

	return nil
}
