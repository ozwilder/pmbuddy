package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ozwilder/pmbuddy/internal/agent"
	"github.com/ozwilder/pmbuddy/internal/audio"
	"github.com/ozwilder/pmbuddy/internal/voice"
)

var (
	voiceEnabled     bool
	voiceNative      bool
	discussionMode   bool
	configPath       string
)

var rootCmd = &cobra.Command{
	Use:   "pmbuddy",
	Short: "PMBuddy - Your AI Product Manager Advisor",
	Long: `PMBuddy is an AI-powered Product Manager advisor that learns from conversations
and helps you make better product decisions. It supports both text and voice interaction.`,
	Run: func(cmd *cobra.Command, args []string) {
		startPMBuddy()
	},
}

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Manage PM Buddy sessions",
}

var sessionStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an interactive PM Buddy session",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting PMBuddy session...")
		startInteractiveSession()
	},
}

var sessionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved sessions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📋 Available sessions:")
		// TODO: Implement session listing
		fmt.Println("  (no sessions saved yet)")
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&voiceEnabled, "voice", "v", false, "Enable voice interaction via Handy")
	rootCmd.Flags().BoolVar(&voiceNative, "voice-native", false, "Enable voice with macOS native speech recognition (no Handy required)")
	rootCmd.Flags().BoolVar(&discussionMode, "discussion", false, "Enable full voice discussion mode (speak questions, hear responses)")
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to config file")

	sessionCmd.AddCommand(sessionStartCmd)
	sessionCmd.AddCommand(sessionListCmd)
	rootCmd.AddCommand(sessionCmd)
}

func startPMBuddy() {
	fmt.Println(`
╔════════════════════════════════════════════════════════════════╗
║                      PMBuddy - PM Advisor                      ║
║            Your AI-powered Product Manager consultant          ║
╚════════════════════════════════════════════════════════════════╝
	`)

	if discussionMode {
		fmt.Println("🎤 Discussion mode enabled (voice input & output)")
		startDiscussionMode()
	} else if voiceEnabled {
		fmt.Println("🎤 Voice interface enabled (Handy)")
		startVoiceSession()
	} else if voiceNative {
		fmt.Println("🎤 Voice interface enabled (macOS Native - voice output only)")
		startNativeVoiceSession()
	} else {
		fmt.Println("📝 Text interface mode")
		startTextSession()
	}
}

func startTextSession() {
	fmt.Println("\n💬 Text Session Started")
	fmt.Println("Type 'help' for commands, 'exit' to quit\n")

	advisor := agent.NewPMAdvisor()

	// Display welcome message
	welcome := advisor.GetWelcomeMessage()
	fmt.Printf("PMBuddy: %s\n\n", welcome)

	// Interactive text loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// Handle special commands
		switch strings.ToLower(input) {
		case "exit", "quit", "bye":
			fmt.Println("\nPMBuddy: Thanks for chatting! Goodbye! 👋")
			return
		case "help":
			printHelp()
			continue
		case "clear":
			advisor.ClearMemory()
			fmt.Println("PMBuddy: Conversation history cleared.")
			continue
		case "status":
			printSessionStatus(advisor)
			continue
		}

		// Process user query
		response, err := advisor.ProcessQuery(input)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		fmt.Printf("\nPMBuddy: %s\n\n", response)
	}
}

func printHelp() {
	fmt.Println(`
Available commands:
  exit, quit, bye    - Leave PMBuddy
  help              - Show this help message
  clear             - Clear conversation history
  status            - Show session information
  
Or just ask me any product management question!
`)
}

