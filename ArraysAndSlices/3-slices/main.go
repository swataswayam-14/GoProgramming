package main

import "fmt"

func getMessagesCost(messages []string) []float64 {

	messageCostArray := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		length := len(messages[i])
		cost := float64(length) * 0.01
		messageCostArray[i] = cost
	}
	return messageCostArray
}

func test(messages []string) {
	fmt.Println("Calculating the total cost of sending these messages")
	costArray := getMessagesCost(messages)
	fmt.Println("Detailed bill: ")
	for i := 0; i < len(messages); i++ {
		fmt.Printf("%v: %.2f", messages[i], costArray[i])
		fmt.Println()
	}
}
func main() {
	test([]string{
		"Swata Swayam Dash is the best",
		"he can achieve anything he wants",
		"he got that potential , determination and desire to bring the goals to physical equivalence",
	},
	)
	test([]string{
		"hello",
		"high",
		"hiyaaa",
	},
	)
}
