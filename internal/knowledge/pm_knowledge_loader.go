package knowledge

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// PMKnowledgeBase represents the Product Management knowledge from markdown files
type PMKnowledgeBase struct {
	Concepts   map[string]*PMConcept // Name -> Concept mapping
	Categories map[string][]string   // Category -> Concept names
	Content    map[string]string     // Concept name -> full content
	RelatedMap map[string][]string   // Concept -> related concepts
}

// LoadPMKnowledge loads PM concepts from the ProductManagement folder
func LoadPMKnowledge(basePath string) (*PMKnowledgeBase, error) {
	kb := &PMKnowledgeBase{
		Concepts:   make(map[string]*PMConcept),
		Categories: make(map[string][]string),
		Content:    make(map[string]string),
		RelatedMap: make(map[string][]string),
	}

	// Construct path to ProductManagement folder
	pmPath := filepath.Join(basePath, "ProductManagement")

	// Check if folder exists
	if _, err := os.Stat(pmPath); os.IsNotExist(err) {
		return kb, fmt.Errorf("ProductManagement folder not found at: %s", pmPath)
	}

	// Read all markdown files
	files, err := ioutil.ReadDir(pmPath)
	if err != nil {
		return kb, fmt.Errorf("failed to read ProductManagement folder: %v", err)
	}

	// Load each markdown file
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			conceptName := strings.TrimSuffix(file.Name(), ".md")
			filePath := filepath.Join(pmPath, file.Name())

			// Read file content
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Warning: failed to read %s: %v\n", file.Name(), err)
				continue
			}

			contentStr := string(content)
			kb.Content[conceptName] = contentStr

			// Extract category from content (look for headers)
			category := extractCategory(contentStr)

			// Create PM concept
			concept := &PMConcept{
				Name:        conceptName,
				Category:    category,
				Description: extractDescription(contentStr),
				RelatedTo:   extractRelated(contentStr),
			}

			kb.Concepts[conceptName] = concept

			// Add to category mapping
			kb.Categories[category] = append(kb.Categories[category], conceptName)

			// Add to related map
			kb.RelatedMap[conceptName] = concept.RelatedTo
		}
	}

	return kb, nil
}

// extractCategory extracts the category from markdown content
func extractCategory(content string) string {
	lines := strings.Split(content, "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "#") {
		return strings.TrimSpace(strings.TrimPrefix(lines[0], "#"))
	}
	return "General"
}

// extractDescription extracts first meaningful line as description
func extractDescription(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "[") {
			if len(line) > 150 {
				return line[:150] + "..."
			}
			return line
		}
	}
	return ""
}

// extractRelated extracts related concepts (wiki-style [[links]])
func extractRelated(content string) []string {
	var related []string
	// Find all [[concept]] references
	start := 0
	for {
		idx := strings.Index(content[start:], "[[")
		if idx == -1 {
			break
		}
		start += idx + 2
		endIdx := strings.Index(content[start:], "]]")
		if endIdx == -1 {
			break
		}
		relatedConcept := strings.TrimSpace(content[start : start+endIdx])
		related = append(related, relatedConcept)
		start += endIdx + 2
	}
	return related
}

// GetConcept returns a specific PM concept
func (kb *PMKnowledgeBase) GetConcept(name string) *PMConcept {
	return kb.Concepts[name]
}

// GetConceptContent returns the full markdown content of a concept
func (kb *PMKnowledgeBase) GetConceptContent(name string) string {
	return kb.Content[name]
}

// GetRelatedConcepts returns concepts related to the given concept
func (kb *PMKnowledgeBase) GetRelatedConcepts(name string) []string {
	return kb.RelatedMap[name]
}

// GetConceptsByCategory returns all concepts in a category
func (kb *PMKnowledgeBase) GetConceptsByCategory(category string) []string {
	return kb.Categories[category]
}

// FindRelevantConcepts finds concepts relevant to a query
func (kb *PMKnowledgeBase) FindRelevantConcepts(query string) []*PMConcept {
	var relevant []*PMConcept
	queryLower := strings.ToLower(query)

	for name, concept := range kb.Concepts {
		nameLower := strings.ToLower(name)
		descLower := strings.ToLower(concept.Description)

		// Check if query matches concept name or description
		if strings.Contains(queryLower, nameLower) || strings.Contains(descLower, queryLower) ||
			strings.Contains(queryLower, strings.ReplaceAll(nameLower, " ", "")) {
			relevant = append(relevant, concept)
		}
	}

	return relevant
}

// GetContextualHint returns relevant context for a conversation topic
func (kb *PMKnowledgeBase) GetContextualHint(topic string) string {
	concepts := kb.FindRelevantConcepts(topic)
	if len(concepts) == 0 {
		return ""
	}

	hint := "Relevant PM concepts: "
	for i, concept := range concepts {
		if i > 0 {
			hint += ", "
		}
		hint += concept.Name
	}
	return hint
}

// GetFullContext returns comprehensive context for integration into advisor
func (kb *PMKnowledgeBase) GetFullContext() string {
	if len(kb.Concepts) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("Product Management Knowledge Base:\n")
	sb.WriteString("=" + strings.Repeat("=", 50) + "\n")

	// Group by category
	for category, concepts := range kb.Categories {
		sb.WriteString("\n" + category + ":\n")
		for _, conceptName := range concepts {
			concept := kb.Concepts[conceptName]
			sb.WriteString("• " + concept.Name + ": " + concept.Description + "\n")
		}
	}

	return sb.String()
}

// ListAllConcepts returns all loaded concepts as strings
func (kb *PMKnowledgeBase) ListAllConcepts() []string {
	var concepts []string
	for name := range kb.Concepts {
		concepts = append(concepts, name)
	}
	return concepts
}
