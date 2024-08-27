package main

import "fmt"

func returnSmsError(cost float64, recipient string) string {
	return fmt.Sprintf("The message having cost $%.2f cannot be sent to %v", cost, recipient)
}

func test(cost float64, recipient string) {
	s := returnSmsError(cost, recipient)
	fmt.Println(s)
	fmt.Println("=====================")
}

func main() {
	test(36.23, "dsgfdshbfshgfs")
	test(32.23, "dsjhfsmnbfsjf")
	test(36.23, "dsgfdshbfshgfs")
	test(32.23, "dsjhfsmnbfsjf")
}
