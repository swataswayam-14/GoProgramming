package main

import "fmt"

func sum(nums ...float64) float64 {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
func test(nums ...float64) {
	fmt.Printf("Calculating sum of %v costs...\n", len(nums))
	total := sum(nums...)
	fmt.Printf("Bill for the entire month: %.2f\n", total)
	fmt.Println("======END REPORT======")
}
func main() {
	test(1.0, 2.0, 21.12)
	test(23.32, 23.53, 65.56, 345.43, 0, 43, 343.4)
}
