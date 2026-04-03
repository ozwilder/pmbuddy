#!/usr/bin/env python3
"""
Simple voice capture utility for PMBuddy
Uses macOS native speech recognition via AppleScript
"""

import subprocess
import sys

def capture_voice_with_applescript():
    """Use AppleScript to capture voice via Dictation"""
    applescript = '''
tell application "System Events"
    set frontmost to true
    set recognizedText to ""
    try
        set recognizedText to text returned of (display dialog "🎤 Speak your question:" default answer "" buttons {"Cancel", "OK"} default button 2 with icon note)
    on error
        set recognizedText to ""
    end try
    return recognizedText
end tell
'''
    
    try:
        result = subprocess.run(['osascript', '-e', applescript], 
                              capture_output=True, 
                              text=True, 
                              timeout=30)
        if result.returncode == 0:
            text = result.stdout.strip()
            if text:
                print(text)
                return 0
        return 1
    except Exception as e:
        return 1

def capture_voice_with_dictation():
    """Use macOS Dictation"""
    applescript = '''
set recognizedText to ""
try
    set recognizedText to text returned of (display dialog "🎤 Speak now (macOS Dictation):" default answer "" buttons {"Cancel", "OK"} default button 2 with icon note)
on error
    set recognizedText to ""
end try
return recognizedText
'''
    
    try:
        result = subprocess.run(['osascript', '-e', applescript], 
                              capture_output=True, 
                              text=True, 
                              timeout=30)
        if result.returncode == 0:
            text = result.stdout.strip()
            if text:
                print(text)
                return 0
        return 1
    except Exception as e:
        return 1

if __name__ == '__main__':
    # Try AppleScript method
    sys.exit(capture_voice_with_applescript())
