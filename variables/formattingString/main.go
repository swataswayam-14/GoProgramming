package main

import "fmt"

func main() {
	const name = "Swata Swayam Dash"
	const openRate = 30.52523

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)
	fmt.Println(msg)
}
