package main

import "fmt"

func main() {
	height := 4
	messageLen := 12
	maxMessageLen := 20

	fmt.Println("Trying to print message of length %v", messageLen, "and a max message length of %v", maxMessageLen)

	if height > 6 {
		fmt.Println("You are damm tall")
	} else if height >= 4 {
		fmt.Println("you are tall enough")
	} else {
		fmt.Println("You are small")
	}

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	}
}
