package main

import "fmt"

type laptopSize int

func (ls laptopSize) getSizeOfLaptop() {
	fmt.Println(ls)
}

func main() {
	var laptop laptopSize = 15
	laptop.getSizeOfLaptop()

}
