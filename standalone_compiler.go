//go:build !adjoining_go_compiler

package goembed

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed internal/golang/bin/*
//go:embed internal/golang/pkg/*
//go:embed internal/golang/src/*
var embedCompiler embed.FS

const allowsCache = true
const rootFolder = "internal/golang"
const cacheFolder = "./gocmp"

func isCached() bool {
	if _, err := os.Stat(cacheFolder); os.IsNotExist(err) {
		return false
	}
	return true
}

func withEmbed(noCache bool, callback func(string) error) error {
	if !isCached() {
		if err := createCache(); err != nil {
			return err
		}
		if noCache {
			defer removeCache()
		}
	}
	return callback(filepath.Clean(cacheFolder))
}

func createCache() error {
	os.Mkdir(cacheFolder, os.ModePerm)
	return fs.WalkDir(embedCompiler, rootFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(rootFolder, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(cacheFolder, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		data, err := fs.ReadFile(embedCompiler, path)
		if err != nil {
			return err
		}

		return os.WriteFile(destPath, data, 0o644)
	})
}

func removeCache() {
	os.RemoveAll(cacheFolder)
}
