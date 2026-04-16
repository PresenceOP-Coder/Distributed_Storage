package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"dfs/internal/config"
	"dfs/internal/node"
)

func main() {
	port := os.Getenv("NODE_PORT")
	if port == "" {
		port = "8001"
	}

	dataDir := filepath.Join(config.DataDir, port)
	mux := http.NewServeMux()
	node.RegisterRoutes(mux, dataDir)

	addr := ":" + port
	log.Printf("[NODE] listening=%s data_dir=%s", addr, dataDir)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("[ERROR] event=node_server_start err=%v", err)
	}
}
