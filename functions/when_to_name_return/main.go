package main

import "errors"

func calc1(a, b int) (mul, div int, err error) { //It is easier to understand than the below function:
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul = a * b
	div = a / b

	return mul, div, nil
}

func calc2(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul := a * b
	div := a / b

	return mul, div, nil
}
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}

//Naked returns should only be used in short and simple functions
