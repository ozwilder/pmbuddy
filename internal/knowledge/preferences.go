package knowledge

import (
"fmt"
"regexp"
"strings"
"time"
)

// PMPreference represents a learned PM preference
type PMPreference struct {
Name        string      `json:"name"`
Category    string      `json:"category"`
Value       interface{} `json:"value"`
Confidence  float32     `json:"confidence"`
Occurrences int         `json:"occurrences"`
FirstSeen   time.Time   `json:"first_seen"`
LastUpdated time.Time   `json:"last_updated"`
}

// PreferenceProfile stores aggregated PM preferences
type PreferenceProfile struct {
Approach          string // "data-driven" or "intuitive"
Thinking          string // "top-down" or "bottom-up"
Style             string // "structured" or "agile"
DecisionFramework string // "consensus" or "decisive"
Focus             string // "business" or "users"
Timeline          string // "long-term" or "short-term"
Confidence        float32
}

// DetectPreferences analyzes text to extract PM preferences
func (kb *KnowledgeBase) DetectPreferences(text string) []PMPreference {
var preferences []PMPreference

// Define preference patterns
patterns := map[string]map[string][]string{
"approach": {
"data-driven": {
`(?i)(data|metric|number|analytics|kpi|measurement)`,
`(?i)(track|measure|quantify|benchmark|analyze)`,
},
"intuitive": {
`(?i)(feel|instinct|gut|sense|vibe|intuition)`,
`(?i)(think|believe|assume|hypothes)`,
},
},
"thinking": {
"top-down": {
`(?i)(vision|strategy|north star|mission|direction|roadmap first)`,
`(?i)(align|cascade|framework|structure)`,
},
"bottom-up": {
`(?i)(user|feedback|iterate|experiment|learn|discover)`,
`(?i)(agile|quick|test|validate|mvp)`,
},
},
"style": {
"structured": {
`(?i)(process|framework|system|formal|document|plan)`,
`(?i)(timeline|phase|gate|stage|sprint)`,
},
"agile": {
`(?i)(flexible|adapt|pivot|responsive|iterate)`,
`(?i)(rapid|continuous|feedback loop)`,
},
},
"decision": {
"consensus": {
`(?i)(team|together|discuss|involve|collaborate|stakeholder)`,
`(?i)(align|agree|consensus|inclusive)`,
},
"decisive": {
`(?i)(decision|decide|move fast|cut|reduce|scope)`,
`(?i)(priorit|focus|say no|tradeoff)`,
},
},
"focus": {
"business": {
`(?i)(revenue|growth|market|competitive|roi|margin|profit)`,
`(?i)(business model|monetiz|scale|acquisition)`,
},
"users": {
`(?i)(user|customer|experience|delight|satisfaction|nps|retention)`,
`(?i)(empathy|listen|need|solve problem)`,
},
},
"timeline": {
"long-term": {
`(?i)(vision|future|years?|roadmap|strategic|quarter|planning)`,
`(?i)(sustainable|build|foundational)`,
},
"short-term": {
`(?i)(sprint|quick|now|immediate|next week|launch|release)`,
`(?i)(urgent|time-sensitive|hot fix)`,
},
},
}

lowerText := strings.ToLower(text)

// Check each preference category
for category, options := range patterns {
for value, patternList := range options {
score := 0.0
matches := 0

for _, pattern := range patternList {
re := regexp.MustCompile(pattern)
if re.MatchString(lowerText) {
score++
matches++
}
}

if matches > 0 {
confidence := float32(float32(score) / float32(len(patternList)))

pref := PMPreference{
Name:        value,
Category:    category,
Value:       value,
Confidence:  confidence,
Occurrences: 1,
FirstSeen:   time.Now(),
LastUpdated: time.Now(),
}

preferences = append(preferences, pref)
}
}
}

// Store preferences
for _, pref := range preferences {
kb.updatePreferenceTracking(pref)
}

return preferences
}

