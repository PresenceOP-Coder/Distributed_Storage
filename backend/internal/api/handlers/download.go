package handlers

import (
	"fmt"
	"net/http"

	"dfs/internal/config"
	"dfs/internal/metadata"
	"dfs/internal/storage"
)

type DownloadHandler struct{}

func (h DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileID := r.URL.Query().Get("file")
	if fileID == "" {
		http.Error(w, "please provide file query param", http.StatusBadRequest)
		return
	}

	meta, err := metadata.Load(fileID, config.MetadataDir)
	if err != nil {
		http.Error(w, "metadata not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+meta.FileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	for _, chunk := range meta.Chunks {
		chunkName := fmt.Sprintf("%s_%d.chunk", fileID, chunk.ID)
		data, err := storage.FetchFromReplicas(chunk.Nodes, chunkName)
		if err != nil {
			http.Error(w, "download failed", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(data); err != nil {
			return
		}
	}
}
