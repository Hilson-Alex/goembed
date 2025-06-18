// Copyright 2025 Hilson Alexandre Wojcikiewicz Junior. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This library embeds a go compiler into your system. It was develop to create a compiler
// that uses go as intermediate code to build the end artifact.
//
// The library has also options to use the system downloaded go compiler when possible, or
// skip it entirely.
//
// If there's no need for the embed compiler, you can add the tag adjoining_go_compiler
// (e.g. -tags=adjoining_go_compiler) when building your application and the embed compiler
// will be left out, at the cost of throwing an error if the computer doesn't have go installed.
// This can reduce the binary size, but will make the code dependent of the system environment.
package goembed

import (
	"io"
	"os/exec"
	"path/filepath"

	"github.com/Hilson-Alex/goembed/props"
)

// Build a go package with the given options
func GoBuild(pkg string, options *CompilerOptions) error {
	if useEmbed(options.bypassSystemCompiler) {
		return withEmbed(options.noCache, func(cacheFolder string) error {
			return execute(
				filepath.Join(cacheFolder, "bin/go.exe"),
				*options.env.Copy().Append("GOROOT", cacheFolder),
				goBuildArgs(pkg, options.args),
				options.stderr, options.stdout,
			)
		})
	}
	return execute("go", *options.env, goBuildArgs(pkg, options.args), options.stderr, options.stdout)
}

// Format a go package with the given options
func GoFmt(pkg string, options *CompilerOptions) error {
	if useEmbed(options.bypassSystemCompiler) {
		return withEmbed(options.noCache, func(cacheFolder string) error {
			return execute(
				filepath.Join(cacheFolder, "bin/gofmt.exe"),
				*options.env.Copy().Append("GOROOT", cacheFolder),
				*options.args.Copy().Append(pkg, ""),
				options.stderr, options.stdout,
			)
		})
	}
	return execute("gofmt", *options.env, *options.args.Copy().Append(pkg, ""), options.stderr, options.stdout)
}

func useEmbed(bypassSystemCompiler bool) bool {
	return (allowsCache && bypassSystemCompiler) || !goExists()
}

func goBuildArgs(pkg string, args *props.Properties) []string {
	return append([]string{"build"}, *args.Copy().Append(pkg, "")...)
}

func execute(command string, env, args []string, stdout, stderr io.Writer) error {
	cmd := exec.Command(command, args...)
	cmd.Env = env
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}

func goExists() bool {
	_, err := exec.LookPath("go")
	return err == nil
}
