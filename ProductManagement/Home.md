# PMBuddy Project - Vault Home

Welcome to the PMBuddy Obsidian vault! This is the central hub for all project knowledge, research, and work documentation.

## 🎯 Current Project Status

**Project**: PMBuddy - AI-powered Product Manager Advisor  
**Repository**: github.com/ozwilder/pmbuddy  
**Status**: Active Development  
**Last Updated**: 2026-04-05

## 📂 Vault Organization

### Core Sections
- **Notes/** - Active research and exploration work
- **Output/** - Completed deliverables and finalized work
- **Topics/** - Feature ideas and future work backlog
- **Root Level** - Core PM knowledge base

### Knowledge Base (Root Level)
- [[Definition]] - Core PM definitions and terminology
- [[Persona]] - User persona framework
- [[User Story]] - User story template and guidelines
- [[Jobs to Be Done]] - JTBD framework
- [[Product discovery]] - Discovery process documentation
- [[Customers Need]] - Customer needs analysis
- [[Reason]] - Core reasoning and frameworks

## 🔧 Architecture Overview

```
PMBuddy
├── cmd/pmbuddy/       → CLI entry point
├── internal/
│   ├── agent/         → PM advisor logic
│   ├── audio/         → Voice/audio handling
│   ├── voice/         → Voice I/O
│   └── knowledge/     → PM knowledge loader
├── scripts/
│   ├── transcribe.py  → Speech-to-text (OpenAI Whisper)
│   └── voice_input_simple.py
└── ProductManagement/ → This vault (PM knowledge base)
```

## 🚀 Latest Work Sessions

### Session: License & Dependency Management (2026-04-05)
**Completed**:
- ✅ Created LICENSE (MIT)
- ✅ Created DEPENDENCIES.txt (11 KB, all licenses documented)
- ✅ Created requirements.txt (pip3-compatible)
- ✅ Added auto-installation of openai-whisper during init()
- ✅ Set up Obsidian vault integration

**Files Created**:
- LICENSE
- DEPENDENCIES.txt
- requirements.txt
- CLAUDE.md (this protocol)
- Home.md (this index)
- Notes/, Output/, Topics/ folders

**See**: [[PMBuddy-License-Setup]] in Output/

### Previous Sessions
- **Session 4**: PM Knowledge Base Integration
  - Created pm_knowledge_loader.go
  - Integrated ProductManagement folder as knowledge source
  - See: [[PM-Knowledge-Integration]]

- **Session 3**: Audio Transcription Bug Fix
  - Fixed Python-Go speech-to-text communication
  - Verified Whisper transcription working
  - See: [[Audio-Transcription-Fix]]

- **Session 2**: Query Validation & Audio Improvements
  - Enhanced domain validation (50+ patterns)
  - Removed broken waveform visualization
  - See: [[Query-Validation]]

- **Session 1**: TTS Interrupt & Basic Validation
  - Added stopCurrentTTS() function
  - Implemented query validation
  - See: [[TTS-Interrupt]]

## 📋 Active Features

| Feature | Status | Notes |
|---------|--------|-------|
| Text Input Mode | ✅ Complete | Core conversation interface |
| Voice Input (Handy) | ✅ Complete | External device support |
| Voice Input (Native) | ✅ Complete | macOS built-in |
| TTS Output | ✅ Complete | Text-to-speech via macOS |
| Query Validation | ✅ Complete | Domain checking + structure |
| PM Knowledge Base | ✅ Complete | Integrated ProductManagement folder |
| Audio Transcription | ✅ Complete | OpenAI Whisper + fallback engines |
| Dependency Management | ✅ Complete | Auto-install on first run |
| Session Management | 🔄 In Progress | Save/resume conversations |

## 🎓 PM Concepts Integrated

The system has access to these PM frameworks:
- **User Stories** - Structured way to capture requirements
- **Personas** - Understanding target users
- **Jobs to Be Done** - Deep motivation understanding
- **Product Discovery** - Discovery methodologies
- **Customer Needs** - Customer-centric thinking
- **Reasoning Frameworks** - Decision-making approaches

## 🔍 Key Files to Know

### Source Code
- `cmd/pmbuddy/main.go` - CLI entry point, initialization
- `internal/agent/advisor.go` - PM advisor logic
- `internal/audio/capture.go` - Audio handling
- `scripts/transcribe.py` - Whisper speech-to-text
- `internal/knowledge/pm_knowledge_loader.go` - Knowledge integration

### Configuration
- `go.mod` - Go dependencies
- `requirements.txt` - Python dependencies
- `LICENSE` - MIT license
- `DEPENDENCIES.txt` - All dependencies with licenses

### Documentation
- `.github/copilot-instructions.md` - AI assistant guide
- `README.md` - Project overview
- `START_HERE.md` - Quick start guide

## 🛠️ Development Workflow

### For Code Changes
1. Review relevant notes in Notes/
2. Check CLAUDE.md for vault conventions
3. Make code changes with clear commits
4. Document in Output/ when complete

### For Research
1. Create dated file in Notes/ (e.g., `20260405-topic.md`)
2. Add [[wiki links]] to related concepts
3. Include findings and decisions
4. Move to Output/ when concluded

### For Features
1. Create in Topics/ as idea
2. Move to Notes/ when researching
3. Implement in code
4. Document in Output/ as deliverable
5. Update Home.md with new status

## 📈 Project Metrics

**Build Status**: ✅ Successful  
**Code Quality**: Production Ready  
**Test Coverage**: Core features validated  
**Documentation**: Comprehensive  
**User Experience**: Zero-friction setup  

## 🎯 Next Steps / Backlog

See `Topics/` folder for:
- [ ] Session persistence feature
- [ ] Advanced analytics
- [ ] Multi-language support
- [ ] Custom PM frameworks
- [ ] Integration with external tools

## 🔗 Important Links

- **Repository**: /Users/ozwilder/Development/PMBuddy
- **Vault Location**: ProductManagement/
- **Main Binary**: pmbuddy
- **Build Command**: `go build -o pmbuddy ./cmd/pmbuddy`

## 💬 Getting Started with CLAUDE

I'm Claude, your AI assistant. I treat this vault as my external memory. When we work together:

1. **I'll reference** previous sessions from Home.md and Output/
2. **I'll research** in Notes/ and link findings with [[wiki links]]
3. **I'll document** completed work in Output/
4. **I'll check** Topics/ for planned features
5. **I'll stay organized** using this vault structure

This ensures continuity across sessions and a clear audit trail of all work.

---

**Quick Links**:
- 📖 [[CLAUDE.md]] - How to use this vault
- 📝 [[Notes/]] - Active research
- ✅ [[Output/]] - Completed work
- 💡 [[Topics/]] - Ideas and backlog
- 📚 [[Definition]] - PM knowledge base

---

**Last Updated**: 2026-04-05 14:47 UTC  
**Version**: 1.0  
**Vault Status**: Ready for Integration
