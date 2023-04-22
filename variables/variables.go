// The most general form to declare a variable in Golang uses the var keyword, an explicit type, and an assignment.

// var name type = expression
package main

import "fmt"

func main() {
	var i int = 10
	var s string = "Canada"

	fmt.Println(i)
	fmt.Println(s)
}
