package main

import "fmt"

type rect struct {
	height int
	width  int
}

func (r rect) area() int {
	return r.height * r.width
}

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return a.username + ":" + a.password
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("===========================================")
}

func main() {
	r := rect{
		width:  12,
		height: 232,
	}
	fmt.Println("the area is ", r.area())
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
	test(authenticationInfo{
		username: "IBM",
		password: "23462384afasfaf",
	})

}
