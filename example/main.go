package main

import (
	"github.com/JessebotX/cog"
)

var cli struct {
	Init    InitCommand `cmd:"init" help:"Initialize a project"`
	Verbose bool        `flag:"verbose" help:"Print debug information"`
}

type InitCommand struct {
	InitThreads int `flag:"init-threads" help:"Max number of threads to use"`
}

func main() {
	args := []string{
		"example (program name)",
		"-h",
		"-verbose",
		"-init-threads",
		"2",
		"init",
	}

	_, _ = cog.Parse(args, &cli)
}
