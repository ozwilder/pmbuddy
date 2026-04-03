package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// StorageManager handles persistence of conversations and learnings
type StorageManager struct {
	BasePath  string
	SessionID string
}

// SessionData represents a complete session
type SessionData struct {
	SessionID       string                 `json:"session_id"`
	CreatedAt       time.Time              `json:"created_at"`
	LastUpdated     time.Time              `json:"last_updated"`
	Conversations   []interface{}          `json:"conversations"`
	LearningHistory []interface{}          `json:"learning_history"`
	Preferences     map[string]interface{} `json:"preferences"`
}

// NewStorageManager creates a new storage manager
func NewStorageManager(basePath string, sessionID string) *StorageManager {
	return &StorageManager{
		BasePath:  basePath,
		SessionID: sessionID,
	}
}

// EnsureStorageDir creates necessary storage directories
func (sm *StorageManager) EnsureStorageDir() error {
	dirs := []string{
		filepath.Join(sm.BasePath, "sessions"),
		filepath.Join(sm.BasePath, "conversations"),
		filepath.Join(sm.BasePath, "learnings"),
		filepath.Join(sm.BasePath, "backups"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	return nil
}

// SaveSession persists a complete session
func (sm *StorageManager) SaveSession(data SessionData) error {
	data.LastUpdated = time.Now()

	filePath := filepath.Join(sm.BasePath, "sessions", sm.SessionID+".json")

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session data: %v", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write session file: %v", err)
	}

	return nil
}

// LoadSession retrieves a previously saved session
func (sm *StorageManager) LoadSession(sessionID string) (*SessionData, error) {
	filePath := filepath.Join(sm.BasePath, "sessions", sessionID+".json")

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read session file: %v", err)
	}

	var data SessionData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session data: %v", err)
	}

	return &data, nil
}

// ListSessions returns all available sessions
func (sm *StorageManager) ListSessions() ([]string, error) {
	sessionsDir := filepath.Join(sm.BasePath, "sessions")

	files, err := os.ReadDir(sessionsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read sessions directory: %v", err)
	}

	var sessions []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			sessionID := file.Name()[:len(file.Name())-5] // Remove .json
			sessions = append(sessions, sessionID)
		}
	}

	return sessions, nil
}

// SaveConversation saves a single conversation
func (sm *StorageManager) SaveConversation(convID string, messages interface{}) error {
	filePath := filepath.Join(sm.BasePath, "conversations", sm.SessionID+"_"+convID+".json")

	jsonData, err := json.MarshalIndent(messages, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal conversation: %v", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write conversation file: %v", err)
	}

	return nil
}

// DeleteSession removes a session
func (sm *StorageManager) DeleteSession(sessionID string) error {
	filePath := filepath.Join(sm.BasePath, "sessions", sessionID+".json")

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}

	return nil
}

// BackupSession creates a backup of a session
func (sm *StorageManager) BackupSession(sessionID string) error {
	sourcePath := filepath.Join(sm.BasePath, "sessions", sessionID+".json")
	backupPath := filepath.Join(sm.BasePath, "backups", sessionID+"_"+time.Now().Format("20060102150405")+".json")

	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read session for backup: %v", err)
	}

	if err := os.WriteFile(backupPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write backup: %v", err)
	}

	return nil
}

// GetStorageStats returns storage statistics
func (sm *StorageManager) GetStorageStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	sessions, err := sm.ListSessions()
	if err != nil {
		return nil, err
	}

	stats["total_sessions"] = len(sessions)
	stats["current_session"] = sm.SessionID
	stats["storage_path"] = sm.BasePath

	return stats, nil
}
