package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
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
	go responseSize("https://go.dev")
	go responseSize("https://go.dev/doc")

	time.Sleep(time.Second * 5)
}