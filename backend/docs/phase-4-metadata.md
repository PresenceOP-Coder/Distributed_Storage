# Phase 4: Metadata System

## Goal
Track where each file chunk is stored.

## Concepts Covered
- Data mapping
- JSON metadata storage
- Chunk location lookup

## Metadata Example

```json
{
  "file": "resume.pdf",
  "chunks": [
    {"id": 1, "node": "8001"},
    {"id": 2, "node": "8002"}
  ]
}
```

## Steps

### 1. Define Metadata Structure
Store:
- File name
- Chunk ID to node mapping

### 2. Persist Metadata
Choose one:
- JSON file
- In-memory map (for early prototype)

### 3. Use Metadata on Download
- Read metadata for requested file
- Fetch each chunk from correct node
- Merge chunks into final file

## Output
- System can locate chunks reliably during download

## Test
- Upload a file
- Confirm metadata entry is created
- Download and verify reconstruction using metadata

## Common Issues
- Missing metadata: file cannot be reconstructed
- Stale mapping: metadata not updated after node/storage change
