package main

import "fmt"

const (
	retry1 = "please sign up"
	retry2 = "please I beg you sign up yaar"
	retry3 = "hum bhi paresan hai , woh bhi paresaan hai karlo abto"
)

func getMessagesWithRetry() [3]string {
	return [3]string{
		retry1,
		retry2,
		retry3,
	}
}

func sendMessages(name string, doneSteps int) {
	fmt.Printf("Sending messages to %v", name)
	fmt.Println()
	messages := getMessagesWithRetry()

	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf("Sending message:  %v", msg)
		fmt.Println()
		if i == doneSteps {
			fmt.Println("They responded")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("Complete failure")
		}
	}
}

func main() {
	sendMessages("swata swayam dash", 2)
	sendMessages("Divya", 0)
	sendMessages("Cadbury", 3)
}
