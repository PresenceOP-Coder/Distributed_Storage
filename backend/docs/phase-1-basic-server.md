# Phase 1: Basic File Server

## Goal
Build a simple HTTP server to upload and download files on a single node.

## Concepts Covered
- HTTP server in Go
- File upload (multipart/form-data)
- File download
- File I/O

## Steps

### 1. Initialize Project

```bash
go mod init dfs
```

### 2. Create Server
- Use net/http
- Run on port 8080

### 3. Upload Endpoint

Endpoint:
- POST /upload

Tasks:
- Parse multipart form data
- Extract uploaded file
- Save file to local disk

### 4. Download Endpoint

Endpoint:
- GET /download?file=filename

Tasks:
- Read file from local disk
- Return file in HTTP response

## Directory

```text
data/
```

## Output
- Uploaded file is saved locally
- Download returns the same file

## Test
- Upload a file
- Download the same file
- Verify content matches

## Common Issues
- File not saving: check write permissions and destination path
- Empty file: verify multipart parsing and stream copy logic
