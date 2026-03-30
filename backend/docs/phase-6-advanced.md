# Phase 6: Advanced Features

## Goal
Make the system more resilient and production-ready.

## Features

### 1. Replication
- Store each chunk on multiple nodes for redundancy

### 2. Failure Handling
- Detect node failure
- Use fallback node or replica for reads
- Retry writes to alternate nodes

### 3. Load Balancing
- Improve distribution strategy beyond simple round-robin
- Consider node capacity and health

### 4. Logging and Observability
- Track upload, chunk placement, retries, and failures
- Add structured logs for troubleshooting

## Testing
- Simulate node crash during upload and download
- Verify data can still be served from replicas
- Validate logs capture key events

## Output
- More robust and fault-tolerant distributed storage

## Bonus Enhancements
- Add authentication and authorization
- Add CLI for upload/download operations
- Deploy nodes and API server to cloud
