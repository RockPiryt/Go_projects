package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	//Embedded structs
	person3 := person{
		firstName: "Jim",
		lastName:  "Drop",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Create pointer and use pointer's method
	person3Pointer := &person3
	person3Pointer.updateName("Tommy")

	// Version without write pointer
	person3.updateName("Placek")

	person3.print()
}

//pointer's method
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//Method of struct
func (p person) print() {
	fmt.Printf("%+v", p)
}
