# Ollama Setup Guide for PMBuddy

## Overview

PMBuddy integrates with **Ollama** for local, private LLM-powered PM advice. Ollama runs language models locally on your machine with zero cloud dependency.

## Installation

### macOS

```bash
# Download and install Ollama
# Visit: https://ollama.ai/download/mac

# Or use Homebrew
brew install ollama

# Verify installation
ollama --version
```

### Windows

```bash
# Download installer from: https://ollama.ai/download/windows
# Run the installer and follow prompts
# Verify installation
ollama --version
```

### Linux

```bash
curl https://ollama.ai/install.sh | sh
```

## Running Ollama

### Start Ollama Service

```bash
# Start Ollama (runs on http://localhost:11434)
ollama serve
```

Keep this terminal open - Ollama will run in the background.

## Choosing a Model

### Quick Start (Recommended)

```bash
# Download and run Mistral (7B, lightweight)
# ~4GB download, ~5GB memory
ollama pull mistral
ollama run mistral
```

### Available Models (by size/speed)

| Model | Size | Memory | Speed | Quality |
|-------|------|--------|-------|---------|
| **mistral** | 4GB | 5GB | Fast | Good |
| neural-chat | 4GB | 5GB | Fast | Good |
| orca-mini | 2GB | 3GB | Very Fast | Fair |
| llama2 | 3.8GB | 5GB | Medium | Excellent |
| dolphin-mixtral | 26GB | 30GB | Slow | Excellent |

### Download a Model

```bash
# Lightweight (recommended for PMBuddy)
ollama pull mistral

# Balanced
ollama pull neural-chat
ollama pull llama2

# Best quality (requires more resources)
ollama pull dolphin-mixtral
```

### Test the Model

```bash
# Starts interactive chat with the model
ollama run mistral

# In the chat, type:
# "What are best practices for product prioritization?"
# Type /bye to exit
```

## Configuring PMBuddy

### Auto-Detection

PMBuddy automatically detects Ollama at `http://localhost:11434`. Once Ollama is running with a model loaded, PMBuddy will:

1. Detect Ollama is available ✓
2. Switch from Mock LLM to Real LLM ✓
3. Use detected preferences for personalization ✓

### Manual Configuration (Optional)

To use a different Ollama URL or model, modify the advisor initialization:

```go
// In internal/agent/advisor.go - NewPMAdvisor()
ollamaClient := NewOllamaClient(
    "http://your-ollama-server:11434",
    "your-model-name",
)
```

## Running PMBuddy with Ollama

### Step 1: Start Ollama (in separate terminal)

```bash
ollama serve
```

### Step 2: Run PMBuddy

```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy
```

### Step 3: Check Status

```
You: status

📊 Session Status:
  ...
  AI Engine: Ollama (mistral)  ← Real LLM!
```

## Troubleshooting

### "Ollama not found"

```bash
# Make sure Ollama is running in another terminal
ollama serve

# Check if it's accessible
curl http://localhost:11434/api/tags

# If this fails, Ollama isn't running or isn't installed
```

### Model Download Fails

```bash
# Check your internet connection
# Models are 2-30GB depending on choice

# Download smaller model first
ollama pull mistral

# Check downloaded models
ollama list
```

### Slow Responses

- You may not have enough RAM
- Try smaller model: `ollama pull orca-mini`
- Increase system resources if possible

### GPU Acceleration

If you have an NVIDIA GPU:

```bash
# Ollama will auto-detect CUDA
# Just ensure NVIDIA drivers are installed
nvidia-smi

# For Mac with Apple Silicon, GPU acceleration is automatic
```

## Model Recommendations for PMBuddy

### Best for PM Work (Balanced)
```bash
ollama pull mistral
ollama pull neural-chat
```

### Best Quality (Higher Resources)
```bash
ollama pull llama2
ollama pull dolphin-mixtral
```

### Fastest (Lower Resources)
```bash
ollama pull orca-mini
```

## Performance Tips

1. **Keep Ollama running** - Don't stop it between PMBuddy sessions
2. **Use fast model** - Mistral is optimal for PM tasks
3. **Close other apps** - Free up RAM for better performance
4. **SSD recommended** - Faster model loading

## Advanced: Custom System Prompt

The system prompt is built dynamically based on your preferences:

```
You are PMBuddy, an expert PM advisor.
The user prefers a data-driven approach.
They think bottom-up.
They prioritize users focus.
Provide practical, actionable advice tailored to their style.
```

PMBuddy automatically customizes this based on detected preferences!

## Uninstalling Ollama

```bash
# macOS
rm -rf /Applications/Ollama.app
rm -rf ~/.ollama

# Remove from brew
brew uninstall ollama

# Linux
curl -fsSL https://ollama.ai/install.sh | sh -s -- uninstall
```

## Support

- **Ollama Docs**: https://github.com/jmorganca/ollama
- **Model Library**: https://ollama.ai/library
- **GitHub Issues**: https://github.com/jmorganca/ollama/issues

---

**Next Steps:**
1. Install Ollama
2. Download a model (`ollama pull mistral`)
3. Start Ollama (`ollama serve`)
4. Run PMBuddy and watch it auto-detect real AI! 🚀
