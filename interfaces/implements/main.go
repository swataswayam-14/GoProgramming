package main

import "fmt"

type employee interface {
	getSalary() float64
	getName() string
}
type contractor struct {
	name        string
	hourlyRate  float64
	hoursInYear int
}
type fulltime struct {
	name   string
	salary float64
}

func (c contractor) getSalary() float64 {
	return c.hourlyRate * float64(c.hoursInYear)
}
func (c contractor) getName() string {
	return c.name
}
func (f fulltime) getSalary() float64 {
	return f.salary
}
func (f fulltime) getName() string {
	return f.name
}
func test(e employee) {
	fmt.Printf("Employee name: %v , Employee salary: %v", e.getName(), e.getSalary())
	fmt.Println("==============================")
}

func main() {
	contractEmployee := contractor{
		name:        "steguf",
		hourlyRate:  30,
		hoursInYear: 30,
	}
	fullTimeEmployee := fulltime{
		name:   "billu",
		salary: 1000,
	}
	test(contractEmployee)
	test(fullTimeEmployee)
}
