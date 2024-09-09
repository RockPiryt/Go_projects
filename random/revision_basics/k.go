package main

import "fmt"

func main() {
	for i := 4; i <= 21; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}

	}
}
