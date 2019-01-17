package main

import (
	"flag"
	"fmt"
	"os"
)

type (
	// Options object is basicly
	// a container for flag values
	Options struct {
		Output   string
		Language string
		Stdin    bool
	}
)

func parseOptions() *Options {
	// Parse flags TODO: These should be better documented with the usage param
	output := flag.String("output", "a.mdc", "Set the output filename")
	language := flag.String("language", "bash", "Set the markdown language flag to look for in input file")
	stdin := flag.Bool("stdin", false, "Set to true if the content to be compiled comes from stdin")
	flag.Parse()

	// Get input filenames from remaining args
	input := flag.Args()
	if len(input) == 0 {
		// TODO: Print usage and give example of usage
		fmt.Println("No input files was specified")
		os.Exit(1)
	}

	return &Options{
		Output:   *output,
		Language: *language,
		Stdin:    *stdin,
	}
}
