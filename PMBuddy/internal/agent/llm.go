package agent

import (
"bytes"
"encoding/json"
"fmt"
"io"
"net/http"
"strings"
"time"
)

// LLMProvider defines the interface for LLM services
type LLMProvider interface {
GenerateResponse(prompt string, context []Message) (string, error)
IsAvailable() bool
GetName() string
}

// OllamaClient handles communication with Ollama local LLM
type OllamaClient struct {
BaseURL      string
Model        string
Temperature  float32
MaxTokens    int
IsConnected  bool
lastCheckTime time.Time
}

// OllamaRequest represents a request to Ollama
type OllamaRequest struct {
Model  string `json:"model"`
Prompt string `json:"prompt"`
Stream bool   `json:"stream"`
}

// OllamaResponse represents a response from Ollama
type OllamaResponse struct {
Model     string `json:"model"`
CreatedAt string `json:"created_at"`
Response  string `json:"response"`
Done      bool   `json:"done"`
}

// NewOllamaClient creates a new Ollama LLM client
func NewOllamaClient(baseURL, model string) *OllamaClient {
if baseURL == "" {
baseURL = "http://localhost:11434"
}
if model == "" {
model = "mistral" // Lightweight, fast model
}

client := &OllamaClient{
BaseURL:      baseURL,
Model:        model,
Temperature:  0.7,
MaxTokens:    2048,
lastCheckTime: time.Now(),
}

// Check connectivity
client.IsConnected = client.checkConnection()

return client
}

// checkConnection verifies Ollama is running
func (oc *OllamaClient) checkConnection() bool {
// Only check every 30 seconds to avoid spam
if time.Since(oc.lastCheckTime) < 30*time.Second && oc.IsConnected {
return true
}

oc.lastCheckTime = time.Now()

client := &http.Client{Timeout: 2 * time.Second}
resp, err := client.Get(oc.BaseURL + "/api/tags")
if err != nil {
return false
}
defer resp.Body.Close()

return resp.StatusCode == 200
}

// IsAvailable checks if Ollama is available
func (oc *OllamaClient) IsAvailable() bool {
if !oc.IsConnected {
oc.IsConnected = oc.checkConnection()
}
return oc.IsConnected
}

// GetName returns the LLM provider name
func (oc *OllamaClient) GetName() string {
return fmt.Sprintf("Ollama (%s)", oc.Model)
}

// GenerateResponse creates an AI response using Ollama
func (oc *OllamaClient) GenerateResponse(prompt string, context []Message) (string, error) {
if !oc.IsAvailable() {
return "", fmt.Errorf("Ollama is not available at %s. Please install and start Ollama", oc.BaseURL)
}

// Build context from conversation
contextStr := oc.buildContext(context)

// Combine context and prompt
fullPrompt := fmt.Sprintf(`%s

User Query: %s

Please provide a helpful PM advisor response:`, contextStr, prompt)

// Make request to Ollama
reqBody := OllamaRequest{
Model:  oc.Model,
Prompt: fullPrompt,
Stream: false,
}

jsonBody, err := json.Marshal(reqBody)
if err != nil {
return "", fmt.Errorf("failed to marshal request: %v", err)
}

client := &http.Client{Timeout: 60 * time.Second}
resp, err := client.Post(
oc.BaseURL+"/api/generate",
"application/json",
bytes.NewReader(jsonBody),
)
if err != nil {
return "", fmt.Errorf("failed to call Ollama: %v", err)
}
defer resp.Body.Close()

if resp.StatusCode != 200 {
body, _ := io.ReadAll(resp.Body)
return "", fmt.Errorf("Ollama returned status %d: %s", resp.StatusCode, string(body))
}

// Parse response
var ollamaResp OllamaResponse
if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
return "", fmt.Errorf("failed to parse response: %v", err)
}

// Clean up response
response := strings.TrimSpace(ollamaResp.Response)
return response, nil
}

// buildContext creates context from conversation history
func (oc *OllamaClient) buildContext(messages []Message) string {
var context []string

// Add PM context
context = append(context, "You are PMBuddy, an expert Product Manager advisor.")
context = append(context, "Provide practical, actionable advice based on PM best practices.")

// Add recent conversation history
if len(messages) > 2 {
context = append(context, "\nRecent conversation:")
start := len(messages) - 4
if start < 0 {
start = 0
}

for i := start; i < len(messages); i++ {
msg := messages[i]
if msg.Role == "user" {
context = append(context, fmt.Sprintf("User: %s", msg.Content))
} else {
context = append(context, fmt.Sprintf("Assistant: %s", msg.Content))
}
}
}

return strings.Join(context, "\n")
}

// MockLLMClient provides mock responses for testing (when LLM unavailable)
type MockLLMClient struct {
Name string
}

// NewMockLLMClient creates a mock LLM for testing
func NewMockLLMClient() *MockLLMClient {
return &MockLLMClient{Name: "Mock LLM (Testing)"}
}

// IsAvailable always returns true for mock
func (mc *MockLLMClient) IsAvailable() bool {
return true
}

// GetName returns mock name
func (mc *MockLLMClient) GetName() string {
return mc.Name
}

// GenerateResponse returns a template response
func (mc *MockLLMClient) GenerateResponse(prompt string, context []Message) (string, error) {
responses := []string{
"That's a great question! Here's my perspective:\n\n1. Start by understanding your users and their needs\n2. Define clear success metrics\n3. Prioritize ruthlessly based on impact and effort\n4. Iterate quickly and measure results\n5. Adapt based on feedback\n\nWhat specific aspect would you like to dive deeper on?",
"Interesting challenge! Consider these approaches:\n\n• Data-driven: Analyze metrics and user behavior\n• User-centric: Focus on solving real problems\n• Strategic: Align with broader product vision\n• Agile: Test hypotheses in short cycles\n\nWhich resonates most with your situation?",
"Great PM thinking! Here's my framework:\n\n1. Validate the problem exists\n2. Define success criteria upfront\n3. Get stakeholder buy-in early\n4. Build minimum viable solution\n5. Measure, learn, iterate\n\nWhere are you in this process?",
}

// Hash the prompt to get consistent mock responses
hash := 0
for _, char := range prompt {
hash = hash*31 + int(char)
}

idx := hash % len(responses)
if idx < 0 {
idx = 0
}

return responses[idx], nil
}
