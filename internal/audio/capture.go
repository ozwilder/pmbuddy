package audio

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// VoiceCapture handles real microphone audio capture on macOS
type VoiceCapture struct {
	TempDir      string
	RecordBinary string // ffmpeg or sox
	Timeout      int    // seconds
}

// NewVoiceCapture creates a new voice capture handler
func NewVoiceCapture(timeoutSeconds int) *VoiceCapture {
	tempDir := os.TempDir()
	
	// Detect which recording tool is available
	recordBinary := detectRecordingBinary()
	
	return &VoiceCapture{
		TempDir:      tempDir,
		RecordBinary: recordBinary,
		Timeout:      timeoutSeconds,
	}
}

// detectRecordingBinary finds the best available audio recording tool
func detectRecordingBinary() string {
	// Try ffmpeg first (most common and reliable)
	if _, err := exec.LookPath("ffmpeg"); err == nil {
		return "ffmpeg"
	}
	
	// Fallback to sox
	if _, err := exec.LookPath("sox"); err == nil {
		return "sox"
	}
	
	// Neither found - will fall back to AppleScript
	return ""
}

// CaptureAudio records audio from the microphone for the specified duration
// Returns the path to the recorded WAV file
func (vc *VoiceCapture) CaptureAudio() (string, error) {
	if vc.RecordBinary == "" {
		return "", fmt.Errorf("no audio recording tool found (ffmpeg or sox required)")
	}

	// Create temporary file path
	timestamp := time.Now().Unix()
	audioFile := filepath.Join(vc.TempDir, fmt.Sprintf("pmbuddy_voice_%d.wav", timestamp))

	fmt.Println("\n🎤 CAPTURING YOUR VOICE...")
	fmt.Println(strings.Repeat("─", 70))
	fmt.Printf("Recording for %d seconds... Speak now!\n", vc.Timeout)
	
	// Show simple animated recording indicator
	for i := 0; i < vc.Timeout; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println()

	var cmd *exec.Cmd
	
	// Auto-detect the best microphone device
	// On macOS, device :3 is often the microphone, but we check :0-:3
	micDevice := findBestMicrophoneDevice()

	switch vc.RecordBinary {
	case "ffmpeg":
		// ffmpeg command to capture microphone on macOS with VOLUME BOOST
		cmd = exec.Command("ffmpeg",
			"-f", "avfoundation",           // macOS audio framework
			"-i", micDevice,                // Microphone device (auto-detected)
			"-t", fmt.Sprintf("%d", vc.Timeout),
			"-af", "loudnorm=I=-20:TP=-1.5:LRA=11", // Audio normalization/volume boost
			"-q:a", "9",                    // Quality setting
			"-acodec", "pcm_s16le",         // WAV codec
			"-ar", "16000",                 // 16kHz sample rate (good for speech recognition)
			"-ac", "1",                     // Mono
			"-y",                           // Overwrite output file
			audioFile,
		)

	case "sox":
		// sox command to capture microphone on macOS
		cmd = exec.Command("sox",
			"-d",                           // Capture from default device
			audioFile,                      // Output file
			"rate", "16000",                // 16kHz sample rate
			"channels", "1",                // Mono
			"trim", "0", fmt.Sprintf("%d", vc.Timeout),
		)
	}

	// Run the command
	if output, err := cmd.CombinedOutput(); err != nil {
		// ffmpeg writes to stderr, so check if file was created despite "error"
		if _, fileErr := os.Stat(audioFile); fileErr == nil {
			fmt.Println("✅ Audio captured successfully")
			return audioFile, nil
		}
		return "", fmt.Errorf("voice capture failed: %v\nOutput: %s", err, string(output))
	}

	fmt.Println("✅ Audio captured successfully")
	return audioFile, nil
}

// findBestMicrophoneDevice auto-detects the best microphone device
// Tries devices :0-:3 and returns the one with the most audio energy
func findBestMicrophoneDevice() string {
	bestDevice := ":0" // Default fallback
	bestRMS := 0.0
	
	// Try each device briefly to find the loudest one (likely the microphone)
	for i := 0; i <= 3; i++ {
		device := fmt.Sprintf(":%d", i)
		testFile := fmt.Sprintf("/tmp/pmbuddy_device_test_%d.wav", i)
		
		// Try to record 0.5 seconds
		cmd := exec.Command("ffmpeg",
			"-f", "avfoundation",
			"-i", device,
			"-t", "0.5",
			"-q:a", "9", "-acodec", "pcm_s16le", "-ar", "16000", "-ac", "1",
			"-y", testFile,
		)
		
		if err := cmd.Run(); err == nil && fileExists(testFile) {
			// Check audio energy
			rms := getAudioEnergy(testFile)
			if rms > bestRMS {
				bestRMS = rms
				bestDevice = device
			}
			// Cleanup
			os.Remove(testFile)
		}
	}
	
	return bestDevice
}

