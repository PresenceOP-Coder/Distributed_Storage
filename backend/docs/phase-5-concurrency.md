# Phase 5: Concurrency

## Goal
Upload chunks in parallel to improve performance.

## Concepts Covered
- Goroutines
- sync.WaitGroup
- Channels for coordination and error collection

## Steps

### 1. Run Uploads in Goroutines
- Launch one goroutine per chunk upload

### 2. Synchronize Completion
- Use sync.WaitGroup to wait until all uploads finish

### 3. Handle Errors
- Collect errors via channel
- Retry failed chunk uploads with bounded retry count

## Output
- Faster upload throughput
- Better resource utilization

## Test
- Upload a large file
- Compare upload time before and after parallelization
- Simulate one failing node and verify retries

## Common Issues
- Race conditions from shared mutable state
- Premature exit due to incorrect WaitGroup usage
- Goroutine leaks from unclosed channels
