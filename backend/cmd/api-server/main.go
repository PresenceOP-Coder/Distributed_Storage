package main

import (
	"log"
	"net/http"

	"dfs/internal/api"
	"dfs/internal/config"
	"dfs/internal/loadbalancer"
)

func main() {
	mux := http.NewServeMux()
	balancer := loadbalancer.New(config.DefaultNodes)
	api.RegisterRoutes(mux, balancer)

	log.Printf("[API] listening=%s", config.APIServerAddr)
	if err := http.ListenAndServe(config.APIServerAddr, mux); err != nil {
		log.Fatalf("[ERROR] event=api_server_start err=%v", err)
	}
}
