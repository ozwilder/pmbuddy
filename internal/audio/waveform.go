package audio

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// WaveformVisualizer creates animated ASCII waveforms
type WaveformVisualizer struct {
	width     int
	height    int
	samples   []float64
	maxSample float64
}

// NewWaveformVisualizer creates a new visualizer
func NewWaveformVisualizer(width, height int) *WaveformVisualizer {
	return &WaveformVisualizer{
		width:     width,
		height:    height,
		samples:   make([]float64, width),
		maxSample: 1.0,
	}
}

// AddSample adds a new audio sample to the waveform
func (w *WaveformVisualizer) AddSample(sample float64) {
	// Shift samples left
	for i := 0; i < len(w.samples)-1; i++ {
		w.samples[i] = w.samples[i+1]
	}
	// Add new sample
	w.samples[len(w.samples)-1] = math.Abs(sample)

	// Update max
	if math.Abs(sample) > w.maxSample {
		w.maxSample = math.Abs(sample)
	}
}

// SimulateAudioInput simulates voice by adding random peaks
func (w *WaveformVisualizer) SimulateAudioInput() {
	// Simulate voice patterns with occasional peaks
	baseLevel := rand.Float64() * 0.2
	if rand.Float64() > 0.7 {
		baseLevel += rand.Float64() * 0.6 // Add peaks
	}
	w.AddSample(baseLevel)
}

// RenderWaveform renders the waveform as ASCII art
func (w *WaveformVisualizer) RenderWaveform() string {
	var output strings.Builder

	// Render from top to bottom (height bars)
	for row := w.height - 1; row >= 0; row-- {
		threshold := float64(row) / float64(w.height)

		for col := 0; col < w.width; col++ {
			// Normalize sample to 0-1 range
			normalized := w.samples[col] / w.maxSample

			// Check if this row should show a bar for this sample
			if normalized >= threshold {
				output.WriteString("█")
			} else {
				output.WriteString(" ")
			}
		}
		output.WriteString("\n")
	}

	return output.String()
}

// SimpleListeningIndicator shows a simple animated indicator
func SimpleListeningIndicator() {
	indicators := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"}

	fmt.Println("\n🎤 CAPTURING YOUR VOICE...")
	fmt.Println(strings.Repeat("─", 70))
	fmt.Println("Audio waveform:\n")

	waveform := NewWaveformVisualizer(70, 6)

	// Quick animation - just 8 frames
	for i := 0; i < 8; i++ {
		// Add simulated audio samples
		for _ = range indicators {
			waveform.SimulateAudioInput()
		}

		// Display waveform
		fmt.Print(waveform.RenderWaveform())

		// Show recording indicator
		fmt.Printf("  %s Recording...\n", indicators[i%len(indicators)])
		time.Sleep(100 * time.Millisecond)

		// Move cursor up to overwrite
		fmt.Print("\033[7A") // Move up 7 lines
	}

	// Clear the animation
	fmt.Print("\033[7B") // Move down
	fmt.Print(strings.Repeat("\n", 2))
}
