# Ollama Tuning & Performance Guide

## Current Setup
- **Model**: Mistral 7B (4.4 GB)
- **Service**: Running on http://localhost:11434
- **Integration**: PMBuddy auto-detects Ollama at startup

## Performance Tuning

### 1. Memory Optimization

Check current memory usage:
```bash
ps aux | grep "ollama serve"
```

Adjust context window (lower = faster, less memory):
```bash
export OLLAMA_NUM_THREAD=4        # Use 4 CPU cores
export OLLAMA_NUM_GPU=1            # Use GPU if available
ollama serve
```

### 2. Response Speed Tuning

If responses are slow, adjust these parameters in your system prompt:

**Fast Mode** (1-2 sec responses):
- Temperature: 0.3 (less creative, more deterministic)
- Top-p: 0.8 (narrow token choices)
- Top-k: 40 (limit vocabulary)

**Balanced Mode** (2-5 sec responses) - **[CURRENT]**
- Temperature: 0.7 (default)
- Top-p: 0.9
- Top-k: 50

**Creative Mode** (5+ sec responses):
- Temperature: 0.9 (more creative)
- Top-p: 0.95
- Top-k: 100

### 3. Monitoring Ollama Health

Check service status:
```bash
curl -s http://localhost:11434/api/tags | jq '.models[].name'
```

Test inference speed:
```bash
curl -s http://localhost:11434/api/generate \
  -d '{"model":"mistral:latest","prompt":"Hello","stream":false}' | jq '.response'
```

View logs:
```bash
tail -f /tmp/ollama.log
```

### 4. Model Selection Trade-offs

| Model | Size | Speed | Quality | Best For |
|-------|------|-------|---------|----------|
| Mistral | 4 GB | Fast | Good | **[CURRENT]** PM advice, balanced |
| Neural-Chat | 4 GB | Fast | Good | Conversation-focused |
| Orca-Mini | 2 GB | Very Fast | Fair | Resource-constrained |
| Llama2 | 3.8 GB | Medium | Excellent | High-quality analysis |

### 5. Mac-Specific Optimizations

**Metal GPU Acceleration** (speeds up by 3-5x):
```bash
# Already enabled by default on ARM64 Macs
# Verify:
ollama list
```

**Memory Management**:
```bash
# Set max memory for Ollama
export OLLAMA_MAX_MEMORY=8000000000  # 8GB
```

**CPU Thread Tuning**:
```bash
# For M1/M2/M3 Macs - use fewer threads for latency
export OLLAMA_NUM_THREAD=2  # Lower latency
# vs
export OLLAMA_NUM_THREAD=8  # Higher throughput
```

## Performance Benchmarks

After tuning, test with:

```bash
# Test response time
curl -s http://localhost:11434/api/generate \
  -d '{
    "model":"mistral:latest",
    "prompt":"As a PM, how do I prioritize features?",
    "stream":false
  }' | jq '.response' | head -c 200
```

## Troubleshooting

### Ollama won't start
```bash
# Check if port 11434 is in use
lsof -i :11434

# Verify process is running
pgrep -f "ollama serve"
```

### Slow responses
- Check CPU/memory usage: `top` or Activity Monitor
- Reduce context window
- Switch to lighter model (Orca-Mini)
- Increase `OLLAMA_NUM_THREAD`

### Out of memory
- Use smaller model (Orca-Mini 2GB)
- Reduce batch size in internal/agent/llm.go
- Set OLLAMA_MAX_MEMORY limit

### Connection refused
- Verify Ollama is running: `curl http://localhost:11434/api/tags`
- Check firewall (should be localhost only)
- Restart service via launchctl or direct command

## Recommended Configuration (Mac)

Add to ~/.zprofile or ~/.bash_profile:
```bash
export OLLAMA_NUM_THREAD=4
export OLLAMA_MAX_MEMORY=6000000000  # 6GB
alias ollama-start='nohup ollama serve > ~/ollama.log 2>&1 &'
alias ollama-status='curl -s http://localhost:11434/api/tags | jq .'
```

Then start with: `ollama-start`

## Next Steps
1. ✅ Ollama installed and running
2. ✅ Mistral model downloaded
3. ✅ PMBuddy connected
4. **→ Monitor performance for 1-2 weeks**
5. **→ Adjust parameters based on real usage**
6. **→ Consider alternative models if needed**
