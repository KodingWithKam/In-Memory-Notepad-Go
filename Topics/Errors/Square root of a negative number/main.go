package main

import (
	"errors"
	"fmt"
	"math"
)

func negativeNumError() error {
	// create a custom error message within the 'err' variable here
	err := errors.New("math: can't calculate square root of negative number")

	return err
}

// DO NOT delete or modify the code within the main() function!
func main() {
	var num float64
	fmt.Scanf("%f", &num)

	if num < 0 {
		err := negativeNumError()
		fmt.Println(err)
		return
	}
	fmt.Println(math.Sqrt(num))
}
