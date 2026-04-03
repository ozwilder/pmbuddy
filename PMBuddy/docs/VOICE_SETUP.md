# Voice Setup Guide for PMBuddy

## Overview

PMBuddy integrates with **Handy** for local speech-to-text and text-to-speech capabilities. This guide covers setup and configuration.

## Prerequisites

- macOS 10.12+ or Windows 11+
- Handy installed locally
- Audio input/output devices configured

## Installation

### Step 1: Verify Handy Installation

```bash
# Check if Handy is installed
which handy

# Check version
handy --version
```

If Handy is not found, install it from: [Handy Documentation](https://example.com/handy)

### Step 2: Configure Handy Path

Edit your PMBuddy configuration (future) or set environment variable:

```bash
export HANDY_PATH="/usr/local/bin/handy"
```

## Usage

### Enable Voice in PMBuddy

```bash
# Start with voice interface
./pmbuddy --voice

# Or with session management
./pmbuddy session start --voice
```

### Voice Commands

Once in voice mode:

- **Speak a question**: PMBuddy listens, transcribes, and responds
- **Voice response**: Answers are spoken back to you
- **Text fallback**: Type if voice input fails

## Voice Parameters

Configure voice behavior in Handy:

```go
// In your code:
voiceInterface := voice.NewHandyInterface()
voiceInterface.SetVoiceParams(
    "en-US",        // Language
    "default",      // Voice profile
    1.0,            // Speed (0.5-2.0)
)
```

## Handy API Reference

### Speech-to-Text
```bash
handy transcribe <audio_file>
handy listen --timeout=30
```

### Text-to-Speech
```bash
handy speak "Hello, World"
handy speak --voice=premium "Better voice output"
```

### Configuration
```bash
handy config --language en-GB
handy config --voice female
handy config --speed 1.2
```

## Troubleshooting

### Handy Not Found
```bash
# Verify installation
which handy

# Add to PATH if needed
export PATH="/usr/local/bin:$PATH"
```

### No Audio Input
- Check microphone permissions
- Test audio separately: `handy listen --timeout=5`
- Verify audio device in Handy settings

### Speech Not Recognized
- Speak clearly and slowly
- Reduce background noise
- Check language/accent settings

### No Audio Output
- Verify speakers are connected
- Test: `handy speak "Test"`
- Check system volume levels

## Advanced Configuration

### Custom Handy Path

Set in environment or config file:

```bash
export HANDY_PATH="/custom/path/to/handy"
```

### Performance Tuning

```go
// Increase timeout for slower speech
voiceInterface.ListenAndRespond(60) // 60 seconds

// Adjust voice speed
voiceInterface.SetVoiceParams("en-US", "default", 0.8) // Slower
```

## Security & Privacy

- All speech processing happens locally via Handy
- No audio is sent to external services
- Conversations stored locally only
- Enable encryption for sensitive data (future feature)

## Support

For Handy-specific issues, refer to: [Handy Docs](https://example.com/handy-support)
