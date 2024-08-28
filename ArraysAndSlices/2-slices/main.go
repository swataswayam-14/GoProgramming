package main

import (
	"errors"
	"fmt"
)

func main() {
	test("Swata Swayam Dash", 2, "free")
	test("Amit", 2, "free")
	test("Amit", 2, "Pro")
}
func test(name string, doneAt int, plan string) {
	fmt.Printf("Sending messages to %v", name)
	fmt.Println()
	messages, err := sendMessages(plan)

	if err != nil {
		fmt.Printf("Failed to send message, error : %v", err)
		fmt.Println()
		return
	}
	for i := 0; i < len(messages); i++ {
		fmt.Printf("Sent message: %v", messages[i])
		fmt.Println()
		if i == doneAt {
			fmt.Println("User responded successfully")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("Complete failure")
			break
		}
	}
}

const (
	planFree = "free"
	planPro  = "Pro"
)

func getMessagesWithRetry() [3]string {
	return [3]string{
		"Please signup",
		"please yaar sign up karle",
		"otsukaresamadesu , please bhai tune bahat kaam kiya hai, abhi toh karle",
	}
}

func sendMessages(plan string) ([]string, error) {

	messages := getMessagesWithRetry()

	if plan == planPro {
		return messages[:], nil
	} else if plan == planFree {
		return messages[0:2], nil
	} else {
		return nil, errors.New("Invalid plan")
	}

}
