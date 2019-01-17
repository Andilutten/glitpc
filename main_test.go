package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type (
	testCase struct {
		desc    string
		options Options
		input   []io.Reader
		output  string
	}
	file struct {
		content string
		cursor  int
	}
)

func (f *file) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		if i+f.cursor < len(f.content) {
			n++
			p[i] = f.content[i+f.cursor]
		} else {
			return n, io.EOF
		}
	}
	f.cursor += len(p) - 1
	return
}

func createTestFile(contents ...string) *file {
	return &file{strings.Join(contents, "\n"), 0}
}

func Test_compile(t *testing.T) {
	options := Options{
		Stdin:    false,
		Language: "bash",
		Output:   "file.sh",
	}
	testCases := []testCase{
		{
			desc:    "No output",
			options: options,
			input: []io.Reader{
				createTestFile("Hello World"),
			},
			output: "",
		},
		{
			desc:    "One insertion",
			options: options,
			input: []io.Reader{
				createTestFile(
					"# Hello world",
					"This is a markdown file",
					"```bash",
					"export FOO=BAR",
					"```",
				),
			},
			output: "export FOO=BAR\n",
		},
		{
			desc:    "More insertions",
			options: options,
			input: []io.Reader{
				createTestFile(
					"# First",
					"```bash",
					"export FOO=BAR",
					"```",
					"# Second",
					"```bash",
					"export BAR=FOO",
					"```",
				),
			},
			output: "export FOO=BAR\nexport BAR=FOO\n",
		},
		{
			desc:    "Multiple files",
			options: options,
			input: []io.Reader{
				createTestFile(
					"# First",
					"```bash",
					"export FOO=BAR",
					"```",
				),
				createTestFile(
					"# Second",
					"```bash",
					"export BAR=FOO",
					"```",
				),
			},
			output: "export FOO=BAR\nexport BAR=FOO\n",
		},
	}
	for _, tC := range testCases {
		buffer := new(bytes.Buffer)
		if err := compile(buffer, tC.options, tC.input); err != nil {
			t.Fatal(err.Error())
		} else {
			if buffer.String() != tC.output {
				t.Fatalf("Got \"%s\" expected \"%s\"", buffer.String(), tC.output)
			}
		}
	}
}
