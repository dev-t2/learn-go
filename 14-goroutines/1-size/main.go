package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string) {
	fmt.Println(url)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(body))
}

func main() {
	responseSize("https://go.dev")
	responseSize("https://go.dev/doc")
}