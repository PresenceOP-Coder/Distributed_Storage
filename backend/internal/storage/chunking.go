package storage

import (
	"fmt"
	"io"
	"os"
)

func ChunkName(fileID string, idx int) string {
	return fmt.Sprintf("%s_%d.chunk", fileID, idx)
}

func MergeDistributed(outputFile string, totalChunks int, fileID string, nodeURLs []string) error {
	if len(nodeURLs) == 0 {
		return fmt.Errorf("no nodes configured")
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	for i := 0; i < totalChunks; i++ {
		node := nodeURLs[i%len(nodeURLs)]
		chunkName := ChunkName(fileID, i)
		data, err := FetchChunk(node, chunkName)
		if err != nil {
			return err
		}
		if _, err := io.Copy(out, bytesReader(data)); err != nil {
			return err
		}
	}
	return nil
}

func bytesReader(data []byte) io.Reader {
	return &byteReader{data: data}
}

type byteReader struct {
	data []byte
	off  int
}

func (r *byteReader) Read(p []byte) (int, error) {
	if r.off >= len(r.data) {
		return 0, io.EOF
	}
	n := copy(p, r.data[r.off:])
	r.off += n
	return n, nil
}
