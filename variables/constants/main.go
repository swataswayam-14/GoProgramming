package main

import "fmt"

func main() {
	const premiumPlanName = "premium plan"
	const basicPlanName = "basic plan"

	const firstName = "swata swayam"
	const lastName = "Dash"

	const secondsInMinute = 60
	const minuteInHour = 60

	//computed constants
	const fullName = firstName + " " + lastName
	const secondsInHour = minuteInHour * secondsInMinute

	fmt.Println("plan: ", premiumPlanName)
	fmt.Println("plan: ", basicPlanName)
	fmt.Println("Full Name: ", fullName)
	fmt.Println("Seconds in 1 hour: ", secondsInHour)

}
