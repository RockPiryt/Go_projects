package main

import (
	"fmt"
	"mojModul/matma"
	"mojModul/mystring"
	"mojModul/api"
)
func main(){
	fmt.Println("wynik dodawania", matma.Add(10,5))
	fmt.Println("wynik odejmowania", matma.Sub(10,2))
	fmt.Println("wynik odejmowania", matma.Multiply(2,4))
	result,err := matma.Divide(20, 2)
	if err != nil{
		fmt.Println("Dzielenie przez 0",err)	
	}else{
		fmt.Println("Wynik Dzielenia wynosi", result)
	}
	

	fmt.Println("odwrrocony string", mystring.Reverse("Leon"))

	api.StartServer()

}