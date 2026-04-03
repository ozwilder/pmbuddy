#!/bin/bash
# Use macOS native Dictation for voice input
# This is the simplest, most reliable method!

osascript << 'APPLESCRIPT'
on run
    activate
    set recognizedText to text returned of (display dialog ¬
        "🎤 Speak to your Mac or type:" ¬
        default answer "" ¬
        buttons {"Cancel", "OK"} ¬
        default button 2 ¬
        with icon note)
    return recognizedText
end run
APPLESCRIPT
