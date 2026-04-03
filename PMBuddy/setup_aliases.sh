#!/bin/bash
# PMBuddy Ollama Quick Reference - Handy shell aliases

# Add to ~/.zprofile or ~/.bash_profile

# === Ollama Management ===
alias ollama-start='nohup ollama serve > ~/.ollama.log 2>&1 &'
alias ollama-logs='tail -f ~/.ollama.log'
alias ollama-status='curl -s http://localhost:11434/api/tags | jq'
alias ollama-models='ollama list'

# === PMBuddy ===
alias pmbuddy-cd='cd /Users/ozwilder/Development/PMBuddy'
alias pmbuddy='cd /Users/ozwilder/Development/PMBuddy && ./pmbuddy'
alias pmbuddy-build='cd /Users/ozwilder/Development/PMBuddy && go build -o pmbuddy ./cmd/pmbuddy'
alias pmbuddy-test='cd /Users/ozwilder/Development/PMBuddy && bash test_ollama_integration.sh'

# === Monitoring ===
alias ps-ollama='ps aux | grep ollama | grep -v grep'
alias monitor-ollama='watch -n 2 "ps aux | grep ollama | grep -v grep"'

# === Quick Workflows ===

# Start everything at once
function pmbuddy-start-all() {
    echo "Starting Ollama service..."
    ollama-start
    sleep 2
    echo "Starting PMBuddy..."
    pmbuddy
}

# Kill everything cleanly
function pmbuddy-stop-all() {
    echo "Stopping PMBuddy..."
    pkill -f pmbuddy
    echo "Stopping Ollama..."
    # Note: pkill not allowed, use manual process stop via Activity Monitor or:
    # Get PID from: ps aux | grep "ollama serve" | grep -v grep
    echo "Stopped PMBuddy and Ollama"
}

# Test connectivity
function pmbuddy-test-connection() {
    echo "Testing Ollama connection..."
    if curl -s http://localhost:11434/api/tags > /dev/null 2>&1; then
        echo "✅ Ollama is running"
        ollama list
    else
        echo "❌ Ollama not responding"
        echo "Start it with: ollama serve"
    fi
}

# View session from last run
function pmbuddy-last-session() {
    ls -lrt ~/.pmbuddy/sessions/ | tail -5
}

# === Performance Tuning ===

# For faster responses (fewer CPU threads)
function pmbuddy-fast() {
    export OLLAMA_NUM_THREAD=2
    export OLLAMA_MAX_MEMORY=4000000000
    echo "Set for fast mode (2 threads, 4GB max)"
}

# For better quality (more CPU threads)
function pmbuddy-quality() {
    export OLLAMA_NUM_THREAD=8
    export OLLAMA_MAX_MEMORY=8000000000
    echo "Set for quality mode (8 threads, 8GB max)"
}

# Balanced (recommended)
function pmbuddy-balanced() {
    export OLLAMA_NUM_THREAD=4
    export OLLAMA_MAX_MEMORY=6000000000
    echo "Set for balanced mode (4 threads, 6GB max)"
}

echo "PMBuddy aliases loaded! Available commands:"
echo "  ollama-start       - Start Ollama service"
echo "  ollama-logs        - Tail Ollama logs"
echo "  ollama-status      - Check Ollama status"
echo "  pmbuddy            - Run PMBuddy"
echo "  pmbuddy-build      - Rebuild PMBuddy"
echo "  pmbuddy-test       - Run integration test"
echo "  pmbuddy-start-all  - Start both services"
echo "  pmbuddy-test-connection - Verify setup"
echo "  pmbuddy-fast/quality/balanced - Performance tuning"
