package main

import (
	"log"
	"os"
)

// Files and directories with examples
// The most important package that allows us to manipulate files and directories as entities is the os package.
// The io package has the io.Reader interface to reads and transfers data from a source into a stream of bytes.
// The io.Writer interface reads data from a provided stream of bytes and writes it as output to a target resource.

// Create an empty file

func main() {
	emptyFile, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}
