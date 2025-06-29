package compiler

import (
	"io"
	"os"

	"github.com/Hilson-Alex/goembed/props"
)

// Options to pass to the command line of the compiler
type CompilerOptions = struct {
	bypassSystemCompiler bool              // If false, uses an already installed go compiler when available
	noCache              bool              // Clean generated files after use.
	env                  *props.Properties // Env flags to use on compilation
	args                 *props.Properties // Command Line Arguments for the compiler
	stdout, stderr       io.Writer         // The Standard output and error for the command
}

// Functions used to customize the command options
type CustomOption func(options *CompilerOptions)

// Build the command options based on the customizations passed.
// Use BuildOptions() to get the default values
func BuildOptions(customOptions ...CustomOption) *CompilerOptions {
	var options = &CompilerOptions{
		env:     props.DefaultEnv(),
		args:    props.Empty(),
		stdout:  os.Stdout,
		stderr:  os.Stderr,
		noCache: false,
	}
	for _, customize := range customOptions {
		customize(options)
	}
	return options
}

// Force to use the embed compiler, even when there's a go compiler already installed
func BypassInstalledCompiler() CustomOption {
	return func(options *CompilerOptions) {
		options.bypassSystemCompiler = true
	}
}

// Delete the embed compiler after use. It'll be recreated and deleted on every use
func NoCache() CustomOption {
	return func(options *CompilerOptions) {
		options.noCache = true
	}
}

// Overwrites the env. Default env is based on [os.Environ()]
//
// If you want to keep the default while adding new variables,
// see [props.DefaultEnv()]
func WithEnv(env *props.Properties) CustomOption {
	return func(options *CompilerOptions) {
		options.env = env
	}
}

// Sets the command line arguments. Default is empty
func WithArgs(args *props.Properties) CustomOption {
	return func(options *CompilerOptions) {
		options.args = args
	}
}

// Use custom stdout for the execution. default is os.Stdout
func WithStdOut(stdout io.Writer) CustomOption {
	return func(options *CompilerOptions) {
		options.stdout = stdout
	}
}

// Use custom stderr for the execution. default is os.Stderr
func WithStdErr(stderr io.Writer) CustomOption {
	return func(options *CompilerOptions) {
		options.stderr = stderr
	}
}
