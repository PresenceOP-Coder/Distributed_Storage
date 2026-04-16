package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"dfs/internal/utils"
)

type StoreHandler struct {
	DataDir string
}

func (h StoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, "invalid chunk upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if err := os.MkdirAll(h.DataDir, os.ModePerm); err != nil {
		http.Error(w, "cannot prepare data directory", http.StatusInternalServerError)
		return
	}

	path := filepath.Join(h.DataDir, filepath.Base(header.Filename))
	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "cannot save chunk", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "cannot write chunk", http.StatusInternalServerError)
		return
	}

	utils.Infof("[STORE] chunk=%s status=success", header.Filename)
	w.Write([]byte("OK"))
}