// updatePreferenceTracking adds or updates preference tracking
func (kb *KnowledgeBase) updatePreferenceTracking(pref PMPreference) {
prefKey := fmt.Sprintf("%s_%s", pref.Category, pref.Name)

if existing, exists := kb.UserPreferences[prefKey]; exists {
if existingPref, ok := existing.(PMPreference); ok {
existingPref.Occurrences++
existingPref.LastUpdated = time.Now()
// Update confidence (weighted average)
existingPref.Confidence = (existingPref.Confidence + pref.Confidence) / 2
kb.UserPreferences[prefKey] = existingPref
}
} else {
kb.UserPreferences[prefKey] = pref
}
}

// GetPreferenceProfile synthesizes learnings into a profile
func (kb *KnowledgeBase) GetPreferenceProfile() PreferenceProfile {
profile := PreferenceProfile{
Approach:          "mixed",
Thinking:          "mixed",
Style:             "mixed",
DecisionFramework: "mixed",
Focus:             "mixed",
Timeline:          "mixed",
Confidence:        0.0,
}

categories := map[string]*string{
"approach":  &profile.Approach,
"thinking":  &profile.Thinking,
"style":     &profile.Style,
"decision":  &profile.DecisionFramework,
"focus":     &profile.Focus,
"timeline":  &profile.Timeline,
}

totalConfidence := 0.0
count := 0

for category, target := range categories {
var highestPref string
var highestConf float32 = 0

for _, val := range kb.UserPreferences {
if pref, ok := val.(PMPreference); ok && pref.Category == category {
if pref.Confidence > highestConf {
highestConf = pref.Confidence
highestPref = pref.Name
}
}
}

if highestPref != "" && highestConf > 0.3 {
*target = highestPref
totalConfidence += float64(highestConf)
count++
}
}

if count > 0 {
profile.Confidence = float32(totalConfidence / float64(count))
}

return profile
}

// ProfileSummary returns a human-readable preference summary
func (profile PreferenceProfile) Summary() string {
var parts []string

if profile.Approach != "mixed" {
parts = append(parts, fmt.Sprintf("**Approach**: %s", profile.Approach))
}
if profile.Thinking != "mixed" {
parts = append(parts, fmt.Sprintf("**Thinking**: %s", profile.Thinking))
}
if profile.Style != "mixed" {
parts = append(parts, fmt.Sprintf("**Style**: %s", profile.Style))
}
if profile.DecisionFramework != "mixed" {
parts = append(parts, fmt.Sprintf("**Decisions**: %s", profile.DecisionFramework))
}
if profile.Focus != "mixed" {
parts = append(parts, fmt.Sprintf("**Focus**: %s", profile.Focus))
}
if profile.Timeline != "mixed" {
parts = append(parts, fmt.Sprintf("**Timeline**: %s", profile.Timeline))
}

if len(parts) == 0 {
return "Still learning your preferences..."
}

return strings.Join(parts, " | ")
}

// GetPreferences returns all tracked preferences
func (kb *KnowledgeBase) GetPreferences() map[string]PMPreference {
result := make(map[string]PMPreference)

for _, val := range kb.UserPreferences {
if pref, ok := val.(PMPreference); ok {
result[pref.Category+"_"+pref.Name] = pref
}
}

return result
}

// LearnFromConversation extracts preferences from a full conversation
func (kb *KnowledgeBase) LearnFromConversation(userMessage string, context string) {
// Combine message and context
fullText := userMessage + " " + context

// Detect and track preferences
prefs := kb.DetectPreferences(fullText)

if len(prefs) > 0 {
// Generate insight about discovered preferences
profile := kb.GetPreferenceProfile()
insight := fmt.Sprintf("PM Style: %s (%.0f%% confidence)", profile.Summary(), profile.Confidence*100)
kb.AddInsight(insight)
}
}
