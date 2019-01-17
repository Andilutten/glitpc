# GO literate programming compiler
Go literate programming compiler (glitpc for short) is a compiler 
used for compiling of markdown documents written in with a literate programming
paradigm. It takes one or several markdown documents as input and compiles
them down into a file with all source targets extracted out.

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