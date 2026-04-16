package handlers

import (
	"net/http"
	"path/filepath"
)

type FetchHandler struct {
	DataDir string
}

func (h FetchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	path := filepath.Join(h.DataDir, filepath.Base(name))
	http.ServeFile(w, r, path)
}
