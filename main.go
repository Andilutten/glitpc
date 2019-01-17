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

func check(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func compile(target io.Writer, opts Options, input []io.Reader) error {
	re := regexp.MustCompile(fmt.Sprintf(`(?msU)^\x60{3}%s$\n(.*)^\x60{3}$`, opts.Language))
	for _, reader := range input {
		buffer := new(bytes.Buffer)
		_, err := io.Copy(buffer, reader)
		check(err, "Could not copy from file to buffer")
		for _, match := range re.FindAllSubmatch(buffer.Bytes(), -1) {
			for _, submatch := range match[1:] {
				fmt.Fprint(target, string(submatch))
			}
		}
	}
	return nil
}

func main() {
	options := parseOptions()
	inputs := []io.Reader{}
	for _, filename := range flag.Args() {
		file, err := os.Open(filename)
		check(err, "Could not open file")
		defer file.Close()
		inputs = append(inputs, file)
	}
	output, err := os.Create(options.Output)
	check(err, "Could not create output file")
	defer output.Close()
	compile(output, *options, inputs)
}
