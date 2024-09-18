package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("Width %0.2f is incorrect", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("height %0.2f is incorrect", height)
	}
	area := width * height
	return area / 10.0, nil // when is good we return value and nil
}

func main() {

	amount, err := paintNeeded(-4.2, 3.0)

	// //Print only err value
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Paint needed %.2f L\n", amount)
	// }

	// fmt.Println("blabla")

	// //Create log and exit status 1
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Paint needed %.2f L\n", amount)
	fmt.Println("blabla")
}
