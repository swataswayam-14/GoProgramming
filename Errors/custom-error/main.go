package main

import "fmt"

type divideCustomError struct {
	dividend float64
}

func (de divideCustomError) Error() string {
	return fmt.Sprintf("%v cannot be divided by zero", de.dividend)
}

func divide(divisor float64, dividend float64) (float64, error) {
	if divisor == 0 {
		return 0.0, divideCustomError{
			dividend: dividend,
		}
	}
	return dividend / divisor, nil
}

func test(divisor, dividend float64) {
	quotient, err := divide(divisor, dividend)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
}

func main() {
	test(12, 76)
	test(9, 45)
	test(0, 76)
}
