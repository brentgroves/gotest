// https://go.dev/doc/tutorial/getting-started
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

// This is your Go code. In this code, you:

// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
// Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
// Implement a main function to print a message to the console. A main function executes by default when you run the main package.
// Run your code to see the greeting.

// $ go run .
// Hello, World!

// Add new module requirements and sums.
// go get rsc.io/quote
// Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module. For more, see Authenticating modules in the Go Modules Reference.

// $ go mod tidy
// go: finding module for package rsc.io/quote
// go: found rsc.io/quote in rsc.io/quote v1.5.2
