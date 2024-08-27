package main

import "fmt"

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = 0.0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v character", len(message))
	}
	return costPerChar * float64(len(message)), nil
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return costSpouse + cost, nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Printf("===========")
	total_cost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", total_cost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}
