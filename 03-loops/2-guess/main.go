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
		fmt.Println("\nYou have", 10 - guesses, "guesses left")
		fmt.Print("Make a guess: ")

		str, err := reader.ReadString('\n')
	
		if err != nil {
			log.Fatal(err)
		}
	
		str = strings.TrimSpace(str)
		num, err := strconv.Atoi(str)
	
		if err != nil {
			log.Fatal(err)
		}
	
		if num < target {
			fmt.Println("Your guess was LOW")
		} else if num > target {
			fmt.Println("Your guess was HIGH")
		} else {
			isSuccess = true

			fmt.Println("\nYou guessed it")

			break
		}
	}

	if !isSuccess {
		fmt.Println("\nIt was:", target)
	}
}