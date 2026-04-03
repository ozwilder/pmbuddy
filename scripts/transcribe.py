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
        print(json.dumps({
            "success": False,
            "error": "speech_recognition not installed. Run: pip3 install SpeechRecognition pydub",
            "text": ""
        }), file=sys.stderr)
        return None

    recognizer = sr.Recognizer()
    recognizer.energy_threshold = 3000  # Lower threshold for quiet audio
    recognizer.dynamic_energy_threshold = True

    try:
        with sr.AudioFile(audio_file) as source:
            print(f"Loading audio from {audio_file}...", file=sys.stderr)
            audio = recognizer.record(source)

        # Check if audio is too quiet
        raw_data = audio.get_raw_data()
        if len(raw_data) < 1000:
            print(json.dumps({
                "success": False,
                "error": "Audio too short or empty (microphone not working?)",
                "text": ""
            }), file=sys.stderr)
            return None

        print("Recognizing speech...", file=sys.stderr)
        
        # Try Google Speech Recognition (no key needed, uses free API)
        try:
            print("  Trying Google Speech API...", file=sys.stderr)
            text = recognizer.recognize_google(audio)
            print(json.dumps({
                "success": True,
                "text": text,
                "engine": "google"
            }))
            return text
        except sr.UnknownValueError:
            print("  Google: Could not understand audio (too quiet?)", file=sys.stderr)
        except sr.RequestError as e:
            print(f"  Google API error: {e}", file=sys.stderr)
        except Exception as e:
            print(f"  Google error: {e}", file=sys.stderr)

        # Try Sphinx (offline, no internet required)
        try:
            print("  Trying Sphinx (offline)...", file=sys.stderr)
            text = recognizer.recognize_sphinx(audio)
            if text and len(text) > 2:  # Sphinx sometimes returns garbage
                print(json.dumps({
                    "success": True,
                    "text": text,
                    "engine": "sphinx"
                }))
                return text
        except sr.RequestError as e:
            print(f"  Sphinx error (may not be installed): {e}", file=sys.stderr)
        except Exception as e:
            print(f"  Sphinx error: {e}", file=sys.stderr)

        # Try Whisper (if installed)
        try:
            print("  Trying Whisper (if installed)...", file=sys.stderr)
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
            print("  Whisper: not installed (optional)", file=sys.stderr)
        except Exception as e:
            print(f"  Whisper error: {e}", file=sys.stderr)

        # If all else fails
        print(json.dumps({
            "success": False,
            "error": "Could not recognize speech with any available engine. Try speaking louder or closer to the microphone.",
            "text": ""
        }), file=sys.stderr)
        return None

    except Exception as e:
        print(json.dumps({
            "success": False,
            "error": f"Audio processing error: {str(e)}",
            "text": ""
        }), file=sys.stderr)
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
        print("Loading Whisper model...", file=sys.stderr)
        model = whisper.load_model("base")  # ~140MB download on first use
        
        print(f"Transcribing {audio_file}...", file=sys.stderr)
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
        print(f"Whisper error: {e}", file=sys.stderr)

    return None


def main():
    if len(sys.argv) < 2:
        print(json.dumps({
            "success": False,
            "error": "Usage: transcribe.py <audio_file>",
            "text": ""
        }), file=sys.stderr)
        sys.exit(1)

    audio_file = sys.argv[1]

    if not os.path.exists(audio_file):
        print(json.dumps({
            "success": False,
            "error": f"Audio file not found: {audio_file}",
            "text": ""
        }), file=sys.stderr)
        sys.exit(1)

    print(f"Transcribing: {audio_file}", file=sys.stderr)

    # Try Whisper first (better quality), then fall back to speech_recognition
    text = transcribe_with_whisper(audio_file)
    if text:
        return
    
    text = transcribe_with_recognizer(audio_file)
    if text:
        return

    # If we get here, all methods failed
    print(json.dumps({
        "success": False,
        "error": "All transcription methods failed",
        "text": ""
    }), file=sys.stderr)
    sys.exit(1)


if __name__ == "__main__":
    main()
