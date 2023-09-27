package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dcaravel/broken-image-reg/internal/env"
	"github.com/google/go-containerregistry/pkg/registry"
	"github.com/google/go-containerregistry/pkg/registryfaker"
)

func main() {
	dirname := getAndPrepStorageDir()
	reg := registryfaker.New(
		registryfaker.WithBlobHandler(registryfaker.NewDiskBlobHandler(dirname)),
	)

	addr := fmt.Sprintf("%s:%s", env.BindHost, env.BindPort)
	log.Printf("Listening on %q", addr)
	if env.CertFile.String() != "" {
		if err := http.ListenAndServeTLS(addr, env.CertFile.String(), env.KeyFile.String(), reg); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := http.ListenAndServe(addr, reg); err != nil {
			log.Fatal(err)
		}
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

func manifestHook(resp http.ResponseWriter, req *http.Request, repo, target string) bool {
	if target == "latest" {
		registry.WriteErr(resp, 403, "UNSUPPORTED", "latest tag is not allowed")
		return true
	}
	return false
}
