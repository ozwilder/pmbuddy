package knowledge

import (
"fmt"
"regexp"
"strings"
"time"
)

// KnowledgeBase manages accumulated PM knowledge and learning
type KnowledgeBase struct {
SessionID        string
CreatedAt        time.Time
LastUpdated      time.Time
LearningHistory  []LearningEntry
ContextMappings  map[string][]string // Maps concepts to related knowledge
UserPreferences  map[string]interface{}
KeyInsights      []string
PMTopics         map[string]int // Tracks topic frequency
}

// LearningEntry represents a single learning from conversation
type LearningEntry struct {
ID         string                 `json:"id"`
Topic      string                 `json:"topic"`
Content    string                 `json:"content"`
Timestamp  time.Time              `json:"timestamp"`
Confidence float32                `json:"confidence"`
Metadata   map[string]interface{} `json:"metadata,omitempty"`
Source     string                 `json:"source"` // conversation_id, user_input, etc
}

// PMConcept represents a product management concept with context
type PMConcept struct {
Name        string
Category    string
Description string
RelatedTo   []string
}

// NewKnowledgeBase creates a new knowledge base
func NewKnowledgeBase(sessionID string) *KnowledgeBase {
kb := &KnowledgeBase{
SessionID:       sessionID,
CreatedAt:       time.Now(),
LastUpdated:     time.Now(),
LearningHistory: []LearningEntry{},
ContextMappings: make(map[string][]string),
UserPreferences: make(map[string]interface{}),
KeyInsights:     []string{},
PMTopics:        make(map[string]int),
}
kb.initializePMConcepts()
return kb
}

// initializePMConcepts sets up core PM concept mappings
func (kb *KnowledgeBase) initializePMConcepts() {
concepts := map[string][]string{
"prioritization": {"roadmap", "strategy", "features", "requirements"},
"strategy":       {"vision", "roadmap", "prioritization", "go-to-market"},
"user-research":  {"discovery", "testing", "metrics", "feedback"},
"roadmap":        {"strategy", "prioritization", "timeline", "resources"},
"metrics":        {"KPIs", "success", "analytics", "user-research"},
"go-to-market":   {"strategy", "positioning", "launch", "messaging"},
"features":       {"prioritization", "requirements", "roadmap", "users"},
}

for concept, related := range concepts {
kb.ContextMappings[concept] = related
kb.PMTopics[concept] = 0
}
}

// ExtractLearnings analyzes text and extracts PM learnings
func (kb *KnowledgeBase) ExtractLearnings(text string, confidence float32) []LearningEntry {
var learnings []LearningEntry

// Convert to lowercase for pattern matching
lowerText := strings.ToLower(text)

// Define PM concept patterns
patterns := map[string]string{
"prioritization": `(?i)(priorit|rank|order|defer|must-have|nice-to-have|mvp)`,
"strategy":       `(?i)(strateg|vision|north star|mission|direction|plan)`,
"roadmap":        `(?i)(roadmap|timeline|quarters?|phases?|milestones?)`,
"metrics":        `(?i)(metrics?|kpi|success|oas|engagement|retention)`,
"user-research":  `(?i)(user research|discovery|interviews|testing|feedback)`,
"go-to-market":   `(?i)(launch|gtm|positioning|messaging|target market)`,
"features":       `(?i)(features?|requirements|capabilities|functionality)`,
}

// Check for each pattern
for topic, pattern := range patterns {
re := regexp.MustCompile(pattern)
if re.MatchString(lowerText) {
// Extract context around the match
context := extractContext(lowerText, topic, 50)

entry := LearningEntry{
ID:         generateLearningID(),
Topic:      topic,
Content:    context,
Timestamp:  time.Now(),
Confidence: confidence,
Source:     "conversation",
Metadata: map[string]interface{}{
"word_count":    len(strings.Fields(text)),
"has_question":  strings.Contains(text, "?"),
"text_length":   len(text),
},
}

learnings = append(learnings, entry)
kb.PMTopics[topic]++
}
}

// Store learnings
kb.LearningHistory = append(kb.LearningHistory, learnings...)
kb.LastUpdated = time.Now()

return learnings
}

// extractContext pulls relevant context around a concept
func extractContext(text string, keyword string, maxLen int) string {
// Find first occurrence of keyword-related content
re := regexp.MustCompile(`(?i)` + keyword)
loc := re.FindStringIndex(text)
if loc == nil {
return text[:min(maxLen, len(text))]
}

start := max(0, loc[0]-maxLen/2)
end := min(len(text), loc[1]+maxLen/2)

return strings.TrimSpace(text[start:end])
}

