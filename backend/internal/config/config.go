package config

const (
	APIServerAddr     = ":8005"
	DataDir           = "data"
	UploadDir         = "upload"
	MetadataDir       = "metadata"
	StorageDir        = "storage"
	ChunkSize         = 1024 * 1024
	MaxRetries        = 3
	ReplicationFactor = 2
	ReconstructedFile = "reconstructed.bin"
)

var DefaultNodes = []string{
	"http://localhost:8001",
	"http://localhost:8002",
	"http://localhost:8003",
}
