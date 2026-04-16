package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"dfs/internal/config"
	"dfs/internal/loadbalancer"
	"dfs/internal/metadata"
	"dfs/internal/storage"
	"dfs/internal/utils"
)

type UploadHandler struct {
	Balancer *loadbalancer.Balancer
}

func (h UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "invalid multipart form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("myfile")
	if err != nil {
		http.Error(w, "missing myfile form field", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := os.MkdirAll(config.UploadDir, os.ModePerm); err != nil {
		http.Error(w, "failed to prepare upload directory", http.StatusInternalServerError)
		return
	}

	dstPath := filepath.Join(config.UploadDir, header.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "could not create destination file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "could not save uploaded file", http.StatusInternalServerError)
		return
	}

	if err := os.MkdirAll(config.DataDir, os.ModePerm); err != nil {
		http.Error(w, "could not prepare chunk directory", http.StatusInternalServerError)
		return
	}
	if err := os.MkdirAll(config.StorageDir, os.ModePerm); err != nil {
		http.Error(w, "could not prepare storage directory", http.StatusInternalServerError)
		return
	}

	fileID := fmt.Sprintf("%d", time.Now().UnixNano())
	meta, err := storage.SplitAndDistribute(dstPath, fileID, config.ChunkSize, h.Balancer, config.ReplicationFactor, config.MaxRetries)
	if err != nil {
		http.Error(w, "could not split uploaded file", http.StatusInternalServerError)
		return
	}

	if err := metadata.Save(fileID, meta, config.MetadataDir); err != nil {
		http.Error(w, "metadata save failed", http.StatusInternalServerError)
		return
	}

	reconstructedPath := filepath.Join(config.StorageDir, config.ReconstructedFile)
	if err := storage.MergeDistributed(reconstructedPath, meta.TotalChunks, fileID, h.Balancer.NodeURLs()); err != nil {
		http.Error(w, "could not reconstruct distributed file", http.StatusInternalServerError)
		return
	}

	utils.LogJSON(map[string]any{
		"event":   "upload_complete",
		"file":    header.Filename,
		"file_id": fileID,
		"chunks":  meta.TotalChunks,
	})

	fmt.Fprintf(w, "uploaded=%s file_id=%s", header.Filename, fileID)
}
