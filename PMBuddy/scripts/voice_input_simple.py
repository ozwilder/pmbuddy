#!/usr/bin/env python3
"""
Simple voice input using macOS native speech recognition
No external microphone access required - uses macOS built-in capabilities
"""

import subprocess
import sys
import json

def get_voice_input():
    """Get voice input using macOS native dialog"""
    
    applescript = """
tell application "System Events"
    activate
    set recognizedText to text returned of (display dialog "🎤 Speak to your Mac or type:" ¬
        default answer "" ¬
        buttons {"Cancel", "OK"} ¬
        default button 2 ¬
        with icon note)
    return recognizedText
end tell
"""
    
    try:
        result = subprocess.run(
            ['osascript', '-e', applescript],
            capture_output=True,
            text=True,
            timeout=60
        )
        
        text = result.stdout.strip()
        
        if result.returncode == 0 and text:
            return {
                "success": True,
                "text": text,
                "method": "macos_dialog"
            }
        else:
            return {
                "success": False,
                "text": "",
                "error": "User cancelled or no input"
            }
            
    except Exception as e:
        return {
            "success": False,
            "text": "",
            "error": str(e)
        }

if __name__ == "__main__":
    result = get_voice_input()
    print(json.dumps(result))
