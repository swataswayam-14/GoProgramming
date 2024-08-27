package main

import "fmt"

type expense interface {
	cost() float64
}
type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}
type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}
type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}
func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}
	return float64(len(s.body)) * 0.03
}
func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)

	if ok {
		return em.toAddress, em.cost()
	}
	sm, ok := e.(sms)

	if ok {
		return sm.toPhoneNumber, sm.cost()
	}
	return "", 0.0
}
func estimateYearlyCost(e expense, averageMessagePerYear int) float64 {
	return e.cost() * float64(averageMessagePerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("======================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("======================")
	case invalid:
		fmt.Printf("Report: Invalid message")
		fmt.Println("======================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "Hey , what's up , how its going??",
		toAddress:    "safjba@gmail.com",
	})
	test(email{
		isSubscribed: false,
		body:         "Its me sending message without subscribing",
		toAddress:    "jhjsdf@gmail.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "Its me sending message (sms) without subscribing",
		toPhoneNumber: "+73 234324241",
	})
	test(sms{
		isSubscribed:  true,
		body:          "Its me sending message (sms) with subscribing",
		toPhoneNumber: "+73 234324241",
	})
	test(invalid{})
}
