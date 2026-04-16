package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Load(fileID, dir string) (Metadata, error) {
	var meta Metadata

	path := filepath.Join(dir, fmt.Sprintf("%s.json", fileID))
	file, err := os.Open(path)
	if err != nil {
		return meta, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&meta)
	return meta, err
}
