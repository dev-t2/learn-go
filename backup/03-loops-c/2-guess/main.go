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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	target := rand.Intn(100) + 1
	
	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	
	reader := bufio.NewReader(os.Stdin)
	isSuccess := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("\nYou have", 10 - guesses, "guesses left")
		fmt.Print("Make a guess: ")

		str, err := reader.ReadString('\n')
	
		checkError(err)
	
		str = strings.TrimSpace(str)
		num, err := strconv.Atoi(str)
	
		checkError(err)
	
		if num < target {
			fmt.Println("Your guess was LOW")
		} else if num > target {
			fmt.Println("Your guess was HIGH")
		} else {
			isSuccess = true

			break
		}
	}

	fmt.Println()

	if isSuccess {
		fmt.Println("You guessed it")
	} else {
		fmt.Println("It was:", target)
	}
}