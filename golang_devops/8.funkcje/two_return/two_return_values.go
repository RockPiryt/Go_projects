package main

import "fmt"

func AddInt(x, y int) int {
	return x + y
}

func Mymsg(m string) {
	fmt.Println(m)
}

func SecretPass(p string) (string, bool) {
	if p == "secret"{
		return "Pass ok", true
	} else {
		return "Not correct", false
	}
}

func main(){
	a := AddInt(2, 3)
	fmt.Println("wynik dodawania", a)

	Mymsg("Moja wiadomość po prostu")

	msg, result := SecretPass("blabla")
	// niepoprawne haslo
	if !result{
		fmt.Println(msg)
	}

	//uzywam ponownie tych samych zmiennych --------------------czesto uzywane
	msg, result = SecretPass("secret")
	// niepoprawne haslo
	if !result{
		fmt.Println(msg)
	} else{
		fmt.Println(msg)
	}

}