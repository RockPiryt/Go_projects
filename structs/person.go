package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

//Short version
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Deklaracja
	// person1 := person{firstName: "Alex", lastName: "Duda"}
	// fmt.Println(person1)

	//Inicjalizacja + update
	// var person2 person
	// person2.firstName = "Kate "
	// fmt.Printf("%+v", person2)

	// //Embedded structs
	// person3 := person{
	// 	firstName: "Jim",
	// 	lastName:  "Drop",
	// 	contact: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }

	//Embedded structs -short
	person3 := person{
		firstName: "Jim",
		lastName:  "Drop",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// fmt.Printf("%+v", person3)
	person3.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
