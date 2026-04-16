package node

import (
	"net/http"

	"dfs/internal/node/handlers"
)

func RegisterRoutes(mux *http.ServeMux, dataDir string) {
	mux.Handle("/store", handlers.StoreHandler{DataDir: dataDir})
	mux.Handle("/chunk", handlers.FetchHandler{DataDir: dataDir})
}
