package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	url string
	size int
}

func responseSize(channel chan Page, url string) {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{url: url, size: len(body)}
}

func main() {
	urls := []string{"https://go.dev", "https://go.dev/doc"}
	page := make(chan Page)

	for _, url := range urls {
		go responseSize(page, url)
	}

	for i := 0; i < len(urls); i++ {
		page := <- page

		fmt.Printf("%s: %d\n", page.url, page.size)
	}
}