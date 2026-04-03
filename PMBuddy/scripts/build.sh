#!/bin/bash

# Build script for PMBuddy - compiles for macOS and Windows

set -e

PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
VERSION=$(git describe --tags --always 2>/dev/null || echo "dev")
BUILD_DATE=$(date -u '+%Y-%m-%d_%H:%M:%S')

echo "🏗️  Building PMBuddy ($VERSION)"

# Default target
TARGET=${1:-darwin}

case $TARGET in
  darwin)
    echo "🍎 Building for macOS..."
    GOOS=darwin GOARCH=arm64 go build -o "$PROJECT_ROOT/pmbuddy-darwin-arm64" \
      -ldflags "-X main.Version=$VERSION -X main.BuildDate=$BUILD_DATE" \
      ./cmd/pmbuddy
    
    GOOS=darwin GOARCH=amd64 go build -o "$PROJECT_ROOT/pmbuddy-darwin-amd64" \
      -ldflags "-X main.Version=$VERSION -X main.BuildDate=$BUILD_DATE" \
      ./cmd/pmbuddy
    
    echo "✅ Created pmbuddy-darwin-arm64 and pmbuddy-darwin-amd64"
    ;;
    
  windows)
    echo "🪟 Building for Windows..."
    GOOS=windows GOARCH=amd64 go build -o "$PROJECT_ROOT/pmbuddy-windows-amd64.exe" \
      -ldflags "-X main.Version=$VERSION -X main.BuildDate=$BUILD_DATE" \
      ./cmd/pmbuddy
    
    echo "✅ Created pmbuddy-windows-amd64.exe"
    ;;
    
  all)
    echo "📦 Building for all platforms..."
    bash "$0" darwin
    bash "$0" windows
    echo "✅ All builds complete!"
    ;;
    
  *)
    echo "❌ Unknown target: $TARGET"
    echo "Usage: ./scripts/build.sh [darwin|windows|all]"
    exit 1
    ;;
esac

echo ""
echo "✨ Build complete!"
ls -lh "$PROJECT_ROOT"/pmbuddy*
