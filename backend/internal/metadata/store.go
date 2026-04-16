package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Save(fileID string, meta Metadata, dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	path := filepath.Join(dir, fmt.Sprintf("%s.json", fileID))
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(meta)
}
