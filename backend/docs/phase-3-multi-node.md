# Phase 3: Multi-Node Storage

## Goal
Distribute chunks across multiple storage nodes.

## Concepts Covered
- Distributed storage basics
- HTTP client-server communication
- Round-robin load distribution

## Steps

### 1. Create Node Server
Each node:
- Runs on a different port
- Stores chunk data

Example ports:
- 8001
- 8002
- 8003

### 2. Define Node APIs
- POST /store: store a chunk
- GET /chunk: fetch a chunk

### 3. Distribute Chunks
Use round-robin:

```go
node := nodes[i%len(nodes)]
```

### 4. Send Data to Nodes
- API server sends each chunk to selected node using HTTP client

## Setup
Run multiple node processes:
- node1
- node2
- node3

## Output
- Chunks are distributed across different nodes

## Test
- Upload a file
- Verify chunks are saved on different node data directories

## Common Issues
- Node unreachable: port or process not running
- Wrong routing: incorrect node selection logic
