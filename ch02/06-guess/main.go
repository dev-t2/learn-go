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
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

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

	fmt.Println(target)

	if guess < target {
		fmt.Println("Your guess was LOW")
	} else if guess > target {
		fmt.Println("Your guess was HIGH")
	}
}