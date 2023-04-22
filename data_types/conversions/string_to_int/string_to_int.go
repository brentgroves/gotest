// https://www.golangprograms.com/go-language/integer-float-string-boolean.html
// How to Convert string to integer type in Go?
// Like most modern languages, Golang includes strings as a built-in type. Let's take an example,
// you may have a string that contains a numeric value "100". However, because this value is represented as a string,
// you can't perform any mathematical calculations on it. You need to explicitly convert this string type into an
// integer type before you can perform any mathematical calculations on it. In order to convert string to integer
// type in Golang, you can use the following methods.

// You can use the strconv package's Atoi() function to convert the string into an integer value.
// Atoi stands for ASCII to integer. The Atoi() function returns two values: the result of the conversion, and the error (if any).

// func Atoi(s string) (int, error)

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

// 100 <nil> int
