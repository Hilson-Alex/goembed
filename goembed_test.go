package goembed

import (
	"testing"
)

func TestEmbed(t *testing.T) {
	fs, _ := embedCompiler.ReadDir("internal/golang")
	var folders = []string{"bin", "pkg", "src"}
	if len(fs) == 0 {
		t.FailNow()
	}
	for index, file := range fs {
		if file.Name() != folders[index] {
			t.FailNow()
		}
	}
}
