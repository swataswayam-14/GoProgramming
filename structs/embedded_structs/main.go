package main

import "fmt"

type car struct {
	make  string
	model string
}

type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
	car
	bedSize int
}

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

func test(s sender) {
	fmt.Println("Sender name: ", s.name)
	fmt.Println("Sender number: ", s.number)
	fmt.Println("Sender rateLimit is: ", s.rateLimit)
	fmt.Println("====================================")
}

func main() {
	myTruck := truck{
		bedSize: 10,
		car: car{
			make:  "Minj ",
			model: "classy X",
		},
	}
	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{
		rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})
	test(sender{
		rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})

	test(sender{
		rateLimit: 872421,
		user: user{
			name:   "swayam",
			number: 8712341,
		},
	})

	fmt.Printf("%v %s %s ", myTruck.bedSize, myTruck.make, myTruck.model)
}
