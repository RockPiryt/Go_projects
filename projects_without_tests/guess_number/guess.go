package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Computer choose random interger between 1-100
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Computer choose random interger between 1-100")
	//fmt.Printf("Computer choose %d \n", target)

	// Create bufio.Reader type
	reader := bufio.NewReader(os.Stdin)
	// Create success variable to end for loop as user guesses
	success := false
	left_attempts := 10

	for guesses_attempts := 0; guesses_attempts < 10; guesses_attempts++ {
		fmt.Printf("Attemps left: %d \n", left_attempts)
		// User inputs a number via keyboard
		fmt.Println("Can you guess?")
		input, err := reader.ReadString('\n') // Read until enter
		// Check if there is no error during ReadSting func
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // remove enter
		guess, err := strconv.Atoi(input) // change string to int

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("To Low!")
			left_attempts--
		} else if guess > target {
			fmt.Println("To High!")
			left_attempts--
		} else {
			success = true
			fmt.Printf("Congratulations! You guess number in %d turn!\n", guesses_attempts)
			fmt.Printf("Value of success is %t", success)
			break
		}
	}
	if !success {
		fmt.Println("You didn't guess.")
		fmt.Printf("Correct answer is: %d", target)
	}

}
