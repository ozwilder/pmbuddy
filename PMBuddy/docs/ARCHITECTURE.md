# PMBuddy Architecture

## System Overview

PMBuddy is a modular AI advisor system for Product Managers with the following core components:

```
┌─────────────────────────────────────────────────────────┐
│                    User Interface Layer                  │
│                  (CLI / Voice / Chat)                    │
└─────────────────────────────────────────────────────────┘
                           ↓
┌─────────────────────────────────────────────────────────┐
│                   Orchestration Layer                    │
│              (Session Management, Routing)               │
└─────────────────────────────────────────────────────────┘
                           ↓
┌──────────────┬──────────────┬──────────────┬──────────────┐
│    Agent     │  Knowledge   │    Voice     │   Storage    │
│   Advisor    │     Base     │   Interface  │   Manager    │
│              │              │              │              │
│ • PM Logic   │ • Learning   │ • Handy STT  │ • Sessions   │
│ • Context    │ • Insights   │ • Handy TTS  │ • Conv History
│ • Reasoning  │ • Preferences│              │ • Analytics  │
└──────────────┴──────────────┴──────────────┴──────────────┘
```

## Component Details

### 1. Agent (advisor.go)
The core PM advisor logic that handles:
- **Session Management**: Creates and tracks unique sessions
- **Query Processing**: Interprets user questions and requests
- **Reasoning**: Provides PM advice based on context and knowledge
- **Memory**: Maintains conversation history for context

**Key Types:**
- `PMAdvisor`: Main advisor instance
- `Message`: Represents individual conversation turns

### 2. Knowledge Base (knowledge.go)
Manages accumulated knowledge and continuous learning:
- **Learning History**: Records insights from conversations
- **Context Mappings**: Links related concepts together
- **User Preferences**: Tracks PM preferences and style
- **Key Insights**: Maintains important learnings

**Key Types:**
- `KnowledgeBase`: Accumulates knowledge over time
- `LearningEntry`: Individual learning records

### 3. Voice Interface (handy.go)
Integrates Handy for voice interaction:
- **Speech-to-Text**: Converts voice to text using Handy
- **Text-to-Speech**: Speaks responses back to user
- **Configuration**: Manages voice parameters
- **Status**: Tracks Handy connection state

**Key Types:**
- `HandyInterface`: Main voice interface manager

### 4. Storage Manager (storage.go)
Persists data for continuity across sessions:
- **Session Storage**: Saves/loads complete sessions
- **Conversation History**: Stores all interactions
- **Learning Records**: Persists knowledge base
- **Backups**: Creates automatic backups

**Key Types:**
- `StorageManager`: Handles all persistence
- `SessionData`: Complete session representation

## Data Flow

### Text Interaction Flow
```
User Input → CLI → Agent.ProcessQuery() → Response → CLI → User
                         ↓
                  Knowledge Base
                  (context, learning)
                         ↓
                  Storage Manager
                  (persistence)
```

### Voice Interaction Flow
```
User Voice → Handy (STT) → Text → Agent.ProcessQuery() → Response 
                                         ↓
                                  Knowledge Base
                                  (learning)
                                         ↓
                                  Handy (TTS) → User Audio
                                         ↓
                                  Storage Manager
                                  (persistence)
```

## Storage Structure

```
~/.pmbuddy/
├── sessions/
│   ├── pmb_1234567890.json    # Complete session data
│   └── ...
├── conversations/
│   ├── pmb_1234567890_conv_1.json
│   └── ...
├── learnings/
│   ├── pmb_1234567890_learn.json
│   └── ...
└── backups/
    ├── pmb_1234567890_20240402180000.json
    └── ...
```

## Session Lifecycle

1. **Initialization**: Create new PMAdvisor, initialize storage
2. **Interaction**: User queries → advisor processes → response
3. **Learning**: Knowledge base captures key insights
4. **Persistence**: Each turn auto-saved to storage
5. **Closure**: Session finalized and backed up

## Extensibility Points

- **AI Engine**: Replace query response logic with LLM integration
- **Data Sources**: Extend KnowledgeBase with external APIs
- **Voice**: Swap Handy for alternative STT/TTS systems
- **Storage**: Implement database backends (currently file-based)
- **Analysis**: Add analytics and reporting capabilities

## Concurrency & Safety

- Sessions are isolated (no concurrent access issues)
- Storage operations use atomic writes
- Future: Add mutex protection for multi-client scenarios
