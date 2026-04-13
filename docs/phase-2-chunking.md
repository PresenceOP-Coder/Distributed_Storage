# Phase 2: File Chunking

## Goal
Split large files into smaller chunks and reconstruct them correctly.

## Concepts Covered
- File buffering
- Chunking strategy
- Sequential file processing

## Steps

### 1. Define Chunk Size
Example:
- 1 MB per chunk

### 2. Read File in Loop
- Use a fixed-size buffer
- Read until EOF

### 3. Save Chunks
Use deterministic naming:
- file_1.chunk
- file_2.chunk
- file_3.chunk

### 4. Merge Chunks
- Read chunk files in order
- Append data into a rebuilt output file

## Structure

```text
data/
|-- file_1.chunk
`-- file_2.chunk
```

## Output
- File is split into chunk files
- Reconstructed file matches original

## Test
- Upload a large file
- Verify chunk files are created
- Merge and compare hash/checksum with original

## Common Issues
- Missing chunks: loop termination or index bug
- Corrupted output: chunks merged in wrong order
