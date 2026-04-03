#!/usr/bin/env python3
"""
PMBuddy Voice Transcriber
Converts recorded audio files to text using local speech recognition.
Supports multiple recognition backends with intelligent fallback.
"""

import sys
import os
import json
import subprocess
from pathlib import Path

def transcribe_with_recognizer(audio_file):
    """
    Transcribe audio using the speech_recognition library.
    Falls back to multiple engines if one fails.
    """
    try:
        import speech_recognition as sr
    except ImportError:
        return None

    recognizer = sr.Recognizer()
    recognizer.energy_threshold = 3000  # Lower threshold for quiet audio
    recognizer.dynamic_energy_threshold = True

    try:
        with sr.AudioFile(audio_file) as source:
            audio = recognizer.record(source)

        # Check if audio is too quiet
        raw_data = audio.get_raw_data()
        if len(raw_data) < 1000:
            return None

        # Try Google Speech Recognition (no key needed, uses free API)
        try:
            text = recognizer.recognize_google(audio)
            print(json.dumps({
                "success": True,
                "text": text,
                "engine": "google"
            }))
            return text
        except sr.UnknownValueError:
            pass
        except sr.RequestError as e:
            pass
        except Exception as e:
            pass

        # Try Sphinx (offline, no internet required)
        try:
            text = recognizer.recognize_sphinx(audio)
            if text and len(text) > 2:  # Sphinx sometimes returns garbage
                print(json.dumps({
                    "success": True,
                    "text": text,
                    "engine": "sphinx"
                }))
                return text
        except sr.RequestError as e:
            pass
        except Exception as e:
            pass

        # Try Whisper (if installed)
        try:
            import whisper
            model = whisper.load_model("base")
            result = model.transcribe(audio_file, language="en")
            text = result["text"].strip()
            if text:
                print(json.dumps({
                    "success": True,
                    "text": text,
                    "engine": "whisper"
                }))
                return text
        except ImportError:
            pass
        except Exception as e:
            pass

        # If all else fails
        return None

    except Exception as e:
        return None


def transcribe_with_whisper(audio_file):
    """
    Transcribe using OpenAI Whisper (if installed).
    Requires: pip3 install openai-whisper
    """
    try:
        import whisper
    except ImportError:
        return None

    try:
        model = whisper.load_model("base")  # ~140MB download on first run
        result = model.transcribe(audio_file, language="en")
        
        text = result["text"].strip()
        if text:
            print(json.dumps({
                "success": True,
                "text": text,
                "engine": "whisper"
            }))
            return text
    except Exception as e:
        pass

    return None


def main():
    if len(sys.argv) < 2:
        return
        
    audio_file = sys.argv[1]

    if not os.path.exists(audio_file):
        return

    # Try Whisper first (better quality), then fall back to speech_recognition
    text = transcribe_with_whisper(audio_file)
    if text:
        return
    
    text = transcribe_with_recognizer(audio_file)
    if text:
        return

    # If we get here, all methods failed - print error JSON
    print(json.dumps({
        "success": False,
        "error": "Transcription failed",
        "text": ""
    }))


if __name__ == "__main__":
    main()
