# CLAUDE.md - AI Assistant Memory Protocol

Welcome! This is your external memory system for the PMBuddy project. I'm Claude, and I'll treat this Obsidian vault as my persistent second brain throughout our work together.

## 📋 Vault Structure

This vault is organized into four main areas:

### 🏠 Root Level
- **CLAUDE.md** - This file (protocol and guidelines)
- **Home.md** - Index and entry point for the vault
- **Definition.md, Persona.md, User Story.md, etc.** - Core PM knowledge base (do not modify unless specified)

### 📝 Notes/ - Research & Exploration
**Purpose**: Temporary notes, research findings, and work-in-progress thoughts
**When to use**: 
- Recording research findings during code exploration
- Documenting decision rationales
- Capturing complex architectural patterns
- Taking notes during debugging sessions

**File naming**: `YYYYMMDD-topic.md` (e.g., `20260405-audio-transcription-bug.md`)

**Example structure**:
```markdown
# Audio Transcription Bug Investigation
Date: 2026-04-05
Status: In Progress

## Findings
- Python script not returning output correctly
- Issue in main() function line 136
- [[Solution-Audio-Fix]] identified

## Related
- [[Whisper-Integration]]
- [[Speech-Recognition-Engine]]
```

### 🎯 Output/ - Completed Work & Deliverables
**Purpose**: Final deliverables, completed analyses, and actionable outcomes
**When to use**:
- After completing a feature or bug fix
- When wrapping up research into conclusions
- Documenting final implementation decisions
- Saving deployment guides or release notes

**File naming**: `Project-Deliverable-Name.md` (e.g., `PMBuddy-License-Setup.md`)

**Example structure**:
```markdown
# PMBuddy License & Dependency Management

## Completed
- ✅ LICENSE file created (MIT)
- ✅ DEPENDENCIES.txt with 10+ packages
- ✅ requirements.txt for pip3
- ✅ Auto-installation during init

## Deliverables
- LICENSE
- DEPENDENCIES.txt
- requirements.txt

## Files Changed
- cmd/pmbuddy/main.go (added dependency check function)
```

### 💡 Topics/ - Topic Ideas & Backlog
**Purpose**: Future features, improvements, and research topics to explore
**When to use**:
- Brainstorming new features
- Recording technical debt
- Planning future enhancements
- Tracking improvement ideas

**File naming**: `Topic-Name.md` (e.g., `Session-Persistence.md`)

**Example structure**:
```markdown
# Session Persistence Feature

## Idea
Allow users to save and resume PM Buddy sessions with conversation history

## Related Features
- [[User-Story-Documentation]]
- [[Persona-Management]]

## Technical Considerations
- Database schema for sessions
- Encryption for sensitive data
- Recovery mechanisms

## Priority
Medium - Nice to have, improves UX
```

## 🔗 Wiki Linking Convention

Use double brackets `[[]]` to create wiki-style links between notes:

```markdown
## Related Concepts
- [[Audio-Transcription]]
- [[OpenAI-Whisper-Integration]]
- [[Voice-Session-Mode]]

## Dependencies
- [[PM-Knowledge-Base]]
```

**Benefits**:
- Navigate relationships between concepts
- Understand project dependencies
- Track connected ideas across sessions
- Build a knowledge map over time

## 📌 Session Workflow

At the start of each session:
1. I'll review Home.md to understand current state
2. Check Notes/ for recent findings
3. Review Topics/ for planned work
4. Update files as work progresses

At the end of each session:
1. Move completed research from Notes/ → Output/
2. Create summary in Output/ documenting work done
3. Update Home.md with latest status
4. Archive old Notes/ entries if needed

## ✍️ File Content Guidelines

### Timestamps
Include dates in research notes:
```markdown
Date: 2026-04-05
Time: 14:47 UTC
Status: In Progress
```

### Code References
When saving code findings:
```markdown
## Code Location
File: cmd/pmbuddy/main.go
Lines: 60-84
Function: checkAndInstallDependencies()
```

### Decision Log
Document why decisions were made:
```markdown
## Decision: Use MIT License
Reasoning:
- Compatible with all dependencies
- Permissive for contributors
- Suitable for commercial use

Alternatives Considered:
- Apache 2.0 (more complex)
- GPL v3 (too restrictive)
```

### Action Items
Track what needs to be done:
```markdown
## Next Steps
- [ ] Implement feature X
- [ ] Test scenario Y
- [ ] Document in README
- [ ] Update copilot-instructions.md
```

## 🔄 Integration with Project

I'll use this vault to:

1. **Remember Context** - Access previous session notes to maintain continuity
2. **Track Decisions** - Reference why choices were made
3. **Document Progress** - Keep clear records of completed work
4. **Plan Future Work** - Review Topics/ for upcoming tasks
5. **Link Concepts** - Use [[wiki links]] to show relationships

## 📚 Knowledge Base Notes

The root-level markdown files are your PM knowledge base:
- **Definition.md** - Core PM definitions and terminology
- **Persona.md** - User persona framework
- **User Story.md** - User story template and guidelines
- **Jobs to Be Done.md** - JTBD framework
- **Product discovery.md** - Discovery process documentation
- **Customers Need.md** - Customer needs analysis
- **Reason.md** - Core reasoning and frameworks

These are linked by [[wiki links]] and inform PMBuddy's intelligence.

## 🎯 Vault as External Memory

This vault serves as:
- **Session Continuity**: I can pick up exactly where we left off
- **Decision Repository**: All important decisions documented with rationale
- **Knowledge Accumulation**: Building a project knowledge map over time
- **Quick Reference**: Fast lookup of completed work and findings
- **Implementation Guide**: Clear records of how things were done

## 💡 Best Practices

1. **Link Everything**: Use [[wiki links]] liberally to connect ideas
2. **Date Your Notes**: Always include date so we know when findings were recorded
3. **Be Specific**: Include file paths, line numbers, function names
4. **Document Why**: Explain decisions, not just what was decided
5. **Keep It Organized**: Use proper folders and naming conventions
6. **Update Status**: Mark notes as "In Progress", "Completed", "On Hold"

## 🚀 Let's Build Together

This vault is our collaborative space. I'll treat it as my external memory, ensuring:
- No context is lost between sessions
- Every decision is documented
- Progress is clearly tracked
- Future work is well-planned

Ready to start using this vault as our persistent second brain!

---

Last Updated: 2026-04-05
Version: 1.0
