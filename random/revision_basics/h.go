package main

import "fmt"

func main() {
	car1, car2, car3 := createCars()
	fmt.Println(car1, car2, car3)
}

func createCars() (string, string, string) {
	return "Toyota", "BMW", "Volvo"

}
