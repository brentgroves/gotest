// Go programming provides a pretty simple error handling framework with inbuilt error interface type of the following declaration −

// type error interface {
//    Error() string
// }

// Functions normally return error as last return value. Use errors.New to construct a basic error message as following −

// func Sqrt(value float64)(float64, error) {
//    if(value < 0){
//       return 0, errors.New("Math: negative number passed to Sqrt")
//    }
//    return math.Sqrt(value), nil
// }

// Use return value and error message.

// result, err:= Sqrt(-1)

// if err != nil {
//    fmt.Println(err)
// }

package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func main() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
