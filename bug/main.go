package main

import "fmt"

func main() {

	messagesFromSwayam := []string{
		"hey you , how you doin",
		"did you had lunch",
		"lets have a beer tonight",
	}
	numMessages := float64(len(messagesFromSwayam))
	constPerMessage := 0.02

	totalCost := constPerMessage * numMessages

	fmt.Printf("Swayam spent %.2f on text messages today\n", totalCost)
}
