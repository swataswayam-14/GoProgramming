package main

import "fmt"

func main() {
	x, _ := getTwoValue() //Even though `getPoint()` returns two values, we can capture the first one and ignore the second.
	fmt.Printf("The value of x is %v", x)

	_, lastName := getName()
	fmt.Printf("The value of last name is %v", lastName)

}

func getTwoValue() (x int, y int) {
	return 2, 4
}

//A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: `_`

//This is crucial to understand because the Go compiler will throw an error if you have unused variable declarations in your code

func getName() (f string, l string) {
	return "Baba", "Black sheep"
}
