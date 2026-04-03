# 🎤 PMBuddy with macOS Native Voice - Quick Start

## What You're Getting

✅ **No additional software needed** - uses macOS built-in "say" command
✅ **Full AI responses** - same Mistral 7B via Ollama
✅ **Voice output** - PMBuddy speaks responses back to you
✅ **Type fallback** - type if you prefer
✅ **Zero setup** - works immediately on macOS

---

## How to Start

### Terminal 1: Ensure Ollama is Running
```bash
ollama serve
```

Or it may already be running. Check:
```bash
curl http://localhost:11434/api/tags
```

### Terminal 2: Launch PMBuddy with Voice
```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --voice-native
```

That's it! 🎉

---

## What You'll See

```
╔════════════════════════════════════════════════════════════════╗
║                      PMBuddy - PM Advisor                      ║
║            Your AI-powered Product Manager consultant          ║
╚════════════════════════════════════════════════════════════════╝

🎤 Voice interface enabled (macOS Native)

PMBuddy: Welcome to PMBuddy! I'm your AI-powered Product Manager advisor...

🎙️ Speak your questions (or type for fallback)
Commands: 'status', 'help', 'clear', 'exit'

🎤 You (type or say):
```

---

## How to Use

### Type Your Questions
```
🎤 You (type or say): What's the best way to prioritize features?

PMBuddy: As a data-driven PM advisor, here are the key frameworks...

(macOS "say" command speaks the response back to you)

🎤 You (type or say):
```

### Commands

| Command | Effect |
|---------|--------|
| Type a question | Sends to AI, response is spoken |
| `status` | Shows what you've learned |
| `help` | Shows available commands |
| `clear` | Clears conversation history |
| `exit` | Leave PMBuddy |

---

## How It Works

```
You type:           "How do I structure go-to-market?"
         ↓
PMBuddy processes:  Sends to Ollama/Mistral 7B
         ↓
AI generates:       "Here are 5 key GTM frameworks..."
         ↓
macOS says:         Speaks the response using "say" command
         ↓
You hear:           Complete AI response in voice
```

---

## Try These Questions

```
What's the best way to prioritize features?
How do I structure a go-to-market plan?
What metrics should I track as a PM?
How do I handle stakeholder conflicts?
What does great product strategy look like?
```

Then type: `status`

---

## Tips

### Volume Control
- Adjust system volume to control PMBuddy voice loudness
- System Preferences → Sound → Output

### Voice Speed
PMBuddy uses the default macOS voice. To customize:
- System Preferences → Accessibility → Spoken Content
- Adjust speaking rate there

### Disable Voice Output
If you don't want PMBuddy to speak (just want text):
```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy
```

Use this instead for text-only mode.

---

## Troubleshooting

### "Command not found"
```bash
# Make sure you're in the right directory
cd /Users/ozwilder/Development/PMBuddy

# Make sure binary exists
ls -la pmbuddy

# Rebuild if needed
go build -o pmbuddy ./cmd/pmbuddy
```

### No sound coming out
- Check system volume is not muted
- Test: `say "Hello"`
- Verify speakers are connected

### Slow responses
This is normal - Mistral 7B takes 2-4 seconds per response.
See OLLAMA_TUNING.md for optimization tips.

### Ollama connection error
```bash
# Check Ollama is running
curl http://localhost:11434/api/tags

# If error, start Ollama
ollama serve
```

---

## Example Session

```
$ cd /Users/ozwilder/Development/PMBuddy
$ ./pmbuddy --voice-native

🎤 Voice interface enabled (macOS Native)

PMBuddy: Welcome to PMBuddy! I'm your AI-powered Product Manager advisor...

🎤 You (type or say): How should I think about user research?

PMBuddy: As an intuitive PM advisor, user research is fundamental to...
(🔊 macOS voice speaks this)

🎤 You (type or say): status

📊 Session Status:
  Messages: 2
  AI Engine: Ollama (mistral)
  
🧠 Learning Stats:
  Topics Tracked: user-research
  Preferences Detected: 1

🎤 You (type or say): exit

PMBuddy: Thanks for chatting! Goodbye! 👋
(🔊 macOS voice says goodbye)
```

---

## What's Different from Text Mode?

| Feature | Text | Voice-Native |
|---------|------|--------------|
| Input | Type questions | Type questions |
| Output | Read text | Hear via speakers |
| Setup | 0 min | 0 min |
| Ollama needed | Yes | Yes |
| AI quality | Same | Same |
| Learning system | Yes | Yes |
| Commands | Type | Type or speak |

---

## Next Steps

**Now:** Use voice mode as is!

**Later:** If you want true speech-to-text (not just type):
```bash
brew install handy
./pmbuddy --voice
```

This would let you speak questions directly (voice input + voice output).

---

## Enjoy!

You now have a full AI PM advisor with voice output!

```bash
./pmbuddy --voice-native
```

Happy product managing! 🚀

---

*Last Updated: 2026-04-02*
