package voice

import (
	"fmt"
	"os/exec"
	"strings"
)

// HandyInterface manages speech-to-text and text-to-speech via Handy
type HandyInterface struct {
	IsInitialized bool
	HandyPath     string
	Config        map[string]string
}

// NewHandyInterface creates a new Handy voice interface
func NewHandyInterface() *HandyInterface {
	return &HandyInterface{
		IsInitialized: false,
		HandyPath:     "/usr/local/bin/handy", // Default path - should be configurable
		Config:        make(map[string]string),
	}
}

// Initialize checks if Handy is available and ready
func (h *HandyInterface) Initialize() error {
	// Check if Handy is installed at default path
	cmd := exec.Command(h.HandyPath, "--version")
	output, err := cmd.Output()
	if err != nil {
		// Try alternative paths
		altPaths := []string{
			"/usr/bin/handy",
			"/opt/handy/bin/handy",
			"/Applications/Handy.app/Contents/MacOS/handy",
		}
		
		found := false
		for _, path := range altPaths {
			cmd := exec.Command(path, "--version")
			if output, err := cmd.Output(); err == nil {
				h.HandyPath = path
				fmt.Printf("✅ Handy found at: %s\n", path)
				fmt.Printf("✅ Handy version: %s\n", strings.TrimSpace(string(output)))
				h.IsInitialized = true
				found = true
				break
			}
		}
		
		if !found {
			return fmt.Errorf("Handy not found. Please install from DMG or add to PATH. Checked: %s and alternatives", h.HandyPath)
		}
		return nil
	}

	fmt.Printf("✅ Handy available: %s\n", strings.TrimSpace(string(output)))
	h.IsInitialized = true
	return nil
}

// SpeechToText converts audio input to text using Handy
func (h *HandyInterface) SpeechToText(audioInput string) (string, error) {
	if !h.IsInitialized {
		return "", fmt.Errorf("Handy not initialized")
	}

	cmd := exec.Command(h.HandyPath, "transcribe", audioInput)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("speech-to-text failed: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// TextToSpeech converts text to audio output using Handy
func (h *HandyInterface) TextToSpeech(text string) error {
	if !h.IsInitialized {
		return fmt.Errorf("Handy not initialized")
	}

	cmd := exec.Command(h.HandyPath, "speak", text)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("text-to-speech failed: %v", err)
	}

	return nil
}

// ListenAndRespond listens for voice input and returns transcribed text
func (h *HandyInterface) ListenAndRespond(timeout int) (string, error) {
	if !h.IsInitialized {
		return "", fmt.Errorf("Handy not initialized")
	}

	cmd := exec.Command(h.HandyPath, "listen", fmt.Sprintf("--timeout=%d", timeout))
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("listen failed: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// SetVoiceParams configures Handy voice parameters
func (h *HandyInterface) SetVoiceParams(language string, voice string, speed float32) {
	h.Config["language"] = language
	h.Config["voice"] = voice
	h.Config["speed"] = fmt.Sprintf("%.1f", speed)
}

// GetStatus returns the status of Handy interface
func (h *HandyInterface) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"initialized": h.IsInitialized,
		"handy_path":  h.HandyPath,
		"config":      h.Config,
	}
}

// Close cleans up Handy resources
func (h *HandyInterface) Close() error {
	// Handy doesn't require explicit cleanup, but we can add cleanup logic here
	h.IsInitialized = false
	return nil
}
