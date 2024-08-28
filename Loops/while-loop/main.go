package main

import "fmt"

func main() {
	test(2, 78)
	test(3, 78)
	test(4, 80)
}

func getMaxMessagesSend(costMultiplier int, maxCost int) (maxMessages int) {

	actualCost := 1
	maxMessages = 0
	for actualCost <= maxCost {
		maxMessages++
		actualCost *= costMultiplier
	}
	return maxMessages
}

func test(costMultiplier int, maxCost int) {
	fmt.Printf("The maximum pennies you can spend it %v", maxCost)
	fmt.Println()
	maxMessages := getMaxMessagesSend(costMultiplier, maxCost)
	fmt.Printf("The max messages is %v", maxMessages)
}
