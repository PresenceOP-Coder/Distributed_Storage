# Distributed File Storage System (Go)

## Overview

This project implements a distributed file storage system in Go, inspired by systems like Google Drive and Dropbox.

It supports:
- Uploading files
- Splitting files into chunks
- Storing chunks across multiple nodes
- Reconstructing files on download

## Goals

- Learn Go concurrency (goroutines, channels)
- Understand distributed systems fundamentals
- Build a production-style backend system
- Strengthen SDE-level system design and implementation skills

## Architecture

Client -> API Server -> Storage Nodes
                |
                v
           Metadata Store

## Components

### 1. API Server
- Handles upload and download requests
- Splits files into chunks
- Distributes chunks to storage nodes

### 2. Storage Nodes
- Store file chunks
- Run on different ports

### 3. Metadata Store
- Tracks where each chunk is stored

## Tech Stack

- Go (Golang)
- net/http
- Goroutines and channels
- File I/O
- JSON

## Folder Structure

```text
project/
|
|-- cmd/
|   |-- server/          # API server
|   `-- node/            # storage nodes
|
|-- internal/
|   |-- chunker/
|   |-- storage/
|   |-- metadata/
|   `-- api/
|
|-- data/                # stored chunks
`-- go.mod
```

## Roadmap

The implementation is divided into six phases:

1. Phase 1: Basic file server
2. Phase 2: File chunking
3. Phase 3: Multi-node storage
4. Phase 4: Metadata system
5. Phase 5: Concurrency
6. Phase 6: Advanced features

Detailed guides are available in the docs folder:

- docs/phase-1-basic-server.md
- docs/phase-2-chunking.md
- docs/phase-3-multi-node.md
- docs/phase-4-metadata.md
- docs/phase-5-concurrency.md
- docs/phase-6-advanced.md

## Phase 1 Quick Start

Initialize the project:

```bash
go mod init dfs
```

Then begin with Phase 1 implementation from the docs.

## Testing Checklist

Use Postman or curl to validate:

- Upload large and small files
- Download and verify file correctness
- Chunk reconstruction integrity
- Node failure behavior

## Resume Line

Built a distributed file storage system in Go with chunking, concurrent uploads, and multi-node storage architecture.

## Interview Preparation

Be ready to explain:

- Why chunking is used
- Why goroutines help throughput
- How to scale to many nodes
- What happens when a node fails
- How consistency is maintained

## Future Improvements

- Add database-backed metadata (for example MongoDB)
- Add authentication and authorization
- Add web UI and CLI support
- Deploy to cloud infrastructure

## Final Advice

Focus on:

- Clean code
- Clear architecture
- Correctness under failures
- Ability to explain design choices
