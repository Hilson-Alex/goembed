//go:build !adjoining_go_compiler

package compiler

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	if len(embedCompiler) == 0 {
		t.Error("No file found")
		t.FailNow()
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
