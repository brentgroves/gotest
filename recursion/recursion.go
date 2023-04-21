// Recursive functions are very useful to solve many mathematical problems such as calculating factorial of a number,
// generating a Fibonacci series, etc.

package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	var i int = 2
	fmt.Printf("Factorial of %d is %d", i, factorial(i))
}
