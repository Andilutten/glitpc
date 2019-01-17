# GO literate programming compiler
Go literate programming compiler (glitpc for short) is a compiler 
used for compiling of markdown documents written in with a literate programming
paradigm. It takes one or several markdown documents as input and compiles
them down into a file with all source targets extracted out.

### Disclaimer
This is just a project i made for my own use.

## Installation
If you have a golang binary installed on your computer and GOPATH+GOBIN setup, just run the following commands
```bash
go get github.com/Andilutten/glitpc
go install github.com/Andilutten/glitpc
```

## Usage
This is an example usage of how it works. Imagine we have a 
markdown file called **bashrc.md** with the following contents.
````markdown
# This is a .bashrc configuration
Written and maintained in markdown.
All it does is exposing some variables.
```bash
export FOO=BAR 
```
````

To compile this down into a bash source file we issue the command:
```bash
glitpc -output bashrc bashrc.md
```

This will create a file called bashrc with following contents:
```bash
export FOO=BAR
```