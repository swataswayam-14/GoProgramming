package main

import (
	"errors"
	"fmt"
)

func main() {
	test(12, 4)
	test(9, 0)
	test(-5, 8)
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Division by 0 is not possible")
	}
	return x / y, nil
}

func test(dividend, divisor float64) {
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
}
