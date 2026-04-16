package storage

import (
	"fmt"
	"io"
	"net/http"

	"dfs/internal/utils"
)

func FetchChunk(nodeURL, chunkName string) ([]byte, error) {
	resp, err := http.Get(nodeURL + "/chunk?name=" + chunkName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s from %s: status %d", chunkName, nodeURL, resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func FetchFromReplicas(nodeURLs []string, chunkName string) ([]byte, error) {
	for _, nodeURL := range nodeURLs {
		data, err := FetchChunk(nodeURL, chunkName)
		if err == nil {
			utils.Infof("[FETCH] chunk=%s node=%s status=success", chunkName, nodeURL)
			return data, nil
		}
		utils.Errorf("[ERROR] chunk=%s node=%s err=%v", chunkName, nodeURL, err)
	}
	return nil, fmt.Errorf("all replicas failed")
}
