package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")

	var userinput int
	fmt.Scanf("%d", &userinput)

	n := userinput * 2

	fmt.Println(n)
	// n := 5

	_, msg := isPrime(n)

	fmt.Println(msg)

}

func isPrime(n int) (bool, string) {

	// If it is not prime number
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	if n < 0 {
		return false, " Negative numbers are not prime, by definition"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	//  If it is prime number
	return true, fmt.Sprintf("%d is a prime number", n)

}
