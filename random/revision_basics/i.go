package main

import "fmt"

func main() {
	//tworze instancje
	nazwakraju := kraj("Polska")
	fmt.Println(nazwakraju.describe())

}

//Tworzę klase
type kraj string

//Tworze metodę klasy
func (k kraj) describe() string {
	return string(k) + " to kraj miodem płynący"
}
