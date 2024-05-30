package main

import "fmt"

type messageToSend struct {
	message     string
	phoneNumber int
}

func test(m messageToSend) {
	fmt.Printf("Sending message from %s to %v", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message:     "Hello , sharp at 6 water bottle and jacket in bbsr",
		phoneNumber: 87542381,
	})
	test(messageToSend{
		message:     "isgue siuf iuesf vsi fsdifhsdf ",
		phoneNumber: 873423874,
	})
}
