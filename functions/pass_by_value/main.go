package main

import "fmt"

func main() {
	sentSoFar := 231
	const sentToAdd = 12
	sentSoFar = increment(sentSoFar, sentToAdd)
	fmt.Printf("The messages sent by you so far is %v", sentSoFar)

}
func increment(sentSoFar, sentToAdd int) int {
	return sentSoFar + sentToAdd
}
