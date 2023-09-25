package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dcaravel/broken-image-reg/internal/env"
	"github.com/google/go-containerregistry/pkg/registry"
)

func main() {
	reg := registry.New()

	addr := fmt.Sprintf("%s:%s", env.BindHost, env.BindPort)
	log.Printf("Listening on %q", addr)
	if err := http.ListenAndServe(addr, reg); err != nil {
		log.Fatal(err)
	}
}
