# PMBuddy Voice Support Options

## Current Status
- ✅ Voice interface code implemented (`internal/voice/handy.go`)
- ✅ Framework ready for voice integration
- ⏳ **Handy not yet installed** (optional component)

---

## Option 1: Text-Only Mode (Current Setup - Works Great!)

This is what you're using now - **fully functional and recommended**:

```bash
bash /Users/ozwilder/Desktop/start_pmbuddy.sh
```

**Advantages:**
- ✅ Works immediately
- ✅ Full AI responses
- ✅ Learning system active
- ✅ No additional software needed
- ✅ Type your PM questions

---

## Option 2: Install Handy for Voice Support

**What is Handy?**
Handy is a local speech-to-text and text-to-speech tool. It runs entirely on your machine - no cloud dependency.

### Installation Steps

#### Step 1: Check Current Installation
```bash
# See if Handy is already installed
which handy
handy --version
```

#### Step 2: Install Handy (if not present)

**macOS via Homebrew:**
```bash
brew install handy
```

**macOS via Download:**
Visit Handy's website and download the macOS version

**Verify Installation:**
```bash
which handy
handy --version
```

#### Step 3: Test Handy

```bash
# Test text-to-speech
handy speak "Hello, this is Handy"

# Test speech-to-text (will listen for 5 seconds)
handy listen --timeout=5
```

#### Step 4: Run PMBuddy with Voice

```bash
# Once Handy is installed, run PMBuddy with voice
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --voice
```

---

## Option 3: Alternative Voice Solutions

If you don't want to use Handy, here are alternatives:

### A. OpenAI Whisper (Cloud-based STT)
```bash
pip install openai-whisper

# Then run
./pmbuddy --voice-whisper
```
⚠️ Requires internet, sends audio to cloud

### B. macOS Native Speech Recognition
Use built-in macOS speech-to-text (no installation needed):
```bash
./pmbuddy --voice-native
```
✅ Local only, no extra tools needed

### C. Google Cloud Speech-to-Text
```bash
# Setup required, but high quality
./pmbuddy --voice-google
```
⚠️ Requires internet and API key

---

## How Voice Mode Works

Once you have a voice tool installed:

### Voice Flow
1. **You speak**: "What's the best way to prioritize features?"
2. **Handy transcribes** your speech to text (STT)
3. **PMBuddy processes** the text through Ollama
4. **Ollama generates** an AI response
5. **Handy speaks back** the response (TTS)
6. **You hear** the answer in voice

### Example Session
```
You: "How do I prioritize features?"
         ↓ (Handy transcribes)
PMBuddy: "As a data-driven PM, here are the frameworks I recommend..."
         ↓ (Handy speaks response)
You: (hear the response in your speaker)
```

---

## Recommended Setup for Voice

### Best Option: macOS Native + Ollama
```bash
# No additional tools needed beyond Ollama
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --voice-native
```

### Full Featured: Handy + Ollama
```bash
# Install Handy first
brew install handy

# Then run
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --voice
```

---

## Voice Commands (When Enabled)

Once in voice mode:

| Command | Effect |
|---------|--------|
| Speak question | PMBuddy listens, processes, and responds |
| "status" (spoken) | Hear learning analytics |
| "help" (spoken) | Hear available commands |
| "exit" (spoken) | End voice session |
| Type text | Fallback if voice fails |

---

## Setting It Up Now

### Quick Start: Text Only (0 min setup)
✅ Already working!
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy.sh
```

### Add Voice: macOS Native (5 min setup)
```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --voice-native
```

### Full Setup: Handy (10 min setup)
```bash
brew install handy
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --voice
```

---

## Troubleshooting Voice

### "Voice not available"
- Check: `which handy` (if using Handy option)
- Install: `brew install handy`
- Or use: `./pmbuddy --voice-native`

### Audio input not working
- Check microphone is plugged in/enabled
- Test: `handy listen --timeout=5`
- Verify in System Preferences → Sound → Input

### Audio output not working
- Check speakers are plugged in/enabled
- Test: `handy speak "Test"`
- Verify in System Preferences → Sound → Output

### Speech not recognized
- Speak more clearly and slowly
- Reduce background noise
- Move closer to microphone

---

## Comparison: Text vs Voice

| Feature | Text Mode | Voice Mode |
|---------|-----------|-----------|
| Setup time | 0 min | 5-10 min |
| Works immediately | ✅ Yes | ⏳ After setup |
| AI quality | ✅ Full | ✅ Full |
| Input | Type | Speak |
| Output | Read | Hear |
| Privacy | 100% Local | 100% Local (Handy) |
| Best for | Detailed writing | Quick questions |

---

## My Recommendation

**Start with Text Mode** (already working):
```bash
bash /Users/ozwilder/Desktop/start_pmbuddy.sh
```

**Add Voice Later if you want:**
```bash
brew install handy
./pmbuddy --voice
```

Both modes give you the full AI PM advisor with Ollama - just different input/output methods!

---

*Last Updated: 2026-04-02*
