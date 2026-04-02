# 📱 PMBuddy - Complete Voice & Text Guide

## Your AI PM Advisor has 4 Modes

---

## 🖥️ Mode 1: Text Only

### Command
```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy
```

### Experience
- ✅ Type your PM questions
- ✅ Read AI responses
- ✅ Full conversation history
- ✅ Learning system active

### When to Use
- Writing detailed PM thoughts
- Documentation work
- Team collaboration (share text)

### Setup Time: **0 minutes**

---

## 🎙️ Mode 2: Voice Output Only (Voice-Native)

### Command
```bash
./pmbuddy --voice-native
```

### Experience
- ✅ Type your PM questions
- ✅ **Hear** AI responses via macOS voice
- ✅ Reading not required
- ✅ Learning system active

### When to Use
- Quick PM questions while working
- Hands-busy situations
- Multitasking (hear responses)

### Setup Time: **0 minutes**

---

## 🗣️ Mode 3: Full Voice Conversation (Discussion)

### Command
```bash
./pmbuddy --discussion
```

### Experience
- ✅ **Speak** your PM questions
- ✅ **Hear** AI responses via macOS voice
- ✅ Natural conversation flow
- ✅ Learning system active

### When to Use
- **Natural conversation** (primary use)
- Team meetings & brainstorming
- Hands-free PM advising
- Building personalized PM style

### Setup Time: **0 minutes**

---

## 🎧 Mode 4: Full Voice + Handy (Future)

### Command
```bash
brew install handy
./pmbuddy --voice
```

### Experience
- Speak questions (Handy STT)
- Hear responses (Handy TTS)
- Higher quality voice recognition
- Professional voice options

### When to Use
- Seeking higher speech recognition accuracy
- After building preference profile

### Setup Time: **5-10 minutes** (install Handy)

---

## 🎯 Recommended: Use Mode 3 (Discussion)

**Why?**
- ✅ Most natural experience
- ✅ No setup needed
- ✅ Full voice I/O
- ✅ Best for learning PM style
- ✅ 100% local & private
- ✅ Same AI quality as text modes

**Start with:**
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
```

---

## 📊 Quick Comparison

| Feature | Text | Voice Output | Discussion | Handy |
|---------|------|--------------|-----------|-------|
| Input | Type | Type | **Speak** | **Speak** |
| Output | Read | **Hear** | **Hear** | **Hear** |
| Setup | 0 min | 0 min | 0 min | 5-10 min |
| AI Quality | ✅ | ✅ | ✅ | ✅ |
| Learning | ✅ | ✅ | ✅ | ✅ |
| Best for | Writing | Quick Q&A | Conversation | Prof. speech |

---

## 🎓 All Modes Share

### AI Engine
- **Mistral 7B** via local Ollama
- 2-4 second response time
- Production-grade advice

### Learning System
Tracks your PM style across:
- **7 PM Topics**: Prioritization, strategy, roadmap, metrics, user research, GTM, features
- **6 Dimensions**: Approach, thinking, style, decisions, focus, timeline

### Privacy
- ✅ 100% local processing
- ✅ No cloud services
- ✅ No data transmission
- ✅ Private conversations

### Personalization
- ✅ Learns your PM preferences
- ✅ Adapts response style
- ✅ Improves over time
- ✅ Shows confidence levels

---

## 🚀 Getting Started Paths

### Path A: Quick Conversation (Recommended)
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
```
**Time:** 30 seconds to first response

### Path B: Text Mode
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy.sh
```
**Time:** 15 seconds to first response

### Path C: Voice Output
```bash
cd /Users/ozwilder/Development/PMBuddy && ./pmbuddy --voice-native
```
**Time:** 15 seconds to first response

### Path D: Manual
```bash
# Terminal 1
ollama serve

# Terminal 2
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --discussion
```

---

## 💡 Sample Questions for Each Mode

### Text Mode
Good for detailed, multi-line responses:
- "Walk me through a product launch checklist"
- "What are the top 10 PM frameworks?"
- "Create a prioritization matrix"

### Voice Output Mode
Good for quick answers:
- "How should I think about metrics?"
- "What's a good feature prioritization approach?"
- "How do I handle stakeholder conflicts?"

### Discussion Mode
Good for natural conversation:
- "How do I prioritize features?"
- "Tell me about your experience with that"
- "What did we learn? What should we do differently?"

---

## 🎙️ Pro Tips

### Maximize Learning (Works in all modes)
1. Ask diverse PM questions
2. Mention your approaches/preferences
3. Run `status` after 5-10 questions
4. Watch preferences get detected
5. Responses become more personalized

### Discussion Mode Specific
1. Speak in complete sentences
2. Use natural, conversational language
3. Don't rush - PMBuddy will wait
4. Ask follow-ups to dig deeper
5. Say "status" to see what's learned

### Voice Quality
- Microphone: Built-in works, external is better
- Noise: Reduce background noise
- Clarity: Speak normally, not slowly
- Volume: System volume at 50-75%

---

## 🔄 Workflow Examples

### Daily PM Standoff
```bash
./pmbuddy --discussion
```
Questions:
1. "What should I prioritize this week?"
2. "What risks should I consider?"
3. "How should I approach stakeholders?"
4. "status"

### Weekly Strategy Session
```bash
./pmbuddy
```
Use text mode to:
1. Write out roadmap options
2. Get PMBuddy's analysis
3. Document decisions
4. Share with team

### Quick Advice
```bash
./pmbuddy --voice-native
```
Type:
1. "How do I handle feature requests?"
2. "What metrics matter most?"
3. Hear responses while working

---

## ✨ What You Have

✅ **Text Mode**: Full conversation with text I/O
✅ **Voice-Native Mode**: Type questions, hear responses
✅ **Discussion Mode**: Speak questions, hear responses ⭐ RECOMMENDED
✅ **Learning System**: Tracks 7 PM topics + 6 preferences
✅ **Personalization**: Adapts to your style
✅ **Privacy**: 100% local
✅ **AI Quality**: Mistral 7B
✅ **Documentation**: Complete guides for each mode

---

## 🎯 Next Steps

### Start Here (Recommended)
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh
```

### Then Explore
```bash
./pmbuddy --voice-native     # Try voice output
./pmbuddy                    # Try text mode
```

### Advanced (Optional)
```bash
brew install handy
./pmbuddy --voice            # Full Handy integration
```

---

## 📖 Documentation

| Guide | Topic |
|-------|-------|
| `START_HERE.md` | 2-minute quick start |
| `README_OLLAMA.md` | Full system overview |
| `DISCUSSION_MODE.md` | Full voice conversation guide |
| `VOICE_NATIVE_QUICKSTART.md` | Voice output mode |
| `VOICE_OPTIONS.md` | Comparison of all options |

---

## 🎉 You're Ready!

Choose a mode and start:

```bash
# Discussion (Recommended)
bash /Users/ozwilder/Desktop/start_pmbuddy_discussion.sh

# Text Only
bash /Users/ozwilder/Desktop/start_pmbuddy.sh

# Voice Output
cd /Users/ozwilder/Development/PMBuddy && ./pmbuddy --voice-native
```

**Your AI PM advisor is waiting!** 🚀

---

*Last Updated: 2026-04-02*
*All modes tested and working*
