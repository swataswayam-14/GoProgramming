// To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:
package main

import "fmt"

type car struct {
	Make   string
	Model  string
	Height string
	Width  string
	// Wheel is a field containing an anonymous struct
	Wheel struct {
		Radius   string
		Material string
	}
}

func main() {
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "tesla",
		Model: "model 3",
	}

	fmt.Printf("The model of the car is %s", myCar.Model)
}
