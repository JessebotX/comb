package main

import (
	"log"

	"github.com/JessebotX/comb"
)

var cli struct {
	Init      InitCommand `cmd:"init" help:"Initialize a project"`
	Verbose   bool        `flag:"verbose" help:"Print debug information"`
	TestFloat float64     `flag:"test-float" help:"Test float value"`
	TestStr   string      `flag:"test-str" help:"Test string value"`
	TestInt   int         `flag:"test-int" help:"Test int value"`
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
		"-test-float",
		"16.25",
		"-test-int",
		"16",
		"-test-str",
		"16",
	}

	_, err := comb.Parse(args, &cli)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", cli)
}
