// Scope of Golang Variables Defined by Brace Brackets
// In the Go language, the scope of a variable refers to the region of the code where the variable is accessible.
// The scope of a variable is determined by where it is declared.

// A variable declared within a block (e.g., inside a function or a control structure such as if, for, or switch)
// is only accessible within that block and any nested blocks. Once the execution leaves the block, the variable goes
// out of scope and cannot be accessed anymore.

package main

import (
	"fmt"
)

var s = "Japan"

func main() {
	fmt.Println(s) // Accessing the global variable 's' and printing its value
	x := true      // Declaring and initializing a local variable 'x'

	if x { // If 'x' is true, execute the block of code inside the if statement
		y := 1 // Declaring and initializing a local variable 'y' inside the if block

		if x != false { // If 'x' is not false, execute the block of code inside the if statement
			fmt.Println(s) // Accessing the global variable 's' and printing its value
			fmt.Println(x) // Printing the value of local variable 'x'
			fmt.Println(y) // Printing the value of local variable 'y'
		}
	}

	fmt.Println(x) // Printing the value of local variable 'x'
}

// In the main function, we have three variables with different scopes:
// s: This is a global variable and can be accessed from any function within the package, including the main function.
// x: This is a local variable declared and initialized inside the main function. It is accessible only within the main function.
// In this case, it is also accessible inside the if blocks within the main function.
// y: This is a local variable declared and initialized inside the first if block. Its scope is limited to the block it is declared in,
// so it is accessible only within that if block and its nested blocks. In this case, it is also accessible inside the second if block,
// which is nested inside the first if block.
