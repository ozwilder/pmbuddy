package agent

import (
"fmt"
"time"

"github.com/ozwilder/pmbuddy/internal/knowledge"
)

// PMAdvisor is the core AI advisor for product management
type PMAdvisor struct {
SessionID      string
CreatedAt      time.Time
ConversationID string
Context        map[string]interface{}
Memory         []Message
KnowledgeBase  *knowledge.KnowledgeBase
LLM            LLMProvider
}

// Message represents a single conversation message
type Message struct {
Role      string                 `json:"role"` // "user" or "assistant"
Content   string                 `json:"content"`
Timestamp time.Time              `json:"timestamp"`
Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// NewPMAdvisor creates a new PM advisor instance
func NewPMAdvisor() *PMAdvisor {
sessionID := generateSessionID()

// Initialize LLM client - try Ollama first
var llmProvider LLMProvider
ollamaClient := NewOllamaClient("http://localhost:11434", "mistral")
if ollamaClient.IsAvailable() {
llmProvider = ollamaClient
} else {
// Fallback to mock if Ollama not available
llmProvider = NewMockLLMClient()
}

return &PMAdvisor{
SessionID:      sessionID,
CreatedAt:      time.Now(),
ConversationID: generateConversationID(),
Context:        make(map[string]interface{}),
Memory:         []Message{},
KnowledgeBase:  knowledge.NewKnowledgeBase(sessionID),
LLM:            llmProvider,
}
}

// GetWelcomeMessage returns the initial welcome message
func (pm *PMAdvisor) GetWelcomeMessage() string {
llmStatus := "🤖 " + pm.LLM.GetName()
if _, isOllama := pm.LLM.(*OllamaClient); isOllama && !pm.LLM.IsAvailable() {
llmStatus = "⚠️  Ollama not found - using mock responses\n   Install Ollama from https://ollama.ai to get real AI responses"
}

return fmt.Sprintf(`Welcome to PMBuddy! I'm your AI-powered Product Manager advisor.
I can help you with:
  • Product strategy and vision
  • Feature prioritization and roadmapping
  • User research and discovery
  • Metrics and KPIs
  • Go-to-market planning
  • Team dynamics and leadership

I learn from our conversations and improve over time.

%s

What would you like to discuss?`, llmStatus)
}

// ProcessQuery handles incoming questions/requests
func (pm *PMAdvisor) ProcessQuery(query string) (string, error) {
// Extract learnings from the query
learnings := pm.KnowledgeBase.ExtractLearnings(query, 0.8)
if len(learnings) > 0 {
pm.Context["last_learnings"] = learnings
}

// Detect preferences from the query
prefs := pm.KnowledgeBase.DetectPreferences(query)
if len(prefs) > 0 {
pm.Context["last_preferences"] = prefs
}

// Add user message to memory
pm.Memory = append(pm.Memory, Message{
Role:      "user",
Content:   query,
Timestamp: time.Now(),
Metadata: map[string]interface{}{
"learnings_count":   len(learnings),
"preferences_count": len(prefs),
},
})

// Learn from conversation
contextStr := fmt.Sprintf("%v", pm.Context)
pm.KnowledgeBase.LearnFromConversation(query, contextStr)

// Generate response using LLM
response, err := pm.generateResponseWithLLM(query)
if err != nil {
// Fallback to basic response on error
response = pm.generateFallbackResponse(query)
}

// Add assistant response to memory
pm.Memory = append(pm.Memory, Message{
Role:      "assistant",
Content:   response,
Timestamp: time.Now(),
})

return response, nil
}

// generateResponseWithLLM uses the LLM to generate a response
func (pm *PMAdvisor) generateResponseWithLLM(query string) (string, error) {
// Build system context based on preferences
profile := pm.KnowledgeBase.GetPreferenceProfile()
contextDirective := pm.buildContextDirective(profile)

// Get LLM response
response, err := pm.LLM.GenerateResponse(contextDirective+"\n\nUser: "+query, pm.Memory)
if err != nil {
return "", err
}

// Append learning context
topics := pm.KnowledgeBase.GetTopics()
if len(topics) > 0 {
response += fmt.Sprintf("\n\n_[Learning: topics discussed - %v]_", topics)
}

if profile.Confidence > 0.3 {
response += fmt.Sprintf("\n_[Adapted to your PM style: %s]_", profile.Summary())
}

return response, nil
}

// buildContextDirective creates a system prompt based on preferences
func (pm *PMAdvisor) buildContextDirective(profile knowledge.PreferenceProfile) string {
directive := "You are PMBuddy, an expert PM advisor."

// Customize based on detected preferences
if profile.Approach != "mixed" {
directive += fmt.Sprintf("\nThe user prefers a %s approach.", profile.Approach)
}

if profile.Thinking != "mixed" {
directive += fmt.Sprintf("\nThey think %s.", profile.Thinking)
}

if profile.Focus != "mixed" {
directive += fmt.Sprintf("\nThey prioritize %s focus.", profile.Focus)
}

directive += "\nProvide practical, actionable advice tailored to their style."

return directive
}

// generateFallbackResponse creates a response when LLM fails
func (pm *PMAdvisor) generateFallbackResponse(query string) string {
topics := pm.KnowledgeBase.GetTopics()
profile := pm.KnowledgeBase.GetPreferenceProfile()

response := fmt.Sprintf("I understand you're asking: '%s'\n\n", query)

// Add some basic guidance
if len(topics) > 0 {
response += fmt.Sprintf("Based on your question about %v, here are some key considerations:\n", topics)
response += "1. Understand your users and their needs\n"
response += "2. Define clear success metrics\n"
response += "3. Prioritize ruthlessly\n"
response += "4. Test and iterate\n"
response += "5. Measure and learn\n\n"
}

if profile.Confidence > 0.3 {
response += fmt.Sprintf("I'm adapting my advice to your style: %s\n", profile.Summary())
}

return response
}

// SaveSession persists the current session
func (pm *PMAdvisor) SaveSession() error {
// TODO: Implement session persistence
return nil
}

// LoadSession loads a previous session
func (pm *PMAdvisor) LoadSession(sessionID string) error {
// TODO: Implement session loading
return nil
}

// GetContext returns current context information
func (pm *PMAdvisor) GetContext() map[string]interface{} {
return pm.Context
}

// SetContext updates context for the advisor
func (pm *PMAdvisor) SetContext(key string, value interface{}) {
pm.Context[key] = value
}

// GetMemory returns conversation history
func (pm *PMAdvisor) GetMemory() []Message {
return pm.Memory
}

// ClearMemory clears the conversation history
func (pm *PMAdvisor) ClearMemory() {
pm.Memory = []Message{}
}

// GetKnowledgeBase returns the knowledge base
func (pm *PMAdvisor) GetKnowledgeBase() *knowledge.KnowledgeBase {
return pm.KnowledgeBase
}

// GetLLM returns the LLM provider
func (pm *PMAdvisor) GetLLM() LLMProvider {
return pm.LLM
}

// generateSessionID creates a unique session ID
func generateSessionID() string {
return fmt.Sprintf("pmb_%d", time.Now().Unix())
}

// generateConversationID creates a unique conversation ID
func generateConversationID() string {
return fmt.Sprintf("conv_%d", time.Now().UnixNano())
}
