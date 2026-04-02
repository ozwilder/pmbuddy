package main

import (
"fmt"
"context"
"time"
"pmbuddy/internal/agent"
"pmbuddy/internal/knowledge"
)

func main() {
// Create LLM client
llm := agent.NewOllamaClient("mistral:latest")

// Test direct LLM call
fmt.Println("Testing direct Ollama call...")
fmt.Println("Checking availability:", llm.IsAvailable())

start := time.Now()
response, err := llm.GenerateResponse("What is product management?", []agent.Message{})
elapsed := time.Since(start)

if err != nil {
fmt.Printf("Error: %v\n", err)
return
}

fmt.Printf("Response time: %dms\n", elapsed.Milliseconds())
fmt.Printf("Response:\n%s\n", response[:min(200, len(response))])
}

func min(a, b int) int {
if a < b {
return a
}
return b
}
