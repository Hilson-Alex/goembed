//go:build !adjoining_go_compiler

//go:generate 7z a -r .\internal\golang\golang.zip .\internal\golang\bin\ .\internal\golang\pkg\ .\internal\golang\src\

package compiler

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"io"
	"os"
	"path/filepath"
)

//go:embed internal/golang/golang.zip
var embedCompiler []byte

const allowsCache = true

// const rootFolder = "internal/golang"
const cacheFolder = "/gocmp"

var CacheRoot = "."

func getCache() string {
	return filepath.Join(CacheRoot, cacheFolder)
}

func isCached() bool {
	if _, err := os.Stat(getCache()); os.IsNotExist(err) {
		return false
	}
	return true
}

func withEmbed(noCache bool, callback func(string) error) error {
	if !isCached() {
		if err := createCache(); err != nil {
			return err
		}
	}
	if noCache {
		defer removeCache()
	}
	return callback(getCache())
}

func createCache() error {
	var bReader = bytes.NewReader(embedCompiler)
	var wkdir = getCache()

	r, err := zip.NewReader(bReader, bReader.Size())
	if err != nil {
		return err
	}

	for _, file := range r.File {
		destPath := filepath.Join(wkdir, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(destPath, os.ModePerm)
			continue
		}
		if err := writeFile(destPath, file); err != nil {
			return err
		}
	}
	return nil
}

func writeFile(destPath string, file *zip.File) error {
	archiveFile, err := file.Open()
	defer func() { archiveFile.Close() }()
	if err != nil {
		return err
	}

	if os.MkdirAll(filepath.Dir(destPath), os.ModePerm) != nil {
		return err
	}
	destinationFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer func() { destinationFile.Close() }()

	_, err = io.Copy(destinationFile, archiveFile)
	return err
}

func removeCache() {
	os.RemoveAll(getCache())
}