func printSessionStatus(advisor *agent.PMAdvisor) {
	memory := advisor.GetMemory()
	context := advisor.GetContext()
	kb := advisor.GetKnowledgeBase()
	llm := advisor.GetLLM()
	stats := kb.GetLearningStats()
	topics := kb.GetTopics()
	insights := kb.GenerateInsights()
	profile := kb.GetPreferenceProfile()
	prefs := kb.GetPreferences()

	fmt.Printf("\n📊 Session Status:\n")
	fmt.Printf("  Session ID: %s\n", advisor.SessionID)
	fmt.Printf("  Messages: %d\n", len(memory))
	fmt.Printf("  Context Keys: %d\n", len(context))
	fmt.Printf("  AI Engine: %s\n\n", llm.GetName())

	fmt.Printf("🧠 Learning Stats:\n")
	fmt.Printf("  Total Learnings: %d\n", stats["total_learnings"])
	fmt.Printf("  Topics Tracked: %d\n", stats["unique_topics"])
	fmt.Printf("  Avg Confidence: %.1f%%\n", stats["avg_confidence"].(float64)*100)
	fmt.Printf("  Preferences Detected: %d\n\n", len(prefs))

	if len(topics) > 0 {
		fmt.Printf("📌 Primary Topics:\n")
		for i, topic := range topics {
			if i >= 3 {
				break
			}
			fmt.Printf("  • %s\n", topic)
		}
		fmt.Printf("\n")
	}

	// Display PM Preference Profile
	if profile.Confidence > 0.2 {
		fmt.Printf("🎯 Your PM Style (%.0f%% confidence):\n", profile.Confidence*100)
		if profile.Approach != "mixed" {
			fmt.Printf("  Approach: %s\n", profile.Approach)
		}
		if profile.Thinking != "mixed" {
			fmt.Printf("  Thinking: %s\n", profile.Thinking)
		}
		if profile.Style != "mixed" {
			fmt.Printf("  Style: %s\n", profile.Style)
		}
		if profile.DecisionFramework != "mixed" {
			fmt.Printf("  Decisions: %s\n", profile.DecisionFramework)
		}
		if profile.Focus != "mixed" {
			fmt.Printf("  Focus: %s\n", profile.Focus)
		}
		if profile.Timeline != "mixed" {
			fmt.Printf("  Timeline: %s\n", profile.Timeline)
		}
		fmt.Printf("\n")
	}

	if len(insights) > 0 {
		fmt.Printf("💡 Insights:\n")
		for _, insight := range insights {
			fmt.Printf("  • %s\n", insight)
		}
		fmt.Printf("\n")
	}
}

func startVoiceSession() {
	fmt.Println("\n🎤 Voice Session Started")
	fmt.Println("Listening for voice input...\n")

	// Initialize voice interface
	voiceInterface := voice.NewHandyInterface()

	if err := voiceInterface.Initialize(); err != nil {
		fmt.Printf("⚠️  Warning: %v\n", err)
		fmt.Println("\n📝 Falling back to text mode...\n")
		startTextSession()
		return
	}

	fmt.Println("✅ Handy interface ready")

	// TODO: Implement voice interaction loop
	fmt.Println("(Voice interface - coming soon)")
}

func startNativeVoiceSession() {
	fmt.Println("\n🎤 Voice Session Started (macOS Native)")
	fmt.Println("Using macOS built-in speech recognition\n")

	// Initialize advisor (handles LLM detection internally)
	advisor := agent.NewPMAdvisor()

	fmt.Println(advisor.GetWelcomeMessage())
	fmt.Println("\n🎙️  Speak your questions (or type for fallback)")
	fmt.Println("Commands: 'status', 'help', 'clear', 'exit'\n")

	scanner := bufio.NewScanner(os.Stdin)

	// Use macOS say command for voice output and keyboard input
	fmt.Print("\n🎤 You (type or say): ")

	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			fmt.Print("🎤 You (type or say): ")
			continue
		}

		// Handle commands
		if input == "exit" || input == "quit" || input == "bye" {
			fmt.Println("\nPMBuddy: Thanks for chatting! Goodbye! 👋")
			// Use macOS say for goodbye
			_ = exec.Command("say", "Thanks for chatting. Goodbye!").Run()
			break
		}

		if input == "help" {
			printHelp()
			fmt.Print("\n🎤 You (type or say): ")
			continue
		}

		if input == "clear" {
			advisor = agent.NewPMAdvisor()
			fmt.Println("\n✨ Conversation cleared\n")
			fmt.Print("🎤 You (type or say): ")
			continue
		}

		if input == "status" {
			printSessionStatus(advisor)
			fmt.Print("\n🎤 You (type or say): ")
			continue
		}

		// Process query
		response, err := advisor.ProcessQuery(input)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			fmt.Print("\n🎤 You (type or say): ")
			continue
		}

		fmt.Printf("\nPMBuddy: %s\n\n", response)

		// Speak response using macOS say command (first 500 chars for brevity)
		responseToSpeak := response
		if len(responseToSpeak) > 500 {
			responseToSpeak = responseToSpeak[:500] + "..."
		}
		_ = exec.Command("say", responseToSpeak).Run()

		fmt.Print("🎤 You (type or say): ")
	}
}

