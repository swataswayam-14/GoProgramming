package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1 + (float64(i) * 0.01)
	}
	return totalCost
}

func test(numMessages int) {
	fmt.Printf("Sending %v messages ", numMessages)
	total_cost := bulkSend(numMessages)
	fmt.Printf("The total cost of sending %v messages is %.2f", numMessages, total_cost)
}

func main() {
	test(30)
	test(10)
	test(20)
}
