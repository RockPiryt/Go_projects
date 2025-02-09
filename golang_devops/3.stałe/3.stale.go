package main

import "fmt"

const MyVar1 = "This is sting"
const MyVar2 = 123
const MyVar3 = -1.23
const MyVar4 int64 = 1234

const (
	one = iota
	two
	three
	four
)


func main() {

	fmt.Printf("stała ma wartość %v i typ %T\n", MyVar1, MyVar1)
	fmt.Printf("stała ma wartość %v i typ %T\n", MyVar2, MyVar2)
	fmt.Printf("stała ma wartość %v i typ %T\n", MyVar3, MyVar3)
	fmt.Printf("stała ma wartość %v i typ %T\n", MyVar4, MyVar4)
	fmt.Printf("stała ma wartość %v i typ %T\n", one, one)
	fmt.Printf("stała ma wartość %v i typ %T\n", two, two)
	fmt.Printf("stała ma wartość %v i typ %T\n", three, three)
	fmt.Printf("stała ma wartość %v i typ %T\n", four, four)

}
