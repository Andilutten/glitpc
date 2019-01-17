package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

const (
	// BlockPatternTemplate is used to create a regexp pattern
	// that matches blocks in specified language
	BlockPatternTemplate = "(?msU)^\x60{3}%s$\n(.*)^\x60{3}$"
)

// Check error and log.fatal if error
// is not nil
func Check(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

// Compile takes the inputs and reads extracts all of the
// source blocks that should be grabbed from the sources
func Compile(target io.Writer, opts Options, input []io.Reader) error {
	// Create a regexp for the block pattern
	re := regexp.MustCompile(fmt.Sprintf(BlockPatternTemplate, opts.Language))
	for _, reader := range input {
		// Copy all of the data from the input file
		// to a buffer to be matched from
		buffer := new(bytes.Buffer)
		_, err := io.Copy(buffer, reader)
		Check(err, "Could not copy from file to buffer")
		// Iterate over all matches
		for _, match := range re.FindAllSubmatch(buffer.Bytes(), -1) {
			// Read from all submatches
			for _, submatch := range match[1:] {
				fmt.Fprint(target, string(submatch))
			}
		}
	}
	return nil
}

func main() {
	// Read options from flags
	options := parseOptions()
	// Open all input files that are supposed
	// to be read by the compiler
	inputs := []io.Reader{}
	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		Check(err, "Could not open file")
		defer file.Close()
		inputs = append(inputs, file)
	}
	// Create the output file
	// TODO: Should try to open the file if it already exists
	output, err := os.Create(options.Output)
	Check(err, "Could not create output file")
	defer output.Close()
	// Compile the targets in to the output file
	Compile(output, *options, inputs)
}
