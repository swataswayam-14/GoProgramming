package main

import "fmt"

type expense interface {
	cost() float64
}
type printer interface {
	print()
}
type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}
	return float64(len(e.body)) * 0.01
}

func (e email) print() {
	fmt.Println(" " + e.body)
}
func test(e expense, p printer) {
	fmt.Printf("printing with cost $%.2f", e.cost())
	p.print()
	fmt.Println("===================================")
}

func main() {
	e1 := email{
		isSubscribed: true,
		body:         "Dilko tumse pyaar hua pehli baar hua",
	}
	test(e1, e1)
	e2 := email{
		isSubscribed: false,
		body:         "Deedar aya hai",
	}
	test(e2, e2)
}
