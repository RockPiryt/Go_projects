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
	fmt.Printf("Computer choose %d \n", target)

	// User input number via keyboard
	fmt.Println("Can you guess?")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

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
	} else if guess > target {
		fmt.Println("To High!")
	} else {
		fmt.Println("Congratulations! You guess number!")
	}

}
