# 🎉 PMBuddy - Complete Setup & Usage Guide

**Version**: 1.0
**Date**: April 2, 2026
**Status**: ✅ Production Ready

---

## 📋 Table of Contents

1. [What is PMBuddy?](#what-is-pmbuddy)
2. [Quick Start](#quick-start)
3. [All 4 Modes](#all-4-modes)
4. [How to Use](#how-to-use)
5. [Features](#features)
6. [Files & Structure](#files--structure)
7. [Troubleshooting](#troubleshooting)

---

## What is PMBuddy?

**Your personal AI-powered Product Manager advisor that listens and learns.**

PMBuddy is a standalone Go application that combines:
- **Ollama/Mistral 7B**: Production-grade AI running locally
- **Learning System**: Extracts PM topics from conversations
- **Preference Detection**: Learns your PM style (6 dimensions)
- **Personalization**: Adapts responses to your approach
- **Multiple Interfaces**: Text, voice input, voice output, or full discussion mode

---

## Quick Start

### The Absolute Fastest Way

```bash
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
```

Done! 🎉

- A dialog box appears
- Speak your PM question
- Hear the response
- Dialog appears again for next question

---

## All 4 Modes

### Mode 1: Text Only
```bash
./pmbuddy
```
- Type questions
- Read responses
- Full learning system
- **Best for:** Writing, documentation

### Mode 2: Voice Output (Voice-Native)
```bash
./pmbuddy --voice-native
```
- Type questions
- Hear responses via macOS
- Full learning system
- **Best for:** Multitasking, hands-busy

### Mode 3: Full Voice Discussion ⭐ RECOMMENDED
```bash
./pmbuddy --discussion
```
- Speak questions (macOS speech recognition)
- Hear responses via macOS voice
- Full learning system
- **Best for:** Natural conversation, building style profile

### Mode 4: Professional Voice (Handy)
```bash
brew install handy
./pmbuddy --voice
```
- Speak questions (Handy speech-to-text)
- Hear responses (Handy text-to-speech)
- Full learning system
- **Best for:** Higher quality voice recognition

---

## How to Use

### Discussion Mode (Recommended)

1. **Start the conversation:**
   ```bash
   bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
   ```

2. **Dialog box appears** asking you to speak

3. **Speak clearly:**
   - "How do I prioritize features?"
   - "What should my roadmap look like?"
   - "How do I conduct user research?"

4. **Listen to the response** (spoken via macOS)

5. **Dialog appears again** for next question

6. **Say "status"** to see what was learned

7. **Say "exit"** to leave

### Text Mode

1. **Start:**
   ```bash
   ./pmbuddy
   ```

2. **Type questions** at the prompt

3. **Read responses** in the terminal

4. **Type "status"** to see learning

5. **Type "exit"** to leave

### Voice Output Mode

1. **Start:**
   ```bash
   ./pmbuddy --voice-native
   ```

2. **Type questions**

3. **Hear responses** spoken by macOS

4. **Type commands** as needed

---

## Features

### AI Engine
- **Model**: Mistral 7B (quantized, runs locally)
- **Source**: Ollama (local LLM framework)
- **Response Time**: 2-4 seconds typically
- **Quality**: Production-grade PM advice

### Learning System

#### Topics Tracked (7 categories)
1. **Prioritization**: Feature prioritization, roadmapping
2. **Strategy**: Vision, positioning, business strategy
3. **Roadmap**: Release planning, timelines
4. **Metrics**: KPIs, success criteria, data analysis
5. **User Research**: Discovery, feedback, validation
6. **Go-to-Market**: Launch, messaging, positioning
7. **Features**: Product capabilities, requirements

#### Preferences Detected (6 dimensions)
- **Approach**: Top-down ↔ Bottom-up
- **Thinking**: Data-driven ↔ Intuitive
- **Style**: Formal ↔ Casual
- **Decisions**: Collaborative ↔ Individual
- **Focus**: Big Picture ↔ Details
- **Timeline**: Long-term ↔ Short-term

### Learning in Action

```
Question 1: "How do I prioritize features?"
  → Detects: Prioritization topic, Approach preference

Question 2: "I like data-driven analysis"
  → Detects: Data-driven thinking preference

Question 3: "What about roadmap?"
  → Detects: Roadmap topic, Long-term focus

Status command shows:
  ✓ Topics: Prioritization (2), Roadmap (1)
  ✓ Preferences: Data-driven (2), Approach unknown (1)
  ✓ Confidence building...
```

### Privacy & Security
- ✅ 100% local processing (no cloud)
- ✅ No data transmission
- ✅ Private conversations
- ✅ Runs on your machine only

---

## Files & Structure

```
/Users/ozwilder/Development/PMBuddy/
├── pmbuddy                              # Compiled binary
├── cmd/pmbuddy/main.go                  # CLI entry point
├── internal/
│   ├── agent/
│   │   ├── advisor.go                   # Core PM logic
│   │   ├── llm.go                       # Ollama integration
│   │   └── voice.go                     # Voice interface
│   ├── knowledge/
│   │   ├── knowledge.go                 # Learning extraction
│   │   └── preferences.go               # Preference detection
│   └── storage/
│       └── storage.go                   # Persistence layer
├── docs/
│   ├── README.md                        # Project overview
│   ├── ARCHITECTURE.md                  # System design
│   ├── OLLAMA_SETUP.md                  # Ollama guide
│   ├── OLLAMA_TUNING.md                 # Performance guide
│   └── VOICE_SETUP.md                   # Voice integration
├── DISCUSSION_MODE.md                   # Full voice guide
├── VOICE_NATIVE_QUICKSTART.md           # Voice output guide
├── ALL_MODES_SUMMARY.md                 # All modes guide
└── START_HERE.md                        # Quick start

/Users/ozwilder/Desktop/
├── start_pmbuddy.sh                     # Text mode launcher
└── start_pmbuddy_discussion.sh          # Discussion mode launcher
```

---

## Troubleshooting

### "Ollama connection refused"
```bash
# Check if Ollama is running
curl http://localhost:11434/api/tags

# Start Ollama
ollama serve
```

### "Speech not recognized" (Discussion Mode)
- Speak louder and more clearly
- Reduce background noise
- Try shorter sentences
- Check System Preferences → Accessibility → Spoken Content

### "No audio output"
- Check system volume is up
- Verify speakers/headphones connected
- Test: `say "Hello"`

### "PMBuddy won't start"
```bash
# Make sure binary exists
ls -la /Users/ozwilder/Development/PMBuddy/pmbuddy

# Rebuild if needed
cd /Users/ozwilder/Development/PMBuddy
go build -o pmbuddy ./cmd/pmbuddy
```

### "Responses are slow"
- This is normal (Mistral takes 2-4 seconds)
- See OLLAMA_TUNING.md for optimization
- Consider lighter model: `ollama pull orca-mini`

---

## Example Conversations

### Discussion Mode Conversation

```
$ bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh

[Dialog: Speak now]
🎤 You: "How do I think about product strategy?"

⏳ PMBuddy is thinking...

PMBuddy: Product strategy is about aligning your vision...
🔊 (Spoken response)

[Dialog: Speak now]
🎤 You: "What about data-driven decisions?"

⏳ PMBuddy is thinking...

PMBuddy: Great question! Here are the key metrics framework...
🔊 (Spoken response)

[Dialog: Speak now]
🎤 You: "status"

📊 Session Status:
  Messages: 2
  AI Engine: Ollama (mistral)
  
🧠 Learning Stats:
  Topics: Strategy, Metrics
  Preferences Detected: 2 (data-driven, strategic focus)
  Confidence: 60%

[Dialog: Speak now]
🎤 You: "exit"

🔊 "Thanks for chatting. Goodbye!"
```

---

## Tips & Best Practices

### For Discussion Mode
1. **Speak clearly** - Normal volume, not rushed
2. **Use complete sentences** - "How do I prioritize?" works better than "prioritize?"
3. **Mention your preferences** - "I prefer data-driven" helps learning
4. **Ask follow-ups** - Deeper conversations build richer profiles
5. **Check status regularly** - See what's being learned

### For All Modes
1. **Diverse questions** - Ask about different PM areas (strategy, metrics, research, etc.)
2. **Be specific** - "How do I handle feature requests from executives?" > "What about features?"
3. **Build context** - Mention your role, company stage, product type when relevant
4. **Use commands** - Type/say "status" periodically to see progress
5. **Test preferences** - Ask questions in different styles to see adaptation

---

## Performance

### Response Times
- **First response**: 3-5 seconds (model warmup)
- **Typical**: 2-4 seconds
- **Complex queries**: Up to 10 seconds

### Resource Usage
- **Memory**: ~3.5 GB (Mistral 7B)
- **CPU**: 4 cores (configurable)
- **Model Size**: 4.4 GB

### Optimization
See `docs/OLLAMA_TUNING.md` for:
- CPU thread tuning
- Memory management
- Model selection
- Performance benchmarking

---

## What Gets Learned

### Session Level (During a conversation)
- Topics discussed
- Your approach/thinking style
- Your focus areas

### Multi-Session Level (Over time)
- Preference profile becomes more accurate
- Confidence increases (0% → 100%)
- PMBuddy adapts responses more

### Example Profile After 20 Interactions
```
🎯 Your PM Style (78% confidence):
  ├─ Approach: Bottom-up (+)
  ├─ Thinking: Data-driven (+++)
  ├─ Style: Collaborative (++)
  ├─ Decisions: Consensus-focused (+)
  ├─ Focus: Details-oriented (+)
  └─ Timeline: Quarterly planning (+)

Top Topics:
  1. Prioritization (8 mentions)
  2. Metrics (6 mentions)
  3. Roadmap (5 mentions)
```

---

## Next Steps

### Immediate (Next 5 minutes)
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
```

### Short-term (This week)
1. Have 5-10 PM conversations
2. Say "status" to see what's learned
3. Try different question types
4. Watch preferences build up

### Longer-term (Optional)
1. Explore other modes (text, voice-native)
2. Install Handy for professional voice: `brew install handy`
3. Use as daily PM advisor
4. Build comprehensive PM style profile

---

## Support & Documentation

| Topic | File |
|-------|------|
| Quick start | START_HERE.md |
| Full system | README_OLLAMA.md |
| Discussion mode | DISCUSSION_MODE.md |
| Voice output | VOICE_NATIVE_QUICKSTART.md |
| All modes | ALL_MODES_SUMMARY.md |
| Ollama setup | docs/OLLAMA_SETUP.md |
| Performance | docs/OLLAMA_TUNING.md |
| Architecture | docs/ARCHITECTURE.md |

---

## Quick Reference Card

```
┌─────────────────────────────────────────────────┐
│ PMBUDDY QUICK REFERENCE                         │
├─────────────────────────────────────────────────┤
│                                                 │
│ START NOW:                                      │
│ $ bash /Users/ozwilder/Desktop/\               │
│   start_pmbuddy_discussion.sh                   │
│                                                 │
│ TEXT MODE:                                      │
│ $ ./pmbuddy                                     │
│                                                 │
│ VOICE OUTPUT:                                   │
│ $ ./pmbuddy --voice-native                      │
│                                                 │
│ FULL VOICE:                                     │
│ $ ./pmbuddy --discussion                        │
│                                                 │
│ HANDY (Optional):                               │
│ $ brew install handy && ./pmbuddy --voice       │
│                                                 │
│ AI ENGINE: Mistral 7B via Ollama               │
│ RESPONSE TIME: 2-4 seconds                      │
│ LEARNING: 7 topics + 6 preferences              │
│ PRIVACY: 100% local                             │
│                                                 │
└─────────────────────────────────────────────────┘
```

---

## You're All Set! 🚀

Your AI PM advisor is ready.

**Start with:**
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
```

**Enjoy building your personalized PM advisor!**

---

*Last Updated: April 2, 2026*
*All features tested and verified*
*Status: Production Ready ✅*
