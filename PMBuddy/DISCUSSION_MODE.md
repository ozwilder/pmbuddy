# 🎤 PMBuddy Discussion Mode - Full Voice Conversation

## What is Discussion Mode?

**True two-way voice conversation:**
- 🗣️ **You speak** a PM question
- 👂 **PMBuddy listens** (macOS speech recognition)
- 🤖 **PMBuddy thinks** (Ollama/Mistral 7B)
- 🔊 **PMBuddy speaks** the response back to you

No typing required - pure voice conversation!

---

## How to Start

### Terminal 1: Ensure Ollama is Running
```bash
ollama serve
```

Check if already running:
```bash
curl http://localhost:11434/api/tags
```

### Terminal 2: Launch Discussion Mode
```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --discussion
```

That's it! 🎉

---

## How It Works

### Step 1: Dialog Box Appears
When you run `--discussion`, macOS shows a dialog box:
```
┌─────────────────────────────┐
│ Speak now:                  │
│ [                        ]  │
│     [Cancel] [OK]           │
└─────────────────────────────┘
```

### Step 2: Speak Your Question
The Mac is listening! Speak clearly:
- "What's the best way to prioritize features?"
- "How do I structure a go-to-market plan?"
- "What metrics should I track?"

### Step 3: PMBuddy Processes
Behind the scenes:
```
Your voice input → Text transcription
     ↓
  Ollama/Mistral 7B processes
     ↓
  AI generates response
```

### Step 4: Hear the Response
PMBuddy speaks back using macOS voice:
```
🔊 "As a data-driven PM, here are the key frameworks..."
```

### Step 5: Next Question
The dialog appears again - keep the conversation going!

---

## Example Session

```
$ cd /Users/ozwilder/Development/PMBuddy
$ ./pmbuddy --discussion

🎤 Discussion Mode Started

PMBuddy: Welcome to PMBuddy...

🎙️ Listening mode active
Speak a PM question...

[Dialog box: Speak now]
🎤 You: "How do I prioritize features?"

⏳ PMBuddy is thinking...

PMBuddy: As a PM, prioritization is about balancing impact, 
effort, and strategic alignment. Here are the frameworks I recommend...

🔊 (macOS voice speaks the response)

[Dialog box: Speak now]
🎤 You: "What about user research?"

⏳ PMBuddy is thinking...

🔊 (macOS voice speaks response about user research)
```

---

## Commands in Discussion Mode

| Command | Effect |
|---------|--------|
| Speak question | Processing and response |
| `exit` (speak or type) | Leave PMBuddy |
| `status` (speak or type) | Show learning analytics |
| `help` (speak or type) | Show available commands |
| `clear` (speak or type) | Clear conversation history |

---

## Tips for Best Results

### Microphone Setup
1. Plug in external microphone (optional, improves accuracy)
2. Reduce background noise
3. Speak clearly and at normal volume
4. Speak at a steady pace

### Voice Quality
- Speak full sentences, not fragments
- Pause between thoughts
- Enunciate clearly (especially acronyms like "KPI")
- Use standard English pronunciation

### System Settings
Adjust macOS speech recognition:
- System Preferences → Accessibility → Spoken Content
- Adjust recognition language if needed
- Adjust voice output speed/volume

---

## Fallback to Text

If speech recognition doesn't work:
- The system falls back to text input automatically
- Type your question instead
- PMBuddy still speaks the response
- Perfect hybrid mode!

---

## Comparison: All Modes

| Feature | Text | Voice-Native | Discussion |
|---------|------|--------------|-----------|
| Input | Type | Type | **Speak** ✨ |
| Output | Read | **Hear** | **Hear** |
| Setup | 0 min | 0 min | 0 min |
| Best for | Detailed writing | Quick questions | Natural conversation |

---

## What Discussion Mode Learns

Like all PMBuddy modes, discussion mode learns:

### Topics Tracked
- Prioritization
- Strategy
- Roadmap
- Metrics
- User research
- Go-to-market
- Features

### Preferences Detected (6D)
- Approach: top-down vs bottom-up
- Thinking: data-driven vs intuitive
- Style: formal vs casual
- Decisions: collaborative vs individual
- Focus: big picture vs details
- Timeline: long-term vs short-term

Example after a conversation:
```
🎯 Your PM Style (65% confidence):
  Approach: bottom-up
  Thinking: data-driven
  Style: collaborative
  Decisions: collaborative
```

---

## Sample Questions to Try

```
"How should I think about my role as a PM?"

"What's the best framework for prioritizing features?"

"How do I structure a go-to-market plan?"

"What key metrics should I track?"

"How do I handle stakeholder conflicts?"

"What does great product strategy look like?"

"How do I conduct effective user research?"

"What's the difference between strategy and tactics?"
```

Then say: "status" to see what was learned!

---

## Troubleshooting

### "Dialog box doesn't appear"
- Make sure you're using a Mac with speech recognition
- Try updating macOS to latest version
- Restart Terminal and try again

### "Speech not recognized"
- Speak louder and more clearly
- Reduce background noise
- Slow down your speech
- Try shorter sentences

### "Ollama connection error"
```bash
# Check Ollama
curl http://localhost:11434/api/tags

# Start if not running
ollama serve
```

### "No audio output"
- Check system volume is not muted
- Verify speakers are connected
- Test: `say "Hello"`

### "Responses are slow"
This is normal! Mistral 7B takes 2-4 seconds per response.
See OLLAMA_TUNING.md for optimization.

---

## Advanced Tips

### Use in Team Meetings
Leave discussion mode running while taking notes. PMBuddy becomes your virtual PM advisor during meetings!

### Brain Dump Mode
Start a conversation and just keep talking. PMBuddy learns your PM philosophy as you explain it.

### Preference Building
The more you interact, the more PMBuddy learns your style. After 10-20 interactions, responses become highly personalized.

### Status Tracking
Say "status" periodically to see:
- What topics you discuss most
- Your emerging PM preferences
- Confidence in preference detection

---

## Workflow Suggestion

### Daily PM Workflow
1. Morning: `./pmbuddy --discussion`
2. Ask questions about priorities/decisions
3. Get instant, personalized PM advice
4. End with `status` to see learning
5. Continue with your day

### Weekly Review
1. Launch: `./pmbuddy --discussion`
2. Ask: "How should I approach next week?"
3. Ask: "What did we learn this week?"
4. Ask: "What should I prioritize?"

---

## The Full Experience

**You get:**
- ✅ Natural voice conversation
- ✅ Instant AI PM advice
- ✅ Learning your preferences
- ✅ Personalized responses
- ✅ 100% local (no cloud)
- ✅ Private conversations
- ✅ Production-grade AI (Mistral 7B)

---

## Quick Reference

```bash
# Start discussion mode
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --discussion

# Or with one-liner
cd /Users/ozwilder/Development/PMBuddy && ./pmbuddy --discussion
```

---

## Next Steps

**Start now:**
```bash
./pmbuddy --discussion
```

**Alternative modes if you prefer:**
```bash
./pmbuddy                  # Text only
./pmbuddy --voice-native   # Type, hear responses
./pmbuddy --discussion     # Speak, hear responses (full voice)
```

---

## Questions?

Check the other guides:
- `START_HERE.md` - Quick overview
- `README_OLLAMA.md` - Full system overview
- `VOICE_NATIVE_QUICKSTART.md` - Voice output mode
- `VOICE_OPTIONS.md` - All voice options

---

**Ready to have a voice conversation with your AI PM advisor?**

```bash
./pmbuddy --discussion
```

Let's talk product management! 🚀

---

*Last Updated: 2026-04-02*
