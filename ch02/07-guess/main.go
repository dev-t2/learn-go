// Guessing Program
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	isSuccess := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println()
		fmt.Println("You have", 10 - guesses, "guesses left")
		fmt.Print("Make a guess: ")

		input, err := reader.ReadString('\n')
	
		if err != nil {
			log.Fatal(err)
		}
	
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
	
		if err != nil {
			log.Fatal(err)
		}
	
		if guess < target {
			fmt.Println("Your guess was low")
		} else if guess > target {
			fmt.Println("Your guess was high")
		} else {
			isSuccess = true

			fmt.Println("You guessed it")

			break
		}
	}

	if !isSuccess {
		fmt.Println()
		fmt.Println("It was:", target)
	}
}