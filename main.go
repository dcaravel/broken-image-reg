package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dcaravel/broken-image-reg/internal/env"
	"github.com/google/go-containerregistry/pkg/registry"
)

func main() {
	dirname := getAndPrepStorageDir()
	reg := registry.New(registry.WithBlobHandler(registry.NewDiskBlobHandler(dirname)))

	addr := fmt.Sprintf("%s:%s", env.BindHost, env.BindPort)
	log.Printf("Listening on %q", addr)
	if err := http.ListenAndServe(addr, reg); err != nil {
		log.Fatal(err)
	}
}

func getAndPrepStorageDir() string {
	dirname := env.BlobDir.Val()

	if dirname == "" {
		homedir, err := os.UserHomeDir()
		dirname = filepath.Join(homedir, "broken-reg")
		if err != nil {
			log.Fatalf("Unable to determine user home dir: %v", err)
		}
	}

	err := os.MkdirAll(dirname, 0700)
	if err != nil {
		log.Fatalf("Unable to create storage dir: %v", err)
	}

	return dirname
}
