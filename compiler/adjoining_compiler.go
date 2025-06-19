//go:build adjoining_go_compiler

package compiler

import "errors"

const allowsCache = false

func isCached() bool {
	return false
}

func createCache() error {
	return errors.New("This build needs a go compiler to work. Try installing a go compiler or get the standalone build")
}

func removeCache() {}

func withEmbed(noCache bool, callback func(string) error) error {
	return createCache()
}
