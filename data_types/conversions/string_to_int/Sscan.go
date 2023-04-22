// Using fmt.Sscan
// The fmt package provides sscan() function which scans string argument and store into variables.
// This function read the string with spaces and assign into consecutive Integer variables.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	strVar := "100"
	intValue := 0
	_, err := fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue))
}

// 100 <nil> int
