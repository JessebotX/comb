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
	InitThreads int            `flag:"init-threads" help:"Max number of threads to use"`
	New         InitNewCommand `cmd:"new" help:"Initialize new project..."`
}

type InitNewCommand struct {
	TestInt int `flag:"test-int" help:"Test int value"`
}

func main() {
	args := []string{
		"example (program name)",
		"new",
		"-h",
		"-verbose",
		"-init-threads",
		"28",
		"init",
		"-test-float",
		"16.25",
		"-test-int",
		"16",
		"-test-str",
		"16",
	}

	context, err := comb.Parse(args, &cli)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", cli)
	log.Printf("%+v\n", context.Rest)
}
