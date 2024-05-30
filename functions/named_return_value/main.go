package main

import "fmt"

func main() {
	test(18)
	test(28)
	test(2)
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	yearsUntilDrinking = 21 - age
	yearsUntilCarRental = 25 - age

	if yearsUntilCarRental < 0 {
		yearsUntilAdult = 0
		yearsUntilCarRental = 0
		yearsUntilDrinking = 0

		return
	}

	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
		yearsUntilAdult = 0
		return
	}

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
		return
	}
	return
}

func test(age int) {
	fmt.Println("Your age is ", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("you are an adult in ", yearsUntilAdult, " years")
	fmt.Println("you are eligible to have alcohol in ", yearsUntilDrinking, " years")
	fmt.Println("You are eligible for car rental in ", yearsUntilCarRental, " years")
	fmt.Println("======================================================")
}

func nakedReturn() (x, y int) {
	return
}

//above is same as

func namedReturn() (int, int) {
	var x int
	var y int

	return x, y
}
