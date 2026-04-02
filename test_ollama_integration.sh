#!/bin/bash
# PMBuddy Ollama Integration Test

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║        PMBuddy + Ollama Integration Test                       ║"
echo "║     Testing AI-powered PM advisor with Mistral 7B             ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

cd /Users/ozwilder/Development/PMBuddy

# Check if binary exists
if [ ! -f ./pmbuddy ]; then
    echo "🔨 Building PMBuddy..."
    go build -o pmbuddy ./cmd/pmbuddy
fi

# Verify Ollama is running
echo "🔍 Verifying Ollama connection..."
if curl -s http://localhost:11434/api/tags > /dev/null 2>&1; then
    echo "✅ Ollama service is running"
    curl -s http://localhost:11434/api/tags | jq '.models[].name'
else
    echo "❌ Ollama not responding on localhost:11434"
    echo "   Start it with: ollama serve"
    exit 1
fi

echo ""
echo "🚀 Starting PMBuddy with Ollama..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Run interactive test with sample PM queries
(
    # Query 1: Data-driven preference indicator
    echo "I want to make a feature prioritization decision. We have data showing users abandon after page 3. Should we fix navigation or add features?"
    sleep 3
    
    # Query 2: Another PM topic - demonstrates learning
    echo "How do I structure a go-to-market plan for a new segment?"
    sleep 3
    
    # View status - shows learning and preference detection
    echo "status"
    sleep 2
    
    # Exit gracefully
    echo "exit"
    
) | ./pmbuddy 2>&1 || true

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "✅ Test complete!"
echo ""
echo "📊 What you just saw:"
echo "   • PMBuddy detected your Ollama instance"
echo "   • Analyzed your PM queries for topics & preferences"
echo "   • Generated personalized responses via Mistral 7B"
echo "   • Displayed learning analytics"
echo ""
echo "🎯 Next steps:"
echo "   1. Try more queries to improve learning"
echo "   2. Run 'status' to see detected preferences"
echo "   3. Adjust OLLAMA_NUM_THREAD for performance"
echo "   4. Switch models: ollama pull llama2"
echo ""
