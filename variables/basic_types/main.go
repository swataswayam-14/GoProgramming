package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerMessage float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerMessage, hasPermission, username)
}
