# PMBuddy CLI Discussion Mode Improvements

## Problem Addressed

The original discussion mode had several UX issues:
- ❌ AppleScript dialog boxes popped up unpredictably
- ❌ No indication that the system was listening
- ❌ Captured input wasn't visible to the user
- ❌ Status/flow got lost when "say" command was speaking
- ❌ Interference between user input and system responses

## Solution Implemented

### 1. **CLI-Only Interface**
- Removed AppleScript dialog boxes entirely
- All I/O now happens in the terminal
- Pure text-based conversation flow

### 2. **Clear Listening Indication**
```
🎙️  LISTENING... (Speak clearly or type):
```
- User sees this prompt and knows to provide input
- Visual feedback that system is ready

### 3. **Echo Captured Input**
```
✅ CAPTURED: "What metrics should I focus on?"
```
- Every input is immediately echoed back
- User confirmation that text was received
- Useful for verifying speech-to-text accuracy

### 4. **Clear Status Separation**
Each phase is separated by dashes for clarity:
```
────────────────────────────────────────────────────────────
⏳ PMBuddy is thinking...
────────────────────────────────────────────────────────────

🤖 PMBuddy Response:
────────────────────────────────────────────────────────────
[Full response text]
────────────────────────────────────────────────────────────
```

### 5. **Non-Blocking Speech Output**
- Speech synthesis now runs in goroutine (async)
- User can immediately type next question
- Doesn't block the listening loop
- `go exec.Command("say", responseText).Run()`

### 6. **Learning Display Integration**
Learning information is shown inline:
```
_[Learning: topics discussed - [metrics]]_
_[Adapted to your PM style: **Approach**: data-driven | **Thinking**: bottom-up]_
```

## Code Changes

### File: `cmd/pmbuddy/main.go`

#### Before
```go
// Old: AppleScript dialog box
appleScript := `display dialog "Speak now:" default answer ""`
cmd := exec.Command("osascript", "-e", appleScript)
output, err := cmd.Output()
input := strings.TrimSpace(string(output))
```

#### After
```go
// New: Pure CLI stdin
fmt.Print("\n🎙️  LISTENING... (Speak clearly or type): ")
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
input := strings.TrimSpace(scanner.Text())
```

#### Echo Captured
```go
fmt.Printf("✅ CAPTURED: \"%s\"\n", input)
```

#### Async Speech
```go
// Non-blocking: spawn goroutine for speech synthesis
go exec.Command("say", responseToSpeak).Run()
```

## User Experience Flow

### Example Interaction
```
🎙️  LISTENING... (Speak clearly or type): What metrics should I focus on?
✅ CAPTURED: "What metrics should I focus on?"

────────────────────────────────────────────────────────────
⏳ PMBuddy is thinking...

🤖 PMBuddy Response:
────────────────────────────────────────────────────────────
As a data-driven Product Manager...
[7-point response about metrics]
_[Learning: topics discussed - [metrics]]_
_[Adapted to your PM style: **Approach**: data-driven]_
────────────────────────────────────────────────────────────

🎙️  LISTENING... (Speak clearly or type):
```

## Benefits

✅ **Better Visual Flow**: Clear separation between input, thinking, and response  
✅ **User Confirmation**: Every input is echoed  
✅ **No Interference**: Speech doesn't block next input  
✅ **Learning Visible**: What was learned displayed inline  
✅ **Professional Look**: CLI-only, no dialog boxes  
✅ **Faster Interaction**: Async speech means instant prompt for next question  
✅ **Debugging**: Input echoed helps debug speech recognition issues  

## Testing

Test the new interface:
```bash
cd /Users/ozwilder/Development/PMBuddy
./pmbuddy --discussion
```

Then type (or voice input):
- `What metrics should I focus on?`
- `How should I prioritize features?`
- `status` (to see learning)
- `exit` (to quit)

## Implementation Details

- **Lines Modified**: cmd/pmbuddy/main.go (startDiscussionMode function)
- **Language**: Go
- **Concurrency**: Goroutines for non-blocking speech output
- **Terminal**: Works in any macOS terminal (zsh, bash, iTerm, etc.)

## Future Improvements (Optional)

1. Add real speech-to-text integration (Whisper API or local models)
2. Configurable listening timeout (auto-transcribe after pause)
3. Save conversation to file with metadata
4. Preference profile evolution tracking
5. Multi-turn context retention

---

**Status**: ✅ Complete and tested  
**Ready for**: Production use  
**Date**: 2026-04-02
