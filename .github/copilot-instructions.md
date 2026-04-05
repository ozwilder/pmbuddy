# PMBuddy - Copilot Instructions

This document guides AI assistants working on the PMBuddy codebase to understand its architecture, conventions, and build processes.

## Build & Run Commands

### Build the project
```bash
go build -o pmbuddy ./cmd/pmbuddy
```

### Run locally (text mode)
```bash
./pmbuddy
```

### Run with voice (macOS native speech recognition)
```bash
./pmbuddy --voice-native
```

### Run in discussion mode (full voice I/O with TTS)
```bash
./pmbuddy --discussion
```

### Run with Handy integration (if installed)
```bash
./pmbuddy --voice
```

### Run a single test
There are no unit tests currently. The project can be tested interactively by running the binary.

### Run all available commands
```bash
./pmbuddy --help
```

## Architecture Overview

PMBuddy is a Go-based CLI application with the following core subsystems:

### 1. **Agent** (`internal/agent/`)
- `advisor.go`: Core `PMAdvisor` struct that handles PM queries and responses
  - `ProcessQuery()`: Main entry point for handling user input
  - `ValidateQuery()`: **ENHANCED** - Validates if input is complete, clear, and relevant to product management
    - Checks for structural completeness (minimum length, complete sentences)
    - Rejects off-domain questions (coding, cooking, weather, etc.)
    - Rejects offensive/inappropriate content
    - Returns rejection message if validation fails instead of processing
    - Only allows PM-domain relevant queries
  - `generateResponseWithLLM()`: Generates responses using configured LLM
  - Memory tracking, context management, and conversation history
- `llm.go`: LLM interface and implementations (Ollama client and mock)

### 2. **Audio** (`internal/audio/`)
- `capture.go`: Real-time microphone capture with device auto-detection
  - Uses ffmpeg or sox for audio recording
  - Auto-detects best microphone device based on audio energy
  - Handles fallback to text input if voice capture fails
  - Simple animated indicator (dots) during recording (waveform visualization removed)

### 3. **Voice** (`internal/voice/`)
- `handy.go`: Integration with Handy speech-to-text/text-to-speech tool
  - Multi-path detection for Handy binary location
  - Speech-to-text transcription
  - Text-to-speech output

### 4. **Knowledge** (`internal/knowledge/`)
- Learning system that extracts preferences and topics from conversations
- Maintains preference profiles (approach, thinking style, focus, etc.)
- Detects user context and learns over time

### 5. **Storage** (`internal/storage/`)
- Persistence layer for conversations and learnings

### 6. **CLI Entry Point** (`cmd/pmbuddy/main.go`)
- Multiple interaction modes: text, voice (native), voice (Handy), discussion
- **NEW TTS Control**: Tracks and cancels text-to-speech when user starts typing or voice input is opened
  - `stopCurrentTTS()`: Stops currently running speech process
  - `speakAsync()`: Runs TTS asynchronously with process tracking

## Key Conventions

### Input Validation (Enhanced Domain Checking)
User queries must pass the `ValidateQuery()` check before processing:

**Structural validation:**
- Minimum 5 characters
- Not incomplete patterns (e.g., "how do i" with nothing after)
- No trailing incomplete punctuation ("...", "- ")
- Sufficient substance beyond pronouns and generic words

**Domain validation (NEW):**
- Query must contain PM-related keywords (product, feature, user, strategy, roadmap, metric, etc.)
- Rejects coding, cooking, weather, sports, entertainment topics
- Rejects offensive/inappropriate content
- If validation fails, returns specific rejection message instead of processing

If validation fails, show rejection message and return to input prompt instead of processing the request.

### Voice Output (TTS)
- Always use `speakAsync()` instead of direct `exec.Command("say", ...)` calls
- TTS is automatically cancelled when:
  - User starts typing in CLI
  - User opens voice input dialog by pressing ENTER
  - New output is being spoken
- Use global `currentTTSLock` mutex for thread-safe process management

