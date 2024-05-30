package main

import "fmt"

func main() {
	subvalue := sub(10, 5)
	fmt.Printf("The value is %v", subvalue)

	if length := getLength("swatasj@gmail.com"); length >= 12 {
		fmt.Println("valid email")
	}
	test("swata swayam", " dash")
	test("hello", " you")
	test("kashmir", " the best")

	fmt.Printf("The value is %v", add(2, 3))

}
func sub(x int, y int) int {
	return x - y
}

func getLength(email string) int {
	return len(email)
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}
func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func add(x, y int) int {
	return x + y
}