func startDiscussionMode() {
	// Welcome banner with new messaging
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("╔" + strings.Repeat("═", 68) + "╗")
	fmt.Println("║" + strings.Repeat(" ", 68) + "║")
	fmt.Println("║" + centerText("🎤 PM Buddy", 68) + "║")
	fmt.Println("║" + centerText("Your AI Partner for Product Tasks & Thinking", 68) + "║")
	fmt.Println("║" + strings.Repeat(" ", 68) + "║")
	fmt.Println("╚" + strings.Repeat("═", 68) + "╝")
	fmt.Println(strings.Repeat("═", 70) + "\n")

	fmt.Println("📋 I'm here to help you with:")
	fmt.Println("  • Product strategy & vision")
	fmt.Println("  • Feature prioritization & roadmapping")
	fmt.Println("  • User research & discovery")
	fmt.Println("  • Metrics, KPIs & analytics")
	fmt.Println("  • Go-to-market planning")
	fmt.Println("  • Product thinking & problem solving\n")

	// Initialize advisor
	advisor := agent.NewPMAdvisor()

	fmt.Println("💡 How it works:")
	fmt.Println("  • Type your question in the CLI and press ENTER")
	fmt.Println("  • Or press ENTER (empty) to open voice/text dialog")
	fmt.Println("  • I'll respond with voice output (you'll hear the answer!)")
	fmt.Println("  • Commands: 'status', 'help', 'exit'")
	fmt.Println(strings.Repeat("─", 70) + "\n")

	// Create scanner once outside loop for consistency
	scanner := bufio.NewScanner(os.Stdin)

	// First question
	firstQuestion := "What would you like to discuss about today?"
	fmt.Printf("\n\033[1m🤖 PM Buddy:\033[0m %s\n", firstQuestion)
	fmt.Println(strings.Repeat("─", 70))

	// Main loop: Listen for text/voice input via CLI
	for {
		// Display input prompt
		fmt.Print("\n⌨️  Type your question (or press ENTER for voice dialog): ")
		os.Stdout.Sync()
		
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		// If empty, show voice/text dialog
		if input == "" {
			fmt.Println("\n" + strings.Repeat("─", 70))
			fmt.Println("🎤 Opening voice/text input dialog...")
			dialogInput := captureVoiceInput()
			if dialogInput != "" {
				input = dialogInput
				fmt.Println(strings.Repeat("─", 70))
				fmt.Printf("  🎤 You: %s\n", input)
				fmt.Println(strings.Repeat("─", 70))
			} else {
				fmt.Println("❌ Dialog cancelled. Try typing instead.\n")
				continue
			}
		} else {
			// Show what was typed
			fmt.Println(strings.Repeat("─", 70))
			fmt.Printf("  📝 You: %s\n", input)
			fmt.Println(strings.Repeat("─", 70))
		}

		// Handle exit commands
		if input == "exit" || input == "quit" || input == "bye" {
			fmt.Println("\n" + strings.Repeat("─", 60))
			fmt.Println("👋 Thanks for chatting! Keep building great products!")
			fmt.Println(strings.Repeat("─", 60))
			go exec.Command("say", "Thanks for chatting. Goodbye!").Run()
			break
		}

		if input == "help" {
			printHelp()
			continue
		}

		if input == "clear" {
			advisor = agent.NewPMAdvisor()
			fmt.Println("\n✨ Conversation cleared")
			go exec.Command("say", "Conversation cleared").Run()
			continue
		}

		if input == "status" {
			printSessionStatus(advisor)
			go exec.Command("say", "Showing status").Run()
			continue
		}

		// Process query
		fmt.Println(strings.Repeat("─", 70))
		fmt.Println("⏳ Thinking...")
		response, err := advisor.ProcessQuery(input)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			fmt.Println(strings.Repeat("─", 70))
			continue
		}

		// Check if response is too short or seems incomplete
		if len(strings.TrimSpace(response)) < 20 {
			fmt.Println("\n\033[1m🤖 PM Buddy:\033[0m Can you please say it again? I'm not sure I understood that clearly.")
			fmt.Println(strings.Repeat("─", 70))
			go exec.Command("say", "Can you please say it again").Run()
			continue
		}

		// Display response with formatting
		fmt.Println("\n\033[1m🤖 PM Buddy:\033[0m")
		fmt.Println(response)
		fmt.Println(strings.Repeat("─", 70))
		
		// Speak response using macOS say command (async to prevent blocking)
		responseToSpeak := response
		if len(responseToSpeak) > 600 {
			responseToSpeak = responseToSpeak[:600] + "..."
		}
		go exec.Command("say", responseToSpeak).Run()
	}
}

func startInteractiveSession() {
	fmt.Println("\n📱 Interactive Session")
	fmt.Println("Choose your interface:")
	fmt.Println("  1) Text")
	fmt.Println("  2) Voice (requires Handy)")

	// TODO: Implement interactive mode selection
}

// captureVoiceInput captures voice directly from microphone without dialogs
func captureVoiceInput() string {
	// Create voice capture handler (5 second recording)
	voiceCapture := audio.NewVoiceCapture(5)

	// Attempt to capture real voice with fallback to text
	text, err := voiceCapture.CaptureVoiceWithFallback()
	if err != nil {
		fmt.Printf("❌ Voice capture error: %v\n", err)
		return ""
	}
	return text
}

// centerText centers text within a given width
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-padding-len(text))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