// getAudioEnergy calculates the RMS energy of audio in a WAV file
func getAudioEnergy(filename string) float64 {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) < 44 {
		return 0
	}
	
	// Skip WAV header (44 bytes) and read audio frames
	audioData := data[44:]
	if len(audioData) < 2 {
		return 0
	}
	
	// Convert bytes to int16 samples (simple, assuming little-endian)
	var sum float64
	for i := 0; i < len(audioData)-1; i += 2 {
		sample := int16(audioData[i]) | (int16(audioData[i+1]) << 8)
		sum += float64(sample*sample)
	}
	
	samples := len(audioData) / 2
	if samples == 0 {
		return 0
	}
	
	return math.Sqrt(sum / float64(samples))
}

// fileExists checks if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// TranscribeAudio converts recorded audio to text using Python speech recognition
func (vc *VoiceCapture) TranscribeAudio(audioFile string) (string, error) {
	// Find transcribe.py script
	scriptPath := findTranscribeScript()
	if scriptPath == "" {
		return "", fmt.Errorf("transcribe.py script not found")
	}

	// Run Python transcriber (suppress stderr to avoid debug clutter)
	cmd := exec.Command("python3", scriptPath, audioFile)
	output, err := cmd.Output()
	
	if err != nil {
		return "", fmt.Errorf("transcription process failed: %v", err)
	}

	text := strings.TrimSpace(string(output))
	if text == "" {
		return "", fmt.Errorf("no output from transcriber")
	}

	// Parse JSON response (from transcriber)
	text = extractTextFromOutput(text)
	
	if text == "" {
		return "", fmt.Errorf("no speech recognized")
	}

	return text, nil
}

// findTranscribeScript locates the transcribe.py helper script
func findTranscribeScript() string {
	possiblePaths := []string{
		"scripts/transcribe.py",
		"./scripts/transcribe.py",
		"/Users/ozwilder/Development/PMBuddy/scripts/transcribe.py",
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}

// extractTextFromOutput extracts transcribed text from Python script JSON output
func extractTextFromOutput(output string) string {
	output = strings.TrimSpace(output)
	
	// If output starts with {, try to parse as JSON
	if strings.HasPrefix(output, "{") {
		var result map[string]interface{}
		if err := json.Unmarshal([]byte(output), &result); err == nil {
			// Check if transcription was successful
			if success, ok := result["success"].(bool); ok && !success {
				// Transcription failed according to the script
				return ""
			}
			
			// Extract text field
			if text, ok := result["text"].(string); ok && text != "" {
				return text
			}
		}
	}
	
	// Fallback: if it doesn't look like JSON, return as-is
	// (but this should not happen with the fixed transcribe.py)
	if !strings.HasPrefix(output, "{") && output != "" {
		return output
	}
	
	return ""
}

// CaptureVoiceWithFallback attempts to capture voice with smart fallback
func (vc *VoiceCapture) CaptureVoiceWithFallback() (string, error) {
	// Try to capture audio from microphone
	audioFile, err := vc.CaptureAudio()
	if err != nil {
		fmt.Printf("⚠️  Microphone capture failed: %v\n", err)
		fmt.Printf("    Issue: %v\n", err)
		fmt.Println("\n💡 Your microphone may not be accessible to ffmpeg.")
		fmt.Println("   Using text input instead (just type your question).\n")
		
		// Fallback to AppleScript dialog (manual typing)
		return vc.TextInputDialog()
	}

	// Check if audio file has real content
	fileInfo, err := os.Stat(audioFile)
	if err != nil || fileInfo.Size() < 1000 {
		fmt.Printf("⚠️  Audio file seems empty or too small\n")
		fmt.Println("    Using text input instead.\n")
		os.Remove(audioFile)
		return vc.TextInputDialog()
	}

	// Try to transcribe the audio
	text, err := vc.TranscribeAudio(audioFile)
	defer os.Remove(audioFile) // Clean up temp file

	if err != nil {
		fmt.Printf("⚠️  Transcription failed: %v\n", err)
		fmt.Println("\n💡 Speech recognition needs clearer audio.")
		fmt.Println("   Using text input instead (just type your question).\n")
		
		// Fallback to text input
		return vc.TextInputDialog()
	}

	if text == "" {
		fmt.Println("⚠️  No speech was recognized (empty audio?)")
		fmt.Println("   Using text input instead.\n")
		return vc.TextInputDialog()
	}

	return text, nil
}

// TextInputDialog provides a fallback text input via AppleScript dialog
func (vc *VoiceCapture) TextInputDialog() (string, error) {
	applescript := `
set recognizedText to ""
try
	set recognizedText to text returned of (display dialog "🎤 Type your question:" default answer "" buttons {"Cancel", "OK"} default button 2 with icon note)
on error
	set recognizedText to ""
end try
return recognizedText
`

	cmd := exec.Command("osascript", "-e", applescript)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("dialog failed: %v", err)
	}

	text := strings.TrimSpace(string(output))
	if text == "" {
		return "", fmt.Errorf("no input provided")
	}

	return text, nil
}

// IsAvailable checks if real voice capture is possible
func (vc *VoiceCapture) IsAvailable() bool {
	return vc.RecordBinary != ""
}

// GetStatus returns current capture status
func (vc *VoiceCapture) GetStatus() map[string]interface{} {
	return map[string]interface{}{
		"recording_tool": vc.RecordBinary,
		"is_available":   vc.IsAvailable(),
		"timeout":        vc.Timeout,
		"temp_dir":       vc.TempDir,
	}
}
