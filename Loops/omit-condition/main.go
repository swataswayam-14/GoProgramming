package main

import "fmt"

func main() {
	test(358)
	test(0)
	test(12)
}

func maxMessages(thresh int) int {
	total_cost := 0.0
	for i := 0; ; i++ {
		total_cost += 1 + float64(i)*0.01
		if total_cost > float64(thresh) {
			return int(i)
		}
	}
}

func test(threshold int) {
	fmt.Printf("Total amount sent is %v", threshold)
	fmt.Println()
	fmt.Printf("The maximum messages that can be sent is %v", maxMessages(threshold))
}
