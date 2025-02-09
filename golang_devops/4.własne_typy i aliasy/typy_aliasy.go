package main

import "fmt"

//własne typy i aliasy deklaracja
type MyTypeInt int64 // zakrycie int64
type MyTypeAlias = MyTypeInt //zakycie MyTypeInt


func main() {
	var X MyTypeInt = 2
	var Y MyTypeAlias = 3
	var Z MyTypeAlias = X
	fmt.Printf("stała ma wartość %v i typ %T\n", X, X)
	fmt.Printf("stała ma wartość %v i typ %T\n", Y, Y)
	fmt.Printf("stała ma wartość %v i typ %T\n", Z, Z)
	
	myTue := true
	myFalse := false
	var myTrue2 bool = true
	var myFalse2 bool
	fmt.Printf("stała ma wartość %t i typ %T\n", myTue, myTue)
	fmt.Printf("stała ma wartość %t i typ %T\n", myFalse, myFalse)

	fmt.Printf("stała ma wartość %t i typ %T\n", myTrue2, myTrue2)
	fmt.Printf("stała ma wartość %t i typ %T\n", myFalse2, myFalse2)

	var i int
	var i8 int8 // max 127 liczba
	var i16 int16
	var i32 int32
	var i64 int64

	var u uint // nie może być ujemna, powyżej zera
	var u8 uint8
	var u32 uint32


	fmt.Printf("stała ma wartość %d i typ %T\n", i, i)
	fmt.Printf("stała ma wartość %d i typ %T\n", i8, i8)
	fmt.Printf("stała ma wartość %d i typ %T\n", i16, i16)
	fmt.Printf("stała ma wartość %d i typ %T\n", i32, i32)
	fmt.Printf("stała ma wartość %d i typ %T\n", i64, i64)

	fmt.Printf("stała ma wartość %d i typ %T\n", u, u)
	fmt.Printf("stała ma wartość %d i typ %T\n", u8, u8)
	fmt.Printf("stała ma wartość %d i typ %T\n", u32, u32)
	

	i32 = -100
	fmt.Printf("Po modyfikacji %d i typ %T\n", i32, i32)

	// i8 = 128//pokazuje że nie można już tej wartośći, max 127
	// u8 = -12//pokazuje że nie można ujemnej i wiekszej od 127


	prec := 1.5 //domyślnie float 64
	var prec2 float32
	prec3 := 1.499999999 // zaokągli do góry ma 1.5
	// sum := prec2 + prec3 // nie uda się bo rózne typy
	sum2 := prec + prec3

	fmt.Printf("stała ma wartość %f i typ %T\n", prec, prec)
	fmt.Printf("stała ma wartość %f i typ %T\n", prec2, prec2)
	fmt.Printf("stała ma wartość %f i typ %T\n", prec3, prec3)
	fmt.Printf("stała ma wartość %f i typ %T\n", sum2, sum2)

	var num int32
	var runNumber rune

	fmt.Printf("num ma wartość %d i typ %T\n", num, num)
	fmt.Printf("runNumber ma wartość %d i typ %T\n", runNumber, runNumber)
	num = 4
	runNumber = 5
	sum5 := num+runNumber
	fmt.Printf("po przypisaniu ma wartość %d i typ %T\n", num, num)
	fmt.Printf("po przypisaniu ma wartość %d i typ %T\n", runNumber, runNumber)
	fmt.Printf("po przypisaniu ma wartość %d i typ %T\n", sum5, sum5)

}