### Response Formatting
- Keep responses short and conversational (2-3 sentences max)
- Use natural, native English
- Only provide detailed output if user explicitly asks for "detailed" or "technical"
- Sound like a colleague, not a textbook

### Error Handling
- Gracefully fall back to text mode if voice fails
- Mock LLM responses if Ollama is unavailable
- Provide user-friendly error messages

### Process State
- `PMAdvisor` maintains conversation `Memory` (array of messages with timestamps)
- `Context` map stores session-specific state
- `KnowledgeBase` tracks user preferences and learnings across conversation
- All three are session-scoped (new instance per conversation)

## Agent Skills (PM Advisor Capabilities)

The `PMAdvisor` provides PM-specific reasoning:
- Product strategy and vision
- Feature prioritization and roadmapping
- User research and discovery
- Metrics and KPIs
- Go-to-market planning
- Team dynamics and leadership

Query validation and clarification is now part of the Agent Skills - if a user's input is incomplete or unclear, the agent will prompt for clarification rather than attempting to process ambiguous requests.

## Important Files

| File | Purpose |
|------|---------|
| `cmd/pmbuddy/main.go` | CLI entry point, input loop, TTS control |
| `internal/agent/advisor.go` | Core PM advisor logic and query validation |
| `internal/agent/llm.go` | LLM integration (Ollama + mock) |
| `internal/audio/capture.go` | Voice capture from microphone |
| `internal/voice/handy.go` | Handy speech-to-text/TTS integration |
| `internal/knowledge/` | Learning system |
| `go.mod` | Dependencies (spf13/cobra, etc.) |

## Dependencies

Main external dependencies:
- `github.com/spf13/cobra`: CLI framework
- ffmpeg or sox: Audio recording (macOS system tools)
- Python 3: For speech recognition transcription scripts
- Ollama: Optional LLM backend (falls back to mock responses)
- Handy: Optional speech-to-text/TTS (can use macOS native instead)

## Testing

The application is primarily tested interactively due to its voice/audio components. To verify changes:

1. Build: `go build -o pmbuddy ./cmd/pmbuddy`
2. Run text mode: `./pmbuddy`
3. Run voice mode: `./pmbuddy --voice-native`
4. Test query validation with incomplete inputs
5. Test TTS cancellation by typing while voice output is playing

## Recent Changes (Current Session)

### Session 1: TTS & Validation Infrastructure
1. **Added Query Validation** (`ValidateQuery` method in `agent/advisor.go`)
   - Checks for incomplete or unclear user input
   - Returns clarifying question if validation fails

2. **Added TTS Interrupt Control** (in `cmd/pmbuddy/main.go`)
   - Tracks text-to-speech process globally
   - Automatically stops TTS when user starts typing
   - Integrated into all voice modes

### Session 2: Enhanced Domain Validation & Waveform Removal
1. **Enhanced Query Validation with Domain Checking** (`agent/advisor.go`)
   - Now rejects off-domain questions (coding, cooking, weather, etc.)
   - Rejects inappropriate/offensive content
   - Only processes PM-related queries
   - Returns domain-specific rejection messages

2. **Removed Audio Waveform Visualization** (`internal/audio/`)
   - Deleted `waveform.go` file
   - Replaced `SimpleListeningIndicator()` with simple dot animation
   - Improves stability and removes non-functional feature

### Session 3: Fixed Audio Transcription Hallucination Bug
1. **Fixed transcribe.py Script** (`scripts/transcribe.py`)
   - **Root Cause**: Script was returning early without printing JSON output
   - **Fix**: Now always prints JSON to stdout (success/failure + transcribed text)
   - Ensures Go can properly parse transcription results
   - Prevents "hallucinated questions" from appearing

2. **Improved capture.go JSON Parsing** (`internal/audio/capture.go`)
   - Enhanced `extractTextFromOutput()` to check "success" field
   - Returns empty string on transcription failure (forces fallback to text input)
   - Better error handling and validation
   - Properly validates transcription engine output

**Result**: Audio input now works reliably - transcribed text is captured correctly, no more made-up questions
