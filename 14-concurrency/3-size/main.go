package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(channel chan int, url string) {
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

	channel <- len(body)
}

func main() {
	urls := []string{"https://go.dev", "https://go.dev/doc"}
	channel := make(chan int)

	for _, url := range urls {
		go responseSize(channel, url)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-channel)
	}
}