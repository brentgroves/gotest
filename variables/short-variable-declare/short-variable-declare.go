// Short Variable Declaration in Golang
// The := short variable assignment operator indicates that short variable declaration is being used.
// There is no need to use the var keyword or declare the variable type.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "John Doe"
	fmt.Println(reflect.TypeOf(name))
}
