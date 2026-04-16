# distributed-storage

## Run API server

```bash
go run ./cmd/api-server
```

## Run node servers

```bash
NODE_PORT=8001 go run ./cmd/node-server
NODE_PORT=8002 go run ./cmd/node-server
NODE_PORT=8003 go run ./cmd/node-server
```