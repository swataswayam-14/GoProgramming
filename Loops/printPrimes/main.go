package main

import "fmt"

func main() {
	test(100)
}

func test(number int) {
	fmt.Printf("Printing prime numbers till %v", number)
	printPrimes(number)
}
func printPrimes(max int) {
	for n := 0; n < max; n++ {
		if n == 2 {
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i <= n/2; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}
