//go:build !adjoining_go_compiler

package compiler

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	fs, _ := embedCompiler.ReadDir(rootFolder)
	var folders = []string{"bin", "pkg", "src"}
	if len(fs) == 0 {
		t.Error("No files found")
		t.FailNow()
	}
	for index, file := range fs {
		if file.Name() != folders[index] {
			t.Errorf("Expecting %q, but found %q instead!", folders[index], file.Name())
			t.FailNow()
		}
	}
}

func TestCache(t *testing.T) {
	createCache()
	if !isCached() {
		t.Error("Cache not found!")
		t.FailNow()
	}
	removeCache()
	if isCached() {
		t.Error("Cache should have been removed!")
		t.FailNow()
	}
}