// LearnUserPreference records a PM preference
func (kb *KnowledgeBase) LearnUserPreference(prefType string, value interface{}) {
kb.UserPreferences[prefType] = value
kb.LastUpdated = time.Now()

// Auto-generate insight about preference
insight := fmt.Sprintf("User prefers %s: %v", prefType, value)
kb.AddInsight(insight)
}

// AddLearning records a new learning
func (kb *KnowledgeBase) AddLearning(topic, content string, confidence float32, metadata map[string]interface{}) LearningEntry {
entry := LearningEntry{
ID:         generateLearningID(),
Topic:      topic,
Content:    content,
Timestamp:  time.Now(),
Confidence: confidence,
Metadata:   metadata,
Source:     "manual",
}

kb.LearningHistory = append(kb.LearningHistory, entry)
kb.PMTopics[topic]++
kb.LastUpdated = time.Now()

return entry
}

// GetRelatedKnowledge retrieves knowledge related to a concept
func (kb *KnowledgeBase) GetRelatedKnowledge(concept string) []LearningEntry {
var related []LearningEntry

// Get related topics from mapping
if topics, exists := kb.ContextMappings[concept]; exists {
for _, topic := range topics {
for _, entry := range kb.LearningHistory {
if entry.Topic == topic && entry.Confidence > 0.5 {
related = append(related, entry)
}
}
}
}

// Also get direct topic matches
for _, entry := range kb.LearningHistory {
if entry.Topic == concept && entry.Confidence > 0.5 {
related = append(related, entry)
}
}

return related
}

// GetTopics returns PM topics discussed sorted by frequency
func (kb *KnowledgeBase) GetTopics() []string {
type topicCount struct {
topic string
count int
}

var topics []topicCount
for topic, count := range kb.PMTopics {
if count > 0 {
topics = append(topics, topicCount{topic, count})
}
}

// Simple sort (highest first)
for i := 0; i < len(topics); i++ {
for j := i + 1; j < len(topics); j++ {
if topics[j].count > topics[i].count {
topics[i], topics[j] = topics[j], topics[i]
}
}
}

var result []string
for _, tc := range topics {
result = append(result, tc.topic)
}
return result
}

// GenerateInsights creates insights from learning patterns
func (kb *KnowledgeBase) GenerateInsights() []string {
var insights []string

topics := kb.GetTopics()
if len(topics) > 0 {
topicStr := strings.Join(topics[:min(3, len(topics))], ", ")
insights = append(insights, fmt.Sprintf("Primary focus areas: %s", topicStr))
}

// Check for preference patterns
if len(kb.UserPreferences) > 0 {
insights = append(insights, fmt.Sprintf("Tracked %d user preferences", len(kb.UserPreferences)))
}

// Calculate confidence average
if len(kb.LearningHistory) > 0 {
var total float32 = 0
for _, entry := range kb.LearningHistory {
total += entry.Confidence
}
avgConf := total / float32(len(kb.LearningHistory))
insights = append(insights, fmt.Sprintf("Learning confidence: %.1f%%", avgConf*100))
}

return insights
}

// UpdatePreference saves user preference
func (kb *KnowledgeBase) UpdatePreference(key string, value interface{}) {
kb.UserPreferences[key] = value
kb.LastUpdated = time.Now()
}

// GetPreference retrieves a user preference
func (kb *KnowledgeBase) GetPreference(key string) (interface{}, bool) {
value, exists := kb.UserPreferences[key]
return value, exists
}

// AddInsight records a key insight
func (kb *KnowledgeBase) AddInsight(insight string) {
kb.KeyInsights = append(kb.KeyInsights, insight)
kb.LastUpdated = time.Now()
}

// GetInsights returns all recorded insights
func (kb *KnowledgeBase) GetInsights() []string {
return kb.KeyInsights
}

// GetLearningStats returns statistics about learning
func (kb *KnowledgeBase) GetLearningStats() map[string]interface{} {
topicCounts := make(map[string]int)
avgConfidence := 0.0

for _, entry := range kb.LearningHistory {
topicCounts[entry.Topic]++
avgConfidence += float64(entry.Confidence)
}

if len(kb.LearningHistory) > 0 {
avgConfidence /= float64(len(kb.LearningHistory))
}

return map[string]interface{}{
"total_learnings":      len(kb.LearningHistory),
"unique_topics":        len(topicCounts),
"topic_distribution":   topicCounts,
"avg_confidence":       avgConfidence,
"key_insights":         len(kb.KeyInsights),
"user_preferences":     len(kb.UserPreferences),
"last_updated":         kb.LastUpdated,
}
}

// Helper functions
func min(a, b int) int {
if a < b {
return a
}
return b
}

func max(a, b int) int {
if a > b {
return a
}
return b
}

// generateLearningID creates a unique learning ID
func generateLearningID() string {
return "learn_" + time.Now().Format("20060102150405")
}
