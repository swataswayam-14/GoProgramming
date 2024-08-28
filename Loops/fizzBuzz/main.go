package main

import "fmt"

func main() {
	test(100)
}
func test(n int) {
	fmt.Printf("Printing fizzbuzz for n = %v ", n)
	fmt.Println()
	fizzBuzz(n)
}

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
