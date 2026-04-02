# PMBuddy - Product Manager AI Advisor

An AI-powered Product Manager advisor that leverages accumulated knowledge and improves over time through conversation and experience.

## Features

- 🤖 **AI-Powered Advice**: Product management guidance based on best practices and accumulated knowledge
- 📚 **Learning System**: Improves over time through ongoing conversations and context
- 🎤 **Multi-Modal Interface**: Text input + voice interaction (Handy speech-to-text/text-to-speech)
- 💾 **Persistent Knowledge**: Stores conversations and learnings for continuous improvement
- 🔄 **Cross-Platform**: Native binaries for macOS and Windows (Go-based)
- 📊 **Context-Aware**: Maintains conversation history and task context

## Quick Start

### Prerequisites
- Go 1.21 or later
- Handy (speech-to-text/text-to-speech) installed locally
- macOS 10.12+ or Windows 11+

### Installation

```bash
cd /Users/ozwilder/Development/PMBuddy
go mod download
go build -o pmbuddy ./cmd/pmbuddy
```

### Running PMBuddy

```bash
# Text interface
./pmbuddy

# With voice enabled
./pmbuddy --voice

# Start interactive session
./pmbuddy session start
```

## Project Structure

```
pmbuddy/
├── cmd/
│   └── pmbuddy/              # Main CLI entry point
├── internal/
│   ├── agent/                # Core PM advisor agent
│   ├── knowledge/            # Knowledge management & learning
│   ├── voice/                # Handy integration (voice I/O)
│   └── storage/              # Persistence layer (conversations, context)
├── scripts/
│   ├── build.sh              # Build for macOS/Windows
│   └── handy_setup.sh        # Handy integration setup
├── docs/
│   ├── ARCHITECTURE.md       # System design
│   ├── VOICE_SETUP.md        # Handy integration guide
│   └── API.md                # Internal API documentation
└── tests/                    # Test files
```

## Development

### Build for macOS
```bash
./scripts/build.sh darwin
```

### Build for Windows
```bash
./scripts/build.sh windows
```

### Run Tests
```bash
go test ./...
```

## Architecture

PMBuddy consists of 4 core subsystems:

1. **Agent** - PM advisor logic and decision making
2. **Knowledge** - Learning system, conversation history, context management
3. **Voice** - Handy integration for speech I/O
4. **Storage** - Persistent data layer for all conversations and learnings

## Handy Integration

PMBuddy uses Handy for local, private voice interaction:
- Speech-to-text: Converts voice input to text
- Text-to-speech: Speaks responses aloud

See [VOICE_SETUP.md](docs/VOICE_SETUP.md) for detailed configuration.

## Roadmap

- [ ] Core text-based PM advisor
- [ ] Voice I/O via Handy
- [ ] Conversation persistence
- [ ] Learning engine (context improvement)
- [ ] Multi-session management
- [ ] Analytics & insights dashboard
- [ ] Plugin system for custom advisors

## License

MIT

## Author

Built for product management excellence 🚀
