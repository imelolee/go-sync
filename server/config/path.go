package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func DefaultPaths() (uploads string) {
	base, err := os.Getwd()
	if err != nil {
		log.Fatal(base)
	}

	dir := filepath.Dir(path.Join(base, "go-sync"))
	uploads = filepath.Join(dir, "uploads")
	return
}

var UploadsDir = DefaultPaths()