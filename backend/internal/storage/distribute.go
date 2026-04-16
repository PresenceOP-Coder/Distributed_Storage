package storage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"dfs/internal/loadbalancer"
	"dfs/internal/metadata"
	"dfs/internal/utils"
)

func SplitAndDistribute(filePath, fileID string, chunkSize int, lb *loadbalancer.Balancer, replicationFactor, maxRetries int) (metadata.Metadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return metadata.Metadata{}, err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)
	var wg sync.WaitGroup
	errChan := make(chan error, 100)

	meta := metadata.Metadata{FileName: filepath.Base(filePath)}
	var mu sync.Mutex
	chunkIndex := 0

	for {
		n, readErr := file.Read(buffer)
		if readErr != nil && readErr != io.EOF {
			return meta, readErr
		}
		if n == 0 {
			break
		}

		data := make([]byte, n)
		copy(data, buffer[:n])

		node, err := lb.GetBestNode()
		if err != nil {
			return meta, err
		}

		chunkName := ChunkName(fileID, chunkIndex)
		wg.Add(1)

		go func(idx int, nodeURL, name string, payload []byte) {
			defer wg.Done()
			lb.TrackLoad(nodeURL, 1)
			defer lb.TrackLoad(nodeURL, -1)

			replicaNodes := lb.GetReplicaNodes(nodeURL, replicationFactor)
			if len(replicaNodes) == 0 {
				errChan <- fmt.Errorf("no replica nodes available")
				return
			}

			var uploadErr error
			for attempt := 1; attempt <= maxRetries; attempt++ {
				uploadErr = sendChunkToNode(nodeURL, name, payload)
				if uploadErr == nil {
					utils.Infof("[UPLOAD] chunk=%d node=%s status=success", idx, nodeURL)
					break
				}
				utils.Infof("[RETRY] chunk=%d attempt=%d node=%s", idx, attempt, nodeURL)
				time.Sleep(500 * time.Millisecond)
			}

			if uploadErr != nil {
				utils.Errorf("[ERROR] chunk=%d err=%v", idx, uploadErr)
				errChan <- fmt.Errorf("failed chunk %d: %v", idx, uploadErr)
				return
			}

			mu.Lock()
			meta.Chunks = append(meta.Chunks, metadata.ChunkInfo{ID: idx, Nodes: replicaNodes})
			mu.Unlock()
		}(chunkIndex, node, chunkName, data)

		chunkIndex++
		if readErr == io.EOF {
			break
		}
	}

	meta.TotalChunks = chunkIndex
	wg.Wait()
	close(errChan)
	for e := range errChan {
		if e != nil {
			return meta, e
		}
	}

	return meta, nil
}

func sendChunkToNode(nodeURL, fileName string, data []byte) error {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("chunk", fileName)
	if err != nil {
		return err
	}
	if _, err := part.Write(data); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, nodeURL+"/store", &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to store %s on %s: status %d", fileName, nodeURL, resp.StatusCode)
	}

	utils.LogJSON(map[string]any{
		"event":  "upload",
		"chunk":  fileName,
		"node":   nodeURL,
		"status": "success",
	})
	return nil
}
