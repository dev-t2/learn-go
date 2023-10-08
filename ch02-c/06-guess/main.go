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
	isSuccess := false

	fmt.Println("\nI've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("\nYou have", 10 - guesses, "guesses left")
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
			fmt.Println("Your guess was LOW")
		} else if guess > target {
			fmt.Println("Your guess was HIGH")
		} else {
			isSuccess = true

			fmt.Println("You guessed it")

			break
		}
	}

	if !isSuccess {
		fmt.Println("\nIt was:", target)
	}
}