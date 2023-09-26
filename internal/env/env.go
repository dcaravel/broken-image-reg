package env

import "os"

var (
	BindHost = registerStringEnv("BIND_HOST", "0.0.0.0")
	BindPort = registerStringEnv("BIND_PORT", "8080")

	BlobDir = registerStringEnv("BLOB_DIR", "")
)

type stringEnv struct {
	key string
	// default value
	def string
}

func registerStringEnv(key string, def string) *stringEnv {
	return &stringEnv{
		key: key,
		def: def,
	}
}

func (s *stringEnv) Key() string {
	return s.key
}

func (s *stringEnv) Val() string {
	v := os.Getenv(s.key)
	if v == "" {
		return s.def
	}

	return v
}

func (s *stringEnv) String() string {
	return s.Val()
}
