package metadata

type ChunkInfo struct {
	ID    int      `json:"id"`
	Nodes []string `json:"nodes"`
}

type Metadata struct {
	FileName    string      `json:"filename"`
	TotalChunks int         `json:"totalChunks"`
	Chunks      []ChunkInfo `json:"chunks"`
}
