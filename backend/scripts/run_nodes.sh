#!/usr/bin/env bash
set -euo pipefail

NODE_PORT=8001 go run ./cmd/node-server &
NODE_PORT=8002 go run ./cmd/node-server &
NODE_PORT=8003 go run ./cmd/node-server &

wait