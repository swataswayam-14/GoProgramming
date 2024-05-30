package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if (len(mToSend.recipient.name) > 0 && mToSend.recipient.number != 0) && (len(mToSend.sender.name) > 0 && mToSend.sender.number != 0) {
		return true
	}
	return false
}

func test(mToSend messageToSend) {
	fmt.Printf(`Sending "%s" from "%s (%v) to %s (%v)....`, mToSend.message, mToSend.sender.name, mToSend.sender.number, mToSend.recipient.name, mToSend.recipient.number)
	fmt.Println()

	if canSendMessage(mToSend) {
		fmt.Println("....send!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("=========================================================")
}

func main() {
	test(messageToSend{
		message: "Hey how are you",
		sender: user{
			name:   "swata swayam dash",
			number: 983683274,
		},
		recipient: user{
			name:   "ayash rath",
			number: 2386432832,
		},
	})

	test(messageToSend{
		message: "lkhsfsifalsfa",
		sender: user{
			number: 876324328,
		},
		recipient: user{
			name:   "sdkjfskff",
			number: 23434124,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
}
