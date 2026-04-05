package agent

import (
	"fmt"
	"strings"
	"time"

	"github.com/ozwilder/pmbuddy/internal/knowledge"
)

// PMAdvisor is the core AI advisor for product management
type PMAdvisor struct {
	SessionID       string
	CreatedAt       time.Time
	ConversationID  string
	Context         map[string]interface{}
	Memory          []Message
	KnowledgeBase   *knowledge.KnowledgeBase
	PMKnowledge     *knowledge.PMKnowledgeBase
	LLM             LLMProvider
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

	// Load PM knowledge base from ProductManagement folder
	pmKnowledge, err := knowledge.LoadPMKnowledge(".")
	if err != nil {
		// If PM knowledge fails to load, continue with empty (non-fatal)
		pmKnowledge = &knowledge.PMKnowledgeBase{
			Concepts:   make(map[string]*knowledge.PMConcept),
			Categories: make(map[string][]string),
			Content:    make(map[string]string),
			RelatedMap: make(map[string][]string),
		}
	}

	return &PMAdvisor{
		SessionID:      sessionID,
		CreatedAt:      time.Now(),
		ConversationID: generateConversationID(),
		Context:        make(map[string]interface{}),
		Memory:         []Message{},
		KnowledgeBase:  knowledge.NewKnowledgeBase(sessionID),
		PMKnowledge:    pmKnowledge,
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

// ValidateQuery checks if user input is valid, complete, and relevant to product management
func (pm *PMAdvisor) ValidateQuery(query string) (bool, string) {
	query = strings.TrimSpace(query)
	
	if len(query) < 5 {
		return false, "That's too brief - could you give me more details about what you need help with?"
	}
	
	incompletePatterns := []string{"how do i", "how to", "what is", "when should", "why", "where", "which"}
	lowerQuery := strings.ToLower(query)
	
	for _, pattern := range incompletePatterns {
		if strings.HasPrefix(lowerQuery, pattern) {
			afterPattern := strings.TrimSpace(strings.TrimPrefix(lowerQuery, pattern))
			if len(afterPattern) < 3 {
				return false, "I need more context - could you complete that thought?"
			}
		}
	}
	
	if strings.HasSuffix(query, "...") || strings.HasSuffix(query, "- ") {
		return false, "It looks like your message got cut off - could you finish what you were saying?"
	}
	
	pmDomainKeywords := map[string]bool{
		"product": true, "feature": true, "user": true, "market": true, "strategy": true,
		"roadmap": true, "priorit": true, "ux": true, "kpi": true, "metric": true,
		"research": true, "discover": true, "launch": true, "gtm": true, "customer": true,
	}
	
	nonPMPatterns := []string{
		"code", "debug", "bug", "fix", "recipe", "bake", "cake", "weather", "homework", "sports",
	}
	
	if strings.Contains(lowerQuery, "porn") || strings.Contains(lowerQuery, "nsfw") {
		return false, "I'm focused on product management topics. Could you ask me something related to product strategy or user needs?"
	}
	
	hasPMKeyword := false
	for keyword := range pmDomainKeywords {
		if strings.Contains(lowerQuery, keyword) {
			hasPMKeyword = true
			break
		}
	}
	
	if !hasPMKeyword {
		for _, pattern := range nonPMPatterns {
			if strings.Contains(lowerQuery, pattern) {
				return false, "I'm specialized in product management. What product-related topic would you like to discuss?"
			}
		}
	}
	
	return true, ""
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
	directive := `You are PMBuddy, an expert PM advisor helping with product tasks and thinking.

IMPORTANT STYLE RULES:
- Keep responses SHORT and conversational (2-3 sentences max)
- Use natural language like a native English speaker
- Be direct and practical, no fluff
- Only provide detailed output if the user explicitly asks for "detailed" or "technical"
- Sound like a colleague having a casual discussion, not a textbook

`

	// Customize based on detected preferences
	if profile.Approach != "mixed" {
		directive += fmt.Sprintf("The user prefers a %s approach. ", profile.Approach)
	}

	if profile.Thinking != "mixed" {
		directive += fmt.Sprintf("They think %s. ", profile.Thinking)
	}

	if profile.Focus != "mixed" {
		directive += fmt.Sprintf("They prioritize %s focus. ", profile.Focus)
	}

	directive += "\nRespond accordingly in a friendly, conversational tone."

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
