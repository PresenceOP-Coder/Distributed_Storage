package api

import (
	"net/http"

	"dfs/internal/api/handlers"
	"dfs/internal/loadbalancer"
)

func RegisterRoutes(mux *http.ServeMux, balancer *loadbalancer.Balancer) {
	mux.Handle("/upload", handlers.UploadHandler{Balancer: balancer})
	mux.Handle("/download", handlers.DownloadHandler{})
}
