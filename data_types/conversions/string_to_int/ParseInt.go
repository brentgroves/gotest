// ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
// This function accepts a string parameter, convert it into a corresponding int type based on a base parameter. By default, it returns Int64 value.
// func ParseInt(s string, base int, bitSize int) (i int64, err error)

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"

	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))
}

// 100 <nil> int64
// 100 <nil> int64
// 100 <nil> int64
// 100 <nil> int64
