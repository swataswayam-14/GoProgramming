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

type superWheel struct {
	radius int
	height int
}

type superCar struct {
	model        string
	chesisNumber int
	frontWheel   superWheel
	rearWheel    superWheel
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
	mySuperCar := superCar{
		model:        "Ferrai",
		chesisNumber: 324,
		frontWheel: superWheel{
			radius: 74,
			height: 88,
		},
		rearWheel: superWheel{
			radius: 63,
			height: 78,
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
	fmt.Printf("%s %v %v %v", mySuperCar.model, mySuperCar.chesisNumber, mySuperCar.frontWheel.radius, mySuperCar.rearWheel.height)
}
